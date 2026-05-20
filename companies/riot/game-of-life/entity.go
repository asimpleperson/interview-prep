package main

// Position represents a coordinate in 64-bit signed integer space.
type Position struct{ X, Y int64 }

// Cell represents a single cell on the board with its position, alive status,
// and the count of live neighbors from the most recent generation.
type Cell struct {
	Position       Position
	Alive          bool
	LiveNeighbors  int
}

// NeighborOffsets are the 8 surrounding cell offsets.
var NeighborOffsets = [8]Position{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}
