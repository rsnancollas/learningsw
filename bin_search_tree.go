// Exercise 4.2

package main

import (
	"fmt"
	"math"
)

type node struct{
	left *node
	right *node
	value int
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
	fmt.Printf("search list len: %v, pivot: %v\n", len(searchList), pivot)
	return &node{
		value:searchList[pivot],
		left:binSearch(searchList[0:pivot]),
		right:binSearch(searchList[pivot+1:]),
	}
}

func main(){
	sl := []int{1,3,4,7,9,12,13,16,17,18,21,23,24,26}
	root := binSearch(sl)
	preTraverse(root)
}
