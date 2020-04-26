package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

// the structs from the other packages
// struct just for the first exercise
type NodeF = student.Node
type NodeFS = correct.Node

// simple struct, just insert a element in the struct
func main() {
	n1 := &NodeFS{}
	n2 := &NodeF{}

	table := [][]int{{{132423}}}

	for i := 0; i < 5; i++ {
		table = append(table, z01.MultRandIntBetween(-1000000, 10000000))
	}

	for _, items := range table {
		for _, item := range items {
			correct.CreateElem(n1, item)
			student.CreateElem(n2, item)
		}

		if n1.Data != n2.Data {
			z01.Fatalf("CreateElem == %d instead of %d\n", n1, n2)
		}
	}
}
