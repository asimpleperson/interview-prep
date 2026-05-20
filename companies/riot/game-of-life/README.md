# Game of Life

Conway's Game of Life over 64-bit integer coordinates. Reads Life 1.06 from stdin, runs 10 iterations, writes result to stdout.

```bash
go run . < input.txt
go test ./...
```

## Files

- **entity.go** - Position, Cell, and neighbor offset definitions
- **life_formatter.go** - parse/format a single cell line in Life 1.06 format
- **game_of_life.go** - Game struct, tick logic, board parsing, output formatting
- **main.go** - reads stdin, ticks 10 times, prints result
