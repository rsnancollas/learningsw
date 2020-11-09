package main

import (
	"fmt"
	"math"
)

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

func leftMost(n *node, v int) int {
	if n == nil {
		return v
	}
	return leftMost(n.left, n.value)
}

func binSearch(searchList []int) *node {
	if len(searchList) == 1 {
		return &node{value: searchList[0], left: nil, right: nil}
	} else if len(searchList) == 0 {
		return nil
	}
	pivot := int(math.Ceil(float64(len(searchList))/2) - 1)
	return &node{
		value: searchList[pivot],
		left:  binSearch(searchList[0:pivot]),
		right: binSearch(searchList[pivot+1:]),
	}
}

func addParent(n *node, p *node) {
	if n == nil {
		return
	}
	fmt.Printf("getting ready to go left: node-%v, parent-%v\n", n.left.value, n.value)
	addParent(n.left, n)
	n.parent = p
	fmt.Printf("node:%+v, parent:%+v\n", n, n.parent)
	addParent(n.right, n)
}

func preOrderTraverse(n *node) {
	if n == nil {
		fmt.Printf("nil\n")
		return
	}
	fmt.Printf("node:%v\n", n.value)
	preOrderTraverse(n.left)
	preOrderTraverse(n.right)
}

func getNext(n *node) int {
	// if right children exist, get leftmost node of right branch
	if n.right != nil {
		return leftMost(n.right, n.value)
	}
	for n.parent.value <= n.value {
		n = n.parent
	}
	return n.value

}

func main() {
	binSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	root := binSearch(binSlice)

	preOrderTraverse(root)
	addParent(root, nil)
	fmt.Printf("value is: %v, next is: %v\n", root.value, getNext(root))
	node5 := root.left.right.right
	fmt.Printf("value is: %v, next is: %v\n", node5.value, getNext(node5))

}
