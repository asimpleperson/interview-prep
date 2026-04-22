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

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Search | [704](https://leetcode.com/problems/binary-search/) | [📖](https://programmercarl.com/0704.二分查找.html) |
| [ ] | Remove Element | [27](https://leetcode.com/problems/remove-element/) | [📖](https://programmercarl.com/0027.移除元素.html) |
| [ ] | Squares of a Sorted Array | [977](https://leetcode.com/problems/squares-of-a-sorted-array/) | [📖](https://programmercarl.com/0977.有序数组的平方.html) |
| [ ] | Minimum Size Subarray Sum | [209](https://leetcode.com/problems/minimum-size-subarray-sum/) | [📖](https://programmercarl.com/0209.长度最小的子数组.html) |
| [ ] | Spiral Matrix II | [59](https://leetcode.com/problems/spiral-matrix-ii/) | [📖](https://programmercarl.com/0059.螺旋矩阵II.html) |

---

## 2. Linked Lists

**Key Concepts:** Pointer manipulation, dummy head node trick, fast/slow pointers, reversal, cycle detection.

**Patterns:**
- [ ] Dummy head node to simplify edge cases
- [ ] Two-pointer (fast/slow) for cycle & nth-from-end
- [ ] In-place reversal
- [ ] Intersection detection via pointer reset

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Remove Linked List Elements | [203](https://leetcode.com/problems/remove-linked-list-elements/) | [📖](https://programmercarl.com/0203.移除链表元素.html) |
| [ ] | Design Linked List | [707](https://leetcode.com/problems/design-linked-list/) | [📖](https://programmercarl.com/0707.设计链表.html) |
| [ ] | Reverse Linked List | [206](https://leetcode.com/problems/reverse-linked-list/) | [📖](https://programmercarl.com/0206.翻转链表.html) |
| [ ] | Swap Nodes in Pairs | [24](https://leetcode.com/problems/swap-nodes-in-pairs/) | [📖](https://programmercarl.com/0024.两两交换链表中的节点.html) |
| [ ] | Remove Nth Node From End | [19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [📖](https://programmercarl.com/0019.删除链表的倒数第N个节点.html) |
| [ ] | Intersection of Two Linked Lists | [160](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [📖](https://programmercarl.com/面试题02.07.链表相交.html) |
| [ ] | Linked List Cycle II | [142](https://leetcode.com/problems/linked-list-cycle-ii/) | [📖](https://programmercarl.com/0142.环形链表II.html) |

---

## 3. Hash Tables

**Key Concepts:** O(1) lookup, collision resolution, choosing the right structure (array vs. set vs. map).

**Patterns:**
- [ ] Frequency counting with arrays (fixed alphabet)
- [ ] Set for existence/intersection
- [ ] Map for complement lookup (Two Sum pattern)
- [ ] Multi-array grouping (Four Sum II)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Valid Anagram | [242](https://leetcode.com/problems/valid-anagram/) | [📖](https://programmercarl.com/0242.有效的字母异位词.html) |
| [ ] | Intersection of Two Arrays | [349](https://leetcode.com/problems/intersection-of-two-arrays/) | [📖](https://programmercarl.com/0349.两个数组的交集.html) |
| [ ] | Happy Number | [202](https://leetcode.com/problems/happy-number/) | [📖](https://programmercarl.com/0202.快乐数.html) |
| [ ] | Two Sum | [1](https://leetcode.com/problems/two-sum/) | [📖](https://programmercarl.com/0001.两数之和.html) |
| [ ] | Four Sum II | [454](https://leetcode.com/problems/4sum-ii/) | [📖](https://programmercarl.com/0454.四数相加II.html) |
| [ ] | Ransom Note | [383](https://leetcode.com/problems/ransom-note/) | [📖](https://programmercarl.com/0383.赎金信.html) |
| [ ] | Three Sum | [15](https://leetcode.com/problems/3sum/) | [📖](https://programmercarl.com/0015.三数之和.html) |
| [ ] | Four Sum | [18](https://leetcode.com/problems/4sum/) | [📖](https://programmercarl.com/0018.四数之和.html) |

---

## 4. Strings

**Key Concepts:** In-place reversal, KMP pattern matching, prefix table (failure function), repeated substring detection.

**Patterns:**
- [ ] Two-pointer reversal (whole string, partial, word-by-word)
- [ ] KMP for O(n) substring search — build the prefix table — [随想录 KMP理论](https://programmercarl.com/0028.实现strStr.html)
- [ ] Sliding window for substring problems

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Reverse String | [344](https://leetcode.com/problems/reverse-string/) | [📖](https://programmercarl.com/0344.反转字符串.html) |
| [ ] | Reverse String II | [541](https://leetcode.com/problems/reverse-string-ii/) | [📖](https://programmercarl.com/0541.反转字符串II.html) |
| [ ] | Reverse Words in a String | [151](https://leetcode.com/problems/reverse-words-in-a-string/) | [📖](https://programmercarl.com/0151.翻转字符串里的单词.html) |
| [ ] | Find the Index of the First Occurrence (strStr) | [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | [📖](https://programmercarl.com/0028.实现strStr.html) |
| [ ] | Repeated Substring Pattern | [459](https://leetcode.com/problems/repeated-substring-pattern/) | [📖](https://programmercarl.com/0459.重复的子字符串.html) |

---

## 5. Two Pointers

**Key Concepts:** Reduces O(n²) brute force to O(n); works on sorted arrays and linked lists.

**Patterns:**
- [ ] Fast/slow pointer (cycle detection, midpoint)
- [ ] Left/right pointer (sorted array problems, kSum)
- [ ] Same-direction pointers (sliding window, in-place overwrite)

> Two-pointer techniques appear across Arrays, Strings, and Linked Lists — treat this as a cross-cutting pattern, not an isolated topic. [随想录 双指针总结](https://programmercarl.com/双指针总结.html)

---

## 6. Stacks & Queues

**Key Concepts:** LIFO vs. FIFO, monotonic deque for sliding window max, priority queue for top-K.

**Patterns:**
- [ ] Stack for bracket matching and expression evaluation
- [ ] Stack for adjacent duplicate removal
- [ ] Monotonic deque for sliding window maximum
- [ ] Heap/priority queue for top-K frequency

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Implement Queue using Stacks | [232](https://leetcode.com/problems/implement-queue-using-stacks/) | [📖](https://programmercarl.com/0232.用栈实现队列.html) |
| [ ] | Implement Stack using Queues | [225](https://leetcode.com/problems/implement-stack-using-queues/) | [📖](https://programmercarl.com/0225.用队列实现栈.html) |
| [ ] | Valid Parentheses | [20](https://leetcode.com/problems/valid-parentheses/) | [📖](https://programmercarl.com/0020.有效的括号.html) |
| [ ] | Remove All Adjacent Duplicates In String | [1047](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) | [📖](https://programmercarl.com/1047.删除字符串中的所有相邻重复项.html) |
| [ ] | Evaluate Reverse Polish Notation | [150](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | [📖](https://programmercarl.com/0150.逆波兰表达式求值.html) |
| [ ] | Sliding Window Maximum | [239](https://leetcode.com/problems/sliding-window-maximum/) | [📖](https://programmercarl.com/0239.滑动窗口最大值.html) |
| [ ] | Top K Frequent Elements | [347](https://leetcode.com/problems/top-k-frequent-elements/) | [📖](https://programmercarl.com/0347.前K个高频元素.html) |

---

## 7. Binary Trees

**Key Concepts:** DFS (pre/in/post-order) and BFS (level-order), recursive vs. iterative traversal, BST properties.

**Patterns:**
- [ ] Recursive DFS for tree properties (depth, balance, paths) — [随想录 二叉树理论](https://programmercarl.com/二叉树理论基础篇.html)
- [ ] BFS/level-order with a queue
- [ ] BST in-order traversal gives sorted sequence
- [ ] Lowest Common Ancestor (LCA) using recursion return values
- [ ] Tree construction from traversal arrays

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Binary Tree Level Order Traversal | [102](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [📖](https://programmercarl.com/0102.二叉树的层序遍历.html) |
| [ ] | Invert Binary Tree | [226](https://leetcode.com/problems/invert-binary-tree/) | [📖](https://programmercarl.com/0226.翻转二叉树.html) |
| [ ] | Symmetric Tree | [101](https://leetcode.com/problems/symmetric-tree/) | [📖](https://programmercarl.com/0101.对称二叉树.html) |
| [ ] | Maximum Depth | [104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [📖](https://programmercarl.com/0104.二叉树的最大深度.html) |
| [ ] | Minimum Depth | [111](https://leetcode.com/problems/minimum-depth-of-binary-tree/) | [📖](https://programmercarl.com/0111.二叉树的最小深度.html) |
| [ ] | Count Complete Tree Nodes | [222](https://leetcode.com/problems/count-complete-tree-nodes/) | [📖](https://programmercarl.com/0222.完全二叉树的节点个数.html) |
| [ ] | Balanced Binary Tree | [110](https://leetcode.com/problems/balanced-binary-tree/) | [📖](https://programmercarl.com/0110.平衡二叉树.html) |
| [ ] | All Root-to-Leaf Paths | [257](https://leetcode.com/problems/binary-tree-paths/) | [📖](https://programmercarl.com/0257.二叉树的所有路径.html) |
| [ ] | Path Sum | [112](https://leetcode.com/problems/path-sum/) | [📖](https://programmercarl.com/0112.路径总和.html) |
| [ ] | Construct from Inorder & Postorder | [106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | [📖](https://programmercarl.com/0106.从中序与后序遍历序列构造二叉树.html) |
| [ ] | Maximum Binary Tree | [654](https://leetcode.com/problems/maximum-binary-tree/) | [📖](https://programmercarl.com/0654.最大二叉树.html) |
| [ ] | Merge Two Binary Trees | [617](https://leetcode.com/problems/merge-two-binary-trees/) | [📖](https://programmercarl.com/0617.合并二叉树.html) |
| [ ] | Lowest Common Ancestor (general) | [236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | [📖](https://programmercarl.com/0236.二叉树的最近公共祖先.html) |
| [ ] | Lowest Common Ancestor (BST) | [235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | [📖](https://programmercarl.com/0235.二叉搜索树的最近公共祖先.html) |
| [ ] | Validate BST | [98](https://leetcode.com/problems/validate-binary-search-tree/) | [📖](https://programmercarl.com/0098.验证二叉搜索树.html) |
| [ ] | Search in BST | [700](https://leetcode.com/problems/search-in-a-binary-search-tree/) | [📖](https://programmercarl.com/0700.二叉搜索树中的搜索.html) |
| [ ] | Insert into BST | [701](https://leetcode.com/problems/insert-into-a-binary-search-tree/) | [📖](https://programmercarl.com/0701.二叉搜索树中的插入操作.html) |
| [ ] | Delete Node in BST | [450](https://leetcode.com/problems/delete-node-in-a-bst/) | [📖](https://programmercarl.com/0450.删除二叉搜索树中的节点.html) |
| [ ] | Trim BST | [669](https://leetcode.com/problems/trim-a-binary-search-tree/) | [📖](https://programmercarl.com/0669.修剪二叉搜索树.html) |
| [ ] | Convert Sorted Array to BST | [108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) | [📖](https://programmercarl.com/0108.将有序数组转换为二叉搜索树.html) |
| [ ] | Convert BST to Greater Tree | [538](https://leetcode.com/problems/convert-bst-to-greater-tree/) | [📖](https://programmercarl.com/0538.把二叉搜索树转换为累加树.html) |

---

## 8. Backtracking

**Key Concepts:** Exhaustive search on a decision tree; prune branches early. Framework: **choose → explore → unchoose**. [随想录 回溯理论](https://programmercarl.com/回溯算法理论基础.html)

**Patterns:**
- [ ] Combinations (order doesn't matter, no reuse)
- [ ] Combinations with reuse (Combination Sum)
- [ ] Deduplication: sort + skip duplicate at same tree level
- [ ] Permutations (order matters, use `used[]` array)
- [ ] Subsets = combinations of all sizes
- [ ] Partitioning (palindrome check, IP address segments)
- [ ] Board problems (N-Queens, Sudoku — 2D constraint propagation)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Combinations | [77](https://leetcode.com/problems/combinations/) | [📖](https://programmercarl.com/0077.组合.html) |
| [ ] | Combination Sum III | [216](https://leetcode.com/problems/combination-sum-iii/) | [📖](https://programmercarl.com/0216.组合总和III.html) |
| [ ] | Letter Combinations of a Phone Number | [17](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [📖](https://programmercarl.com/0017.电话号码的字母组合.html) |
| [ ] | Combination Sum | [39](https://leetcode.com/problems/combination-sum/) | [📖](https://programmercarl.com/0039.组合总和.html) |
| [ ] | Combination Sum II | [40](https://leetcode.com/problems/combination-sum-ii/) | [📖](https://programmercarl.com/0040.组合总和II.html) |
| [ ] | Palindrome Partitioning | [131](https://leetcode.com/problems/palindrome-partitioning/) | [📖](https://programmercarl.com/0131.分割回文串.html) |
| [ ] | Restore IP Addresses | [93](https://leetcode.com/problems/restore-ip-addresses/) | [📖](https://programmercarl.com/0093.复原IP地址.html) |
| [ ] | Subsets | [78](https://leetcode.com/problems/subsets/) | [📖](https://programmercarl.com/0078.子集.html) |
| [ ] | Subsets II | [90](https://leetcode.com/problems/subsets-ii/) | [📖](https://programmercarl.com/0090.子集II.html) |
| [ ] | Non-decreasing Subsequences | [491](https://leetcode.com/problems/non-decreasing-subsequences/) | [📖](https://programmercarl.com/0491.递增子序列.html) |
| [ ] | Permutations | [46](https://leetcode.com/problems/permutations/) | [📖](https://programmercarl.com/0046.全排列.html) |
| [ ] | Permutations II | [47](https://leetcode.com/problems/permutations-ii/) | [📖](https://programmercarl.com/0047.全排列II.html) |
| [ ] | N-Queens | [51](https://leetcode.com/problems/n-queens/) | [📖](https://programmercarl.com/0051.N皇后.html) |
| [ ] | Sudoku Solver | [37](https://leetcode.com/problems/sudoku-solver/) | [📖](https://programmercarl.com/0037.解数独.html) |

---

## 9. Greedy Algorithms

**Key Concepts:** Make the locally optimal choice at each step; no rollback. Correctness requires proof — not just intuition. [随想录 贪心理论](https://programmercarl.com/贪心算法理论基础.html)

**Patterns:**
- [ ] Interval scheduling (sort by end time, count non-overlapping)
- [ ] Two-pass greedy (candy: left→right then right→left)
- [ ] Jump game (track max reachable index)
- [ ] Stock problems (accumulate positive daily diffs)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Assign Cookies | [455](https://leetcode.com/problems/assign-cookies/) | [📖](https://programmercarl.com/0455.分发饼干.html) |
| [ ] | Wiggle Subsequence | [376](https://leetcode.com/problems/wiggle-subsequence/) | [📖](https://programmercarl.com/0376.摆动序列.html) |
| [ ] | Maximum Subarray | [53](https://leetcode.com/problems/maximum-subarray/) | [📖](https://programmercarl.com/0053.最大子序和.html) |
| [ ] | Best Time to Buy and Sell Stock II | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II.html) |
| [ ] | Jump Game | [55](https://leetcode.com/problems/jump-game/) | [📖](https://programmercarl.com/0055.跳跃游戏.html) |
| [ ] | Jump Game II | [45](https://leetcode.com/problems/jump-game-ii/) | [📖](https://programmercarl.com/0045.跳跃游戏II.html) |
| [ ] | Maximize Sum After K Negations | [1005](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/) | [📖](https://programmercarl.com/1005.K次取反后最大化的数组和.html) |
| [ ] | Gas Station | [134](https://leetcode.com/problems/gas-station/) | [📖](https://programmercarl.com/0134.加油站.html) |
| [ ] | Candy | [135](https://leetcode.com/problems/candy/) | [📖](https://programmercarl.com/0135.分发糖果.html) |
| [ ] | Lemonade Change | [860](https://leetcode.com/problems/lemonade-change/) | [📖](https://programmercarl.com/0860.柠檬水找零.html) |
| [ ] | Queue Reconstruction by Height | [406](https://leetcode.com/problems/queue-reconstruction-by-height/) | [📖](https://programmercarl.com/0406.根据身高重建队列.html) |
| [ ] | Minimum Arrows to Burst Balloons | [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | [📖](https://programmercarl.com/0452.用最少数量的箭引爆气球.html) |
| [ ] | Non-overlapping Intervals | [435](https://leetcode.com/problems/non-overlapping-intervals/) | [📖](https://programmercarl.com/0435.无重叠区间.html) |
| [ ] | Partition Labels | [763](https://leetcode.com/problems/partition-labels/) | [📖](https://programmercarl.com/0763.划分字母区间.html) |
| [ ] | Merge Intervals | [56](https://leetcode.com/problems/merge-intervals/) | [📖](https://programmercarl.com/0056.合并区间.html) |
| [ ] | Monotone Increasing Digits | [738](https://leetcode.com/problems/monotone-increasing-digits/) | [📖](https://programmercarl.com/0738.单调递增的数字.html) |
| [ ] | Binary Tree Cameras | [968](https://leetcode.com/problems/binary-tree-cameras/) | [📖](https://programmercarl.com/0968.监控二叉树.html) |

---

## 10. Dynamic Programming

**Key Concepts:** Optimal substructure + overlapping subproblems. Define state → write recurrence → determine traversal order → initialize base cases. [随想录 DP理论](https://programmercarl.com/动态规划理论基础.html)

### Basic DP

**Patterns:**
- [ ] 1D DP (Fibonacci-style, climbing stairs)
- [ ] 2D DP (grid paths)
- [ ] Math-based DP (integer break, unique BST count)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Fibonacci Number | [509](https://leetcode.com/problems/fibonacci-number/) | [📖](https://programmercarl.com/0509.斐波那契数.html) |
| [ ] | Climbing Stairs | [70](https://leetcode.com/problems/climbing-stairs/) | [📖](https://programmercarl.com/0070.爬楼梯.html) |
| [ ] | Min Cost Climbing Stairs | [746](https://leetcode.com/problems/min-cost-climbing-stairs/) | [📖](https://programmercarl.com/0746.使用最小花费爬楼梯.html) |
| [ ] | Unique Paths | [62](https://leetcode.com/problems/unique-paths/) | [📖](https://programmercarl.com/0062.不同路径.html) |
| [ ] | Unique Paths II | [63](https://leetcode.com/problems/unique-paths-ii/) | [📖](https://programmercarl.com/0063.不同路径II.html) |
| [ ] | Integer Break | [343](https://leetcode.com/problems/integer-break/) | [📖](https://programmercarl.com/0343.整数拆分.html) |
| [ ] | Unique BSTs | [96](https://leetcode.com/problems/unique-binary-search-trees/) | [📖](https://programmercarl.com/0096.不同的二叉搜索树.html) |

### 0/1 Knapsack

**Patterns:**
- [ ] 2D dp table → compress to 1D rolling array — [随想录 0/1背包理论](https://programmercarl.com/背包理论基础01背包-1.html)
- [ ] Iterate items outer, capacity inner (right to left for 0/1)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Partition Equal Subset Sum | [416](https://leetcode.com/problems/partition-equal-subset-sum/) | [📖](https://programmercarl.com/0416.分割等和子集.html) |
| [ ] | Last Stone Weight II | [1049](https://leetcode.com/problems/last-stone-weight-ii/) | [📖](https://programmercarl.com/1049.最后一块石头的重量II.html) |
| [ ] | Target Sum | [494](https://leetcode.com/problems/target-sum/) | [📖](https://programmercarl.com/0494.目标和.html) |
| [ ] | Ones and Zeroes | [474](https://leetcode.com/problems/ones-and-zeroes/) | [📖](https://programmercarl.com/0474.一和零.html) |

### Complete Knapsack

**Patterns:**
- [ ] Iterate capacity inner left to right (items can repeat) — [随想录 完全背包理论](https://programmercarl.com/背包问题理论基础完全背包.html)
- [ ] Order matters (permutations) → iterate capacity outer

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Coin Change II | [518](https://leetcode.com/problems/coin-change-ii/) | [📖](https://programmercarl.com/0518.零钱兑换II.html) |
| [ ] | Combination Sum IV | [377](https://leetcode.com/problems/combination-sum-iv/) | [📖](https://programmercarl.com/0377.组合总和Ⅳ.html) |
| [ ] | Coin Change | [322](https://leetcode.com/problems/coin-change/) | [📖](https://programmercarl.com/0322.零钱兑换.html) |
| [ ] | Perfect Squares | [279](https://leetcode.com/problems/perfect-squares/) | [📖](https://programmercarl.com/0279.完全平方数.html) |
| [ ] | Word Break | [139](https://leetcode.com/problems/word-break/) | [📖](https://programmercarl.com/0139.单词拆分.html) |

### House Robber Series

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | House Robber | [198](https://leetcode.com/problems/house-robber/) | [📖](https://programmercarl.com/0198.打家劫舍.html) |
| [ ] | House Robber II | [213](https://leetcode.com/problems/house-robber-ii/) | [📖](https://programmercarl.com/0213.打家劫舍II.html) |
| [ ] | House Robber III (Tree DP) | [337](https://leetcode.com/problems/house-robber-iii/) | [📖](https://programmercarl.com/0337.打家劫舍III.html) |

### Stock Trading

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Best Time to Buy and Sell Stock | [121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [📖](https://programmercarl.com/0121.买卖股票的最佳时机.html) |
| [ ] | Best Time — Multiple Transactions | [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [📖](https://programmercarl.com/0122.买卖股票的最佳时机II（动态规划）.html) |
| [ ] | Best Time — At Most 2 Transactions | [123](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [📖](https://programmercarl.com/0123.买卖股票的最佳时机III.html) |
| [ ] | Best Time — At Most K Transactions | [188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [📖](https://programmercarl.com/0188.买卖股票的最佳时机IV.html) |
| [ ] | Best Time with Cooldown | [309](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [📖](https://programmercarl.com/0309.最佳买卖股票时机含冷冻期.html) |
| [ ] | Best Time with Transaction Fee | [714](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [📖](https://programmercarl.com/0714.买卖股票的最佳时机含手续费（动态规划）.html) |

### Subsequence / String DP

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Longest Increasing Subsequence | [300](https://leetcode.com/problems/longest-increasing-subsequence/) | [📖](https://programmercarl.com/0300.最长上升子序列.html) |
| [ ] | Longest Continuous Increasing Subsequence | [674](https://leetcode.com/problems/longest-continuous-increasing-subsequence/) | [📖](https://programmercarl.com/0674.最长连续递增序列.html) |
| [ ] | Longest Common Subsequence | [1143](https://leetcode.com/problems/longest-common-subsequence/) | [📖](https://programmercarl.com/1143.最长公共子序列.html) |
| [ ] | Longest Repeating Subarray | [718](https://leetcode.com/problems/maximum-length-of-repeated-subarray/) | [📖](https://programmercarl.com/0718.最长重复子数组.html) |
| [ ] | Is Subsequence | [392](https://leetcode.com/problems/is-subsequence/) | [📖](https://programmercarl.com/0392.判断子序列.html) |
| [ ] | Distinct Subsequences | [115](https://leetcode.com/problems/distinct-subsequences/) | [📖](https://programmercarl.com/0115.不同的子序列.html) |
| [ ] | Delete Operation for Two Strings | [583](https://leetcode.com/problems/delete-operation-for-two-strings/) | [📖](https://programmercarl.com/0583.两个字符串的删除操作.html) |
| [ ] | Edit Distance | [72](https://leetcode.com/problems/edit-distance/) | [📖](https://programmercarl.com/0072.编辑距离.html) |
| [ ] | Palindromic Substrings | [647](https://leetcode.com/problems/palindromic-substrings/) | [📖](https://programmercarl.com/0647.回文子串.html) |
| [ ] | Longest Palindromic Subsequence | [516](https://leetcode.com/problems/longest-palindromic-subsequence/) | [📖](https://programmercarl.com/0516.最长回文子序列.html) |

---

## 11. Monotonic Stack

**Key Concepts:** Maintain a stack in monotonically increasing or decreasing order to answer "next greater/smaller element" queries in O(n). [随想录 单调栈理论](https://programmercarl.com/0739.每日温度.html)

**Patterns:**
- [ ] Next greater element (iterate left→right, pop when current > top)
- [ ] Circular arrays (iterate 2× with modulo index)
- [ ] Histogram / rain water trapping (area under bounded containers)

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Daily Temperatures | [739](https://leetcode.com/problems/daily-temperatures/) | [📖](https://programmercarl.com/0739.每日温度.html) |
| [ ] | Next Greater Element I | [496](https://leetcode.com/problems/next-greater-element-i/) | [📖](https://programmercarl.com/0496.下一个更大元素I.html) |
| [ ] | Next Greater Element II | [503](https://leetcode.com/problems/next-greater-element-ii/) | [📖](https://programmercarl.com/0503.下一个更大元素II.html) |
| [ ] | Trapping Rain Water | [42](https://leetcode.com/problems/trapping-rain-water/) | [📖](https://programmercarl.com/0042.接雨水.html) |
| [ ] | Largest Rectangle in Histogram | [84](https://leetcode.com/problems/largest-rectangle-in-histogram/) | [📖](https://programmercarl.com/0084.柱状图中最大的矩形.html) |

---

## 12. Graph Theory

**Key Concepts:** DFS/BFS on graphs, Union-Find for connectivity, shortest paths (Dijkstra/Bellman-Ford/Floyd), minimum spanning tree (Prim/Kruskal), topological sort. [随想录 图论理论](https://programmercarl.com/图论理论基础.html)

### DFS / BFS / Island Problems

**Patterns:**
- [ ] Treat 2D grid as a graph; DFS/BFS from each unvisited land cell
- [ ] Mark visited in-place (flood fill) or with a visited set

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Number of Islands | [200](https://leetcode.com/problems/number-of-islands/) | [📖](https://programmercarl.com/0200.岛屿数量.深搜版.html) |
| [ ] | Max Area of Island | [695](https://leetcode.com/problems/max-area-of-island/) | [📖](https://programmercarl.com/0695.岛屿的最大面积.html) |
| [ ] | Pacific Atlantic Water Flow | [417](https://leetcode.com/problems/pacific-atlantic-water-flow/) | [📖](https://programmercarl.com/0417.太平洋大西洋水流问题.html) |
| [ ] | Surrounded Regions | [130](https://leetcode.com/problems/surrounded-regions/) | [📖](https://programmercarl.com/0130.被围绕的区域.html) |
| [ ] | All Paths From Source to Target | [797](https://leetcode.com/problems/all-paths-from-source-to-target/) | [📖](https://programmercarl.com/0797.所有可能的路径.html) |

### Union-Find

**Patterns:**
- [ ] Path compression + union by rank for near-O(1) operations — [随想录 并查集理论](https://programmercarl.com/图论-并查集理论基础.html)
- [ ] Detect cycle in undirected graph

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

### Topological Sort

| Done | Problem | LC | 随想录 |
|------|---------|-----|--------|
| [ ] | Course Schedule | [207](https://leetcode.com/problems/course-schedule/) | [📖](https://programmercarl.com/0207.课程表.html) |
| [ ] | Course Schedule II | [210](https://leetcode.com/problems/course-schedule-ii/) | [📖](https://programmercarl.com/0210.课程表II.html) |

---

## Study Order

```
Arrays → Linked Lists → Hash Tables → Strings
       → Two Pointers (cross-cutting)
→ Stacks & Queues → Binary Trees → Backtracking
→ Greedy → Dynamic Programming → Monotonic Stack → Graph Theory
```

> **Note:** DP is the heaviest section — budget extra time for the knapsack framework and subsequence problems. Binary Trees is the prerequisite for both Backtracking (recursive tree thinking) and Graph Theory (DFS/BFS).
