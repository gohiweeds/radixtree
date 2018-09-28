package radixtree

import (
	"testing"
)

func TestRadixTree(t *testing.T) {
	strs := []string{
		"hello",
		"h",
		"hollo",
		"hellol",
		"hellw",
		"helle",
		"hello",
	}

	tree := NewRadixTree()
	for k := range strs {
		tree.Add(strs[k])
	}

	expected := true
	for k := range strs {
		result := tree.Find(strs[k])
		if result != expected {
			t.Errorf("Expected value [%s] in tree but instead not found %v!", strs[k], result)
		}
	}

}
