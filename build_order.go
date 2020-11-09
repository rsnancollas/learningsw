package main

import "fmt"

type graph struct {
	nodes map[string]*node
}

type node struct {
	name     string
	incoming []*node
	outgoing []*node
	visited  bool
}

type depOrder struct {
	order []string
}

func buildGraph(nodeList []string, depList [][]string) (*graph, error) {
	g := graph{}
	g.nodes = make(map[string]*node)
	for _, pair := range depList {
		if len(pair) != 2 {
			return nil, fmt.Errorf("dependency pair contains more than two nodes - cannot produce graph")
		}
		primName := pair[0]
		secName := pair[1]
		primNode := g.addNode(primName)
		secNode := g.addNode(secName)
		primNode.outgoing = append(primNode.outgoing, secNode)
		secNode.incoming = append(secNode.incoming, primNode)
	}
	// Add any nodes that don't have dependencies
	for _, n := range nodeList {
		g.addNode(n)
	}
	return &g, nil
}

func (g *graph) addNode(name string) *node {
	n, ok := g.nodes[name]
	if !ok {
		n = &node{name: name}
		g.nodes[name] = n
	}
	return n
}

func (g *graph) print() {
	for name, n := range g.nodes {
		fmt.Printf("node: %v", name)
		fmt.Printf("| incoming: ")
		for _, n := range n.incoming {
			fmt.Printf("%v, ", n.name)
		}
		fmt.Printf("| outgoing: ")
		for _, n := range n.outgoing {
			fmt.Printf("%v, ", n.name)
		}
		fmt.Printf("\n")
	}
}

func (g *graph) findStarts() []*node {
	starts := []*node{}
	for _, n := range g.nodes {
		if len(n.incoming) == 0 {
			starts = append(starts, n)
		}
	}
	return starts
}

func (do *depOrder) makeDepOrder(n, lastNode *node) {
	// Remove lastNode from incoming list
	for i, iNode := range n.incoming {
		if iNode == lastNode {
			// remove element (does not preserve order)
			n.incoming[i] = n.incoming[len(n.incoming)-1]
			n.incoming = n.incoming[:len(n.incoming)-1]
			break
		}
	}
	if len(n.outgoing) == 0 {
		do.order = append(do.order, n.name)
		//fmt.Printf("%v\n", n.name)
		return
	} else if len(n.incoming) > 0 {
		return
	}
	do.order = append(do.order, n.name)
	//fmt.Printf("%v\n", n.name)
	for _, nextNode := range n.outgoing {
		do.makeDepOrder(nextNode, n)
	}
}

func main() {
	nodeList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	depList := [][]string{{"f", "a"}, {"a", "d"}, {"d", "c"}, {"f", "b"}, {"b", "g"}, {"g", "e"}, {"e", "h"}, {"b", "h"}, {"h", "d"}, {"j", "k"}, {"j", "k"}}

	g, _ := buildGraph(nodeList, depList)
	g.print()

	do := depOrder{}
	starts := g.findStarts()
	for _, n := range starts {
		do.makeDepOrder(n, nil)
	}
	if len(do.order) != len(nodeList) {
		fmt.Printf("Error: cannot create dependency list\n")
	} else {
		fmt.Printf("%v\n", do.order)
	}

}
