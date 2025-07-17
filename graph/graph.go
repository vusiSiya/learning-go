package graph

import (
	"slices"

	queue "github.com/vusiSiya/data-structures/queue"
)

type Graph[T comparable] struct {
	Nodes []*Node[T]
	Size  int
}

type Node[T comparable] struct {
	Value      T
	Successors []*Node[T]
	Weights    []*float32
}

func AddNode[T comparable](graph *Graph[T], u T, weight float32) {
	var node = &Node[T]{
		Value:      u,
		Successors: []*Node[T]{},
		Weights:    []*float32{&weight},
	}

	if !Contains(graph.Nodes, *node) {
		graph.Nodes = append(graph.Nodes, node)
	}
	graph.Size++
}

func AddSuccessor[T comparable](graph *Graph[T], u T, v T) {
	var U = GetNode(graph.Nodes, u)
	var V = GetNode(graph.Nodes, v)

	if U != nil && V != nil {
		U.Successors = append(U.Successors, V)
		U.Weights = append(U.Weights, V.Weights[0])
	}
}

func GetNode[T comparable](nodes []*Node[T], u T) *Node[T] {
	for _, node := range nodes {
		if node.Value == u {
			return node
		}
	}
	return nil
}

func GetSuccessors[T comparable](array []*Node[T], u T) []*Node[T] {
	var U = GetNode(array, u)
	if U != nil {
		return U.Successors
	}
	return []*Node[T]{}
}

// connected graph BFT
func ListNodesBFT[T comparable](graph *Graph[T]) []T {
	var list []T
	var qList queue.Queue[T]
	queue.Enqueue(&qList, graph.Nodes[0].Value)

	for qList.Head != nil {
		var currentItem = queue.Dequeue(&qList)
		if !slices.Contains(list, currentItem) {
			list = append(list, currentItem)
			var successors = GetSuccessors(graph.Nodes, currentItem)
			for _, value := range successors {
				queue.Enqueue(&qList, value.Value)
			}
		}
	}
	return list
}

// connected graph DFT
func ListNodesDFT[T comparable](graph *Graph[T]) []T {
	var lstNodes = &[]*Node[T]{}
	var list []T
	listNodesDFT(graph, lstNodes, graph.Nodes[0].Value)
	for _, node := range *lstNodes {
		list = append(list, node.Value)
	}
	return list
}

func listNodesDFT[T comparable](graph *Graph[T], lstNodes *[]*Node[T], u T) {
	var nodeExists = false
	for _, node := range *lstNodes {
		if node.Value == u {
			nodeExists = true
			break
		}
	}

	if !nodeExists {
		var node = GetNode(graph.Nodes, u)
		*lstNodes = append(*lstNodes, node)
		for _, successor := range node.Successors {
			listNodesDFT(graph, lstNodes, successor.Value)
		}
	}
}

func Contains[T comparable](list []*Node[T], V Node[T]) bool {
	for _, item := range list {
		if item.Value == V.Value {
			return true
		}
	}
	return false
}
