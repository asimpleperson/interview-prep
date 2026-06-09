# ML Design — Personalized Feed Ranking System

> Reusable across recommendation/ranking rounds (Reddit, Meta, TikTok/ByteDance, Netflix, Uber).
> First mocked for Reddit Home feed. Companion to `design-ml-serving-platform.md`.
> This is an **ML modeling** design — spend time on objective, candidate gen, model, training,
> and the experimentation loop, not raw infra.

## Mental model
Core problem: given user `u`, time `t`, context `c`, **select + order posts to maximize
expected long-term retention-driving engagement**, subject to hard safety/policy filters,
freshness, and diversity. Ranking at scale = a **funnel**: cheap-and-wide → expensive-and-narrow.

Winning framing (say early): *"Retention isn't optimizable per-request, so I predict calibrated
near-term events that correlate with retention and combine them with a business-tunable value
function — and I keep 'what we predict' separate from 'how we value it'. The make-or-break is
getting the objective right, not the architecture."*

---

## Requirements

**Objective:** north star = long-term retention (D7/D30, sessions/week) + healthy engagement
(comments, upvotes, community joins). **Do NOT single-objective on dwell time** → clickbait/churn.

**Functional:** ordered paginated feed leaving ad slots; blend subscribed + discovery;
cold-start (new user/post/community); hard filters (safety/NSFW/community rules/seen) + soft
constraints (diversity, freshness, fatigue).

**Non-functional:** ranking latency tens of ms (batched candidate scoring on the serving platform);
new posts rankable within minutes; always return something (degrade → heuristic sort).

**Objective design (most important):** multi-objective proxy — predict per (user,post):
p(click), p(long-dwell), p(upvote), p(comment), p(share), p(subscribe), and **negatives**
p(hide/downvote/report). `score = Σ wᵢ·pᵢ`, weights calibrated to long-term value. Weight
*contribution* signals (comment/vote/subscribe) heavily; passive low; negatives strongly negative.

---

## High-level architecture (funnel)

```
Millions of posts
  → Stage 0  Candidate Generation / Retrieval  → ~thousands
       sources unioned: subscribed-recency pool · two-tower EBR (ANN, discovery)
       · collaborative/co-engagement · trending/fresh pool
       hard filters applied here (safety/NSFW/seen/blocked)
  → Stage 1  Light ranker (two-tower dot / small GBDT)  → ~hundreds
       optimize for RECALL of what heavy ranker likes (distillation target)
  → Stage 2  Heavy ranker (multi-task DNN, full features)  → scored ~hundreds
  → Stage 3  Re-ranking / policy  → final page
       combine multi-task → value · diversity (MMR) · freshness boost · fatigue/seen-decay
       · community balance · integrity demotions · leave ad slots
```

Rationale: each stage trades recall for precision under the latency budget. **Retrieval errors
are permanent** — a post not retrieved can never be ranked → bias retrieval toward recall.
Keep product constraints in Stage 3 (out of the model) so PM/policy tune without retraining.

**Heavy ranker:** multi-task MMoE DNN (shared bottom + per-objective heads so conflicting tasks
don't fight). Embeddings for sparse IDs; **sequence encoder (DIN/transformer attention over
recent history conditioned on candidate)** = biggest accuracy lever. Calibrated heads.

**Features:** user (long/short-term interest, activity, subs) · post (community, author, age,
media, topic emb, **early-engagement velocity**) · community (size/health/affinity/membership) ·
author (reputation) · user×post cross (affinity) · context (time, session depth, position).

**Training:** logged impressions+engagement → multi-label targets; impressed-not-engaged = hard
negatives + random easy negatives for retrieval. Retrain daily/continuous (trends move fast),
warm-start. Weighted BCE per head + calibration.

---

## Follow-ups + model answers

### 1. Objective & weights — how to set them, how to know it's retention not clicks
Separate **(a) what the model predicts** (calibrated near-term probs, stable/reusable) from
**(b) how we value them** (weights/value model, business-tunable). Lets PM shift emphasis
without retraining; lets you audit weights→retention.
- V1: hand-set weights from **retention-correlation analysis** (comment/subscribe correlate far more than passive click).
- Validate the proxy is real: **A/B where treatment changes weights, read out D7/D30 retention** (not in-session). If dwell lifts but D30 flat/down, proxy is lying → retune.
- Mature: tune weights via Bayesian-opt/bandit over a **long-term value (LTV) model**; RL framing (cumulative reward) is the endpoint but operationally heavy — earn it after supervised multi-task is solid.

