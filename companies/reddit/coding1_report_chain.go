package reddit

import (
	"fmt"
	"strings"
)

// ============================================================
// Coding 1 — Report Chain / Manager Hierarchy
// Input: ["A,B,C", "C,D", "B,E"]
// A manages B,C; C manages D; B manages E
// ============================================================

type OrgTree struct {
	children  map[string][]string // parent -> [children]
	parent    map[string]string   // child -> parent
	root      string
}

func BuildOrgTree(input []string) *OrgTree {
	t := &OrgTree{
		children: make(map[string][]string),
		parent:   make(map[string]string),
	}
	allNodes := make(map[string]bool)

	for _, entry := range input {
		parts := strings.Split(entry, ",")
		manager := parts[0]
		allNodes[manager] = true
		for _, report := range parts[1:] {
			t.children[manager] = append(t.children[manager], report)
			t.parent[report] = manager
			allNodes[report] = true
		}
	}

	for node := range allNodes {
		if _, hasParent := t.parent[node]; !hasParent {
			t.root = node
			break
		}
	}
	return t
}

// Part 1 — 打印树
func (t *OrgTree) PrintTree() string {
	var sb strings.Builder
	t.printDFS(t.root, 0, &sb)
	return sb.String()
}

func (t *OrgTree) printDFS(node string, depth int, sb *strings.Builder) {
	sb.WriteString(strings.Repeat("....", depth))
	sb.WriteString(node)
	sb.WriteString("\n")
	for _, child := range t.children[node] {
		t.printDFS(child, depth+1, sb)
	}
}

// Part 2 — Skip-Level Meeting Pairs
// grandparent 和 grandchild 的配对（中间至少隔一层）
func (t *OrgTree) SkipLevelPairs() [][2]string {
	var result [][2]string
	t.collectSkipPairs(t.root, nil, &result)
	return result
}

// ancestors 从近到远存储当前节点的所有祖先
func (t *OrgTree) collectSkipPairs(node string, ancestors []string, result *[][2]string) {
	// ancestors[0] = parent, ancestors[1] = grandparent, ...
	// 跳过 parent(index 0)，从 grandparent 开始配对
	for i := 1; i < len(ancestors); i++ {
		*result = append(*result, [2]string{ancestors[i], node})
	}
	for _, child := range t.children[node] {
		t.collectSkipPairs(child, append([]string{node}, ancestors...), result)
	}
}

// Part 3 — 查询某人的上级链 + 下属树
func (t *OrgTree) PrintChain(person string) string {
	// 向上收集 manager 链
	chain := []string{person}
	cur := person
	for {
		p, ok := t.parent[cur]
		if !ok {
			break
		}
		chain = append(chain, p)
		cur = p
	}

	// chain 现在是 [person, ..., root]，反转后从 root 开始打印
	var sb strings.Builder
	for i := len(chain) - 1; i >= 0; i-- {
		depth := len(chain) - 1 - i
		sb.WriteString(strings.Repeat("....", depth))
		sb.WriteString(chain[i])
		sb.WriteString("\n")
	}

	// 打印 person 的下属子树（person 本身已打印，从 children 开始）
	baseDepth := len(chain) - 1
	for _, child := range t.children[person] {
		t.printDFS(child, baseDepth+1, &sb)
	}
	return sb.String()
}

// Part 4 — Lowest Common Manager (LCA)
func (t *OrgTree) LowestCommonManager(a, b string) string {
	ancestorsA := make(map[string]bool)
	cur := a
	for cur != "" {
		ancestorsA[cur] = true
		cur = t.parent[cur] // map 返回 zero value "" 如果不存在
	}

	cur = b
	for cur != "" {
		if ancestorsA[cur] {
			return cur
		}
		cur = t.parent[cur]
	}
	return ""
}

func DemoCoding1() {
	input := []string{"A,B,C", "C,D", "B,E"}
	tree := BuildOrgTree(input)

	fmt.Println("=== Part 1: Print Tree ===")
	fmt.Print(tree.PrintTree())

	fmt.Println("=== Part 2: Skip-Level Pairs ===")
	for _, pair := range tree.SkipLevelPairs() {
		fmt.Printf("(%s, %s)\n", pair[0], pair[1])
	}

	fmt.Println("=== Part 3: Chain for B ===")
	fmt.Print(tree.PrintChain("B"))

	fmt.Println("=== Part 4: LCA(C, E) ===")
	fmt.Println(tree.LowestCommonManager("C", "E"))
}
