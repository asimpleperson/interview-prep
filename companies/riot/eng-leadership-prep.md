# Engineering Leadership Prep — Ola Mogstad, 5/21 2:30pm

## What They're Evaluating (from interview description)

Three pillars:
1. **Technical influence** — driving direction beyond your own team
2. **People growth** — mentoring and elevating others
3. **Culture/craft** — engineering practices, processes, standards

## Principal vs Senior — The Key Difference

| Dimension | Senior Answer | Principal Answer |
|---|---|---|
| Scope | "My team did X" | "I drove X across N teams" |
| People | "I helped a teammate" | "I built a system that grows engineers" |
| Decisions | "I chose technology X" | "I wrote the decision framework teams use" |
| Impact | "We shipped on time" | "The architecture/practice is still in place N years later" |
| Conflict | "I convinced them" | "I made them co-owners so they became advocates" |
| Failure | "It didn't go well" | "Here's what I changed systemically to prevent recurrence" |

---

## Story Map — Which Story for Which Question

| Question Theme | Primary Story | Backup Story |
|---|---|---|
| Technical direction beyond team | ByteDance re-architecture (12 services, 3 storage, 20M creators) | Snap ranking re-architecture (ML+rules framework) |
| Mentoring / growing engineers | ByteDance team 4→12, 100+ interviews, coached to senior promos | Snap code review and pairing culture |
| Architectural decisions that last | Snap unified ranking framework (invariant vs variable separation) | ByteDance data-driven campaign engine |
| Led team to success | Snap TikTok shutdown + NYE (8 engineers, 99.5% availability) | ByteDance GDPR deadline |
| Engineering culture | Snap observability + experiment launch framework | ByteDance hiring bar and interview calibration |
| Recent learning | ML training data pipeline integrity at Snap | — |
| Under pressure | ByteDance GDPR (legal deadline, org-wide, non-negotiable) | Snap TikTok traffic spike |
| Conflict resolution | Snap ranking vs infra team during incident (decision framework) | ByteDance ops platform API pushback (made them co-owners) |

---

## Q1: Technical Direction Beyond Team

**Story: ByteDance Growth Center Re-architecture**

- 20M+ TikTok creators, rigid legacy backend → 12-service platform, 3 storage systems
- Key insight: separate campaign config engine from execution layer → operations teams create campaigns via UI, no code changes
- 4-phase migration: campaign config → user data transition → dual read/write → cutover
- Cross-team influence: aligned 3 teams (e-commerce ops, reward disbursement, creator analytics) on API contracts
- Ops platform team pushed back → brought them into design, co-designed API that simplified their code → became advocates
- Result: campaign ROI 0.4 → 2.1 (5x), architecture still in production

**Key follow-up answers:**
- Buy-in from disagreers: show the math (2-3 sprints per campaign → UI config), reframe from "rework" to "unblocks your roadmap", offer to co-own the migration
- What survived: service decomposition + data-driven campaign engine. Migration pattern became a template for other teams. Would change: consolidate 3 storage systems earlier to avoid cross-store consistency issues.

## Q2: Mentoring / Growing Engineers

**Story: ByteDance Team Scaling 4→12**

- 100+ interviews, personally selected every member
- Moved junior engineers from "execute my design" to "own the design with me as reviewer"
- Specific example: paired with junior on dual read/write migration module, asked design questions instead of giving answers, let them make decisions
- Both junior engineers promoted to senior within ~2 quarters
- Systematic: every major project paired senior + junior, mandatory thorough code reviews modeled by me

**Key follow-up answers:**
- Balance mentoring vs IC: "Growing the team IS my IC work at principal level. My leverage comes from team throughput, not personal throughput."
- Struggling engineer: someone overloaded on 3 projects → coached them on the prioritization conversation with their manager → raised it as systemic issue → implemented policy: max 2 projects per engineer

## Q3: Architectural Decisions That Outlast You

**Story: Snap Ranking Re-architecture**

