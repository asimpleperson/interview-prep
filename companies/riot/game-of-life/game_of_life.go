package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Game struct {
	Board      map[Position]Cell
	Generation int
}

func NewGame(scanner *bufio.Scanner) (*Game, error) {
	board, err := parseBoard(scanner)
	if err != nil {
		return nil, err
	}
	return &Game{Board: board, Generation: 0}, nil
}

func (g *Game) Tick() {
	// count live neighbors for every cell adjacent to a live cell
	candidates := make(map[Position]Cell, len(g.Board)*4)

	for pos := range g.Board {
		for _, off := range NeighborOffsets {
			nx, okX := safeAdd(pos.X, off.X)
			ny, okY := safeAdd(pos.Y, off.Y)
			if !okX || !okY {
				continue
			}

			npos := Position{X: nx, Y: ny}
			cell, exists := candidates[npos]
			if !exists {
				cell = Cell{Position: npos, Alive: false, LiveNeighbors: 0}
				if _, alive := g.Board[npos]; alive {
					cell.Alive = true
				}
			}
			cell.LiveNeighbors++
			candidates[npos] = cell
		}
	}

	// apply rules: birth at 3, survival at 2 or 3
	next := make(map[Position]Cell, len(g.Board))
	for pos, cell := range candidates {
		if cell.LiveNeighbors == 3 || (cell.LiveNeighbors == 2 && cell.Alive) {
			cell.Alive = true
			next[pos] = cell
		}
	}

	g.Board = next
	g.Generation++
}

func safeAdd(a, b int64) (int64, bool) {
	r := a + b
	if (b > 0 && r < a) || (b < 0 && r > a) {
		return 0, false
	}
	return r, true
}

func (g *Game) BoardStatusLife106() string {
	positions := make([]Position, 0, len(g.Board))
	for pos := range g.Board {
		positions = append(positions, pos)
	}

	sort.Slice(positions, func(i, j int) bool {
		if positions[i].X != positions[j].X {
			return positions[i].X < positions[j].X
		}
		return positions[i].Y < positions[j].Y
	})

	var sb strings.Builder
	sb.WriteString("#Life 1.06\n")
	for _, pos := range positions {
		sb.WriteString(FormatCell(g.Board[pos]))
		sb.WriteByte('\n')
	}
	return sb.String()
}

func parseBoard(scanner *bufio.Scanner) (map[Position]Cell, error) {
	board := make(map[Position]Cell)
	lineNum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		// handle BOM on windows
		line = strings.TrimPrefix(line, "\xef\xbb\xbf")

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		cell, err := ParseCell(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum, err)
		}
		board[cell.Position] = cell
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	return board, nil
}
