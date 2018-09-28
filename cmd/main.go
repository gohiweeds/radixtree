package main

import (
	"fmt"

	"github.com/gohiweeds/radixtree"
)

func main() {
	fmt.Println("Radix Test Add/Find")
	strs := []string{
		"/hello",
		"/h",
		"/hollo",
		"/hell",
		"/hellol",
		"/hellw",
		"/helle",
		"/hello",
	}

	tree := radixtree.NewRadixTree()
	for k := range strs {
		tree.Add(strs[k])
	}

	tree.WalkAll()

	//	fmt.Println("Find", tree.Find("/h"))
	//	fmt.Println("Find---", tree.Find("/hellw"))
	//	fmt.Printf("Remove %v\n", tree.Delete("/hellw"))

	for k := range strs {
		fmt.Printf("remove=%s, %v\n", strs[k], tree.Delete(strs[k]))
	}

	tree.WalkAll()
}
