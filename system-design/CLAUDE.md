# system-design/

## Purpose

Study materials and design write-ups for large-scale distributed systems interviews. Staff-level system design rounds evaluate architectural judgment, not just knowledge — tradeoffs and reasoning matter more than memorized answers.

## Planned Contents

- Topic notes (e.g., consistent hashing, CAP theorem, rate limiting, distributed consensus)
- Design write-ups for common interview problems (URL shortener, news feed, chat system, etc.)
- A reusable design framework / template

## Conventions

- One file per design problem: `design-<system>.md`
- Each design file should follow the structure: Requirements → Capacity Estimation → High-Level Design → Deep Dive → Tradeoffs
- Topic notes go in `notes/` subdirectory when that folder is created.

## Key Areas to Cover

- Storage: SQL vs NoSQL, sharding, replication, indexing
- Caching: eviction policies, cache invalidation, CDN
- Messaging: Kafka, pub/sub, exactly-once delivery
- Consistency models and distributed transactions
- API design, rate limiting, auth at scale
- ML system design (relevant for AI company rounds)
