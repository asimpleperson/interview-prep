package main

type Position struct{ X, Y int64 }

type Cell struct {
	Position      Position
	Alive         bool
	LiveNeighbors int
}

// 8 surrounding cell offsets
var NeighborOffsets = [8]Position{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}
