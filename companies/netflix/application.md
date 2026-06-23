# Netflix — Application

## Position

| Field | Value |
|---|---|
| Role | Distributed Systems Engineer 5 (L5) — Decisioning & Optimization |
| Team / Org | Netflix Ads — Decisioning & Optimization (Ad Serving & Decisioning org) |
| Location | California |
| Account | huzhihao505@gmail.com |
| Recruiter | — |
| Applied | — |
| Status | `active` |
| Decision | — |

## Timeline

| Date | Event | Outcome / Notes |
|---|---|---|
| — | Applied | Distributed Systems Engineer 5 — Decisioning & Optimization (L5, California) |
| 2026-06-11 | Recruiter call | Passed |
| 2026-06-24 11:00–12:00 PDT | L5 Technical Screen | **Confirmed** |

## Prep Checklist for This Role

### Coding
Netflix SE5 interviews are rigorous. Expect hard problems with emphasis on clean, production-quality code.
- [ ] All Must-Have problems across every topic
- [ ] Dynamic programming
- [ ] Graph algorithms
- [ ] Data Structure Design
- [ ] Heap / priority queue

### System Design — key signals from the JD (Decisioning & Optimization)
Real-time ad decisioning + low-latency ML serving. The must-know areas:
- [ ] Real-time decisioning path: ranking → scoring → bidding → pacing under strict latency/throughput
- [ ] ML model serving infra: many concurrent hot-path models, **sub-20ms P99**, model routing, lifecycle, fallback tiers, calibration
- [ ] Auction mechanics: first-/second-price, reserve pricing, bid shading
- [ ] Budget pacing & goal-based delivery: dynamic budget/inventory allocation across demand channels
- [ ] Multi-stage ranking (retrieval → scoring → reranking); podding & ad-break planning
- [ ] Simulation / counterfactual testing of marketplace & auction changes (offline validation before live rollout)
- [ ] Ad-serving fundamentals: inventory mgmt, frequency capping, supply-demand, member ad-experience quality
- [ ] Nice-to-have: CTV / server-side ad insertion (SSAI), live-event ad serving; experimentation (A/B, holdouts)

Profile they want: 7+ yrs distributed systems, 2+ yrs ads domain, productionizing ML into low-latency serving paths.

### Behavioral
Netflix culture is high-autonomy, high-accountability. Frame answers around ownership and outcomes.
- [ ] "Tell me about a time you made a significant decision with incomplete information"
- [ ] "Describe a system you owned end-to-end — what would you do differently?"
- [ ] "How do you handle a situation where you disagree with your manager's direction?"
- [ ] "Give an example of a time you raised the bar for the team"

### Project Deep Dive
- [ ] Lead with any ads, monetization, or pricing systems work
- [ ] Emphasize end-to-end ownership and measurable revenue / business impact

### Company Research
- [ ] See [CLAUDE.md](CLAUDE.md) for Netflix-specific notes

## Notes / Observations

<!-- Add recruiter conversations, interview feedback, impressions here -->
