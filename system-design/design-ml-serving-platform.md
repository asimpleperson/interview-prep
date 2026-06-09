# ML Design — ML Serving Platform

> Reusable across AI/infra rounds (Reddit, OpenAI, Anthropic, Meta, Databricks, Uber).
> First mocked for Reddit GenAI Platform. Companion to `design-ml-feed-ranking.md`
> and `../companies/reddit/ml-design-llm-gateway.md`.

## Mental model
A serving platform = a **uniform control plane over a heterogeneous data plane**. Any ML
team deploys a model via config; the platform owns routing, multi-tenancy, autoscaling,
rollout, and observability. Value is *not* a new inference kernel — you operate open-source
runtimes (Triton, vLLM/TGI, TF-Serving, KServe) and absorb the operational complexity.

Winning framing (say early): *"Three workload classes (CPU ranking, GPU deep, LLM gen) have
fundamentally different performance profiles, so one runtime is wrong. The platform's job is
one control plane, three data planes — and the single most important decision is keeping the
control plane out of the request path."*

---

## Requirements

**Functional**
- Deploy from registry → live, versioned, autoscaled endpoint via one config ("model as config").
- Uniform API over **(A) low-latency CPU ranking/ads, (B) GPU deep/embeddings, (C) LLM generation**.
- Online (sync) + batch (async) inference.
- Safe rollouts: canary, shadow, A/B, instant rollback.
- Multi-tenant isolation (30–50 teams, hundreds of models).

**Non-functional (SLOs drive design)**
- Latency: ranking p99 ≤ 50ms server-side; LLM = TTFT + tokens/sec, *not* a single p99.
- Availability 99.9%+ for feed-critical tier; graceful degradation (cached/fallback score, never fail the feed).
- Scale: millions QPS aggregate.
- Cost top-three: maximize GPU/CPU utilization.
- Velocity: self-serve deploy in minutes.

---

## High-level architecture

```
CONTROL PLANE (low QPS, stateful):
  Deploy API/CLI → Model Registry → Rollout Controller (canary/shadow, autoscale policy)
       (config)     (immutable,        → reconciles desired→actual via K8s/KServe CRDs
                     versioned,
                     schema/signature)

DATA PLANE (hot path, millions QPS):
  Client → Inference Gateway → Router → [ Tier A CPU | Tier B GPU | Tier C LLM ]
           (auth, quota,       (version           Triton    Triton    vLLM/TGI
            validate, timeout,  pinning,          micro-    dynamic   continuous
            fallback)           load/cache-aware)  batch     batch     batch+paged-attn
                                              └─ Online Feature Store (low-latency join)
Cross-cutting: observability (metrics/traces/model-quality), GPU pool/scheduler, model cache/CDN
```

**Why control/data split is the make-or-break:** lets one deploy API + registry + rollout +
observability serve a microsecond CPU ranker and a token-streaming LLM, while each tier picks
the right runtime and scaling signal. Control plane must never be in the request path.

---

## Follow-ups + model answers

### 1. The 50ms budget (ranking + features)
Budget it explicitly: gateway 1–2ms / feature fetch+join 10–15ms / inference 5–15ms /
ser+net 3–5ms / ~10ms tail headroom. Key insight: **ranking scores N candidates per request**,
so it's a batch-of-candidates call and feature fan-out dominates.
- Batch the whole candidate set into one inference call (never N calls).
- Two-sided fetch: user features once/request; item features per-candidate but heavily cached (slow-changing).
- Strict feature-store timeout → degrade (last-known-good cache → default/mean-imputed flagged → cheaper model/precomputed score). **Feed always returns a ranking.**
- Hedged requests at p95 to cut tail (~5% extra load).
- Persistent gRPC conns, pre-warmed/compiled models; cold starts are the autoscaler's job, never on the hot path.
- Co-location/sidecar mode as a *config option* of the same platform for the hottest path (kills a hop) — keep one control plane.

