package ch4

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func TreeSort() {
	values := []int{0, 9, 1, 3, 6, 2, 8, 7, 5, 4}
	Sort(values)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	printTree(root)
}

func appendValues(values []int, t *tree) []int  {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func printTree(t *tree) {
	if t == nil {
		return
	}
	printTree(t.left)
	fmt.Printf("%d ", t.value)
	printTree(t.right)
}