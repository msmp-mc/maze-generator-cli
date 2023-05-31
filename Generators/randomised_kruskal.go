package Generators

type kruskal struct {
	*Maze
}

// NewRandomisedKruskal is a func generating a new maze with the Randomized Kruskal's Algorithm
func NewRandomisedKruskal(b *Maze) error {
	m := kruskal{b}
	m.generateWalls()
	m.generateCells()
	return nil
}
