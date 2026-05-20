package main

import "testing"

func TestFormatCell(t *testing.T) {
	tests := []struct {
		cell Cell
		want string
	}{
		{Cell{Position: Position{0, 0}, Alive: true}, "0 0"},
		{Cell{Position: Position{-500, -999}, Alive: true}, "-500 -999"},
		{Cell{Position: Position{9000000000000, 7000000000000}, Alive: true}, "9000000000000 7000000000000"},
		{Cell{Position: Position{-1, 4503599627370496}, Alive: true}, "-1 4503599627370496"},
	}

	for _, tt := range tests {
		got := FormatCell(tt.cell)
		if got != tt.want {
			t.Errorf("FormatCell(%v) = %q, want %q", tt.cell, got, tt.want)
		}
	}
}

func TestParseCell(t *testing.T) {
	// valid cases
	valid := []struct {
		line string
		want Position
	}{
		{"10 20", Position{10, 20}},
		{"-300 -400", Position{-300, -400}},
		{"4611686018427387903 -4611686018427387903", Position{4611686018427387903, -4611686018427387903}},
		{"  5   10  ", Position{5, 10}},
	}
	for _, tt := range valid {
		cell, err := ParseCell(tt.line)
		if err != nil {
			t.Errorf("ParseCell(%q): %v", tt.line, err)
			continue
		}
		if cell.Position != tt.want {
			t.Errorf("ParseCell(%q) = %v, want %v", tt.line, cell.Position, tt.want)
		}
	}

	// error cases
	bad := []string{"42", "1 2 3", "abc def", "99999999999999999999 0"}
	for _, line := range bad {
		if _, err := ParseCell(line); err == nil {
			t.Errorf("ParseCell(%q): expected error", line)
		}
	}
}