### 2. Multi-tenancy + GPU cost
Enemy = GPU fragmentation (whole A100 doing 20 QPS). Four fronts:
- **Multi-model packing:** Triton instance groups + MIG to slice big GPUs; scheduler bin-packs by mem+compute. 5–10x util on the long tail.
- **Tiered autoscaling + scale-to-zero** for the tail (warm pool bounds cold start). LLMs scale on batch occupancy / KV-cache util, not CPU%.
- **GPU broker + priority classes:** Guaranteed (feed-critical, reserved, never preempted) / Burstable (spare capacity, preempted first) / Batch (leftover GPU, fills troughs). Preemption stops one team's spike starving another.
- **Isolation for the critical path:** MIG hardware isolation, per-tenant gateway quotas, physically separate feed-critical from experimental.
Tradeoff stated explicitly: *utilization-maximizing for the tail, isolation-maximizing for the critical path* — exposed as the **service tier** knob picked at deploy. Add **showback/chargeback** (GPU-seconds per model/team) — changes behavior.

### 3. LLM tier — what's different
- **Two latency metrics:** TTFT vs TPOT (throughput) decoupled → streaming (SSE/gRPC) end-to-end, no buffering.
- **Continuous batching changes routing:** autoscale on running+waiting sequences & KV-cache util; router is **load/KV-aware** and **prefix-cache-aware** (route a conversation to the pod holding its KV cache). Round-robin thrashes caches.
- **Memory is binding:** huge weights → warm replica floor + stream weights from fast store/NVMe/CDN; quantization (FP8/INT8/AWQ); tensor/pipeline parallelism; KV-cache admission control + max-context limits.
- **New platform features:** semantic/prompt caching; guided/structured decoding; **LoRA/adapter multiplexing** (many fine-tunes off one base = LLM analog of model packing); token-based quotas/cost; small→large cascade routing.
- Shares the *control plane*, needs a different *data plane + scaling/routing signals* — exactly why control/data were split up front.

### 4. Safe rollouts — silent quality regression (healthy pod, garbage predictions)
The scary failure isn't a crash. Defense in depth:
- **Before traffic:** schema/signature validation at registration; **shadow** (mirror live traffic, compare output distribution offline); **canary 1→5→25→100%** with automated analyzer on ops + prediction-distribution + proxy-business metrics, auto-rollback on breach.
- **Detect bad-but-healthy:** prediction-distribution monitoring as a platform primitive; **online/offline skew** detection (replay sampled online (features,pred) through offline model); gateway sanity guardrails (range, not-all-identical, not-NaN — variance check catches "same score for every candidate").
- **Containment:** instant rollback = config flip (prev version warm, deterministic via immutability); blast-radius caps via canary %; kill switch → known-good fallback model.
Staff point: liveness/readiness catch crashes, say nothing about quality. Platform's differentiated value = make prediction-quality a monitored, gated, auto-rollback-able signal.

### 5. One decision + V1 scope
Make-or-break = **control/data split, uniform control plane over heterogeneous data plane.**
V1 (one quarter): control plane (registry + declarative deploy + KServe rollout) + **Tier A ranking only** (batched candidate scoring + feature degradation) + gateway (auth/quota/timeout/fallback) + canary/shadow/rollback + prediction monitoring + showback.
Cut/defer: LLM tier (fast-follow Q2 on same control plane), scale-to-zero/MIG/semantic cache/LoRA (cost opts come *after* reliability), auto-bandit AB.
Sequencing logic: ranking team trusting you with the main feed is the adoption gate; lead with cost/LLM flash and nobody trusts you with prod traffic.

---

## Staff signals
1. Drive with SLOs and explicit latency budgets, not component lists.
2. Name the core tension early (heterogeneous workloads, one platform) and reuse it as the spine.
3. Make tradeoffs a config knob (Guaranteed/Burstable/Batch service tiers) not hand-waving.
4. Treat ML reliability ≠ service reliability (silent quality failures, degradation, skew).
5. Sequence V1 to earn trust before chasing cost/LLM flash.

## Reference systems to name-drop
KServe/Knative, Triton (instance groups, dynamic batching, MIG), vLLM/TGI (continuous batching,
paged attention), Ray Serve, BentoML, Seldon. NVIDIA MIG. AWQ/GPTQ quantization, LoRA.
