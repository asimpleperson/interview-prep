package reddit

import (
	"fmt"
	"strconv"
	"strings"
)

// ============================================================
// Coding 3 — Moderator List System
//
// Log 格式: "user1,user2 | added | user3 | 123.0"
//   字段: mod_given_new_access | action | mod_take_action | timestamp
//
// Part 1: parse log, can_remove_mod, get_mod_list
// Part 2: 多 community 支持
// Part 3: demote(user) — 将 user 在列表中后移一位
// ============================================================

type ModInfo struct {
	Name      string
	AddedTime float64
	Order     int // demote 操作后用于维护相对顺序
}

type ModeratorSystem struct {
	mods     map[string]*ModInfo // name -> info
	modList  []string            // 有序列表，按加入时间
	orderSeq int
}

func NewModeratorSystem(logs []string) *ModeratorSystem {
	ms := &ModeratorSystem{
		mods: make(map[string]*ModInfo),
	}
	for _, log := range logs {
		ms.applyLog(log)
	}
	return ms
}

func (ms *ModeratorSystem) applyLog(log string) {
	parts := strings.Split(log, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	users := strings.Split(parts[0], ",")
	action := parts[1]
	// actor := parts[2]  // mod_take_action，本题暂不需要
	timestamp, _ := strconv.ParseFloat(parts[3], 64)

	for _, u := range users {
		u = strings.TrimSpace(u)
		switch action {
		case "added":
			if _, exists := ms.mods[u]; !exists {
				ms.mods[u] = &ModInfo{
					Name:      u,
					AddedTime: timestamp,
					Order:     ms.orderSeq,
				}
				ms.orderSeq++
				ms.modList = append(ms.modList, u)
			}
		case "removed":
			if _, exists := ms.mods[u]; exists {
				delete(ms.mods, u)
				ms.removeFromList(u)
			}
		}
	}
}

func (ms *ModeratorSystem) removeFromList(user string) {
	for i, u := range ms.modList {
		if u == user {
			ms.modList = append(ms.modList[:i], ms.modList[i+1:]...)
			return
		}
	}
}

// CanRemoveMod: user1 能否移除 user2
// 规则: user1 必须比 user2 更早成为 moderator
func (ms *ModeratorSystem) CanRemoveMod(user1, user2 string) bool {
	info1, ok1 := ms.mods[user1]
	info2, ok2 := ms.mods[user2]
	if !ok1 || !ok2 {
		return false
	}
	return info1.AddedTime < info2.AddedTime
}

func (ms *ModeratorSystem) GetModList() []string {
	return append([]string{}, ms.modList...)
}

// ============================================================
// Part 2 — Multi-Community
// ============================================================

type MultiCommunityModSystem struct {
	communities map[string]*ModeratorSystem
}

func NewMultiCommunityModSystem() *MultiCommunityModSystem {
	return &MultiCommunityModSystem{
		communities: make(map[string]*ModeratorSystem),
	}
}

// ParseLog: log 格式增加第5个字段 community
// "user1 | added | user3 | 123.0 | community_name"
func (mc *MultiCommunityModSystem) ParseLog(log string) {
	parts := strings.Split(log, "|")
	community := strings.TrimSpace(parts[4])

	if _, ok := mc.communities[community]; !ok {
		mc.communities[community] = NewModeratorSystem(nil)
	}
	// 把前4个字段重新拼接传给单个 community
	singleLog := strings.Join(parts[:4], "|")
	mc.communities[community].applyLog(singleLog)
}

func (mc *MultiCommunityModSystem) CanRemoveMod(community, user1, user2 string) bool {
	sys, ok := mc.communities[community]
	if !ok {
		return false
	}
	return sys.CanRemoveMod(user1, user2)
}

func (mc *MultiCommunityModSystem) GetModList(community string) []string {
	sys, ok := mc.communities[community]
	if !ok {
		return nil
	}
	return sys.GetModList()
}

// ============================================================
// Part 3 — Demote: 将 user 在 mod list 中后移一位
// [u1, u2, u3] demote(u1) → [u2, u1, u3]
// ============================================================

func (ms *ModeratorSystem) Demote(user string) {
	idx := -1
	for i, u := range ms.modList {
		if u == user {
			idx = i
			break
		}
	}
	if idx == -1 || idx == len(ms.modList)-1 {
		return // 不存在或已在末尾
	}

	// 交换位置
	ms.modList[idx], ms.modList[idx+1] = ms.modList[idx+1], ms.modList[idx]

	// 同步交换 Order 和 AddedTime，保持 CanRemoveMod 语义一致
	infoA := ms.mods[ms.modList[idx]]
	infoB := ms.mods[ms.modList[idx+1]]
	infoA.AddedTime, infoB.AddedTime = infoB.AddedTime, infoA.AddedTime
	infoA.Order, infoB.Order = infoB.Order, infoA.Order
}

func DemoCoding3() {
	logs := []string{
		"u1 | added | admin | 100.0",
		"u2 | added | u1 | 200.0",
		"u3 | added | u1 | 300.0",
	}
	ms := NewModeratorSystem(logs)

	fmt.Println("Mod list:", ms.GetModList())
	fmt.Println("u1 can remove u2?", ms.CanRemoveMod("u1", "u2")) // true
	fmt.Println("u2 can remove u1?", ms.CanRemoveMod("u2", "u1")) // false

	ms.Demote("u1")
	fmt.Println("After demote u1:", ms.GetModList()) // [u2, u1, u3]
	fmt.Println("u1 can remove u2?", ms.CanRemoveMod("u1", "u2")) // false now
}
