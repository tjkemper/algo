package main

import (
	"fmt"

	"github.com/tjkemper/algo/adj"
)

//BFS runs breadth-first search starting at Node n.
//Returns map of visited nodes with their level from n.
func BFS(n *adj.Node) map[*adj.Node]int {

	frontier := make([]*adj.Node, 1)
	frontier[0] = n

	visited := make(map[*adj.Node]int)
	visited[n] = 0

	level := 0

	for len(frontier) > 0 {
		next := make([]*adj.Node, 0)
		level++
		for _, curr := range frontier {
			for _, neighbor := range curr.Neighbors {
				if _, ok := visited[neighbor]; !ok {
					visited[neighbor] = level
					next = append(next, neighbor)
				}
			}
		}
		frontier = next

	}
	return visited
}

func main() {
	vertices := adj.BuildGraph()
	m := BFS(vertices[1])
	fmt.Println(m)
}
