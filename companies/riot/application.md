# Riot Games — Application

## Position

| Field | Value |
|---|---|
| Role | Principal Software Engineer |
| Team / Org | — |
| Recruiter | — |
| Applied | — |
| Status | `pending` |
| Decision | — (onsite complete, awaiting final decision — Riot comparing other candidates) |

## Timeline

| Date | Event | Outcome / Notes |
|---|---|---|
| 2026-04-28 10:00 | Recruiter call | Passed |
| 2026-05-08 | Round 1 changed to OA | Format simplified from live interview |
| 2026-05-11 | OA completed | Passed |
| 2026-05-13 | Onsite availabilities submitted | — |
| 2026-05-14 10:30 | Round 2 interview | Passed |
| 2026-05-21 11:00–15:30 | Onsite — Day 1 | Completed (Broad Tech · Deep Tech · Eng Leadership) |
| 2026-05-26 13:00–17:00 | Onsite — Day 2 | Completed (Player & Mission · Innovative · Communicate) |
| 2026-06-21 | Awaiting decision | Riot comparing a few other candidates before the final call |

## Prep Checklist for This Role

### Broad Tech (Game of Life Review) — David Eilering, 5/21 11:00am
- [x] Code walkthrough ready — sparse hashmap design, O(N) tick, Life 1.06 I/O
- [x] Integer overflow handling — safeAdd() explanation, boundary behavior
- [x] Scalability story — single-machine bottlenecks, GC pressure, map allocation
- [x] Parallelization answer — spatial tiling, ghost cells, double-buffering, sync.WaitGroup
- [x] Infinite board answer — distributed sharding with coordinator/worker/cache architecture
- [x] Hashlife algorithm — quadtree memoization, RESULT operation, binary decomposition for arbitrary ticks
- [x] Distributed Hashlife system design — coordinator task tree, stateless workers, Redis content-addressed cache, stability detection
- [x] Testing strategy — known patterns (blinker/block/glider), overflow edge cases, parser tests
- [x] "What would you improve" — Hashlife, parallel tick, configurable rules, streaming I/O
- [x] Go-specific questions — map internals (bucket hashing, incremental rehash), why Go for this problem
- See [onsite-prep-guide.md](onsite-prep-guide.md) for full question/answer prep

### Deep Tech — John Kline, 5/21 1:00pm
- [x] LoL end-to-end architecture — full service map: auth → social → lobby → matchmaking → champ select → game server → post-game
- [x] Service communication patterns — gRPC sync, Kafka async, WebSocket push, UDP gameplay
- [x] Game server networking — UDP both directions, redundant inputs, delta-compressed state, custom selective reliability
- [x] Reconnection flow — TCP handshake + re-auth, full state snapshot, UDP resume
- [x] Matchmaking deep dive — data structure, scoring function, party balance, MMR bands, queue health monitoring
- [x] Champion select — server-authoritative state machine, crash recovery (log replay)
- [x] Post-game pipeline — Kafka event-driven, independent consumers, per-player MMR serialization
- [x] Auth across services — gateway validation, internal identity context, match-scoped game tokens
- [x] Global infrastructure — regional isolation for gameplay, global shared for identity/commerce, region transfer saga
- [x] Security/anti-cheat — client agent, server-authoritative state, statistical anomaly detection, rate limiting
- [x] Profanity filter — Rabin-Karp rolling hash, evasion handling, hot-reloaded blacklist, system placement
- [x] Observability at principal level — 4-layer model (player-facing → service → infra → business)
- [x] Debugging story structure — context, symptom, wrong hypothesis, investigation, root cause, fix, prevention
- [x] Go internals — GC (tricolor mark-and-sweep, GOGC), goroutine scheduler (GMP model), compilation pipeline
- See [onsite-prep-guide.md](onsite-prep-guide.md) for full question/answer prep

### Engineering Leadership — Ola Mogstad, 5/21 2:30pm
- [x] 3-4 leadership stories with extreme specifics (who, what, when, numbers)
- [x] Conflict resolution stories with full arc
- [x] "What did you learn recently?" — genuine, specific answer
- [x] Planning/roadmap story — concrete timeline and prioritization
- [x] Frame as team success, not personal heroics
- See [onsite-prep-guide.md](onsite-prep-guide.md) for predicted questions
- See [eng-leadership-prep.md](eng-leadership-prep.md) for resume-based answers and story map

### Player & Mission Focus — Maxfield Stewart, 5/26 1:00pm
- [ ] "Why Riot?" — genuine answer beyond "I like games"
- [ ] Know your own company's culture and how you've embodied it
- [ ] User/player empathy stories — prioritizing UX over easy technical paths
- [ ] Study Riot values: https://www.riotgames.com/en/who-we-are/values

### Innovative Problem Solving — Kenneth Guillory + Arun Ramakrishnan, 5/26 2:00pm
- [ ] Root cause diagnosis stories — wrong turns, adaptation, learning
- [ ] Creative/unconventional solution examples
- [ ] Challenging assumptions stories
- [ ] Two interviewers — prepare for cross-examination style

### Communicate & Collaborate — Mike Horstmanshof, 5/26 4:00pm
- [ ] Giving negative feedback story
- [ ] Receiving hard feedback story
- [ ] Cross-team collaboration challenge story
- [ ] Missed deadline / scope trade-off story
- [ ] Practice explaining technical concepts simply (evaluated LIVE)

### Company Research
- [x] See [CLAUDE.md](CLAUDE.md) for Riot Games-specific notes
- [x] 11 interview experiences logged in [interview-experiences.md](interview-experiences.md)

## Notes / Observations

<!-- Add recruiter conversations, interview feedback, impressions here -->
