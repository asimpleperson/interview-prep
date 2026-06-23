# Amazon Phone Screen — Coding Prep Log

> Phone screen: **2026-06-23 10:00 PT** · Interviewer: Jon Friesen · Zoom (possible shadow)
> Language for the round: **Go**
> Practiced in interviewer/interviewee format — problem framing + solution + full follow-up volley.

## Meta-strategy for the round

- **Name the pattern out loud first**, then code. Stating the decomposition buys credit before code compiles.
- **Talk while typing** — Amazon weights communication heavily; silence reads as a red flag.
- **Maintain aggregates incrementally** (running sums/counts), never recompute `O(n)` per query.
- **Tie trade-offs to Leadership Principles**: ship simple first (*Bias for Action*), don't over-engineer (*Frugality*), verify requirements (*Are Right, A Lot / Dive Deep*).
- Always volunteer **test cases**: empty input, `k`/`q` out of range, ties, duplicates, single element, boundaries.

---

## Problem 1 — Graph traversal: top-K nodes by rating

**Framing.** Bounded BFS from a source (friends-of-friends / related products up to `maxDepth` hops), return the `k` highest-rated reachable nodes.

**Decomposition:** *bounded BFS to gather candidates* **+** *top-k selection via a size-`k` min-heap*.

**Key insight:** top-k uses a **min-heap of size k**, not a max-heap + full sort. Root is the weakest survivor — cheap `O(log k)` eviction. Time `O(V + E + V log k)`, space `O(V + k)`.

**Go heap gotcha (came up explicitly):** with `container/heap`, the `Push`/`Pop` *methods* you implement are just append/truncate hooks; the package functions `heap.Push`/`heap.Pop` do the sift-up/sift-down. **Never call `h.Push`/`h.Pop` directly.** Be ready to hand-roll an explicit `up`/`down` min-heap if asked to "implement the heap yourself."

**Follow-ups:** min-heap-of-k vs sort trade-off (`V ≫ k`); graph too big → per-shard local top-k merged (top-k is a monoid); real-time rating changes → indexed PQ / Space-Saving; weighted edges → Dijkstra/best-first; cycles & memory → `visited` set, Bloom filter, iterative deepening.

---

## Problem 2 — Top-K pages by click count (scales up)

**Framing.** Count clicks per page from logs, return top `k`. Same core (count map + size-`k` min-heap) wrapped in escalating constraints.

**Key insight:** raw log size ≠ bottleneck — **cardinality (distinct pages)** is. Memory is `O(D + k)`, independent of `N`. Stream the input (`io.Reader`), never slurp.

**⚠️ The senior trap (most important point):** "local top-k then merge" is **only correct if you partition by key**. Arbitrary chunk → a page that's #(k+1) in every chunk can be #1 globally; dropping below local-k loses the counts. Fixes:
1. Merge **full partial counts** (not local top-k), then one global top-k; or
2. **Hash-partition by page id** so each page's full count lands in one reducer → local top-k is then safe to merge. This is MapReduce: `map (page,1)` → shuffle-by-key → `reduce` sum → top-k.

**Follow-ups:** URL logs → stream + retry + gzip + HTTP Range resume; counts don't fit RAM → external hash-partition into spill files **or** approximate (Count-Min Sketch / Space-Saving heavy hitters — *ask if approximate is OK, huge lever*); streaming/sliding window → Space-Saving; skew/hot keys → combiner + key salting.

---

## Problem 3 — Two age-balanced queues, running score sums (LC 295 variant)

**Framing.** Stream of `(age, score)`. Maintain younger/older halves split **by age**; sizes differ ≤1; odd → **younger** queue gets the extra. After each insert, output `(youngSum, oldSum)`.

**Decomposition:** two-heap median skeleton (LC 295), but **partition key (age) ≠ aggregated value (score)**, and we output **running sums, not the median**.

- `young` = max-heap by age (top = oldest of young = boundary); `old` = min-heap by age.
- Invariant `young.size ∈ {old.size, old.size+1}`.
- **Maintain `youngSum`/`oldSum` incrementally** — every cross-boundary move adjusts both sums by the mover's score. `O(log n)`/insert; summing the heap each time would be the `O(n²)` trap.

**Follow-ups:** report median age (falls out of boundary tops); flip which queue gets the extra (one line); removals → lazy deletion + logical sizes; bounded ages / k-way percentile splits → **Fenwick (BIT) over age storing count + score-sum**, binary-search for rank boundary, prefix query for half-sum.

---

## Problem 4 — Count time points with exact coverage, in a range

**Framing.** Intervals `[start,end]` (one-to-one). `coverage[t]` = # intervals covering integer `t`. Query `(q, qstart, qend)` → count of `t ∈ [qstart,qend]` with `coverage[t] == q`. **Verify the interpretation against all sample outputs before coding** — statement is ambiguous.

Worked example: intervals `[1,5],[0,10],[10,12]`; queries gave `[3,1,0]` — confirmed.

**Decomposition:** *difference array + prefix sum* → coverage; *bucket positions by coverage value (sorted)* → each query is a **binary search** for count in `[qstart,qend]`. Time `O(N + T + Q log T)`, space `O(T)`. Handles `q=0` (gaps) and `q` > max coverage (empty bucket) naturally.

**Follow-ups:** time up to 1e9 → **coordinate compression + sweep line** into constant-coverage segments, bucket segments by value with prefix point-counts, clip boundary segments (`O((N+Q) log N)`, space `O(N)`); continuous time → length measure instead of point count; dynamic add/remove intervals → segment tree w/ lazy propagation (exact-value counting is genuinely hard — push back on requirements); `≥q`/`≤q` queries → 2D dominance count via merge-sort/wavelet/persistent segment tree.

---

## Pattern recognition cheat-sheet

| Trigger phrase | Core structure |
|---|---|
| "top k by <metric>" / "k most frequent" | count/collect → **size-k min-heap** |
| "traverse … find best k" | **bounded BFS/DFS** + size-k heap |
| "huge data / distributed top-k" | **partition by key** (MapReduce) or approximate (Count-Min / Space-Saving) |
| "split stream into balanced halves / running median" | **two heaps**; partition key may ≠ aggregate; incremental sums |
| "coverage at a point" + "count points with property in range" | **difference array / sweep line** → bucket + binary search |
