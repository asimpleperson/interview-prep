# ML Design Prep — Design the LLM Gateway (Reddit GenAI Platform)

> Most likely ML/systems design prompt for the GenAI Platform role. Onsite: Mon 2026-06-08.
>
> Part of the cross-cutting ML system design set — see `../../system-design/ml-system-design-index.md`.
> Broader platform context: `../../system-design/design-ml-serving-platform.md` (this gateway is a
> specialization of that platform's Tier C LLM data plane).

## Mental model
An LLM Gateway = a **reverse proxy / control plane for LLM traffic**. Every team's app calls
the gateway instead of OpenAI/Anthropic/internal vLLM directly. It centralizes cross-cutting
concerns: auth, routing, rate limiting, caching, fallbacks, cost tracking, observability,
safety. **Platform/infra problem, not a modeling problem.**

Winning framing (say early): *"My job is to make N product teams faster and safer while
controlling cost and reliability — clean unified API, tenant isolation, graceful degradation."*

---

## Answer structure

### 1. Clarify requirements (don't skip)
- Clients: internal teams (feed, search, ads, safety) + maybe external → multi-tenant.
- Providers: external (OpenAI, Anthropic, Google) + internal vLLM on K8s → abstract heterogeneous backends.
- Modalities: text chat/completions + embeddings; streaming required.
- Scale: state a number (~10K RPS, tens of M tokens/min, p99 budget). Reddit scale.
- Problems solved: cost control, reliability/failover, governance/safety, dev velocity, observability.

| Functional | Non-functional |
|---|---|
| Unified API across providers | Low added latency (<10–20ms p50) |
| Routing & model selection | High availability (no SPOF) |
| Streaming (SSE) | Horizontal scalability |
| Rate limiting & quotas/tenant | Tenant isolation (noisy neighbor) |
| Caching | Observability + cost attribution |
| Fallback / retries | Security (key custody, PII) |

### 2. API design
OpenAI-compatible schema (pragmatic — SDKs already speak it). One endpoint, model in body:
```
POST /v1/chat/completions
{ "model": "router/fast" | "gpt-4o" | "internal/llama-3-70b",
  "messages": [...], "stream": true, "metadata": {"team":"ads","feature":"..."} }
```
- Logical model names/aliases decoupled from physical backends (swap providers w/o client change).
- Streaming via SSE, normalized across providers.
- `metadata` drives cost attribution + rate limits.

### 3. High-level architecture
```
Clients → L7 LB → [ Gateway (stateless, scaled) ]
  AuthN/Z → Rate limiter → Cache lookup → Router → Provider adapters
                              │                          │
                       exact/semantic cache     OpenAI/Anthropic/vLLM adapters
  → async emit → Kafka (usage) → billing/dashboards/audit
Stores: Redis (rate limit + cache), Config svc (policies), Vault (keys), Vector store (semantic cache)
```
**Key principle: gateway is STATELESS.** All shared state in external stores → trivial horizontal scaling. This is the answer to "how do you scale it."

### 4. Request path, component by component
- **Auth:** API key/mTLS/JWT → tenant + quota. Provider keys injected from Vault, never held by clients (centralized custody + rotation).
- **Rate limiting:** Redis token bucket per tenant. Limit RPS **and** tokens/min (cost is token-driven). Priority tiers (prod > batch).
- **Caching:** exact-match (hash prompt+params; great for embeddings) + semantic (embed → ANN → return if cosine > threshold). Semantic is opt-in, risky for correctness; respect temperature/freshness.
- **Router:** logical→backend by cost/latency/capability/health/load. Mention latency-aware/least-loaded + optional model-based routing (cheap model for easy prompts).
- **Provider adapters:** normalize request/response/stream/error per provider; token counting.
- **Resilience:** retries w/ backoff+jitter, fallback chains (primary 429 → secondary → internal), circuit breakers per backend, timeouts + hedging. Graceful degradation, never hard-fail.
- **Observability:** async-emit usage event/request (tokens, cost, latency, model, tenant, cache-hit) to Kafka. Never block request path on logging. Tracing + audit log.

### 5. Cross-cutting (Platform role bonus)
- **Guardrails:** pluggable pre/post middleware (PII redaction, prompt-injection, toxicity, policy). Every team gets safety for free.
- **Prompt/config versioning + A/B routing** (weighted/canary).
- **Cost governance:** per-team budgets, alerts, hard caps.

### 6. Tradeoffs / bottlenecks
- Extra hop → keep thin/stateless, pool connections, fast path for streaming.
- Not a SPOF: multi-AZ, fast failover, degraded passthrough mode.
- Redis is hot dependency → cluster; define fail-open (cache) vs fail-lenient (rate limit) behavior.

---

## Follow-ups + model answers

**Latency overhead low?** Stateless thin proxy; conn pooling + HTTP/2 keep-alive; Redis pipelining or L1 in-mem cache; async logging off hot path; stream first byte immediately. Measure TTFT for streaming.

**Rate limiting across instances?** Central Redis counters (atomic Lua token bucket) = global limit. To avoid Redis per request: local bucket + periodic sync (sliding window approx) — trade accuracy for latency.

**Semantic cache + risks?** Embed → ANN (HNSW) → return if sim > threshold. Risks: false hits, ignores context, bad for high-temp/time-sensitive. Mitigate: opt-in, conservative threshold, exclude temperature>0, scope by tenant. Exact-match is safe default.

**Provider down / 429?** Circuit breaker trips → router reroutes → fallback chain → retries w/ backoff+jitter. Health checks + half-open recovery. Surface degradation in metadata. Degrade gracefully.

**Cost attribution / budget control?** Tag requests w/ tenant/feature → Kafka → per-team aggregation. Token quotas + budget caps in Redis; soft alert → hard throttle. Dashboards.

**Streaming across providers?** Adapters translate each provider's chunks to normalized SSE; proxy as they arrive (no buffering); tally tokens for end-of-stream usage event; handle mid-stream errors + client disconnects (cancel upstream).

**Roll out new model / A/B?** Config supports weighted/percentage routing per logical name. Shadow traffic → compare offline → canary 1%→10%→100% w/ auto-rollback on regression. Decouples model change from client deploy.

**Add safety/guardrails?** Pluggable middleware chain pre/post call; centralized so policy ships once; configurable per route to manage latency.

**Multi-region/HA?** Gateway per region, nearest routing, regional Redis, replicated config, stateless scaling, multi-AZ + health LB.

**Long/batch requests?** Async path: submit → job ID → poll/webhook; lower-priority queue + cheaper backends so batch doesn't starve interactive prod.

---

## Staff signals (not senior)
1. Lead with requirements, tenants, tradeoffs before drawing.
2. Tie every decision to cost / reliability / developer velocity.
3. Name failure modes + graceful degradation unprompted (SPOF, Redis down, provider outage, noisy neighbor).

## Reference systems to name-drop
LiteLLM, Cloudflare AI Gateway, Portkey, Kong AI Gateway, vLLM (continuous batching, KV cache).
