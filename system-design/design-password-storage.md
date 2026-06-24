# Design: User Password Storage & Verification (where are the bottlenecks?)

> **First mocked for:** Snowflake — System Design screen (2026-06-23).
> **Prompt:** Design a system to store user passwords; then, based on that system, identify the
> bottlenecks. *(The second half is the real question — reason about limits, don't just recite
> bcrypt.)*

## Requirements

**Functional:** password-at-rest storage + a verification (login) path. Register / change-password.
(Not full SSO/session mgmt, but note where it connects.)

**Non-functional / threat model (drives everything):**
- Must stay secure **even if the DB is fully exfiltrated**. This single requirement drives the design.
- Login is interactive: p99 a few hundred ms acceptable — and we *deliberately spend* most of it.

**Scale assumptions:** 500M users, 100M DAU; peak 50–100K login attempts/sec (incl. credential-
stuffing spikes, post-outage thundering herd).

## Capacity Estimation (the twist)
KDF tuned to ~50–100ms CPU **and** ~64MB RAM per hash (on purpose). At 50K logins/sec:
- CPU: 50K × 75ms ≈ **3,750 cores** just for hashing.
- RAM: 50K concurrent × 64MB ≈ **3.2 TB** transient at peak.
→ The auth fleet is sized by your **threat model**, not your user count.

## High-Level Design

### Storage (non-negotiables)
Never plaintext, never reversible encryption, never *fast* hashes (MD5/SHA-256 → offline brute
force / rainbow tables fall instantly). **Store a slow, salted, memory-hard hash.**

```
users_credentials:
  user_id (PK)
  algo          e.g. "argon2id"        ← store the algorithm
  params        e.g. m=64MB,t=3,p=1     ← store the cost params (lets you migrate cost upward)
  salt          32B, unique per user, CSPRNG
  hash          derived key
  updated_at, needs_rehash
```
- **Argon2id** (memory-hard; resists GPU/ASIC) — or scrypt, or bcrypt (conservative fallback).
- **Per-user random salt** — defeats rainbow tables; equal passwords → different hashes. Not secret,
  just unique.
- **Store algo+params with each hash** → raise cost over time via **rehash-on-login** (`needs_rehash`).
- **Pepper (staff add):** secret key in **HSM/KMS, not the DB**. HMAC password pre-hash, or encrypt
  the stored hash. A DB-only leak (no HSM key) is then useless. Separates the secret from the data.

**Deliberate-slowness paradox:** slow + memory-heavy is a *security feature* against attackers — and
it is exactly what creates our bottleneck.

### Verification (login) path
```
client → gateway → Auth svc → fetch cred by user_id
                            → KDF(password, salt, params)   ← the expensive step
                            → constant-time compare         ← no timing side channel
                            → if params stale → rehash & store
                            → issue session token / JWT
```
- **Constant-time compare** (no early-exit `==`).
- **Don't leak user existence:** same error both cases; run a dummy KDF on unknown user so response
  time doesn't reveal account existence (costs CPU — bottleneck theme again).

## Deep Dive — Bottlenecks (the main event)

1. **KDF CPU+RAM — THE bottleneck, by design.** ~3,750 cores + ~3.2TB RAM at peak (above). Every
   notch of security (higher t/m) linearly raises fleet size/cost. Mitigate: CPU/RAM-aware
   autoscaling provisioned for *attack* peaks; admission control / load shedding; **cap concurrency
   per node to its RAM** (128GB node ⇒ ~2K concurrent 64MB hashes, else thrash).
2. **Credential stuffing turns the security feature into a DoS amplifier.** Each guess = one cheap
   HTTP request for attacker, one full 75ms/64MB op for us. **Defend *before* the KDF:** rate limit
   (IP/account/device), lockout + exponential backoff, bot defense (CAPTCHA/PoW/fingerprint),
   breached-password checks, push to MFA/passkeys.
