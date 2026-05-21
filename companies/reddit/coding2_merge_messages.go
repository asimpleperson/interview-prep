package reddit

import (
	"container/heap"
	"fmt"
)

// ============================================================
// Coding 2 — Merge Chat Messages (Merge K Sorted Lists 变形)
//
// get_chat_messages(id) 返回该消息前后5条（含自身），按 ID 排序
// merge_messages(ids[]) 合并所有结果，去重，按 ID 排序
// 不能直接用 set 去重 → 利用有序性，相邻去重
// ============================================================

type Message struct {
	ID      int
	Content string
	Version int
}

// --- Min Heap 实现 ---

type heapItem struct {
	msg      Message
	listIdx  int // 来自哪个 list
	innerIdx int // 在该 list 中的 index
}

type minHeap []heapItem

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool   { return h[i].msg.ID < h[j].msg.ID }
func (h minHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{})  { *h = append(*h, x.(heapItem)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// getChatMessages 模拟：返回 id 前后 5 条消息
func getChatMessages(id int) []Message {
	// 面试时这是已有函数，这里只是 stub
	var result []Message
	for i := id - 5; i <= id+5; i++ {
		if i > 0 {
			result = append(result, Message{ID: i, Content: fmt.Sprintf("msg_%d", i)})
		}
	}
	return result
}

// MergeMessages 合并多个 id 对应的消息窗口
// 时间复杂度: O(N log K), N=总消息数, K=ids 数量
func MergeMessages(ids []int) []Message {
	lists := make([][]Message, 0, len(ids))
	for _, id := range ids {
		msgs := getChatMessages(id)
		if len(msgs) > 0 {
			lists = append(lists, msgs)
		}
	}
	if len(lists) == 0 {
		return nil
	}

	h := &minHeap{}
	heap.Init(h)

	// 每个 list 的第一个元素入堆
	for i, lst := range lists {
		heap.Push(h, heapItem{msg: lst[0], listIdx: i, innerIdx: 0})
	}

	var result []Message
	for h.Len() > 0 {
		item := heap.Pop(h).(heapItem)

		// 去重：利用有序性，只跟 result 最后一个比
		if len(result) == 0 || result[len(result)-1].ID != item.msg.ID {
			result = append(result, item.msg)
		}

		// 推入同一个 list 的下一个元素
		nextIdx := item.innerIdx + 1
		if nextIdx < len(lists[item.listIdx]) {
			heap.Push(h, heapItem{
				msg:      lists[item.listIdx][nextIdx],
				listIdx:  item.listIdx,
				innerIdx: nextIdx,
			})
		}
	}
	return result
}

// ============================================================
// Follow-up: Versioning — 消息支持编辑，保存历史版本
// ============================================================

type VersionedMessage struct {
	ID       int
	Content  string
	Version  int
	History  []MessageSnapshot
}

type MessageSnapshot struct {
	Content   string
	Version   int
	Timestamp int64
}

func (m *VersionedMessage) Edit(newContent string, timestamp int64) {
	m.History = append(m.History, MessageSnapshot{
		Content:   m.Content,
		Version:   m.Version,
		Timestamp: timestamp,
	})
	m.Content = newContent
	m.Version++
}

// 合并时用 (ID, Version) 去重，保留最新版本
// 如果同一个 ID 出现多次且 version 不同，取 version 最大的
func MergeVersionedMessages(ids []int) []Message {
	merged := MergeMessages(ids)
	// 实际场景中 getChatMessages 可能返回不同版本
	// 用 map 按 ID 保留最高 version
	best := make(map[int]Message)
	for _, m := range merged {
		if existing, ok := best[m.ID]; !ok || m.Version > existing.Version {
			best[m.ID] = m
		}
	}
	result := make([]Message, 0, len(best))
	for _, m := range merged {
		if b, ok := best[m.ID]; ok {
			result = append(result, b)
			delete(best, m.ID) // 保证顺序且不重复
		}
	}
	return result
}
