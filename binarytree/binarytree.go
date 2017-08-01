package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Ch, t2Ch := make(chan int), make(chan int)
	go Walk(t1, t1Ch)
	go Walk(t2, t2Ch)

	for i := 0; i < 10; i++ {
		if <-t1Ch != <-t2Ch {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