### 2. Feedback loops & position bias (closed-loop / selection bias — existential)
Untreated → popular communities dominate, niche/new starved, world-view narrows, metrics look fine while ecosystem rots.
- **Position bias:** position-as-feature at train, neutralized at serve; **IPW** (weight by 1/p(examined)).
- **Exploration** (gather counterfactual data): ε / controlled randomization on a slot/session fraction; **uncertainty-aware (Thompson/UCB)** on high-variance items; log with propensity for unbiased train+eval.
- **Cold start:** new posts get freshness/early-velocity boost + guaranteed exploratory impressions; new users get onboarding + popularity priors + in-session adaptation; new communities get exploration budget.
- **Anti-monopoly constraints:** re-rank diversity caps; monitor ecosystem health (impression Gini, new-creator/community success) → feed back into exploration budget.
- **Offline eval respects bias:** counterfactual/off-policy (IPS, replay), paired with online A/B.
Staff point: it's a *systemic* risk → attack at 3 levels: debias labels, inject fresh info, constrain outcome. No exploration = silently dying, shows up in retention a quarter later.

### 3. Candidate generation depth (two-tower + freshness + recall)
- **Train:** user tower vs item tower, dot/cosine; **in-batch sampled softmax + logQ correction** (fixes popularity bias) + hard negatives (impressed-not-engaged) + random negatives. Negative-sampling quality *is* retrieval quality. Strong positives = comment/upvote/subscribe/long-dwell.
- **Freshness (hard for Reddit):** **content-based item tower** so a new post is embedded at creation (streaming) into ANN within seconds — no waiting for retrain (works day-zero). **Two pools:** big ANN index (evergreen) ∪ real-time fresh buffer (time-bounded). User emb near-real-time for session intent. Periodic full rebuild + incremental upserts; HNSW/IVF-PQ tuned.
- **Recall (not just fast):** measure recall@K vs actually-engaged held-out; multiple sources unioned (each covers others' failure modes); over-retrieve then prune (retrieval mistakes unrecoverable); logQ + impression-distribution monitoring vs popularity collapse.
Staff point: retrieval errors are permanent → optimize for recall + robustness via a *portfolio* of sources; solve freshness with content-based tower + fresh pool, not one periodically-rebuilt index.

### 4. Evaluation & rollout — three gates
- **Offline (cheap filter, not ship decision):** per-head AUC/PR-AUC + **calibration** (heads must be comparable for value combo); NDCG/MAP/recall@K; counterfactual IPS/replay. Offline gains routinely don't translate.
- **Online A/B (real verdict):** primary = retention/sessions/healthy-engagement, adequate duration (retention is slow, watch novelty decay), powered for retention variance, long-running **holdback** for cumulative effect.
- **Guardrails (do-no-harm, non-negotiable):** integrity/safety exposure, ecosystem health (Gini, new-creator success, diversity), negative engagement (hide/downvote/report), operational (p99, recall, fallback rate), fairness across cohorts (power vs casual/new).
- **Mechanics:** shadow→1→5→25→100% with automated analysis + auto-rollback; prediction-distribution/calibration monitoring between retrains; instant rollback via versioned model + config flip (rides the serving platform).
Staff point: offline metrics, in-session engagement, and the true goal (retention + healthy ecosystem) are 3 different things — bridge them explicitly. Most real regressions are "optimized the proxy, true goal quietly suffered" → guardrails make that loud.

### 5. One decision + V1 scope
Make-or-break = **right objective: calibrated multi-task model + business-tunable value function validated against long-term retention** (predict contribution not just consumption; separate predict-vs-value; prove proxy→retention in experiments).
V1 (one team, one quarter): funnel skeleton (subscribed pool + v1 two-tower + trending → light ranker → one multi-task heavy ranker: click/upvote/comment/hide → re-rank diversity+freshness+safety) + hand-tuned value weights from retention-correlation + A/B vs Hot baseline on retention+healthy+safety guardrails + daily retrain + basic position-bias + **small exploration slice from day one** + offline→A/B→guardrail eval stack.
Defer: sequence/transformer encoder, MMoE, learned/value-model weights, RL; real-time emb streaming + fancy ANN; deep cold-start personalization.
Logic: ship a *correctly-shaped* system (right objective + exploration + guardrails), beat heuristic on the *real* metric, then climb modeling sophistication against a trustworthy eval harness.

---

## Staff signals
1. Lead with the objective, not the architecture — "retention vs dwell" is the central judgment.
2. Separate predictions from value/weights (tune product without retraining; audit the goal).
3. Hit the two staff traps: feedback-loop/position-bias (exploration+IPW+ecosystem guardrails) and proxy-vs-true-goal evaluation.
4. "Retrieval errors are permanent" → recall-biased portfolio retrieval.
5. Make constraints separable (re-rank/policy out of the model).

## Reference concepts to name-drop
Two-tower / EBR, ANN (HNSW, IVF-PQ, Faiss/ScaNN), MMoE, DIN/transformer history encoder,
in-batch sampled softmax + logQ, IPW/IPS off-policy eval, MMR diversity, Thompson/UCB exploration,
calibration (Platt/isotonic), holdback experiments. (Meta "Beyond CTR", YouTube MMoE, TikTok Monolith.)
