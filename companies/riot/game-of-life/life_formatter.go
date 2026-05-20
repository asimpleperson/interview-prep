package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatCell(cell Cell) string {
	return fmt.Sprintf("%d %d", cell.Position.X, cell.Position.Y)
}

// ParseCell parses a line like "2 3" into a Cell.
func ParseCell(line string) (Cell, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return Cell{}, fmt.Errorf("expected 2 coordinates, got %d", len(parts))
	}

	x, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return Cell{}, fmt.Errorf("invalid x coordinate %q: %w", parts[0], err)
	}

	y, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return Cell{}, fmt.Errorf("invalid y coordinate %q: %w", parts[1], err)
	}

	return Cell{Position: Position{X: x, Y: y}, Alive: true}, nil
}
