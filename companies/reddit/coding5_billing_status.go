package reddit

import (
	"fmt"
	"sort"
)

// ============================================================
// Coding 5 — Billing Status System
//
// 从交易日志重建每个用户的账单状态
// Part 1: 累加
// Part 2: overwrite 覆盖
// Part 3: undo/redo
// ============================================================

type Transaction struct {
	ID        string
	UserID    int
	Columns   map[string]int // monetary column -> value
	Timestamp int
	Overwrite bool
	UndoLast  bool
	RedoLast  bool
}

type BillingStatus struct {
	Values    map[string]int   // column -> current value
	Applied   []Transaction    // undo 栈: 已应用的常规交易
	Undone    []Transaction    // redo 栈: 已撤销的交易
}

func NewBillingStatus(columns []string) *BillingStatus {
	bs := &BillingStatus{
		Values: make(map[string]int),
	}
	for _, col := range columns {
		bs.Values[col] = 0
	}
	return bs
}

func (bs *BillingStatus) String() string {
	return fmt.Sprintf("BillingStatus(%v)", bs.Values)
}

// applyTransaction 正向应用一笔交易（累加或覆盖）
func (bs *BillingStatus) applyTransaction(tx Transaction) {
	if tx.Overwrite {
		for col, val := range tx.Columns {
			bs.Values[col] = val
		}
	} else {
		for col, val := range tx.Columns {
			bs.Values[col] += val
		}
	}
}

// reverseTransaction 反向撤销一笔交易
func (bs *BillingStatus) reverseTransaction(tx Transaction) {
	if tx.Overwrite {
		// overwrite 的逆操作需要知道之前的值
		// 但我们用重新 replay 的方式来实现 undo，见下方说明
		// 简单实现：对 overwrite 的 undo，减去覆盖值并不正确
		// 所以 Part 3 用 replay 方式更安全
		panic("overwrite undo should use replay approach")
	}
	for col, val := range tx.Columns {
		bs.Values[col] -= val
	}
}

// ============================================================
// Part 1 — 简单累加
// ============================================================

func ProcessTransactionsPart1(
	monetaryColumns []string,
	rawTransactions map[string]map[string]interface{},
) map[int]*BillingStatus {

	sorted := sortTransactions(rawTransactions)
	result := make(map[int]*BillingStatus)

	for _, tx := range sorted {
		bs, ok := result[tx.UserID]
		if !ok {
			bs = NewBillingStatus(monetaryColumns)
			result[tx.UserID] = bs
		}
		for col, val := range tx.Columns {
			bs.Values[col] += val
		}
	}
	return result
}

// ============================================================
// Part 2 — 支持 overwrite
// ============================================================

func ProcessTransactionsPart2(
	monetaryColumns []string,
	rawTransactions map[string]map[string]interface{},
) map[int]*BillingStatus {

	sorted := sortTransactions(rawTransactions)
	result := make(map[int]*BillingStatus)

	for _, tx := range sorted {
		bs, ok := result[tx.UserID]
		if !ok {
			bs = NewBillingStatus(monetaryColumns)
			result[tx.UserID] = bs
		}
		bs.applyTransaction(tx)
	}
	return result
}

// ============================================================
// Part 3 — 支持 undo_last / redo_last
//
// undo: 撤销最近一笔「常规交易」（无 undo/redo 标记的）
// redo: 重新应用最近一笔被撤销的交易
//
// 对于非 overwrite 的交易，undo = 减去对应值
// 对于 overwrite 交易的 undo，需要 replay 所有剩余交易
// 这里用 replay 方式实现，最通用也最安全
// ============================================================

func ProcessTransactionsPart3(
	monetaryColumns []string,
	rawTransactions map[string]map[string]interface{},
) map[int]*BillingStatus {

	sorted := sortTransactions(rawTransactions)

	// 按 user 分组
	userTxs := make(map[int][]Transaction)
	for _, tx := range sorted {
		userTxs[tx.UserID] = append(userTxs[tx.UserID], tx)
	}

	result := make(map[int]*BillingStatus)

	for userID, txs := range userTxs {
		bs := NewBillingStatus(monetaryColumns)
		result[userID] = bs

		var applied []Transaction // 已应用的常规交易栈
		var undone []Transaction  // 已撤销的交易栈

		for _, tx := range txs {
			if tx.UndoLast {
				if len(applied) > 0 {
					// 弹出最近的常规交易
					last := applied[len(applied)-1]
					applied = applied[:len(applied)-1]
					undone = append(undone, last)
					// replay 所有剩余 applied 交易
					bs.replayAll(monetaryColumns, applied)
				}
			} else if tx.RedoLast {
				if len(undone) > 0 {
					// 弹出最近被撤销的交易，重新应用
					last := undone[len(undone)-1]
					undone = undone[:len(undone)-1]
					applied = append(applied, last)
					bs.replayAll(monetaryColumns, applied)
				}
			} else {
				// 常规交易（包括 overwrite）
				applied = append(applied, tx)
				bs.applyTransaction(tx)
			}
		}
	}
	return result
}