- Old system: each feed surface had its own ranking path, duplicated logic, inconsistent training data logging
- Design principle: separate invariant (ML pipeline) from variable (surface-specific rules)
- Shared pipeline + surface-specific config, NOT a generic plugin system (constrained flexibility = harder to break)
- Co-developed with 1 BE + 1 MLE, documented WHY behind each trade-off
- Result: multiple new feed surfaces onboarded without architectural changes

**Key follow-up answers:**
- Trade-offs: flexibility vs complexity, chose constrained middle path. Kept surface-specific feature stores for isolation (would consolidate in hindsight).
- How it held up: core framework still in production, new surfaces onboarded successfully. Documented the conditions under which decisions should be revisited.

## Q4: Led Team to Success

**Story: Snap TikTok Shutdown + NYE (8 engineers)**

- Context: TikTok shutdown → massive traffic surge to Snapchat content feeds (400M+ DAU)
- Plan: capacity review 2 days before, mapped every service bottleneck, pre-scaled 40%, assigned specific engineers to specific systems
- Critical moment: traffic 2x normal peak, ranking p99 latency elevated, RocksDB read amplification under load
- Decision: activate popularity-based ranking fallback for 10% of traffic in affected regions
- Conflict: ranking team (keep ML quality) vs infra team (protect availability) → set concrete threshold: below 99.5% = degrade, above = keep full ranking. Data showed impact was acceptable.
- What I did vs delegated: owned coordination + stakeholder comms (hourly to Ads/PM/VP). Delegated technical mitigation to assigned engineers.
- Result: 99.5% availability, zero ad serving interruptions, 20% sustained Spotlight view time increase

## Q5: Engineering Culture

**Story: Snap Observability + Experiment Framework**

- Problem: observability was reactive, found problems when users complained
- Built monitoring dashboards + alerting, but more importantly: every A/B experiment launch requires pre-defined observability criteria (metrics, rollback thresholds, monitoring owner for first 24 hours)
- Cultural shift: "launch and see" → "launch with eyes open"
- Result: caught training data logging issue in 5% Spotlight experiment that would have corrupted ML model at 100%
- Also at ByteDance: raised org hiring bar, shared interview rubric, advocated for design component in interview loop

## Q6: Recent Learning

**Story: ML Training Data Pipeline Integrity at Snap**

- Learned how training data logging integrity directly impacts model quality
- Discovered logging inconsistency between Discover and Spotlight feature pipelines → model learning WHERE content was shown, not WHAT was good
- Changed mental model: now treat training data pipeline as first-class production system, not an afterthought
- At principal level: when designing serving architecture, explicitly consider what training signals are generated and whether they're consistent

## Q7: Under Pressure

**Story: ByteDance GDPR Project**

- Legal compliance deadline, non-negotiable, org-wide scope
- Audited all 12 services and 3 storage systems, triaged into: no work / medium / hard redesign
- Key decision: targeted EU data routing (scoped solution for deadline) rather than general-purpose data residency framework. Explicit with manager: "scoped now, generalize later"
- Coordinated 3 partner teams, weekly sync, shared tracking. When one team fell behind, reassigned my engineers to unblock them.
- Result: shipped on time, zero EU data violations. Targeted solution later generalized by another team — exactly the plan.

**Key follow-up — keeping team calm:**
"Transparency and decisiveness. Don't hide severity — people handle hard truths, not uncertainty. Make decisions quickly, even imperfect ones. A team with a clear imperfect plan outperforms a team waiting for someone to decide."

---

## Riot-Specific Framing Reminders

- **Team success > personal heroics** — every story ends with the team's outcome
- **Player/user impact** — connect technical decisions to end-user experience (400M DAU, 20M creators)
- **Growth mindset** — include what you learned, what you'd do differently
- **Don't BS specifics** — "I don't remember the exact number but it was in the range of..." is better than fabricating
- **Collaborative culture** — "I made them co-owners" not "I convinced them I was right"
