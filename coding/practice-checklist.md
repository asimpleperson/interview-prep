# Coding Practice Checklist (代码随想录)

---

## 1. Arrays

**Key Concepts:** Contiguous memory, index-based access, in-place operations, prefix sums, sliding window, binary search.

**Patterns:**
- [ ] Binary search (left/right boundary variants)
- [ ] Two-pointer for in-place removal
- [ ] Sliding window for subarray problems
- [ ] Prefix sums for range queries
- [ ] Simulation (spiral fill)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Binary Search | [LC 704](https://leetcode.com/problems/binary-search/) |
| [ ] | Remove Element | [LC 27](https://leetcode.com/problems/remove-element/) |
| [ ] | Squares of a Sorted Array | [LC 977](https://leetcode.com/problems/squares-of-a-sorted-array/) |
| [ ] | Minimum Size Subarray Sum | [LC 209](https://leetcode.com/problems/minimum-size-subarray-sum/) |
| [ ] | Spiral Matrix II | [LC 59](https://leetcode.com/problems/spiral-matrix-ii/) |

---

## 2. Linked Lists

**Key Concepts:** Pointer manipulation, dummy head node trick, fast/slow pointers, reversal, cycle detection.

**Patterns:**
- [ ] Dummy head node to simplify edge cases
- [ ] Two-pointer (fast/slow) for cycle & nth-from-end
- [ ] In-place reversal
- [ ] Intersection detection via pointer reset

| Done | Problem | Link |
|------|---------|------|
| [ ] | Remove Linked List Elements | [LC 203](https://leetcode.com/problems/remove-linked-list-elements/) |
| [ ] | Design Linked List | [LC 707](https://leetcode.com/problems/design-linked-list/) |
| [ ] | Reverse Linked List | [LC 206](https://leetcode.com/problems/reverse-linked-list/) |
| [ ] | Swap Nodes in Pairs | [LC 24](https://leetcode.com/problems/swap-nodes-in-pairs/) |
| [ ] | Remove Nth Node From End | [LC 19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) |
| [ ] | Intersection of Two Linked Lists | [LC 160](https://leetcode.com/problems/intersection-of-two-linked-lists/) |
| [ ] | Linked List Cycle II | [LC 142](https://leetcode.com/problems/linked-list-cycle-ii/) |

---

## 3. Hash Tables

**Key Concepts:** O(1) lookup, collision resolution, choosing the right structure (array vs. set vs. map).

**Patterns:**
- [ ] Frequency counting with arrays (fixed alphabet)
- [ ] Set for existence/intersection
- [ ] Map for complement lookup (Two Sum pattern)
- [ ] Multi-array grouping (Four Sum II)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Valid Anagram | [LC 242](https://leetcode.com/problems/valid-anagram/) |
| [ ] | Intersection of Two Arrays | [LC 349](https://leetcode.com/problems/intersection-of-two-arrays/) |
| [ ] | Happy Number | [LC 202](https://leetcode.com/problems/happy-number/) |
| [ ] | Two Sum | [LC 1](https://leetcode.com/problems/two-sum/) |
| [ ] | Four Sum II | [LC 454](https://leetcode.com/problems/4sum-ii/) |
| [ ] | Ransom Note | [LC 383](https://leetcode.com/problems/ransom-note/) |
| [ ] | Three Sum | [LC 15](https://leetcode.com/problems/3sum/) |
| [ ] | Four Sum | [LC 18](https://leetcode.com/problems/4sum/) |

---

## 4. Strings

**Key Concepts:** In-place reversal, KMP pattern matching, prefix table (failure function), repeated substring detection.

**Patterns:**
- [ ] Two-pointer reversal (whole string, partial, word-by-word)
- [ ] KMP for O(n) substring search — build the prefix table
- [ ] Sliding window for substring problems

| Done | Problem | Link |
|------|---------|------|
| [ ] | Reverse String | [LC 344](https://leetcode.com/problems/reverse-string/) |
| [ ] | Reverse String II | [LC 541](https://leetcode.com/problems/reverse-string-ii/) |
| [ ] | Reverse Words in a String | [LC 151](https://leetcode.com/problems/reverse-words-in-a-string/) |
| [ ] | Find the Index of the First Occurrence (strStr) | [LC 28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) |
| [ ] | Repeated Substring Pattern | [LC 459](https://leetcode.com/problems/repeated-substring-pattern/) |

---

## 5. Two Pointers

**Key Concepts:** Reduces O(n²) brute force to O(n); works on sorted arrays and linked lists.

**Patterns:**
- [ ] Fast/slow pointer (cycle detection, midpoint)
- [ ] Left/right pointer (sorted array problems, kSum)
- [ ] Same-direction pointers (sliding window, in-place overwrite)

> Two-pointer techniques appear across Arrays, Strings, and Linked Lists — treat this as a cross-cutting pattern, not an isolated topic.

---

## 6. Stacks & Queues

**Key Concepts:** LIFO vs. FIFO, monotonic deque for sliding window max, priority queue for top-K.

**Patterns:**
- [ ] Stack for bracket matching and expression evaluation
- [ ] Stack for adjacent duplicate removal
- [ ] Monotonic deque for sliding window maximum
- [ ] Heap/priority queue for top-K frequency

| Done | Problem | Link |
|------|---------|------|
| [ ] | Implement Queue using Stacks | [LC 232](https://leetcode.com/problems/implement-queue-using-stacks/) |
| [ ] | Implement Stack using Queues | [LC 225](https://leetcode.com/problems/implement-stack-using-queues/) |
| [ ] | Valid Parentheses | [LC 20](https://leetcode.com/problems/valid-parentheses/) |
| [ ] | Remove All Adjacent Duplicates In String | [LC 1047](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) |
| [ ] | Evaluate Reverse Polish Notation | [LC 150](https://leetcode.com/problems/evaluate-reverse-polish-notation/) |
| [ ] | Sliding Window Maximum | [LC 239](https://leetcode.com/problems/sliding-window-maximum/) |
| [ ] | Top K Frequent Elements | [LC 347](https://leetcode.com/problems/top-k-frequent-elements/) |

---

## 7. Binary Trees

**Key Concepts:** DFS (pre/in/post-order) and BFS (level-order), recursive vs. iterative traversal, BST properties.

**Patterns:**
- [ ] Recursive DFS for tree properties (depth, balance, paths)
- [ ] BFS/level-order with a queue
- [ ] BST in-order traversal gives sorted sequence
- [ ] Lowest Common Ancestor (LCA) using recursion return values
- [ ] Tree construction from traversal arrays

| Done | Problem | Link |
|------|---------|------|
| [ ] | Binary Tree Level Order Traversal | [LC 102](https://leetcode.com/problems/binary-tree-level-order-traversal/) |
| [ ] | Invert Binary Tree | [LC 226](https://leetcode.com/problems/invert-binary-tree/) |
| [ ] | Symmetric Tree | [LC 101](https://leetcode.com/problems/symmetric-tree/) |
| [ ] | Maximum Depth | [LC 104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) |
| [ ] | Minimum Depth | [LC 111](https://leetcode.com/problems/minimum-depth-of-binary-tree/) |
| [ ] | Count Complete Tree Nodes | [LC 222](https://leetcode.com/problems/count-complete-tree-nodes/) |
| [ ] | Balanced Binary Tree | [LC 110](https://leetcode.com/problems/balanced-binary-tree/) |
| [ ] | All Root-to-Leaf Paths | [LC 257](https://leetcode.com/problems/binary-tree-paths/) |
| [ ] | Path Sum | [LC 112](https://leetcode.com/problems/path-sum/) |
| [ ] | Construct from Inorder & Postorder | [LC 106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) |
| [ ] | Maximum Binary Tree | [LC 654](https://leetcode.com/problems/maximum-binary-tree/) |
| [ ] | Merge Two Binary Trees | [LC 617](https://leetcode.com/problems/merge-two-binary-trees/) |
| [ ] | Lowest Common Ancestor (general) | [LC 236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) |
| [ ] | Lowest Common Ancestor (BST) | [LC 235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) |
| [ ] | Validate BST | [LC 98](https://leetcode.com/problems/validate-binary-search-tree/) |
| [ ] | Search in BST | [LC 700](https://leetcode.com/problems/search-in-a-binary-search-tree/) |
| [ ] | Insert into BST | [LC 701](https://leetcode.com/problems/insert-into-a-binary-search-tree/) |
| [ ] | Delete Node in BST | [LC 450](https://leetcode.com/problems/delete-node-in-a-bst/) |
| [ ] | Trim BST | [LC 669](https://leetcode.com/problems/trim-a-binary-search-tree/) |
| [ ] | Convert Sorted Array to BST | [LC 108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) |
| [ ] | Convert BST to Greater Tree | [LC 538](https://leetcode.com/problems/convert-bst-to-greater-tree/) |

---

## 8. Backtracking

**Key Concepts:** Exhaustive search on a decision tree; prune branches early. Framework: **choose → explore → unchoose**.

**Patterns:**
- [ ] Combinations (order doesn't matter, no reuse)
- [ ] Combinations with reuse (Combination Sum)
- [ ] Deduplication: sort + skip duplicate at same tree level
- [ ] Permutations (order matters, use `used[]` array)
- [ ] Subsets = combinations of all sizes
- [ ] Partitioning (palindrome check, IP address segments)
- [ ] Board problems (N-Queens, Sudoku — 2D constraint propagation)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Combinations | [LC 77](https://leetcode.com/problems/combinations/) |
| [ ] | Combination Sum III | [LC 216](https://leetcode.com/problems/combination-sum-iii/) |
| [ ] | Letter Combinations of a Phone Number | [LC 17](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) |
| [ ] | Combination Sum | [LC 39](https://leetcode.com/problems/combination-sum/) |
| [ ] | Combination Sum II | [LC 40](https://leetcode.com/problems/combination-sum-ii/) |
| [ ] | Palindrome Partitioning | [LC 131](https://leetcode.com/problems/palindrome-partitioning/) |
| [ ] | Restore IP Addresses | [LC 93](https://leetcode.com/problems/restore-ip-addresses/) |
| [ ] | Subsets | [LC 78](https://leetcode.com/problems/subsets/) |
| [ ] | Subsets II | [LC 90](https://leetcode.com/problems/subsets-ii/) |
| [ ] | Non-decreasing Subsequences | [LC 491](https://leetcode.com/problems/non-decreasing-subsequences/) |
| [ ] | Permutations | [LC 46](https://leetcode.com/problems/permutations/) |
| [ ] | Permutations II | [LC 47](https://leetcode.com/problems/permutations-ii/) |
| [ ] | N-Queens | [LC 51](https://leetcode.com/problems/n-queens/) |
| [ ] | Sudoku Solver | [LC 37](https://leetcode.com/problems/sudoku-solver/) |

---

## 9. Greedy Algorithms

**Key Concepts:** Make the locally optimal choice at each step; no rollback. Correctness requires proof — not just intuition.

**Patterns:**
- [ ] Interval scheduling (sort by end time, count non-overlapping)
- [ ] Two-pass greedy (candy: left→right then right→left)
- [ ] Jump game (track max reachable index)
- [ ] Stock problems (accumulate positive daily diffs)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Assign Cookies | [LC 455](https://leetcode.com/problems/assign-cookies/) |
| [ ] | Wiggle Subsequence | [LC 376](https://leetcode.com/problems/wiggle-subsequence/) |
| [ ] | Maximum Subarray | [LC 53](https://leetcode.com/problems/maximum-subarray/) |
| [ ] | Best Time to Buy and Sell Stock II | [LC 122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) |
| [ ] | Jump Game | [LC 55](https://leetcode.com/problems/jump-game/) |
| [ ] | Jump Game II | [LC 45](https://leetcode.com/problems/jump-game-ii/) |
| [ ] | Maximize Sum After K Negations | [LC 1005](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/) |
| [ ] | Gas Station | [LC 134](https://leetcode.com/problems/gas-station/) |
| [ ] | Candy | [LC 135](https://leetcode.com/problems/candy/) |
| [ ] | Lemonade Change | [LC 860](https://leetcode.com/problems/lemonade-change/) |
| [ ] | Queue Reconstruction by Height | [LC 406](https://leetcode.com/problems/queue-reconstruction-by-height/) |
| [ ] | Minimum Arrows to Burst Balloons | [LC 452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) |
| [ ] | Non-overlapping Intervals | [LC 435](https://leetcode.com/problems/non-overlapping-intervals/) |
| [ ] | Partition Labels | [LC 763](https://leetcode.com/problems/partition-labels/) |
| [ ] | Merge Intervals | [LC 56](https://leetcode.com/problems/merge-intervals/) |
| [ ] | Monotone Increasing Digits | [LC 738](https://leetcode.com/problems/monotone-increasing-digits/) |
| [ ] | Binary Tree Cameras | [LC 968](https://leetcode.com/problems/binary-tree-cameras/) |

---

## 10. Dynamic Programming

**Key Concepts:** Optimal substructure + overlapping subproblems. Define state → write recurrence → determine traversal order → initialize base cases.

### Basic DP

**Patterns:**
- [ ] 1D DP (Fibonacci-style, climbing stairs)
- [ ] 2D DP (grid paths)
- [ ] Math-based DP (integer break, unique BST count)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Fibonacci Number | [LC 509](https://leetcode.com/problems/fibonacci-number/) |
| [ ] | Climbing Stairs | [LC 70](https://leetcode.com/problems/climbing-stairs/) |
| [ ] | Min Cost Climbing Stairs | [LC 746](https://leetcode.com/problems/min-cost-climbing-stairs/) |
| [ ] | Unique Paths | [LC 62](https://leetcode.com/problems/unique-paths/) |
| [ ] | Unique Paths II | [LC 63](https://leetcode.com/problems/unique-paths-ii/) |
| [ ] | Integer Break | [LC 343](https://leetcode.com/problems/integer-break/) |
| [ ] | Unique BSTs | [LC 96](https://leetcode.com/problems/unique-binary-search-trees/) |

### 0/1 Knapsack

**Patterns:**
- [ ] 2D dp table → compress to 1D rolling array
- [ ] Iterate items outer, capacity inner (right to left for 0/1)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Partition Equal Subset Sum | [LC 416](https://leetcode.com/problems/partition-equal-subset-sum/) |
| [ ] | Last Stone Weight II | [LC 1049](https://leetcode.com/problems/last-stone-weight-ii/) |
| [ ] | Target Sum | [LC 494](https://leetcode.com/problems/target-sum/) |
| [ ] | Ones and Zeroes | [LC 474](https://leetcode.com/problems/ones-and-zeroes/) |

### Complete Knapsack

**Patterns:**
- [ ] Iterate capacity inner left to right (items can repeat)
- [ ] Order matters (permutations) → iterate capacity outer

| Done | Problem | Link |
|------|---------|------|
| [ ] | Coin Change II | [LC 518](https://leetcode.com/problems/coin-change-ii/) |
| [ ] | Combination Sum IV | [LC 377](https://leetcode.com/problems/combination-sum-iv/) |
| [ ] | Coin Change | [LC 322](https://leetcode.com/problems/coin-change/) |
| [ ] | Perfect Squares | [LC 279](https://leetcode.com/problems/perfect-squares/) |
| [ ] | Word Break | [LC 139](https://leetcode.com/problems/word-break/) |

### House Robber Series

| Done | Problem | Link |
|------|---------|------|
| [ ] | House Robber | [LC 198](https://leetcode.com/problems/house-robber/) |
| [ ] | House Robber II | [LC 213](https://leetcode.com/problems/house-robber-ii/) |
| [ ] | House Robber III (Tree DP) | [LC 337](https://leetcode.com/problems/house-robber-iii/) |

### Stock Trading

| Done | Problem | Link |
|------|---------|------|
| [ ] | Best Time to Buy and Sell Stock | [LC 121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) |
| [ ] | Best Time — Multiple Transactions | [LC 122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) |
| [ ] | Best Time — At Most 2 Transactions | [LC 123](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) |
| [ ] | Best Time — At Most K Transactions | [LC 188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) |
| [ ] | Best Time with Cooldown | [LC 309](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) |
| [ ] | Best Time with Transaction Fee | [LC 714](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) |

### Subsequence / String DP

| Done | Problem | Link |
|------|---------|------|
| [ ] | Longest Increasing Subsequence | [LC 300](https://leetcode.com/problems/longest-increasing-subsequence/) |
| [ ] | Longest Continuous Increasing Subsequence | [LC 674](https://leetcode.com/problems/longest-continuous-increasing-subsequence/) |
| [ ] | Longest Common Subsequence | [LC 1143](https://leetcode.com/problems/longest-common-subsequence/) |
| [ ] | Longest Repeating Subarray | [LC 718](https://leetcode.com/problems/maximum-length-of-repeated-subarray/) |
| [ ] | Is Subsequence | [LC 392](https://leetcode.com/problems/is-subsequence/) |
| [ ] | Distinct Subsequences | [LC 115](https://leetcode.com/problems/distinct-subsequences/) |
| [ ] | Delete Operation for Two Strings | [LC 583](https://leetcode.com/problems/delete-operation-for-two-strings/) |
| [ ] | Edit Distance | [LC 72](https://leetcode.com/problems/edit-distance/) |
| [ ] | Palindromic Substrings | [LC 647](https://leetcode.com/problems/palindromic-substrings/) |
| [ ] | Longest Palindromic Subsequence | [LC 516](https://leetcode.com/problems/longest-palindromic-subsequence/) |

---

## 11. Monotonic Stack

**Key Concepts:** Maintain a stack in monotonically increasing or decreasing order to answer "next greater/smaller element" queries in O(n).

**Patterns:**
- [ ] Next greater element (iterate left→right, pop when current > top)
- [ ] Circular arrays (iterate 2× with modulo index)
- [ ] Histogram / rain water trapping (area under bounded containers)

| Done | Problem | Link |
|------|---------|------|
| [ ] | Daily Temperatures | [LC 739](https://leetcode.com/problems/daily-temperatures/) |
| [ ] | Next Greater Element I | [LC 496](https://leetcode.com/problems/next-greater-element-i/) |
| [ ] | Next Greater Element II | [LC 503](https://leetcode.com/problems/next-greater-element-ii/) |
| [ ] | Trapping Rain Water | [LC 42](https://leetcode.com/problems/trapping-rain-water/) |
| [ ] | Largest Rectangle in Histogram | [LC 84](https://leetcode.com/problems/largest-rectangle-in-histogram/) |

---

## 12. Graph Theory

**Key Concepts:** DFS/BFS on graphs, Union-Find for connectivity, shortest paths (Dijkstra/Bellman-Ford/Floyd), minimum spanning tree (Prim/Kruskal), topological sort.

### DFS / BFS / Island Problems

**Patterns:**
- [ ] Treat 2D grid as a graph; DFS/BFS from each unvisited land cell
- [ ] Mark visited in-place (flood fill) or with a visited set

| Done | Problem | Link |
|------|---------|------|
| [ ] | Number of Islands | [LC 200](https://leetcode.com/problems/number-of-islands/) |
| [ ] | Max Area of Island | [LC 695](https://leetcode.com/problems/max-area-of-island/) |
| [ ] | Pacific Atlantic Water Flow | [LC 417](https://leetcode.com/problems/pacific-atlantic-water-flow/) |
| [ ] | Surrounded Regions | [LC 130](https://leetcode.com/problems/surrounded-regions/) |
| [ ] | All Paths From Source to Target | [LC 797](https://leetcode.com/problems/all-paths-from-source-to-target/) |

### Union-Find

**Patterns:**
- [ ] Path compression + union by rank for near-O(1) operations
- [ ] Detect cycle in undirected graph

| Done | Problem | Link |
|------|---------|------|
| [ ] | Redundant Connection | [LC 684](https://leetcode.com/problems/redundant-connection/) |
| [ ] | Redundant Connection II | [LC 685](https://leetcode.com/problems/redundant-connection-ii/) |

### Shortest Path

| Done | Problem | Link |
|------|---------|------|
| [ ] | Network Delay Time (Dijkstra) | [LC 743](https://leetcode.com/problems/network-delay-time/) |
| [ ] | Cheapest Flights Within K Stops (Bellman-Ford) | [LC 787](https://leetcode.com/problems/cheapest-flights-within-k-stops/) |
| [ ] | Find the City (Floyd-Warshall) | [LC 1334](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) |

### Topological Sort

| Done | Problem | Link |
|------|---------|------|
| [ ] | Course Schedule | [LC 207](https://leetcode.com/problems/course-schedule/) |
| [ ] | Course Schedule II | [LC 210](https://leetcode.com/problems/course-schedule-ii/) |

---

## Study Order

```
Arrays → Linked Lists → Hash Tables → Strings
       → Two Pointers (cross-cutting)
→ Stacks & Queues → Binary Trees → Backtracking
→ Greedy → Dynamic Programming → Monotonic Stack → Graph Theory
```

> **Note:** DP is the heaviest section — budget extra time for the knapsack framework and subsequence problems. Binary Trees is the prerequisite for both Backtracking (recursive tree thinking) and Graph Theory (DFS/BFS).
