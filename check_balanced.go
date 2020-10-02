// Exercise 4.4

package main

import(
	"fmt"
	"math"
)

type node struct {
	val int
	left *node
	right *node
	leftHeight float64
	rightHeight float64
	visited bool
}

func (n *node) addHeight() {
	if n.left==nil && n.right==nil {
		n.leftHeight = 0
		n.rightHeight = 0
		return
	}
	if n.left != nil {
		n.left.addHeight()
		n.leftHeight = 1 + math.Max(n.left.leftHeight, n.left.rightHeight)
	} else {
		n.leftHeight = 0
	}
	if n.right != nil {
		n.right.addHeight()
		n.rightHeight = 1 + math.Max(n.right.leftHeight, n.right.rightHeight)
	}
}

func (n *node) searchPrint() {
	if n == nil {
		return
	}
	if math.Abs(n.leftHeight - n.rightHeight) > 1 {
		fmt.Printf("This tree is not balanced! Node: %v has leftHeight: %v and rightHeight: %v\n", n.val, n.leftHeight, n.rightHeight)
		return
	}
	fmt.Printf("node: %v, leftHeight: %v, rightHeight: %v\n", n.val, n.leftHeight, n.rightHeight)
	n.visited = true
	if n.left == nil || n.left.visited == false {
		n.left.searchPrint()
	}
	if n.right == nil || n.right.visited == false {
		n.right.searchPrint()
	}

}

func (n *node) heightDiff() float64{
	if n == nil {
		return 0
	}
	leftBranch := n.left.heightDiff()
	if leftBranch == -100 {
		return -100
	}
	rightBranch := n.right.heightDiff()
	fmt.Printf("Node: %v, leftHeight: %v, rightHeight: %v\n", n.val, leftBranch, rightBranch)
	if rightBranch == -100 {
		return -100
	}
	if math.Abs(leftBranch - rightBranch) > 1{
		return -100
	}
	return 1 + math.Max(leftBranch, rightBranch)
}

func (n *node) isBalanced() bool {
	return n.heightDiff() != -100
}

func main() {
	n12 := node{val:12}
	n10 := node{val:10, left:&n12}
	n11 := node{val:11}
	n7 := node{val:7, left:&n10, right:&n11}
	n4 := node{val:4, right:&n7}
	n2 := node{val:2, left:&n4}
	n8 := node{val:8}
	n5 := node{val:5, left:&n8}
	n9 := node{val:9}
	n6 := node{val:6, right:&n9}
	n3 := node{val:3, left:&n5, right:&n6}
	n1 := node{val:1, left:&n2, right:&n3}

	n1.addHeight()
	n1.searchPrint()

	if n1.isBalanced() {
		fmt.Printf("Tree is balanced\n")
	} else {
		fmt.Printf("Tree is unbalanced\n")
	}




}