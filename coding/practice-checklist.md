# Coding Practice Checklist — 120 Problems

---
## Parts
| # | Part | Sections |
|---|------|----------|
| I | Search & Indexing | Binary Search · Hash Tables & Prefix Sum |
| II | Linear Data Structures | Arrays & Strings · Two Pointers · Sliding Window · Linked Lists |
| III | Stack & Queue Patterns | Stack & Monotonic Stack · Heap / Priority Queue |
| IV | Tree Structures | Binary Trees · Binary Search Trees · Trie |
| V | Graph Algorithms | Graph Traversal · Graph Algorithms |
| VI | Algorithmic Paradigms | Backtracking · Greedy & Intervals · Dynamic Programming |
| VII | Advanced Topics | Data Structure Design · Bit Manipulation · Math · Segment Tree / Fenwick Tree |

---

# Part I — Search & Indexing

## 1. Binary Search

**Key Concepts:** Maintain a search invariant on a sorted space. Left/right boundary variants for finding exact values. Generalizes to any monotone predicate over a numeric range (binary search on the answer).

### Textbook — sorted array search

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Binary Search | [704](https://leetcode.com/problems/binary-search/) | [📖](https://programmercarl.com/0704.二分查找.html) |

### On Answer Space — binary search the answer, validate with a helper

> **Key idea:** When `feasible(x)` is monotone (once true, always true), binary search over the answer range. Template: `lo, hi = min_ans, max_ans; while lo < hi: mid=(lo+hi)//2; lo=mid+1 if feasible(mid) else hi=mid`.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Koko Eating Bananas | [875](https://leetcode.com/problems/koko-eating-bananas/) | — |

---

## 2. Hash Tables & Prefix Sum

**Key Concepts:** O(1) lookup via hashing. Choose the right structure: array (fixed alphabet), set (existence), map (complement/count). Prefix sum turns range-sum into a lookup problem; combining with a hash map finds subarrays with a target sum in O(n).

### Frequency Counting — fixed-size alphabet array

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Valid Anagram | [242](https://leetcode.com/problems/valid-anagram/) | [📖](https://programmercarl.com/0242.有效的字母异位词.html) |

### Set — existence check / cycle detection

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Intersection of Two Arrays | [349](https://leetcode.com/problems/intersection-of-two-arrays/) | [📖](https://programmercarl.com/0349.两个数组的交集.html) |
| [x] | Happy Number | [202](https://leetcode.com/problems/happy-number/) | [📖](https://programmercarl.com/0202.快乐数.html) |

### Map — complement lookup (Two Sum pattern)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Two Sum | [1](https://leetcode.com/problems/two-sum/) | [📖](https://programmercarl.com/0001.两数之和.html) |

### Prefix Sum + Map — subarray with target sum in O(n)

> **Key idea:** `prefix[i] = prefix[i-1] + nums[i-1]`. Store `prefix_sum → count` in a map. For each position, look up `prefix[i] - target` to count valid subarrays ending here.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Subarray Sum Equals K | [560](https://leetcode.com/problems/subarray-sum-equals-k/) | — |

---

# Part II — Linear Data Structures

## 3. Arrays & Strings

**Key Concepts:** Simulation with careful boundary invariants; in-place string reversal with two pointers; KMP for O(n) pattern matching via a prefix-failure table.

### Simulation — boundary-invariant traversal

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Spiral Matrix II | [59](https://leetcode.com/problems/spiral-matrix-ii/) | [📖](https://programmercarl.com/0059.螺旋矩阵II.html) |

### Two-Pointer Reversal — whole string / word-by-word

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Reverse String | [344](https://leetcode.com/problems/reverse-string/) | [📖](https://programmercarl.com/0344.反转字符串.html) |
| [ ] | Reverse Words in a String | [151](https://leetcode.com/problems/reverse-words-in-a-string/) | [📖](https://programmercarl.com/0151.翻转字符串里的单词.html) |

### KMP — O(n) pattern matching via prefix-failure table — [随想录 KMP理论](https://programmercarl.com/0028.实现strStr.html)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Find the Index of First Occurrence (strStr) | [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | [📖](https://programmercarl.com/0028.实现strStr.html) |

---

## 4. Two Pointers

**Key Concepts:** Eliminate a dimension of brute force by maintaining two indices that converge toward each other (left/right on sorted input) or advance together (same-direction overwrite). [随想录 双指针总结](https://programmercarl.com/双指针总结.html)

### Same-Direction — in-place overwrite

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Remove Element | [27](https://leetcode.com/problems/remove-element/) | [📖](https://programmercarl.com/0027.移除元素.html) |

### Left / Right — converge on sorted array

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Squares of a Sorted Array | [977](https://leetcode.com/problems/squares-of-a-sorted-array/) | [📖](https://programmercarl.com/0977.有序数组的平方.html) |

### kSum Reduction — sort then shrink with deduplication

> **Key idea:** Fix outer element(s), reduce to 2-Sum with a left/right pointer. Skip duplicate values at each pointer level to avoid repeated answers. Three Sum → O(n²); Four Sum → O(n³) → same pattern one level deeper.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Three Sum | [15](https://leetcode.com/problems/3sum/) | [📖](https://programmercarl.com/0015.三数之和.html) |
| [x] | Four Sum | [18](https://leetcode.com/problems/4sum/) | [📖](https://programmercarl.com/0018.四数之和.html) |

---

## 5. Sliding Window

**Key Concepts:** Maintain a valid window `[left, right]` over a sequence. Expand `right` to add elements; shrink `left` when the window violates the constraint. Each element enters and exits the window once → O(n). Track window state in a counter or hash map. [随想录 滑动窗口](https://programmercarl.com/0209.长度最小的子数组.html)

### Fixed-Alphabet Window — anagram / permutation check

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Minimum Size Subarray Sum | [209](https://leetcode.com/problems/minimum-size-subarray-sum/) | [📖](https://programmercarl.com/0209.长度最小的子数组.html) |

### Variable Window — longest substring under constraint

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Substring Without Repeating Characters | [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | — |
| [ ] | Longest Repeating Character Replacement | [424](https://leetcode.com/problems/longest-repeating-character-replacement/) | — |

### Variable Window — shortest substring covering constraint

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Minimum Window Substring | [76](https://leetcode.com/problems/minimum-window-substring/) | — |

---

## 6. Linked Lists

**Key Concepts:** Dummy head eliminates edge cases at the front. In-place reversal via pointer rethreading. Fast/slow pointer finds cycle entry and nth-from-end in one pass. Two lists reach their intersection at equal total distance.

### Dummy Head — simplify deletion at list head

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Remove Linked List Elements | [203](https://leetcode.com/problems/remove-linked-list-elements/) | [📖](https://programmercarl.com/0203.移除链表元素.html) |

### In-Place Reversal — rethread pointers

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Reverse Linked List | [206](https://leetcode.com/problems/reverse-linked-list/) | [📖](https://programmercarl.com/0206.翻转链表.html) |
| [ ] | Swap Nodes in Pairs | [24](https://leetcode.com/problems/swap-nodes-in-pairs/) | [📖](https://programmercarl.com/0024.两两交换链表中的节点.html) |

### Fast / Slow Pointers — nth-from-end & cycle detection

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove Nth Node From End | [19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [📖](https://programmercarl.com/0019.删除链表的倒数第N个节点.html) |
| [x] | Linked List Cycle II | [142](https://leetcode.com/problems/linked-list-cycle-ii/) | [📖](https://programmercarl.com/0142.环形链表II.html) |

### Pointer Reset — intersection via equal-distance traversal

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Intersection of Two Linked Lists | [160](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [📖](https://programmercarl.com/面试题02.07.链表相交.html) |

---

# Part III — Stack & Queue Patterns

## 7. Stack & Monotonic Stack

**Key Concepts:** Stack enforces LIFO processing — ideal for matching, evaluation, and deferred decisions. Monotonic stack (increasing or decreasing) answers "next greater/smaller" in O(n) by evicting stale elements on each push.

### Stack — mutual implementation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Implement Queue using Stacks | [232](https://leetcode.com/problems/implement-queue-using-stacks/) | [📖](https://programmercarl.com/0232.用栈实现队列.html) |

### Stack — bracket matching / expression evaluation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Valid Parentheses | [20](https://leetcode.com/problems/valid-parentheses/) | [📖](https://programmercarl.com/0020.有效的括号.html) |
| [ ] | Evaluate Reverse Polish Notation | [150](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | [📖](https://programmercarl.com/0150.逆波兰表达式求值.html) |

### Stack — adjacent duplicate removal

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove All Adjacent Duplicates In String | [1047](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) | [📖](https://programmercarl.com/1047.删除字符串中的所有相邻重复项.html) |

### Monotonic Deque — sliding window maximum — [随想录 单调栈理论](https://programmercarl.com/0739.每日温度.html)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Sliding Window Maximum | [239](https://leetcode.com/problems/sliding-window-maximum/) | [📖](https://programmercarl.com/0239.滑动窗口最大值.html) |

### Monotonic Stack — next greater element

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Daily Temperatures | [739](https://leetcode.com/problems/daily-temperatures/) | [📖](https://programmercarl.com/0739.每日温度.html) |
| [ ] | Next Greater Element II (circular) | [503](https://leetcode.com/problems/next-greater-element-ii/) | [📖](https://programmercarl.com/0503.下一个更大元素II.html) |

### Monotonic Stack — bounded area (histogram / rain water)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Trapping Rain Water | [42](https://leetcode.com/problems/trapping-rain-water/) | [📖](https://programmercarl.com/0042.接雨水.html) |
| [ ] | Largest Rectangle in Histogram | [84](https://leetcode.com/problems/largest-rectangle-in-histogram/) | [📖](https://programmercarl.com/0084.柱状图中最大的矩形.html) |

---

## 8. Heap / Priority Queue

**Key Concepts:** Min-heap of size K answers "K largest" queries in O(n log K). For k-way merge, push `(value, list_idx)` tuples and always pop the minimum. Two heaps (max-heap of lower half + min-heap of upper half) give streaming median in O(log n) per insert.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Top K Frequent Elements | [347](https://leetcode.com/problems/top-k-frequent-elements/) | [📖](https://programmercarl.com/0347.前K个高频元素.html) |
| [ ] | Merge K Sorted Lists | [23](https://leetcode.com/problems/merge-k-sorted-lists/) | — |

---

# Part IV — Tree Structures

## 9. Binary Trees

**Key Concepts:** Two traversal modes — BFS (level-order via queue) and DFS (pre/in/post-order, recursive or iterative). Most tree problems are DFS where return values bubble answers up the call stack. [随想录 二叉树理论](https://programmercarl.com/二叉树理论基础篇.html)

### BFS / Level-Order — queue-based layer processing

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Binary Tree Level Order Traversal | [102](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [📖](https://programmercarl.com/0102.二叉树的层序遍历.html) |

### DFS — tree properties (symmetry, depth, balance, paths)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Invert Binary Tree | [226](https://leetcode.com/problems/invert-binary-tree/) | [📖](https://programmercarl.com/0226.翻转二叉树.html) |
| [ ] | Symmetric Tree | [101](https://leetcode.com/problems/symmetric-tree/) | [📖](https://programmercarl.com/0101.对称二叉树.html) |
| [ ] | Maximum Depth | [104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [📖](https://programmercarl.com/0104.二叉树的最大深度.html) |
| [ ] | Balanced Binary Tree | [110](https://leetcode.com/problems/balanced-binary-tree/) | [📖](https://programmercarl.com/0110.平衡二叉树.html) |
| [ ] | Path Sum | [112](https://leetcode.com/problems/path-sum/) | [📖](https://programmercarl.com/0112.路径总和.html) |

### DFS — tree construction from traversal arrays

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Construct from Inorder & Postorder | [106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | [📖](https://programmercarl.com/0106.从中序与后序遍历序列构造二叉树.html) |

### DFS — Lowest Common Ancestor (LCA)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Lowest Common Ancestor (general) | [236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | [📖](https://programmercarl.com/0236.二叉树的最近公共祖先.html) |

---

## 10. Binary Search Trees

**Key Concepts:** BST invariant: left subtree < node < right subtree. In-order traversal yields a sorted sequence — use this to validate, search, and accumulate. Recursive structure means insert/delete/trim follow the same "go left or right" decision at each node.

### LCA & Validation — exploiting BST ordering

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Lowest Common Ancestor (BST) | [235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | [📖](https://programmercarl.com/0235.二叉搜索树的最近公共祖先.html) |
| [ ] | Validate BST | [98](https://leetcode.com/problems/validate-binary-search-tree/) | [📖](https://programmercarl.com/0098.验证二叉搜索树.html) |

### Structural Modification — insert, delete, build

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Insert into BST | [701](https://leetcode.com/problems/insert-into-a-binary-search-tree/) | [📖](https://programmercarl.com/0701.二叉搜索树中的插入操作.html) |
| [ ] | Delete Node in BST | [450](https://leetcode.com/problems/delete-node-in-a-bst/) | [📖](https://programmercarl.com/0450.删除二叉搜索树中的节点.html) |
| [ ] | Convert Sorted Array to BST | [108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) | [📖](https://programmercarl.com/0108.将有序数组转换为二叉搜索树.html) |

---

## 11. Trie (Prefix Tree)

**Key Concepts:** Each node holds one character; paths spell words. `is_end` flag marks complete words. `insert` and `search` are O(L). Build the `TrieNode` skeleton first — all trie problems reuse it. Combine with backtracking to prune word-search on a board.

> Not covered by 代码随想录.

### Implementation & Prefix Search

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Implement Trie (Prefix Tree) | [208](https://leetcode.com/problems/implement-trie-prefix-tree/) | — |
| [ ] | Design Add and Search Words | [211](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | — |

### Trie + Backtracking — prune board search with prefix check

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Word Search II | [212](https://leetcode.com/problems/word-search-ii/) | — |

---

# Part V — Graph Algorithms

## 12. Graph Traversal

**Key Concepts:** Model problems as graphs and walk them. DFS flood fill marks connected components. Multi-source BFS seeds all sources at layer 0 for simultaneous shortest distances. Reverse-BFS from the boundary identifies cells reachable from the edge. BFS on implicit graphs (strings, states) finds minimum transformations. [随想录 图论理论](https://programmercarl.com/图论理论基础.html)

### DFS Flood Fill — connected components from each unvisited cell

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Number of Islands | [200](https://leetcode.com/problems/number-of-islands/) | [📖](https://programmercarl.com/0200.岛屿数量.深搜版.html) |

### Reverse DFS / BFS — flood inward from the boundary

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Pacific Atlantic Water Flow | [417](https://leetcode.com/problems/pacific-atlantic-water-flow/) | [📖](https://programmercarl.com/0417.太平洋大西洋水流问题.html) |
| [ ] | Surrounded Regions | [130](https://leetcode.com/problems/surrounded-regions/) | [📖](https://programmercarl.com/0130.被围绕的区域.html) |

### Multi-Source BFS — simultaneous wavefront from all start nodes

> **Key idea:** Seed the BFS queue with *all* source nodes at distance 0. BFS naturally computes shortest distance from the nearest source to every reachable node in O(V+E).

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Rotting Oranges | [994](https://leetcode.com/problems/rotting-oranges/) | — |

### BFS on Implicit Graph — state as node, transformation as edge

> **Key idea:** Each unique state (string, board config) is a graph node. BFS finds the minimum number of transformations. A `visited` set prevents re-enqueuing.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Word Ladder | [127](https://leetcode.com/problems/word-ladder/) | — |

### DFS on DAG — enumerate all paths

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | All Paths From Source to Target | [797](https://leetcode.com/problems/all-paths-from-source-to-target/) | [📖](https://programmercarl.com/0797.所有可能的路径.html) |

---

## 13. Graph Algorithms

**Key Concepts:** Union-Find detects cycles and merges components in near-O(1) with path compression + union by rank. Dijkstra finds shortest paths in non-negative weighted graphs using a min-heap. Topological sort (Kahn's BFS) resolves dependencies by repeatedly removing in-degree-0 nodes. [随想录 并查集理论](https://programmercarl.com/图论-并查集理论基础.html)

### Union-Find — cycle detection / connectivity

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Redundant Connection | [684](https://leetcode.com/problems/redundant-connection/) | [📖](https://programmercarl.com/0684.冗余连接.html) |

### Dijkstra — single-source shortest path (non-negative weights)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Network Delay Time | [743](https://leetcode.com/problems/network-delay-time/) | [📖](https://programmercarl.com/0743.网络延迟时间.html) |

### Topological Sort — Kahn's BFS on in-degree

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Course Schedule | [207](https://leetcode.com/problems/course-schedule/) | [📖](https://programmercarl.com/0207.课程表.html) |
| [ ] | Course Schedule II | [210](https://leetcode.com/problems/course-schedule-ii/) | [📖](https://programmercarl.com/0210.课程表II.html) |

---

# Part VI — Algorithmic Paradigms

## 14. Backtracking

**Key Concepts:** Build a solution incrementally on a decision tree. At each node: **choose → recurse → unchoose**. Prune branches that cannot lead to a valid solution. [随想录 回溯理论](https://programmercarl.com/回溯算法理论基础.html)

### Combinations — no reuse, fixed pool

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combinations | [77](https://leetcode.com/problems/combinations/) | [📖](https://programmercarl.com/0077.组合.html) |
| [ ] | Letter Combinations of a Phone Number | [17](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [📖](https://programmercarl.com/0017.电话号码的字母组合.html) |

### Combinations — with reuse (unlimited picks per element)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combination Sum | [39](https://leetcode.com/problems/combination-sum/) | [📖](https://programmercarl.com/0039.组合总和.html) |

### Deduplication — sort + skip duplicate at same tree level

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combination Sum II | [40](https://leetcode.com/problems/combination-sum-ii/) | [📖](https://programmercarl.com/0040.组合总和II.html) |

### Subsets — collect every node of the decision tree

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Subsets | [78](https://leetcode.com/problems/subsets/) | [📖](https://programmercarl.com/0078.子集.html) |

### Permutations — order matters, `used[]` array

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Permutations | [46](https://leetcode.com/problems/permutations/) | [📖](https://programmercarl.com/0046.全排列.html) |
| [ ] | Permutations II | [47](https://leetcode.com/problems/permutations-ii/) | [📖](https://programmercarl.com/0047.全排列II.html) |

### Partitioning — split string into valid segments

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Palindrome Partitioning | [131](https://leetcode.com/problems/palindrome-partitioning/) | [📖](https://programmercarl.com/0131.分割回文串.html) |

### Board Problems — 2D constraint propagation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | N-Queens | [51](https://leetcode.com/problems/n-queens/) | [📖](https://programmercarl.com/0051.N皇后.html) |
| [ ] | Sudoku Solver | [37](https://leetcode.com/problems/sudoku-solver/) | [📖](https://programmercarl.com/0037.解数独.html) |

---

## 15. Greedy & Intervals

**Key Concepts:** Make the locally optimal choice at each step; no backtracking. Proof of correctness is required — not just intuition. Interval problems almost always start with sorting; minimum-room problems use a min-heap on end times. [随想录 贪心理论](https://programmercarl.com/贪心算法理论基础.html)

### Local Optimum — single-pass greedy decision

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Maximum Subarray | [53](https://leetcode.com/problems/maximum-subarray/) | [📖](https://programmercarl.com/0053.最大子序和.html) |
| [ ] | Best Time to Buy and Sell Stock II | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II.html) |
| [ ] | Gas Station | [134](https://leetcode.com/problems/gas-station/) | [📖](https://programmercarl.com/0134.加油站.html) |

### Jump Game — track maximum reachable index

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Jump Game | [55](https://leetcode.com/problems/jump-game/) | [📖](https://programmercarl.com/0055.跳跃游戏.html) |
| [ ] | Jump Game II | [45](https://leetcode.com/problems/jump-game-ii/) | [📖](https://programmercarl.com/0045.跳跃游戏II.html) |

### Two-Pass Greedy — left→right then right→left

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Candy | [135](https://leetcode.com/problems/candy/) | [📖](https://programmercarl.com/0135.分发糖果.html) |
| [ ] | Queue Reconstruction by Height | [406](https://leetcode.com/problems/queue-reconstruction-by-height/) | [📖](https://programmercarl.com/0406.根据身高重建队列.html) |

### Interval Scheduling — sort by end time

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Minimum Arrows to Burst Balloons | [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | [📖](https://programmercarl.com/0452.用最少数量的箭引爆气球.html) |
| [ ] | Non-overlapping Intervals | [435](https://leetcode.com/problems/non-overlapping-intervals/) | [📖](https://programmercarl.com/0435.无重叠区间.html) |
| [ ] | Merge Intervals | [56](https://leetcode.com/problems/merge-intervals/) | [📖](https://programmercarl.com/0056.合并区间.html) |

### Interval Insertion & Intersection — linear scan

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Insert Interval | [57](https://leetcode.com/problems/insert-interval/) | — |
| [ ] | Interval List Intersections | [986](https://leetcode.com/problems/interval-list-intersections/) | — |

### Interval Scheduling — minimum rooms via heap on end times

> **Key idea:** Sort by start; maintain a min-heap of end times. If `heap[0] <= current.start`, pop and reuse that room. Otherwise push a new room. Heap size = answer.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Meeting Rooms II | [253](https://leetcode.com/problems/meeting-rooms-ii/) | — |

### Tree Greedy — bottom-up post-order decision

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Tree Cameras | [968](https://leetcode.com/problems/binary-tree-cameras/) | [📖](https://programmercarl.com/0968.监控二叉树.html) |

---

## 16. Dynamic Programming

**Key Concepts:** Optimal substructure + overlapping subproblems. Steps: define state → write recurrence → determine traversal order → initialize base cases. [随想录 DP理论](https://programmercarl.com/动态规划理论基础.html)

### 1D DP — adjacent-state recurrence

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [x] | Climbing Stairs | [70](https://leetcode.com/problems/climbing-stairs/) | [📖](https://programmercarl.com/0070.爬楼梯.html) |
| [ ] | House Robber | [198](https://leetcode.com/problems/house-robber/) | [📖](https://programmercarl.com/0198.打家劫舍.html) |
| [ ] | House Robber II (circular array) | [213](https://leetcode.com/problems/house-robber-ii/) | [📖](https://programmercarl.com/0213.打家劫舍II.html) |

### 2D DP — grid path counting

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Unique Paths | [62](https://leetcode.com/problems/unique-paths/) | [📖](https://programmercarl.com/0062.不同路径.html) |
| [ ] | Unique Paths II | [63](https://leetcode.com/problems/unique-paths-ii/) | [📖](https://programmercarl.com/0063.不同路径II.html) |

### 0/1 Knapsack — each item used at most once — [随想录 0/1背包理论](https://programmercarl.com/背包理论基础01背包-1.html)

> Traverse items outer, capacity **right → left** to prevent item reuse.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Partition Equal Subset Sum | [416](https://leetcode.com/problems/partition-equal-subset-sum/) | [📖](https://programmercarl.com/0416.分割等和子集.html) |

### Complete Knapsack — items reusable — [随想录 完全背包理论](https://programmercarl.com/背包问题理论基础完全背包.html)

> Traverse items outer, capacity **left → right** to allow item reuse.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Coin Change II | [518](https://leetcode.com/problems/coin-change-ii/) | [📖](https://programmercarl.com/0518.零钱兑换II.html) |
| [ ] | Coin Change | [322](https://leetcode.com/problems/coin-change/) | [📖](https://programmercarl.com/0322.零钱兑换.html) |
| [ ] | Word Break | [139](https://leetcode.com/problems/word-break/) | [📖](https://programmercarl.com/0139.单词拆分.html) |

### Stock Trading — state machine DP (hold / not-hold transitions)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Best Time to Buy and Sell Stock (1 tx) | [121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [📖](https://programmercarl.com/0121.买卖股票的最佳时机.html) |
| [ ] | Best Time — unlimited transactions | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II（动态规划）.html) |
| [ ] | Best Time — at most 2 transactions | [123](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [📖](https://programmercarl.com/0123.买卖股票的最佳时机III.html) |
| [ ] | Best Time with Cooldown | [309](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [📖](https://programmercarl.com/0309.最佳买卖股票时机含冷冻期.html) |
| [ ] | Best Time with Transaction Fee | [714](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [📖](https://programmercarl.com/0714.买卖股票的最佳时机含手续费（动态规划）.html) |

### Subsequence DP — 1D (LIS-style)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Increasing Subsequence | [300](https://leetcode.com/problems/longest-increasing-subsequence/) | [📖](https://programmercarl.com/0300.最长上升子序列.html) |

### Subsequence DP — 2D two-string comparison

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Common Subsequence | [1143](https://leetcode.com/problems/longest-common-subsequence/) | [📖](https://programmercarl.com/1143.最长公共子序列.html) |
| [ ] | Is Subsequence | [392](https://leetcode.com/problems/is-subsequence/) | [📖](https://programmercarl.com/0392.判断子序列.html) |
| [ ] | Distinct Subsequences | [115](https://leetcode.com/problems/distinct-subsequences/) | [📖](https://programmercarl.com/0115.不同的子序列.html) |
| [ ] | Edit Distance | [72](https://leetcode.com/problems/edit-distance/) | [📖](https://programmercarl.com/0072.编辑距离.html) |

### Palindrome DP — expand from center / 2D interval

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Palindromic Substrings | [647](https://leetcode.com/problems/palindromic-substrings/) | [📖](https://programmercarl.com/0647.回文子串.html) |

### Interval DP — optimize over all ways to split `[i..j]`

> **Key idea:** `dp[i][j]` = optimal cost for subproblem on `[i..j]`. Enumerate split point `k` in a third loop. Build from length-1 intervals up to the full range.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Burst Balloons | [312](https://leetcode.com/problems/burst-balloons/) | — |

---

# Part VII — Advanced Topics

## 17. Data Structure Design

**Key Concepts:** Combine base structures to hit all required complexities simultaneously. LRU = doubly-linked list (O(1) move-to-front) + hash map (O(1) lookup). Two-heap design (max-heap of lower half + min-heap of upper half) streams the median in O(log n) per insert.

> Not covered by 代码随想录. Interviewers assess *which structures you combine and why*, not just the final code.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Min Stack | [155](https://leetcode.com/problems/min-stack/) | — |
| [ ] | LRU Cache | [146](https://leetcode.com/problems/lru-cache/) | — |
| [ ] | Find Median from Data Stream | [295](https://leetcode.com/problems/find-median-from-data-stream/) | — |
| [ ] | Insert Delete GetRandom O(1) | [380](https://leetcode.com/problems/insert-delete-getrandom-o1/) | — |

---

## 18. Bit Manipulation

**Key Concepts:** Core idioms — `n & (n-1)` clears the lowest set bit; `a ^ a = 0` cancels duplicates; `a ^ 0 = a` preserves; `n & (-n)` isolates lowest set bit (used in Fenwick Tree update). These appear short in code but require knowing the trick cold.

> Not covered by 代码随想录.

### XOR — cancellation and identity

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Single Number | [136](https://leetcode.com/problems/single-number/) | — |
| [ ] | Missing Number | [268](https://leetcode.com/problems/missing-number/) | — |

### Bit Counting

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Number of 1 Bits | [191](https://leetcode.com/problems/number-of-1-bits/) | — |
| [ ] | Counting Bits | [338](https://leetcode.com/problems/counting-bits/) | — |

---

## 19. Math

**Key Concepts:** Fast exponentiation reduces `pow(x, n)` to O(log n) by halving the exponent each step. Sieve of Eratosthenes marks all non-primes in O(n log log n). Big-integer multiplication simulates grade-school column arithmetic.

> Not covered by 代码随想录. These surface as sub-problems in harder questions.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Pow(x, n) — repeated squaring | [50](https://leetcode.com/problems/powx-n/) | — |
| [ ] | Count Primes — Sieve of Eratosthenes | [204](https://leetcode.com/problems/count-primes/) | — |
| [ ] | Multiply Strings — column multiplication | [43](https://leetcode.com/problems/multiply-strings/) | — |

---

## 20. Segment Tree / Fenwick Tree (BIT)

**Key Concepts:** Both support point-update + prefix-query in O(log n). Fenwick Tree update: `i += i & (-i)`; query: `i -= i & (-i)` — ~10 lines total. Use when naive prefix sum breaks due to mutations. Segment Tree is more flexible (range update, range min/max) but more code.

> Not covered by 代码随想录. Appears at Databricks and Google for mutable range queries. Practice writing Fenwick Tree from scratch.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Range Sum Query — Mutable | [307](https://leetcode.com/problems/range-sum-query-mutable/) | — |
| [ ] | Count of Smaller Numbers After Self | [315](https://leetcode.com/problems/count-of-smaller-numbers-after-self/) | — |

---

## Study Order & Priority

```
Part I  : Binary Search → Hash Tables & Prefix Sum
Part II : Arrays & Strings → Two Pointers → Sliding Window → Linked Lists
Part III: Stack & Monotonic Stack → Heap
Part IV : Binary Trees → BST → Trie
Part V  : Graph Traversal → Graph Algorithms
Part VI : Backtracking → Greedy & Intervals → Dynamic Programming
Part VII: Data Structure Design → Bit Manipulation → Math → Segment Tree
```

| Tier | Sections | Target Companies |
|------|----------|-----------------|
| 🔴 Must | Binary Search, Hash Tables, Two Pointers, Sliding Window, Linked Lists, Stack, Trees, Backtracking, DP, Graph Traversal | All |
| 🟠 High | Heap, BST, Greedy & Intervals, Graph Algorithms, Data Structure Design, Trie | All |
| 🟡 Differentiator | Bit Manipulation, Interval DP, Math, Segment Tree | Google, Databricks |
