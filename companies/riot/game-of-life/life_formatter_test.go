package main

import (
	"testing"
)

func TestFormatCell(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
		want string
	}{
		{
			name: "origin",
			cell: Cell{Position: Position{0, 0}, Alive: true},
			want: "0 0",
		},
		{
			name: "negative coordinates",
			cell: Cell{Position: Position{-500, -999}, Alive: true},
			want: "-500 -999",
		},
		{
			name: "large positive coordinates",
			cell: Cell{Position: Position{9000000000000, 7000000000000}, Alive: true},
			want: "9000000000000 7000000000000",
		},
		{
			name: "mixed sign coordinates",
			cell: Cell{Position: Position{-1, 4503599627370496}, Alive: true},
			want: "-1 4503599627370496",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatCell(tt.cell)
			if got != tt.want {
				t.Errorf("FormatCell(%v) = %q, want %q", tt.cell, got, tt.want)
			}
		})
	}
}

func TestParseCell(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantPos Position
		wantErr bool
	}{
		{
			name:    "simple positive",
			line:    "10 20",
			wantPos: Position{10, 20},
		},
		{
			name:    "both negative",
			line:    "-300 -400",
			wantPos: Position{-300, -400},
		},
		{
			name:    "large 64-bit values",
			line:    "4611686018427387903 -4611686018427387903",
			wantPos: Position{4611686018427387903, -4611686018427387903},
		},
		{
			name:    "extra whitespace between coords",
			line:    "  5   10  ",
			wantPos: Position{5, 10},
		},
		{
			name:    "single number missing second coord",
			line:    "42",
			wantErr: true,
		},
		{
			name:    "three numbers",
			line:    "1 2 3",
			wantErr: true,
		},
		{
			name:    "non-numeric input",
			line:    "abc def",
			wantErr: true,
		},
		{
			name:    "overflow beyond int64",
			line:    "99999999999999999999 0",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cell, err := ParseCell(tt.line)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseCell(%q): expected error, got nil", tt.line)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseCell(%q): unexpected error: %v", tt.line, err)
			}
			if cell.Position != tt.wantPos {
				t.Errorf("ParseCell(%q): got position %v, want %v", tt.line, cell.Position, tt.wantPos)
			}
			if !cell.Alive {
				t.Errorf("ParseCell(%q): expected cell to be alive", tt.line)
			}
		})
	}
}
