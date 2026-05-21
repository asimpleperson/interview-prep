# Reddit — Interview Experiences & Coding Problems

## Sources

Collected from 地里 (1point3acres) and other interview sharing platforms.

---

## Coding 1 — Report Chain / Manager Hierarchy (高频)

**File:** [coding1_report_chain.go](coding1_report_chain.go)

**Input:** `["A,B,C", "C,D", "B,E"]` — each string is `manager,report1,report2,...`

### Parts

| Part | Task | Core Technique |
|------|------|---------------|
| 1 | Build tree, print with indentation | Graph→Tree, DFS |
| 2 | Find skip-level meeting pairs (grandparent+ ↔ descendant) | Ancestor traversal |
| 3 | Given a person, print all managers above + all reports below | Parent chain + subtree DFS |
| 4 | Lowest common manager of two people | LCA via ancestor set |
| 5 | Unknown | — |

### Key Points
- Maintain both `parent` map and `children` map for bidirectional traversal
- Root = node with no parent
- LCA: walk both nodes up, use ancestor set intersection

---

## Coding 2 — Merge Chat Messages (Merge K Sorted 变形)

**File:** [coding2_merge_messages.go](coding2_merge_messages.go)

**Setup:** `get_chat_messages(id)` returns ~11 sorted messages around `id`. Implement `merge_messages(ids[])` — merge all, deduplicate, sort by ID. Cannot use set.

### Key Points
- Classic merge K sorted lists with min-heap
- Dedup by comparing with `result[last]` (sorted order guarantees adjacent duplicates)
- Go: `container/heap` requires 5 interface methods — practice writing from memory
- Time: O(N log K)

### Follow-up: Versioning
- Messages can be edited; store version history
- Merge uses `(ID, Version)` as key, keep highest version

---

## Coding 3 — Moderator List System (数据结构设计)

**File:** [coding3_moderator_system.go](coding3_moderator_system.go)

**Input:** Logs like `"user1,user2 | added | user3 | 123.0"`

### Parts

| Part | Task | Core Technique |
|------|------|---------------|
| 1 | Parse logs, `can_remove_mod(u1, u2)`, `get_mod_list()` | Ordered map by timestamp |
| 2 | Multi-community support | Map of community → ModeratorSystem |
| 3 | `demote(user)` — move user back one position in list | Swap + sync timestamps |

### Key Points
- `can_remove_mod`: u1 can remove u2 only if u1 joined **earlier**
- Demote swaps position AND timestamps to keep `can_remove` semantics consistent
- If asked about optimization: `OrderedDict` / doubly-linked list for O(1) removal

---

## Coding 4 — Tennis Game / Score System (OOD)

**File:** [coding4_tennis_game.go](coding4_tennis_game.go)

**API:** `AddScore(player)`, `GetScore()`, `GetResult()`

### Rules
- First to 4 points with 2-point lead wins
- If both ≥ 3 and tied → reset to 3:3 (Deuce)
- Game over → no more scoring allowed

### Follow-up: Human Readable Score

| Score | Term |
|-------|------|
| 0 | Love |
| 1 | Fifteen |
| 2 | Thirty |
| 3 | Forty |
| 3:3 (deuce zone) | Deuce |
| 4:3 after deuce | Advantage |

### Key Points
- State machine: order matters — increment, check deuce reset, check win
- Advantage is always 4:3 (4:4 resets, 5:3 already won)
- Potential follow-up: Set/Match layer — design for composability

---

## Coding 5 — Billing Status System (事务处理 + Undo/Redo)

**File:** [coding5_billing_status.go](coding5_billing_status.go)

**Scenario:** Rebuild per-user billing status from transaction logs after a database drop.

### Parts

| Part | Task | Core Technique |
|------|------|---------------|
| 1 | Accumulate monetary columns per user | Sort by timestamp + group by user |
| 2 | Support `overwrite` transactions | Overwrite replaces only columns present in that tx |
| 3 | Support `undo_last` / `redo_last` | Dual-stack undo/redo pattern |

### Key Points
- **Must sort by timestamp first** — input is a map (unordered)
- Overwrite only affects columns **present in that transaction**, not all columns
- Undo target: only "regular" transactions (not undo/redo themselves)
- Undo + overwrite combo: simple subtraction doesn't work → use **replay** approach
- Replay: reset to zero, re-apply all remaining transactions in the applied stack
- Trade-off: replay is O(n) per undo — discuss snapshot optimization if asked

---

## Frequency & Difficulty Summary

| Problem | Frequency | Difficulty | Time Pressure |
|---------|-----------|------------|---------------|
| Report Chain | 极高 | Medium | 45 min for all 4 parts |
| Merge Messages | Medium | Medium | 30 min |
| Moderator System | Medium | Medium | 35 min |
| Tennis Game | Medium | Easy-Medium | 25 min |
| Billing Status | Medium | Medium-Hard | 40 min |
