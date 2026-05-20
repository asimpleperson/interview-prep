package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	game, err := NewGame(scanner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		game.Tick()
	}

	fmt.Print(game.BoardStatusLife106())
}
