package reddit

import (
	"errors"
	"fmt"
)

// ============================================================
// Coding 4 — Tennis Game / Score System
//
// 规则:
// - 先到 4 分且领先 2 分胜利
// - 双方都 >= 3 且平分 → 重置为 3:3 (Deuce)
// - 比赛结束后不能继续加分
//
// Follow-up: Human readable score (love, deuce, advantage)
// ============================================================

type TennisGame struct {
	scores [2]int
	winner int // -1 = 进行中, 0 or 1 = 赢家
}

func NewTennisGame() *TennisGame {
	return &TennisGame{winner: -1}
}

func (g *TennisGame) AddScore(player int) error {
	if player != 0 && player != 1 {
		return errors.New("invalid player")
	}
	if g.winner != -1 {
		return errors.New("game already over")
	}

	g.scores[player]++

	// Deuce 重置: 双方都 >= 3 且平分
	if g.scores[0] >= 3 && g.scores[1] >= 3 && g.scores[0] == g.scores[1] {
		g.scores = [2]int{3, 3}
	}

	// 胜利判定: >= 4 且领先 2
	if g.scores[player] >= 4 && g.scores[player]-g.scores[1-player] >= 2 {
		g.winner = player
	}

	return nil
}

func (g *TennisGame) GetScore() [2]int {
	return g.scores
}

func (g *TennisGame) GetResult() int {
	return g.winner // -1 = in progress
}

// ============================================================
// Follow-up: Human Readable Score
// ============================================================

var scoreNames = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

func (g *TennisGame) GetHumanScore() string {
	if g.winner != -1 {
		return fmt.Sprintf("Player %d wins", g.winner)
	}

	a, b := g.scores[0], g.scores[1]

	// Deuce 区间
	if a >= 3 && b >= 3 {
		if a == b {
			return "Deuce"
		}
		if a > b {
			return "Advantage Player 0"
		}
		return "Advantage Player 1"
	}

	return fmt.Sprintf("%s - %s", scoreNames[a], scoreNames[b])
}

func DemoCoding4() {
	g := NewTennisGame()

	// 模拟一场到 Deuce 的比赛
	points := []int{0, 1, 0, 1, 0, 1} // 3:3
	for _, p := range points {
		g.AddScore(p)
		fmt.Printf("Score: %v  Human: %s\n", g.GetScore(), g.GetHumanScore())
	}

	// Deuce 后 player 0 拿 advantage
	g.AddScore(0)
	fmt.Printf("Score: %v  Human: %s\n", g.GetScore(), g.GetHumanScore())

	// Player 1 追平，回到 Deuce (4:4 → 3:3)
	g.AddScore(1)
	fmt.Printf("Score: %v  Human: %s\n", g.GetScore(), g.GetHumanScore())

	// Player 0 连拿两分赢
	g.AddScore(0)
	fmt.Printf("Score: %v  Human: %s\n", g.GetScore(), g.GetHumanScore())
	g.AddScore(0)
	fmt.Printf("Score: %v  Human: %s  Result: %d\n", g.GetScore(), g.GetHumanScore(), g.GetResult())

	// 尝试在结束后加分
	err := g.AddScore(0)
	fmt.Println("Add after game over:", err)
}
