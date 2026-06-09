# ML System Design — Index

Cross-cutting tracker for all ML/AI system-design prep in this repo. ML design rounds show up
across AI-forward companies (Reddit, OpenAI, Anthropic, Meta, TikTok/ByteDance, Netflix, Uber,
Databricks), so the reusable write-ups live here in `system-design/` rather than in any one
company folder. Company-specific framing is linked from each entry.

## Write-ups

| # | Topic | File | Type | First mocked for | Status |
|---|-------|------|------|------------------|--------|
| 1 | LLM Gateway / GenAI platform | `../companies/reddit/ml-design-llm-gateway.md` | Infra / control plane | Reddit GenAI Platform | Done |
| 2 | ML Serving Platform (CPU/GPU/LLM tiers) | `design-ml-serving-platform.md` | Infra / platform | Reddit (mock) | Done |
| 3 | Personalized Feed Ranking | `design-ml-feed-ranking.md` | ML modeling / recsys | Reddit (mock) | Done |

> #1 is intentionally kept under `companies/reddit/` because it was written as role-specific
> onsite prep; #2 and #3 are generalized here for reuse. All three are mutually cross-linked.

## How these connect
- **#1 LLM Gateway** is a *specialization* of #2's Tier C (LLM) data plane + gateway — narrower,
  focused on multi-provider routing/caching/cost. Read #2 for the broader platform, #1 for the
  GenAI-specific deep dive.
- **#3 Feed Ranking** *consumes* #2: the ranking model is deployed as a Tier A (low-latency CPU,
  batched candidate scoring) workload on the serving platform, and its rollout/monitoring ride
  the platform's canary/shadow/quality-monitoring machinery.

## Recurring staff-level themes (transfer across any ML design prompt)
- **Control plane vs data plane** separation; control plane never in the request path.
- **Latency budgets** stated explicitly; for LLMs, TTFT + tokens/sec, never a single p99.
- **ML reliability ≠ service reliability:** silent quality regressions, online/offline skew,
  prediction-distribution monitoring, graceful degradation (never hard-fail the user surface).
- **Objective design:** optimize a proxy validated against the true long-horizon goal; separate
  *what the model predicts* from *how the business values it*.
- **Feedback loops / selection bias:** exploration + propensity correction (IPW/IPS) +
  outcome constraints (diversity, ecosystem health).
- **Evaluation as three gates:** offline filter → online A/B on the real metric → guardrails.
- **Cost as a first-class axis:** GPU packing/MIG, scale-to-zero, priority/preemption, chargeback.
- **Sequencing V1:** ship a correctly-shaped, modest system that earns trust, then climb.

## To add later (candidate prompts)
- Embedding / vector search infra (ANN index lifecycle, freshness, sharding).
- RAG pipeline (chunking, retrieval, re-ranking, generation, eval).
- Content safety / abuse classification at scale (real-time + human-in-the-loop).
- Ads CTR/CVR prediction & auction.
- Real-time feature store / streaming features.
