// Exercise 4.5

package main

import (
	"fmt"
	"math"
)

type node struct{
	value int
	left *node
	right *node
	visited bool
}

func preTraverse(n *node){
	if n == nil{
		fmt.Println("nil")
		return
	}
	fmt.Println(n.value)
	preTraverse(n.left)
	preTraverse(n.right)
}

func binSearch(searchList []int) *node {
	if len(searchList) == 1 {
		return &node{value:searchList[0], left:nil, right:nil}
	} else if len(searchList) == 0 {
		return nil
	}
	pivot := int(math.Ceil(float64(len(searchList))/2) -1)
	return &node{
		value:searchList[pivot],
		left:binSearch(searchList[0:pivot]),
		right:binSearch(searchList[pivot+1:]),
	}
}

func isBST(n *node, lt, gt int) bool {
	if n == nil {
		return true
	}
	if n.visited == true {
		return true
	}
	n.visited = true
	fmt.Printf("node: %v. %v < %v <= %v\n", n.value, gt, n.value, lt)
	// Don't forget = edge case!
	if n.value <= gt || n.value > lt {
		return false
	}
	return isBST(n.left, n.value, gt) && isBST(n.right, lt, n.value)

}

func main() {
	sl := []int{1,3,4,7,9,12,13,16,17,17,21,23,24,26}
	root := binSearch(sl)
	preTraverse(root)

	// Good to use max/min int here (but consider what the range is)
	bstBool := isBST(root, math.MaxInt32, math.MinInt32)

	fmt.Printf("Binary tree is a BST: %v\n", bstBool)

}