// replayAll 从零开始重新应用所有交易（处理 overwrite undo 的安全方式）
func (bs *BillingStatus) replayAll(columns []string, txs []Transaction) {
	for _, col := range columns {
		bs.Values[col] = 0
	}
	for _, tx := range txs {
		bs.applyTransaction(tx)
	}
}

// ============================================================
// 辅助: 解析并按 timestamp 排序
// ============================================================

func sortTransactions(raw map[string]map[string]interface{}) []Transaction {
	txs := make([]Transaction, 0, len(raw))

	for id, data := range raw {
		tx := Transaction{
			ID:      id,
			Columns: make(map[string]int),
		}

		for key, val := range data {
			switch key {
			case "user_id":
				tx.UserID = toInt(val)
			case "transaction_timestamp":
				tx.Timestamp = toInt(val)
			case "overwrite":
				tx.Overwrite = val.(bool)
			case "undo_last":
				tx.UndoLast = val.(bool)
			case "redo_last":
				tx.RedoLast = val.(bool)
			default:
				// 所有其他字段都是 monetary columns
				tx.Columns[key] = toInt(val)
			}
		}
		txs = append(txs, tx)
	}

	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Timestamp < txs[j].Timestamp
	})
	return txs
}

func toInt(v interface{}) int {
	switch n := v.(type) {
	case int:
		return n
	case float64:
		return int(n)
	default:
		return 0
	}
}

// ============================================================
// Demo / 验证
// ============================================================

func DemoCoding5() {
	columns := []string{"ad_delivery_pennies", "payment_pennies"}

	// --- Part 1 ---
	fmt.Println("=== Part 1 ===")
	txs1 := map[string]map[string]interface{}{
		"tx1": {"user_id": 1, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000001},
		"tx2": {"user_id": 1, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000002},
		"tx3": {"user_id": 1, "payment_pennies": 500, "transaction_timestamp": 1500000003},
		"tx4": {"user_id": 1, "ad_delivery_pennies": 1000, "payment_pennies": 500, "transaction_timestamp": 1500000004},
	}
	for uid, bs := range ProcessTransactionsPart1(columns, txs1) {
		fmt.Printf("  User %d: %s\n", uid, bs) // ad=3000, pay=1000
	}

	// --- Part 2 ---
	fmt.Println("=== Part 2 ===")
	txs2 := map[string]map[string]interface{}{
		"t01": {"user_id": 1, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000001, "overwrite": false},
		"t02": {"user_id": 2, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000004},
		"t03": {"user_id": 2, "payment_pennies": 600, "transaction_timestamp": 1500000007, "overwrite": false},
		"t04": {"user_id": 1, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000002, "overwrite": false},
		"t05": {"user_id": 2, "ad_delivery_pennies": 1000, "payment_pennies": 500, "transaction_timestamp": 1500000003, "overwrite": false},
		"t06": {"user_id": 2, "payment_pennies": 2000, "transaction_timestamp": 1500000005, "overwrite": true},
		"t07": {"user_id": 1, "payment_pennies": 500, "transaction_timestamp": 1500000003, "overwrite": false},
		"t08": {"user_id": 1, "ad_delivery_pennies": 1000, "payment_pennies": 500, "transaction_timestamp": 1500000004, "overwrite": true},
		"t09": {"user_id": 2, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000001},
		"t10": {"user_id": 2, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000002},
		"t11": {"user_id": 1, "payment_pennies": 100, "transaction_timestamp": 1500000013},
	}
	for uid, bs := range ProcessTransactionsPart2(columns, txs2) {
		fmt.Printf("  User %d: %s\n", uid, bs)
		// User 1: ad=1000, pay=600
		// User 2: ad=4000, pay=2600
	}

	// --- Part 3 ---
	fmt.Println("=== Part 3 ===")
	txs3 := map[string]map[string]interface{}{
		"t01": {"user_id": 1, "ad_delivery_pennies": 1000, "transaction_timestamp": 1500000001},
		"t02": {"user_id": 1, "undo_last": true, "transaction_timestamp": 1500000002},
		"t03": {"user_id": 1, "payment_pennies": 500, "transaction_timestamp": 1500000003},
		"t04": {"user_id": 1, "ad_delivery_pennies": 1000, "payment_pennies": 500, "transaction_timestamp": 1500000004},
	}
	for uid, bs := range ProcessTransactionsPart3(columns, txs3) {
		fmt.Printf("  User %d: %s\n", uid, bs) // ad=1000, pay=1000
	}
}
