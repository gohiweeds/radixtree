package main

import (
	"fmt"

	"github.com/gohiweeds/radixtree"
)

func main() {
	fmt.Println("Radix Test Add/Find")
	strs := []string{
		"hello",
		"h",
		"hollo",

		"hellol",
		"hellw",
		"helle",
		"hello",
	}

	tree := radixtree.NewRadixTree()
	for k := range strs {
		tree.Add(strs[k])
	}

	//	tree.WalkAll()

	fmt.Println("Find", tree.Find("h"))
	fmt.Println("Find---", tree.Find("hellw"))
}
