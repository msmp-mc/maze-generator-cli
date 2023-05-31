package Generators

type kruskal struct {
	*Maze
}

// NewRandomisedKruskal is a func generating a new maze with the Randomized Kruskal's Algorithm
func NewRandomisedKruskal(b *Maze) error {
	m := kruskal{b}
	m.firstWalls()
	return nil
}

func (m *kruskal) firstWalls() {
	m.HorizontalWalls = generateWalls(m.Width, m.Height, false)
	m.VerticalWalls = generateWalls(m.Height, m.Width, true)
}
