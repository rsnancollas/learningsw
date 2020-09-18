package main

import (
	"fmt"
)

type binNode struct{
	value int
	left *binNode
	right *binNode
}

type llNode struct{
	value *binNode
	next *llNode
}

func llGenerator(llUp *llNode) *llNode {
	llCur := &llNode{}
	var prev *llNode
	root := llCur
	for llUp != nil {
		if llUp.value.left != nil {
			llCur.value = llUp.value.left
			if prev != nil {
				prev.next = llCur
			}
			prev = llCur
			llCur = &llNode{}
		}
		if llUp.value.right != nil {
			llCur.value = llUp.value.right
			if prev != nil {
				prev.next = llCur
			}
			prev = llCur
			llCur = &llNode{}
		}
		llUp = llUp.next
	}
	return root
}

func generateBinTree(intList []int) *binNode {
	nodeList := []*binNode{}
	for _, n := range intList {
		newNode := &binNode{value:n}
		nodeList = append(nodeList, newNode)
	}
	l := len(nodeList)
	j := 0
	for _, node := range nodeList {
		if j+1 < l {
			node.left = nodeList[j+1]
		} else {
			node.left = nil
		}
		if j+2 < l {
			node.right = nodeList[j+2]
		} else {
			node.right = nil
		}
		j += 2
	}
	return nodeList[0]
}

func llPrinter(ll *llNode) {
	for ll != nil {
		fmt.Printf("%v, ", ll.value)
		ll = ll.next
	}
}

func main() {
	testSlice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15, 16}
	root := generateBinTree(testSlice)
	fmt.Println(root.value)

	llSlice := []*llNode{}
	newll := &llNode{value:root}
	for newll.value !=nil {
		llSlice = append(llSlice, newll)
		newll = llGenerator(newll)
	}

	for i, ll := range llSlice{
		fmt.Printf("Line %v: ", i)
		llPrinter(ll)
		fmt.Printf("\n")
	}
}
