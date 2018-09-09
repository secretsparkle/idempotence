package structures

type Tree struct {
	Board    [8][8]string
	Children []*Tree // to check if it is a leaf, we can check if Children is nil
	Parent   *Tree
	Score    int
}
