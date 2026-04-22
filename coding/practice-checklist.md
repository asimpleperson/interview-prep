# Coding Practice Checklist (代码随想录 + Staff-Level Supplements)

---

## 1. Arrays

**Key Concepts:** Contiguous memory, index-based access, in-place operations, prefix sums, sliding window, binary search.

### Binary Search — textbook (left/right boundary variants)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Search | [704](https://leetcode.com/problems/binary-search/) | [📖](https://programmercarl.com/0704.二分查找.html) |

### Binary Search on Answer Space — predicate over a monotonic range

> **Key idea:** When you can define a monotone predicate `feasible(x)` (e.g., "can we do it with capacity x?"), binary search the answer directly instead of the array. Template: `lo, hi = min_ans, max_ans; while lo < hi: mid = (lo+hi)//2; lo = mid+1 if feasible(mid) else hi = mid`.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Sqrt(x) | [69](https://leetcode.com/problems/sqrtx/) | — |
| [ ] | Koko Eating Bananas | [875](https://leetcode.com/problems/koko-eating-bananas/) | — |
| [ ] | Capacity to Ship Packages Within D Days | [1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) | — |
| [ ] | Split Array Largest Sum | [410](https://leetcode.com/problems/split-array-largest-sum/) | — |
| [ ] | Find Peak Element | [162](https://leetcode.com/problems/find-peak-element/) | — |

### Two Pointers — in-place removal / sorted merge

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove Element | [27](https://leetcode.com/problems/remove-element/) | [📖](https://programmercarl.com/0027.移除元素.html) |
| [ ] | Squares of a Sorted Array | [977](https://leetcode.com/problems/squares-of-a-sorted-array/) | [📖](https://programmercarl.com/0977.有序数组的平方.html) |

### Prefix Sum — range query in O(1) after O(n) build

> **Key idea:** `prefix[i] = prefix[i-1] + nums[i-1]`. Range sum `[l, r]` = `prefix[r+1] - prefix[l]`. Combine with a hash map to find subarrays with a target sum in O(n): store `prefix → count` and look up `prefix[i] - target`.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Range Sum Query — Immutable | [303](https://leetcode.com/problems/range-sum-query-immutable/) | — |
| [ ] | Product of Array Except Self | [238](https://leetcode.com/problems/product-of-array-except-self/) | — |
| [ ] | Subarray Sum Equals K | [560](https://leetcode.com/problems/subarray-sum-equals-k/) | — |
| [ ] | Subarray Sums Divisible by K | [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) | — |

### Sliding Window — minimum/maximum subarray

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Minimum Size Subarray Sum | [209](https://leetcode.com/problems/minimum-size-subarray-sum/) | [📖](https://programmercarl.com/0209.长度最小的子数组.html) |

### Simulation (boundary invariant)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Spiral Matrix II | [59](https://leetcode.com/problems/spiral-matrix-ii/) | [📖](https://programmercarl.com/0059.螺旋矩阵II.html) |

---

## 2. Linked Lists

**Key Concepts:** Pointer manipulation, dummy head node trick, fast/slow pointers, reversal, cycle detection.

### Dummy Head Node — simplify edge cases

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove Linked List Elements | [203](https://leetcode.com/problems/remove-linked-list-elements/) | [📖](https://programmercarl.com/0203.移除链表元素.html) |
| [ ] | Design Linked List | [707](https://leetcode.com/problems/design-linked-list/) | [📖](https://programmercarl.com/0707.设计链表.html) |

### In-place Reversal

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Reverse Linked List | [206](https://leetcode.com/problems/reverse-linked-list/) | [📖](https://programmercarl.com/0206.翻转链表.html) |
| [ ] | Swap Nodes in Pairs | [24](https://leetcode.com/problems/swap-nodes-in-pairs/) | [📖](https://programmercarl.com/0024.两两交换链表中的节点.html) |

### Fast / Slow Pointers — nth-from-end, cycle

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove Nth Node From End | [19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [📖](https://programmercarl.com/0019.删除链表的倒数第N个节点.html) |
| [ ] | Linked List Cycle II | [142](https://leetcode.com/problems/linked-list-cycle-ii/) | [📖](https://programmercarl.com/0142.环形链表II.html) |

### Intersection Detection — pointer reset trick

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Intersection of Two Linked Lists | [160](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [📖](https://programmercarl.com/面试题02.07.链表相交.html) |

---

## 3. Hash Tables

**Key Concepts:** O(1) lookup, collision resolution, choosing the right structure (array vs. set vs. map).

### Frequency Counting with Array (fixed alphabet)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Valid Anagram | [242](https://leetcode.com/problems/valid-anagram/) | [📖](https://programmercarl.com/0242.有效的字母异位词.html) |
| [ ] | Ransom Note | [383](https://leetcode.com/problems/ransom-note/) | [📖](https://programmercarl.com/0383.赎金信.html) |

### Set — existence check / intersection

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Intersection of Two Arrays | [349](https://leetcode.com/problems/intersection-of-two-arrays/) | [📖](https://programmercarl.com/0349.两个数组的交集.html) |
| [ ] | Happy Number | [202](https://leetcode.com/problems/happy-number/) | [📖](https://programmercarl.com/0202.快乐数.html) |

### Map — complement lookup (Two Sum pattern)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Two Sum | [1](https://leetcode.com/problems/two-sum/) | [📖](https://programmercarl.com/0001.两数之和.html) |
| [ ] | Four Sum II | [454](https://leetcode.com/problems/4sum-ii/) | [📖](https://programmercarl.com/0454.四数相加II.html) |

### kSum — sort + two-pointer with deduplication

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Three Sum | [15](https://leetcode.com/problems/3sum/) | [📖](https://programmercarl.com/0015.三数之和.html) |
| [ ] | Four Sum | [18](https://leetcode.com/problems/4sum/) | [📖](https://programmercarl.com/0018.四数之和.html) |

---

## 4. Strings

**Key Concepts:** In-place reversal, KMP pattern matching, prefix table (failure function), repeated substring detection.

### Two-Pointer Reversal (whole / partial / word-by-word)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Reverse String | [344](https://leetcode.com/problems/reverse-string/) | [📖](https://programmercarl.com/0344.反转字符串.html) |
| [ ] | Reverse String II | [541](https://leetcode.com/problems/reverse-string-ii/) | [📖](https://programmercarl.com/0541.反转字符串II.html) |
| [ ] | Reverse Words in a String | [151](https://leetcode.com/problems/reverse-words-in-a-string/) | [📖](https://programmercarl.com/0151.翻转字符串里的单词.html) |

### KMP — O(n) substring search via prefix table — [随想录 KMP理论](https://programmercarl.com/0028.实现strStr.html)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Find the Index of the First Occurrence (strStr) | [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | [📖](https://programmercarl.com/0028.实现strStr.html) |
| [ ] | Repeated Substring Pattern | [459](https://leetcode.com/problems/repeated-substring-pattern/) | [📖](https://programmercarl.com/0459.重复的子字符串.html) |

---

## 5. Two Pointers & Sliding Window

**Key Concepts:** Two pointers reduce O(n²) brute force to O(n). Sliding window maintains a valid substructure under a shrink-or-expand constraint. [随想录 双指针总结](https://programmercarl.com/双指针总结.html)

> Two-pointer problems are distributed across Arrays, Strings, Linked Lists, and Hash Tables above. This section focuses on the **variable-size sliding window** pattern as its own first-class topic.

### Sliding Window — longest/shortest substring under a constraint

> **Key idea:** Expand `right` to include new chars; shrink `left` when the window violates the constraint. Track the window state in a hash map or counter. O(n) because each element enters and exits the window once.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Substring Without Repeating Characters | [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | — |
| [ ] | Permutation in String | [567](https://leetcode.com/problems/permutation-in-string/) | — |
| [ ] | Find All Anagrams in a String | [438](https://leetcode.com/problems/find-all-anagrams-in-a-string/) | — |
| [ ] | Longest Repeating Character Replacement | [424](https://leetcode.com/problems/longest-repeating-character-replacement/) | — |
| [ ] | Minimum Window Substring | [76](https://leetcode.com/problems/minimum-window-substring/) | — |

| Pattern | Representative Problems |
|---------|------------------------|
| Fast / slow (cycle, midpoint) | LC 142, LC 19 |
| Left / right on sorted array (kSum) | LC 977, LC 15, LC 18 |
| Same-direction overwrite (in-place remove) | LC 27, LC 209 |

---

## 6. Stacks & Queues

**Key Concepts:** LIFO vs. FIFO, monotonic deque for sliding window max, priority queue for top-K.

### Mutual Implementation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Implement Queue using Stacks | [232](https://leetcode.com/problems/implement-queue-using-stacks/) | [📖](https://programmercarl.com/0232.用栈实现队列.html) |
| [ ] | Implement Stack using Queues | [225](https://leetcode.com/problems/implement-stack-using-queues/) | [📖](https://programmercarl.com/0225.用队列实现栈.html) |

### Stack — bracket matching / expression evaluation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Valid Parentheses | [20](https://leetcode.com/problems/valid-parentheses/) | [📖](https://programmercarl.com/0020.有效的括号.html) |
| [ ] | Evaluate Reverse Polish Notation | [150](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | [📖](https://programmercarl.com/0150.逆波兰表达式求值.html) |

### Stack — adjacent duplicate removal

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove All Adjacent Duplicates In String | [1047](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) | [📖](https://programmercarl.com/1047.删除字符串中的所有相邻重复项.html) |

### Monotonic Deque — sliding window maximum

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Sliding Window Maximum | [239](https://leetcode.com/problems/sliding-window-maximum/) | [📖](https://programmercarl.com/0239.滑动窗口最大值.html) |

### Heap / Priority Queue — top-K and streaming median

> **Key idea:** A min-heap of size K answers "K largest" queries in O(n log K). Two heaps (max-heap of lower half + min-heap of upper half, kept balanced) give streaming median in O(log n) per insert.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Top K Frequent Elements | [347](https://leetcode.com/problems/top-k-frequent-elements/) | [📖](https://programmercarl.com/0347.前K个高频元素.html) |
| [ ] | Last Stone Weight | [1046](https://leetcode.com/problems/last-stone-weight/) | — |
| [ ] | Kth Largest Element in an Array | [215](https://leetcode.com/problems/kth-largest-element-in-an-array/) | — |
| [ ] | Merge K Sorted Lists | [23](https://leetcode.com/problems/merge-k-sorted-lists/) | — |
| [ ] | Find Median from Data Stream | [295](https://leetcode.com/problems/find-median-from-data-stream/) | — |
| [ ] | Kth Smallest Element in Sorted Matrix | [378](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/) | — |
| [ ] | Task Scheduler | [621](https://leetcode.com/problems/task-scheduler/) | — |

---

## 7. Binary Trees

**Key Concepts:** DFS (pre/in/post-order) and BFS (level-order), recursive vs. iterative traversal, BST properties. [随想录 二叉树理论](https://programmercarl.com/二叉树理论基础篇.html)

### BFS / Level-Order (queue)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Tree Level Order Traversal | [102](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [📖](https://programmercarl.com/0102.二叉树的层序遍历.html) |

### Recursive DFS — tree properties (depth, balance, paths, merge)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Invert Binary Tree | [226](https://leetcode.com/problems/invert-binary-tree/) | [📖](https://programmercarl.com/0226.翻转二叉树.html) |
| [ ] | Symmetric Tree | [101](https://leetcode.com/problems/symmetric-tree/) | [📖](https://programmercarl.com/0101.对称二叉树.html) |
| [ ] | Maximum Depth | [104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [📖](https://programmercarl.com/0104.二叉树的最大深度.html) |
| [ ] | Minimum Depth | [111](https://leetcode.com/problems/minimum-depth-of-binary-tree/) | [📖](https://programmercarl.com/0111.二叉树的最小深度.html) |
| [ ] | Count Complete Tree Nodes | [222](https://leetcode.com/problems/count-complete-tree-nodes/) | [📖](https://programmercarl.com/0222.完全二叉树的节点个数.html) |
| [ ] | Balanced Binary Tree | [110](https://leetcode.com/problems/balanced-binary-tree/) | [📖](https://programmercarl.com/0110.平衡二叉树.html) |
| [ ] | All Root-to-Leaf Paths | [257](https://leetcode.com/problems/binary-tree-paths/) | [📖](https://programmercarl.com/0257.二叉树的所有路径.html) |
| [ ] | Path Sum | [112](https://leetcode.com/problems/path-sum/) | [📖](https://programmercarl.com/0112.路径总和.html) |
| [ ] | Merge Two Binary Trees | [617](https://leetcode.com/problems/merge-two-binary-trees/) | [📖](https://programmercarl.com/0617.合并二叉树.html) |

### Tree Construction from Traversal Arrays

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Construct from Inorder & Postorder | [106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | [📖](https://programmercarl.com/0106.从中序与后序遍历序列构造二叉树.html) |
| [ ] | Maximum Binary Tree | [654](https://leetcode.com/problems/maximum-binary-tree/) | [📖](https://programmercarl.com/0654.最大二叉树.html) |

### Lowest Common Ancestor (LCA) — recursion return value

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Lowest Common Ancestor (general) | [236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | [📖](https://programmercarl.com/0236.二叉树的最近公共祖先.html) |
| [ ] | Lowest Common Ancestor (BST) | [235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | [📖](https://programmercarl.com/0235.二叉搜索树的最近公共祖先.html) |

### BST — in-order gives sorted sequence

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Validate BST | [98](https://leetcode.com/problems/validate-binary-search-tree/) | [📖](https://programmercarl.com/0098.验证二叉搜索树.html) |
| [ ] | Search in BST | [700](https://leetcode.com/problems/search-in-a-binary-search-tree/) | [📖](https://programmercarl.com/0700.二叉搜索树中的搜索.html) |
| [ ] | Convert BST to Greater Tree | [538](https://leetcode.com/problems/convert-bst-to-greater-tree/) | [📖](https://programmercarl.com/0538.把二叉搜索树转换为累加树.html) |

### BST — structural modification (insert, delete, trim, build)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Insert into BST | [701](https://leetcode.com/problems/insert-into-a-binary-search-tree/) | [📖](https://programmercarl.com/0701.二叉搜索树中的插入操作.html) |
| [ ] | Delete Node in BST | [450](https://leetcode.com/problems/delete-node-in-a-bst/) | [📖](https://programmercarl.com/0450.删除二叉搜索树中的节点.html) |
| [ ] | Trim BST | [669](https://leetcode.com/problems/trim-a-binary-search-tree/) | [📖](https://programmercarl.com/0669.修剪二叉搜索树.html) |
| [ ] | Convert Sorted Array to BST | [108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) | [📖](https://programmercarl.com/0108.将有序数组转换为二叉搜索树.html) |

---

## 8. Backtracking

**Key Concepts:** Exhaustive search on a decision tree; prune branches early. Framework: **choose → explore → unchoose**. [随想录 回溯理论](https://programmercarl.com/回溯算法理论基础.html)

### Combinations — no reuse, fixed pool

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combinations | [77](https://leetcode.com/problems/combinations/) | [📖](https://programmercarl.com/0077.组合.html) |
| [ ] | Combination Sum III | [216](https://leetcode.com/problems/combination-sum-iii/) | [📖](https://programmercarl.com/0216.组合总和III.html) |
| [ ] | Letter Combinations of a Phone Number | [17](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [📖](https://programmercarl.com/0017.电话号码的字母组合.html) |

### Combinations — with reuse (unlimited picks per element)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combination Sum | [39](https://leetcode.com/problems/combination-sum/) | [📖](https://programmercarl.com/0039.组合总和.html) |

### Deduplication — sort + skip duplicate at same tree level

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combination Sum II | [40](https://leetcode.com/problems/combination-sum-ii/) | [📖](https://programmercarl.com/0040.组合总和II.html) |
| [ ] | Subsets II | [90](https://leetcode.com/problems/subsets-ii/) | [📖](https://programmercarl.com/0090.子集II.html) |
| [ ] | Permutations II | [47](https://leetcode.com/problems/permutations-ii/) | [📖](https://programmercarl.com/0047.全排列II.html) |

### Subsets — collect every node of the decision tree

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Subsets | [78](https://leetcode.com/problems/subsets/) | [📖](https://programmercarl.com/0078.子集.html) |
| [ ] | Non-decreasing Subsequences | [491](https://leetcode.com/problems/non-decreasing-subsequences/) | [📖](https://programmercarl.com/0491.递增子序列.html) |

### Permutations — order matters, `used[]` array

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Permutations | [46](https://leetcode.com/problems/permutations/) | [📖](https://programmercarl.com/0046.全排列.html) |

### Partitioning — split string into valid segments

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Palindrome Partitioning | [131](https://leetcode.com/problems/palindrome-partitioning/) | [📖](https://programmercarl.com/0131.分割回文串.html) |
| [ ] | Restore IP Addresses | [93](https://leetcode.com/problems/restore-ip-addresses/) | [📖](https://programmercarl.com/0093.复原IP地址.html) |

### Board Problems — 2D constraint propagation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | N-Queens | [51](https://leetcode.com/problems/n-queens/) | [📖](https://programmercarl.com/0051.N皇后.html) |
| [ ] | Sudoku Solver | [37](https://leetcode.com/problems/sudoku-solver/) | [📖](https://programmercarl.com/0037.解数独.html) |

---

## 9. Greedy Algorithms

**Key Concepts:** Make the locally optimal choice at each step; no rollback. Correctness requires proof — not just intuition. [随想录 贪心理论](https://programmercarl.com/贪心算法理论基础.html)

### Local Max / Greedy Scan — single pass decision

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Assign Cookies | [455](https://leetcode.com/problems/assign-cookies/) | [📖](https://programmercarl.com/0455.分发饼干.html) |
| [ ] | Wiggle Subsequence | [376](https://leetcode.com/problems/wiggle-subsequence/) | [📖](https://programmercarl.com/0376.摆动序列.html) |
| [ ] | Maximum Subarray | [53](https://leetcode.com/problems/maximum-subarray/) | [📖](https://programmercarl.com/0053.最大子序和.html) |
| [ ] | Maximize Sum After K Negations | [1005](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/) | [📖](https://programmercarl.com/1005.K次取反后最大化的数组和.html) |
| [ ] | Gas Station | [134](https://leetcode.com/problems/gas-station/) | [📖](https://programmercarl.com/0134.加油站.html) |
| [ ] | Lemonade Change | [860](https://leetcode.com/problems/lemonade-change/) | [📖](https://programmercarl.com/0860.柠檬水找零.html) |
| [ ] | Monotone Increasing Digits | [738](https://leetcode.com/problems/monotone-increasing-digits/) | [📖](https://programmercarl.com/0738.单调递增的数字.html) |

### Stock — accumulate every positive daily difference

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Best Time to Buy and Sell Stock II | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II.html) |

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

### Interval Scheduling — sort by end time, count non-overlapping

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Minimum Arrows to Burst Balloons | [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | [📖](https://programmercarl.com/0452.用最少数量的箭引爆气球.html) |
| [ ] | Non-overlapping Intervals | [435](https://leetcode.com/problems/non-overlapping-intervals/) | [📖](https://programmercarl.com/0435.无重叠区间.html) |
| [ ] | Partition Labels | [763](https://leetcode.com/problems/partition-labels/) | [📖](https://programmercarl.com/0763.划分字母区间.html) |
| [ ] | Merge Intervals | [56](https://leetcode.com/problems/merge-intervals/) | [📖](https://programmercarl.com/0056.合并区间.html) |

### Tree Greedy — post-order traversal, bottom-up decision

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Tree Cameras | [968](https://leetcode.com/problems/binary-tree-cameras/) | [📖](https://programmercarl.com/0968.监控二叉树.html) |

---

## 10. Dynamic Programming

**Key Concepts:** Optimal substructure + overlapping subproblems. Define state → write recurrence → determine traversal order → initialize base cases. [随想录 DP理论](https://programmercarl.com/动态规划理论基础.html)

### 1D DP — Fibonacci-style (adjacent states)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Fibonacci Number | [509](https://leetcode.com/problems/fibonacci-number/) | [📖](https://programmercarl.com/0509.斐波那契数.html) |
| [ ] | Climbing Stairs | [70](https://leetcode.com/problems/climbing-stairs/) | [📖](https://programmercarl.com/0070.爬楼梯.html) |
| [ ] | Min Cost Climbing Stairs | [746](https://leetcode.com/problems/min-cost-climbing-stairs/) | [📖](https://programmercarl.com/0746.使用最小花费爬楼梯.html) |

### 2D DP — grid path counting

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Unique Paths | [62](https://leetcode.com/problems/unique-paths/) | [📖](https://programmercarl.com/0062.不同路径.html) |
| [ ] | Unique Paths II | [63](https://leetcode.com/problems/unique-paths-ii/) | [📖](https://programmercarl.com/0063.不同路径II.html) |

### Math-based DP — combinatorial counting

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Integer Break | [343](https://leetcode.com/problems/integer-break/) | [📖](https://programmercarl.com/0343.整数拆分.html) |
| [ ] | Unique BSTs | [96](https://leetcode.com/problems/unique-binary-search-trees/) | [📖](https://programmercarl.com/0096.不同的二叉搜索树.html) |

### 0/1 Knapsack — each item used at most once — [随想录 0/1背包理论](https://programmercarl.com/背包理论基础01背包-1.html)

> Traverse items outer, capacity **right → left** to prevent reuse.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Partition Equal Subset Sum | [416](https://leetcode.com/problems/partition-equal-subset-sum/) | [📖](https://programmercarl.com/0416.分割等和子集.html) |
| [ ] | Last Stone Weight II | [1049](https://leetcode.com/problems/last-stone-weight-ii/) | [📖](https://programmercarl.com/1049.最后一块石头的重量II.html) |
| [ ] | Target Sum | [494](https://leetcode.com/problems/target-sum/) | [📖](https://programmercarl.com/0494.目标和.html) |
| [ ] | Ones and Zeroes | [474](https://leetcode.com/problems/ones-and-zeroes/) | [📖](https://programmercarl.com/0474.一和零.html) |

### Complete Knapsack — unordered (combinations) — [随想录 完全背包理论](https://programmercarl.com/背包问题理论基础完全背包.html)

> Traverse items outer, capacity **left → right** to allow reuse.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Coin Change II | [518](https://leetcode.com/problems/coin-change-ii/) | [📖](https://programmercarl.com/0518.零钱兑换II.html) |
| [ ] | Coin Change | [322](https://leetcode.com/problems/coin-change/) | [📖](https://programmercarl.com/0322.零钱兑换.html) |
| [ ] | Perfect Squares | [279](https://leetcode.com/problems/perfect-squares/) | [📖](https://programmercarl.com/0279.完全平方数.html) |
| [ ] | Word Break | [139](https://leetcode.com/problems/word-break/) | [📖](https://programmercarl.com/0139.单词拆分.html) |

### Complete Knapsack — ordered (permutations)

> Traverse capacity outer, items inner so orderings count as distinct.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combination Sum IV | [377](https://leetcode.com/problems/combination-sum-iv/) | [📖](https://programmercarl.com/0377.组合总和Ⅳ.html) |

### House Robber — 1D DP with skip constraint

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | House Robber | [198](https://leetcode.com/problems/house-robber/) | [📖](https://programmercarl.com/0198.打家劫舍.html) |
| [ ] | House Robber II (circular array) | [213](https://leetcode.com/problems/house-robber-ii/) | [📖](https://programmercarl.com/0213.打家劫舍II.html) |
| [ ] | House Robber III (tree DP) | [337](https://leetcode.com/problems/house-robber-iii/) | [📖](https://programmercarl.com/0337.打家劫舍III.html) |

### Stock Trading — state machine DP (hold / not-hold)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Best Time to Buy and Sell Stock (1 tx) | [121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [📖](https://programmercarl.com/0121.买卖股票的最佳时机.html) |
| [ ] | Best Time — unlimited transactions | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II（动态规划）.html) |
| [ ] | Best Time — at most 2 transactions | [123](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [📖](https://programmercarl.com/0123.买卖股票的最佳时机III.html) |
| [ ] | Best Time — at most K transactions | [188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [📖](https://programmercarl.com/0188.买卖股票的最佳时机IV.html) |
| [ ] | Best Time with Cooldown | [309](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [📖](https://programmercarl.com/0309.最佳买卖股票时机含冷冻期.html) |
| [ ] | Best Time with Transaction Fee | [714](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [📖](https://programmercarl.com/0714.买卖股票的最佳时机含手续费（动态规划）.html) |

### Subsequence DP — 1D LIS-style

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Increasing Subsequence | [300](https://leetcode.com/problems/longest-increasing-subsequence/) | [📖](https://programmercarl.com/0300.最长上升子序列.html) |
| [ ] | Longest Continuous Increasing Subsequence | [674](https://leetcode.com/problems/longest-continuous-increasing-subsequence/) | [📖](https://programmercarl.com/0674.最长连续递增序列.html) |

### Subsequence DP — 2D two-string comparison

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Common Subsequence | [1143](https://leetcode.com/problems/longest-common-subsequence/) | [📖](https://programmercarl.com/1143.最长公共子序列.html) |
| [ ] | Longest Repeating Subarray | [718](https://leetcode.com/problems/maximum-length-of-repeated-subarray/) | [📖](https://programmercarl.com/0718.最长重复子数组.html) |
| [ ] | Is Subsequence | [392](https://leetcode.com/problems/is-subsequence/) | [📖](https://programmercarl.com/0392.判断子序列.html) |
| [ ] | Distinct Subsequences | [115](https://leetcode.com/problems/distinct-subsequences/) | [📖](https://programmercarl.com/0115.不同的子序列.html) |
| [ ] | Delete Operation for Two Strings | [583](https://leetcode.com/problems/delete-operation-for-two-strings/) | [📖](https://programmercarl.com/0583.两个字符串的删除操作.html) |
| [ ] | Edit Distance | [72](https://leetcode.com/problems/edit-distance/) | [📖](https://programmercarl.com/0072.编辑距离.html) |

### Palindrome DP

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Palindromic Substrings | [647](https://leetcode.com/problems/palindromic-substrings/) | [📖](https://programmercarl.com/0647.回文子串.html) |
| [ ] | Longest Palindromic Subsequence | [516](https://leetcode.com/problems/longest-palindromic-subsequence/) | [📖](https://programmercarl.com/0516.最长回文子序列.html) |

### Interval DP — optimize over all ways to split a subarray

> **Key idea:** `dp[i][j]` = optimal answer for the subarray/subproblem `[i..j]`. Enumerate split point `k` in a third loop. Build solutions from small intervals up to the full range. Order: length 1 → length 2 → … → length n.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Burst Balloons | [312](https://leetcode.com/problems/burst-balloons/) | — |
| [ ] | Minimum Insertion Steps to Make a String Palindrome | [1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) | — |
| [ ] | Strange Printer | [664](https://leetcode.com/problems/strange-printer/) | — |

---

## 11. Monotonic Stack

**Key Concepts:** Maintain a monotonically increasing or decreasing stack to answer "next greater/smaller element" queries in O(n). [随想录 单调栈理论](https://programmercarl.com/0739.每日温度.html)

### Next Greater Element — iterate left→right, pop when current > top

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Daily Temperatures | [739](https://leetcode.com/problems/daily-temperatures/) | [📖](https://programmercarl.com/0739.每日温度.html) |
| [ ] | Next Greater Element I | [496](https://leetcode.com/problems/next-greater-element-i/) | [📖](https://programmercarl.com/0496.下一个更大元素I.html) |

### Circular Array — iterate 2× with modulo index

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Next Greater Element II | [503](https://leetcode.com/problems/next-greater-element-ii/) | [📖](https://programmercarl.com/0503.下一个更大元素II.html) |

### Bounded Area — histogram / rain water trapping

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Trapping Rain Water | [42](https://leetcode.com/problems/trapping-rain-water/) | [📖](https://programmercarl.com/0042.接雨水.html) |
| [ ] | Largest Rectangle in Histogram | [84](https://leetcode.com/problems/largest-rectangle-in-histogram/) | [📖](https://programmercarl.com/0084.柱状图中最大的矩形.html) |

---

## 12. Graph Theory

**Key Concepts:** DFS/BFS on graphs, Union-Find for connectivity, shortest paths, topological sort. [随想录 图论理论](https://programmercarl.com/图论理论基础.html)

### DFS Flood Fill — mark visited in-place from each unvisited cell

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Number of Islands | [200](https://leetcode.com/problems/number-of-islands/) | [📖](https://programmercarl.com/0200.岛屿数量.深搜版.html) |
| [ ] | Max Area of Island | [695](https://leetcode.com/problems/max-area-of-island/) | [📖](https://programmercarl.com/0695.岛屿的最大面积.html) |

### BFS / Reverse DFS — flood from boundary inward

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Pacific Atlantic Water Flow | [417](https://leetcode.com/problems/pacific-atlantic-water-flow/) | [📖](https://programmercarl.com/0417.太平洋大西洋水流问题.html) |
| [ ] | Surrounded Regions | [130](https://leetcode.com/problems/surrounded-regions/) | [📖](https://programmercarl.com/0130.被围绕的区域.html) |

### Multi-Source BFS — simultaneous wavefront expansion

> **Key idea:** Seed the queue with *all* starting nodes at once (layer 0). BFS naturally computes shortest distance from the nearest source to every reachable cell. Used for "rotting spreads", "distance to nearest 0", etc.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Rotting Oranges | [994](https://leetcode.com/problems/rotting-oranges/) | — |
| [ ] | 01 Matrix | [542](https://leetcode.com/problems/01-matrix/) | — |

### BFS on Implicit Graph — state as node, transitions as edges

> **Key idea:** Each unique state (string, tuple, board) is a graph node. BFS finds the minimum number of transformations. Use a `visited` set to avoid re-enqueuing. Common at Google and Amazon.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Word Ladder | [127](https://leetcode.com/problems/word-ladder/) | — |
| [ ] | Minimum Genetic Mutation | [433](https://leetcode.com/problems/minimum-genetic-mutation/) | — |

### DFS on DAG — enumerate all paths

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | All Paths From Source to Target | [797](https://leetcode.com/problems/all-paths-from-source-to-target/) | [📖](https://programmercarl.com/0797.所有可能的路径.html) |

### Union-Find — cycle detection / connectivity — [随想录 并查集理论](https://programmercarl.com/图论-并查集理论基础.html)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Redundant Connection | [684](https://leetcode.com/problems/redundant-connection/) | [📖](https://programmercarl.com/0684.冗余连接.html) |
| [ ] | Redundant Connection II | [685](https://leetcode.com/problems/redundant-connection-ii/) | [📖](https://programmercarl.com/0685.冗余连接II.html) |

### Shortest Path

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Network Delay Time (Dijkstra) | [743](https://leetcode.com/problems/network-delay-time/) | [📖](https://programmercarl.com/0743.网络延迟时间.html) |
| [ ] | Cheapest Flights Within K Stops (Bellman-Ford) | [787](https://leetcode.com/problems/cheapest-flights-within-k-stops/) | [📖](https://programmercarl.com/0787.K站中转内最便宜的航班.html) |
| [ ] | Find the City (Floyd-Warshall) | [1334](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) | [📖](https://programmercarl.com/1334.阈值距离内邻居最少的城市.html) |

### Topological Sort — Kahn's algorithm (BFS + in-degree)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Course Schedule | [207](https://leetcode.com/problems/course-schedule/) | [📖](https://programmercarl.com/0207.课程表.html) |
| [ ] | Course Schedule II | [210](https://leetcode.com/problems/course-schedule-ii/) | [📖](https://programmercarl.com/0210.课程表II.html) |

---

## 13. Trie (Prefix Tree)

**Key Concepts:** Each node stores one character; edges spell out strings. `insert` and `search` are O(L) where L = word length. Two node flags needed: `is_end` (marks a complete word) and children array/map.

> Not covered by 代码随想录. Build the `TrieNode` class first, then implement `insert / search / startsWith` — every other trie problem reuses this skeleton.

### Trie Implementation + Prefix Search

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Implement Trie (Prefix Tree) | [208](https://leetcode.com/problems/implement-trie-prefix-tree/) | — |
| [ ] | Design Add and Search Words Data Structure | [211](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | — |
| [ ] | Replace Words | [648](https://leetcode.com/problems/replace-words/) | — |

### Trie + Backtracking — prune search space with prefix check

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Word Search II | [212](https://leetcode.com/problems/word-search-ii/) | — |

---

## 14. Bit Manipulation

**Key Concepts:** Operate directly on binary representations for O(1) tricks. Core idioms: `n & (n-1)` clears lowest set bit, `n & (-n)` isolates lowest set bit, `a ^ a = 0` cancels duplicates, `a ^ 0 = a` preserves value.

> Not covered by 代码随想录. These problems are short to code but require knowing the idiom cold — drill them as flash cards.

### XOR — duplicate cancellation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Single Number | [136](https://leetcode.com/problems/single-number/) | — |
| [ ] | Single Number II (bit counting mod 3) | [137](https://leetcode.com/problems/single-number-ii/) | — |
| [ ] | Missing Number | [268](https://leetcode.com/problems/missing-number/) | — |

### Bit Counting

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Number of 1 Bits | [191](https://leetcode.com/problems/number-of-1-bits/) | — |
| [ ] | Counting Bits | [338](https://leetcode.com/problems/counting-bits/) | — |
| [ ] | Reverse Bits | [190](https://leetcode.com/problems/reverse-bits/) | — |

### Bit Arithmetic — simulate operations without operators

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Sum of Two Integers (no `+`) | [371](https://leetcode.com/problems/sum-of-two-integers/) | — |

---

## 15. Intervals

**Key Concepts:** Sort by start time for most problems. For minimum rooms / maximum overlap: use a min-heap tracking end times, or a sweep-line difference array. Insert interval requires finding the merge zone by comparing start/end with existing entries.

> Partially covered by 代码随想录 under Greedy (merge/non-overlapping). The problems below focus on the **insertion** and **scheduling** variants that appear frequently at Amazon and Google.

### Insert / Overlap — merge a new interval into a sorted list

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Insert Interval | [57](https://leetcode.com/problems/insert-interval/) | — |
| [ ] | Interval List Intersections | [986](https://leetcode.com/problems/interval-list-intersections/) | — |

### Scheduling — minimum resources / maximum concurrent events

> **Key idea (Meeting Rooms II):** Sort by start time; maintain a min-heap of end times. If the earliest-ending meeting is done before the next one starts, reuse that room — else add a new room.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Meeting Rooms II | [253](https://leetcode.com/problems/meeting-rooms-ii/) | — |
| [ ] | Minimum Number of Platforms Required | [LC variant](https://leetcode.com/problems/minimum-number-of-platforms-required-for-a-railway-station/) | — |

---

## 16. Data Structure Design

**Key Concepts:** Combine multiple base structures to hit all required complexities simultaneously. LRU = doubly-linked list (O(1) move-to-front) + hash map (O(1) lookup). LFU adds a frequency map and a per-frequency LRU list.

> Not covered by 代码随想录. These problems are common at Meta and Databricks staff rounds — the interviewer cares about *which structures you combine and why*, not just the final code.

### Classic Designs

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Min Stack | [155](https://leetcode.com/problems/min-stack/) | — |
| [ ] | LRU Cache | [146](https://leetcode.com/problems/lru-cache/) | — |
| [ ] | LFU Cache | [460](https://leetcode.com/problems/lfu-cache/) | — |
| [ ] | Insert Delete GetRandom O(1) | [380](https://leetcode.com/problems/insert-delete-getrandom-o1/) | — |
| [ ] | Find Median from Data Stream | [295](https://leetcode.com/problems/find-median-from-data-stream/) | — |
| [ ] | All O`one Data Structure | [432](https://leetcode.com/problems/all-oone-data-structure/) | — |

---

## 17. Math

**Key Concepts:** Fast exponentiation (O(log n) via repeated squaring), Sieve of Eratosthenes (O(n log log n) primes), modular arithmetic, integer overflow handling.

> Not covered by 代码随想录. These rarely appear as the *main* topic but show up as sub-problems in harder questions — e.g., `pow` inside a crypto problem, or prime checking inside a combinatorics DP.

### Fast Exponentiation — repeated squaring

> `pow(x, n)`: if `n` is even → `pow(x*x, n/2)`; if odd → `x * pow(x, n-1)`. Handles negative `n` with `1/pow(x, -n)`.

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Pow(x, n) | [50](https://leetcode.com/problems/powx-n/) | — |

### Sieve of Eratosthenes — enumerate primes up to n

> Mark all multiples of each prime starting from p² as composite. Time O(n log log n), space O(n).

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Count Primes | [204](https://leetcode.com/problems/count-primes/) | — |

### Integer / String Arithmetic

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Multiply Strings | [43](https://leetcode.com/problems/multiply-strings/) | — |
| [ ] | Plus One | [66](https://leetcode.com/problems/plus-one/) | — |
| [ ] | Reverse Integer | [7](https://leetcode.com/problems/reverse-integer/) | — |

---

## 18. Segment Tree / Fenwick Tree (Binary Indexed Tree)

**Key Concepts:** Both support point-update + range-query in O(log n). Fenwick Tree is simpler to code (BIT update: `i += i & (-i)`; query: `i -= i & (-i)`). Segment Tree is more flexible (supports range-update, range-min/max). Use when the naive prefix-sum approach breaks due to mutations.

> Not covered by 代码随想录. Appears at Databricks and Google for problems requiring mutable range queries. Build Fenwick Tree from scratch in interview — it's ~10 lines.

### Fenwick Tree (BIT) — mutable prefix sums

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Range Sum Query — Mutable | [307](https://leetcode.com/problems/range-sum-query-mutable/) | — |
| [ ] | Count of Smaller Numbers After Self | [315](https://leetcode.com/problems/count-of-smaller-numbers-after-self/) | — |

### Segment Tree — range queries with lazy propagation

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | My Calendar I | [729](https://leetcode.com/problems/my-calendar-i/) | — |
| [ ] | My Calendar III (max overlap) | [732](https://leetcode.com/problems/my-calendar-iii/) | — |

---

## Study Order

```
Arrays (incl. Binary Search on Answer, Prefix Sum)
→ Linked Lists → Hash Tables → Strings
→ Two Pointers & Sliding Window
→ Stacks & Queues (incl. Heap)
→ Binary Trees → Backtracking
→ Greedy (incl. Intervals) → Dynamic Programming (incl. Interval DP)
→ Monotonic Stack → Trie
→ Graph Theory (incl. Advanced BFS)
→ Bit Manipulation → Data Structure Design
→ Math → Segment Tree / Fenwick Tree
```

> **Priority tiers for staff-level target:**
> - 🔴 **Must master:** Arrays, Sliding Window, Binary Search on Answer, Hash Tables, Trees, Backtracking, DP (knapsack + subsequence), Graph (BFS/DFS + topo), Heap, Trie, Data Structure Design
> - 🟠 **High value:** Prefix Sum, Intervals, Monotonic Stack, Greedy, Bit Manipulation, Union-Find
> - 🟡 **Differentiators at Google/Databricks:** Interval DP, Segment Tree / BIT, Math, Advanced BFS
