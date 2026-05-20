# Conway's Game of Life

An implementation of Conway's Game of Life in 64-bit signed integer space. Reads live cell coordinates from stdin in Life 1.06 format, runs 10 iterations, and prints the result to stdout.

## Usage

```bash
go run . < input.txt
```

## File Overview

### main.go
Program entry point. Reads from stdin, creates a Game instance, ticks 10 generations, and prints the final board state. Pure IO glue with no game logic.

### entity.go
Core data types shared across the project:
- `Position` — an (X, Y) coordinate using int64 to support the full 64-bit signed integer range.
- `Cell` — a cell with its position, alive status, and live neighbor count from the most recent tick.
- `NeighborOffsets` — the 8 directional offsets for adjacent cell lookup.

### game_of_life.go
Game controller that owns the board state and generation counter:
- `NewGame()` — constructs a Game by parsing Life 1.06 input.
- `Tick()` — advances one generation in two rounds: (1) build candidate cells with neighbor counts, (2) apply Conway's rules to determine survivors. Includes `safeAdd()` to guard against int64 overflow at coordinate boundaries.
- `BoardStatusLife106()` — returns the current board as a sorted Life 1.06 formatted string.

### life_formatter.go
Single-cell serialization between Cell structs and Life 1.06 format:
- `FormatCell()` — converts one Cell to a coordinate line (e.g. `"2 3"`).
- `ParseCell()` — parses one coordinate line into a Cell. Returns an error for malformed input or values that overflow int64.

### go.mod
Go module definition.