3. **Datastore — counterintuitively the *least* stressed tier.** Indexed `SELECT by user_id` is
   sub-ms; KDF dwarfs it ~1000×. "Shard the DB!" is the *wrong* instinct. Replicate for HA;
   **don't cache hashes** (security regression + pointless — the fetch isn't the cost).
4. **Pepper / KMS / HSM = SPOF on the hot path** if you call KMS per login. Keep pepper **in memory**
   on auth nodes (fetched once at boot); HMAC locally. Envelope encryption w/ cached data key — no
   per-login KMS round trip.
5. **Thundering herd / correlated load** (post-outage, forced mass reset). Expensive per-login makes
   herds far worse. Jittered retries, queueing, headroom, sane session TTLs.
6. **Operational:** rehash-on-login can't reach dormant accounts (no plaintext); rolling a leaked
   pepper is hard.

**Bottleneck hierarchy:** KDF CPU+RAM dominates and sizes the fleet → credential stuffing weaponizes
it (defend upstream) → datastore is nearly free → KMS/pepper is a SPOF to keep off the per-request
path. **Meta-point: security parameters *are* scalability parameters — the same dial.**

## Follow-ups practiced (with answers)
- **Cut the ~3,750 cores without weakening security?** Core lever: pay the KDF only when you must.
  (1) **Sessions, not re-auth** — KDF once per session, not per request; DAU hashes ~once/day. The
  3,750 figure is *peak concurrent fresh logins*, not steady state. (2) Shed attack traffic upstream.
  (3) Right-size params to the highest the *organic* p99 tolerates; lean on rate-limiting for the
  online fight (hash cost mainly protects the *offline* post-leak case). (4) Spot/preemptible +
  autoscale (stateless, bursty). **Not** by lowering memory-hardness or caching hashes — win via
  *fewer KDF invocations*, not a cheaper KDF.
- **Quantify what salt+slow-hash buys vs SHA-256? Why 64MB not 1GB?** Salt kills precomputation
  (cost scales users×guesses, not a one-time table). Slowness sets per-guess cost: SHA-256 ~10^10
  guesses/sec on a GPU (8-char falls in seconds); Argon2id 64MB×t=3 ~75ms + memory-bound ⇒ ~6–9
  orders of magnitude fewer guesses, ASICs ineffective (memory-bandwidth bound). **64MB not 1GB**
  because *we* pay the same memory on the legit path at full QPS — 1GB ⇒ ~16× our RAM/latency for
  marginal attacker pain. Tune memory to the largest value that fits latency+fleet budget (same dial
  as cost). Caveat: doesn't save `password123` — slow hash buys time for *decent* passwords + limits
  blast radius; not a substitute for breach detection + MFA.
- **Where does a pepper actually help vs theater?** Helps on **DB-only** leak (SQLi, leaked backup,
  exposed snapshot, insider DB read) — pepper in separate trust domain ⇒ hashes uncrackable. Theater
  if the **same breach roots the app servers** (pepper read from process memory). Value = raises bar
  from "leaked one system" to "leaked two independent systems." Only real if pepper truly lives in a
  separate trust domain (real HSM, separate custody) — not in the same DB/config.
- **Passkeys / WebAuthn — how does it change the analysis?** Dissolves the bottleneck: server stores
  a **public key**, verification = one asymmetric signature check (~µs–ms, no memory-hardness — a
  leaked public-key DB is worthless). Bottleneck #1 disappears; #2 (stuffing DoS) largely gone
  (nothing expensive to weaponize, phishing/replay beaten by challenge-response + origin binding).
  Bottlenecks **shift** to **account recovery** (new weakest link — email/SMS reset reintroduces a
  password-class secret), multi-device key sync/attestation, and the password+passkey **migration**
  period (still carry the KDF fleet for the long tail). Takeaway: the whole analysis stems from
  storing a **shared secret**; asymmetric credentials invert the cost model.

## Tradeoffs / Meta-takeaway
In password systems, your **security parameters are your scalability parameters** — security cost,
fleet size, and latency budget are one dial. The strategic direction is to retire the shared secret
(passkeys), which converts an expensive verification into a cheap one and moves the hard problems to
recovery and migration.
