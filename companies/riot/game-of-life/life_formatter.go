package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatCell converts a single Cell to its Life 1.06 line representation.
// Example output: "2 3"
func FormatCell(cell Cell) string {
	return fmt.Sprintf("%d %d", cell.Position.X, cell.Position.Y)
}

// ParseCell parses a single Life 1.06 coordinate line into a Cell.
// The line should contain exactly two space-separated integers (e.g. "2 3").
// Returns the Cell with Alive set to true.
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

	pos := Position{X: x, Y: y}
	return Cell{Position: pos, Alive: true}, nil
}
