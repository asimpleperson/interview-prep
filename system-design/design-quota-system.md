# Design: Usage Quota System (high concurrency, low latency)

> **First mocked for:** Snowflake — System Design screen (2026-06-23), framed for the Billing
> Platform role. Directly relevant to Snowflake usage-governance / credit-tracking.
> **Prompt:** Design a quota system that tracks user usage across backend services. e.g. user has
> 100GB storage quota; Drive used 30GB, Photos used 20GB. Prevent users from exceeding quota.
> **Emphasis:** high concurrency, low latency.

## Requirements

**Functional**
- `CheckAndReserve(account, service, bytes) → granted | denied`
- `Commit(reservationId)` / `Release(reservationId)`
- `GetUsage(account)` — breakdown per service.
- Admin: set/update limits; hierarchical rollups (org → user → service).

**Non-functional (the real meat)**
- High concurrency: many services update the same account's usage concurrently → **race conditions
  are the core problem**.
- Low latency: p99 single-digit ms on the enforcement (write) path.
- Availability: quota system down ≠ business hard-stop → explicit fail-open/closed policy.
- Correctness: must *converge*; **bounded** temporary overshoot acceptable, unbounded is not.

**Key assumptions to state**
- Storage quota tolerates small, bounded, temporary overshoot but must converge (most important
  decision — drives everything).
- Hierarchical: account total with per-service sub-buckets.
- Per-account write QPS modest; fleet-wide huge (millions of users, 100K+ checks/sec).
- We are on the **critical path of every upload** → tight latency budget (p99 < 5–10ms).

## Capacity Estimation (rough)
- 500M accounts, 100M DAU; peak ~100K reserve/sec fleet-wide.
- Counter row tiny (~tens of bytes/account-service). Hot path must avoid a DB round-trip.

## High-Level Design

```
 Drive ─┐     ┌──────────── Quota Service (stateless) ───────────┐
 Photos ┼──▶  │  CheckAndReserve / Commit / Release / GetUsage    │
 other ─┘     └───────────────────────┬───────────────────────────┘
                  ┌───────────────────┼────────────────────┐
                  ▼                   ▼                      ▼
          In-process lease      Distributed counter     Durable source of truth
          (local budget,        (Redis Lua CAS /         (Spanner/DynamoDB/FDB,
           µs hot path)         sharded in-mem)           atomic conditional write)
                                        │
                                        ▼
                          Async reconciler vs object-store metadata
                          (object store = ground truth; counter = accelerator)
```

**Core idea:** never do a read-modify-write against the durable store on the hot path — that is
the source of both the latency *and* the concurrency problem. Use a **two-tier
reservation/lease (escrow) model** plus an authoritative async reconciler.

## Deep Dive

### The race condition (why naive fails)
`SELECT usage; if usage+bytes<=limit UPDATE` → two concurrent writes both read 90GB, both think
+5GB fits, both commit → 100GB busted. Classic read-modify-write race, *and* a DB hop on the hot
path.

### Three escalating mechanisms
- **A — Single atomic counter w/ conditional update.** Redis `INCRBY` in a Lua script that checks
  limit, or DynamoDB conditional write, or Spanner RW txn. One round trip, atomic, no race.
  *Problem:* per-account **hot key** serializes a heavy account through one shard.
- **B — Escrow / lease (default).** Quota node grabs a **lease** (a chunk, e.g. 1GB) from the
  central counter, then serves many `Reserve` calls **locally in-memory** (µs, no network). Refill
  asynchronously at a low watermark. → central counter sees ~1 req per chunk, not per write;
  throughput ×chunk-size; p99 = µs except occasional refill. *Tradeoff:* idle leases **trap
  budget**; overshoot ≤ Σ(outstanding leases) — **bounded**, which we allowed. TTL + reconciler
  handle crashes.
- **C — Strict serialization** (Spanner/FDB txn) for hard-limit accounts. ~10ms+ but exact.

→ **Tiered enforcement:** escrow for the 99% (fast, bounded overshoot), strict txn for the few
that need exactness. *Don't pay for strong consistency where it isn't required.*

### Data model
```
Quota:       (account, service) → limit_bytes, policy
Usage:       (account, service) → used_bytes        (authoritative, durable)
Reservation: (resId, account, service, bytes, state, ttl, lease_node)
Lease:       (leaseId, node, account, granted, consumed, ttl)
```
`used = committed + Σ(outstanding reservations)`. Enforce against `used`, not committed, so
in-flight uploads can't collectively bust the limit. Lifecycle: Reserve → upload → Commit / Release
/ TTL-expiry (TTL = safety net for crashed clients).

### Sharding & hot keys
Shard by `account` (consistent hashing) → co-locates services for cheap totals. Extreme account:
split counter into K sub-counters (`account:shard0..K`), sum on read, reserve hits a random one.

### Storage
Hot tier Redis (Lua atomic check-decrement) write-through over Spanner/DynamoDB/FDB. Object store =
ground truth; reconciler recomputes authoritative usage from object metadata and corrects drift.

### Failure / consistency policy
- Quota service down → soft-limit accounts **fail-open + circuit breaker + later reconciliation**;
  hard-limit accounts **fail-closed**. Explicit product choice.
- Lease-holding node dies → TTL reclaims un-consumed bytes; self-heals.
- Reconciler diffs counter vs object metadata; alerts on large drift (bug/abuse signal).

## Tradeoffs / Staff-level bets
1. Reservation/escrow kills the race condition **and** the hot-key latency in one move.
2. Tiered consistency — strong only where required.
3. Object store = ground truth; counter = reconciled accelerator.
4. **Lease size is THE knob:** bigger lease = lower latency but more overshoot + fragmentation.

## Follow-ups practiced (with answers)
- **Zero overshoot for hard-limit accounts?** Can't have {0 overshoot, no hot counter, low latency}
  for the same account. Classify by policy → route hard-limit to Option C (atomic conditional write,
  ~10ms). Overshoot bound = Σ lease remainder; set lease→0 ⇒ overshoot→0. Escrow is a *tunable knob*,
  not binary.
- **Counter vs laggy object store / double counting?** Enforce against the **counter** (object store
  too laggy as an oracle). Idempotency keyed on `resId`; Commit idempotent; reconciler matches files
  to resId via object metadata, finds orphans both directions. At any instant counter ≥
  object-visible (reserve precedes physical write); converges after lag window.
- **Hierarchical budget fragmentation (1TB shared by 10K users; budget trapped in idle leases)?**
  (1) Dynamic/proportional lease sizing (AIMD — idle nodes tiny leases, hot nodes big). (2) Short
  TTL + aggressive reclaim + proactive release on heartbeat. (3) **Adaptive mode switch:** below a
  watermark (near limit) switch org out of escrow into direct-counter mode — exactness matters more
  than latency when tight. (4) Hierarchical leasing (org→regional aggregator→node). Cure =
  small, short-lived, demand-proportional leases + strict mode as utilization → 100%.
- **p99 tail under synchronous refill?** (1) **Low-watermark prefetch** (async refill at ~20% left)
  — biggest tail-killer, converts stall to background op. (2) Jitter watermark + randomize lease size
  to avoid synchronized refill / thundering herd. (3) Hedge / fast timeout on refill → fail-open
  grant for soft accounts. (4) Bound counter latency (in-mem, co-located, shard hot accounts).
  Tail = f(refill frequency) = f(lease size) → tune per account class from real traffic.

## Meta-takeaway
Security/consistency parameters *are* the scalability parameters — the same dial. Lease size trades
latency ↔ overshoot/fragmentation; pick per account class, measured from real traffic.
