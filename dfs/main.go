package main

import (
	"fmt"

	"github.com/tjkemper/algo/adj"
)

//DFS runs depth-first search for each vertex.
//Use for disconnected graphs (forests).
func DFS(vertices []*adj.Node) {

	visited := make(map[*adj.Node]bool)

	for _, v := range vertices {
		if _, ok := visited[v]; !ok {
			DFSVisit(v, visited)
		}
	}
}

//DFSVisit runs depth-first search starting at Node v.
func DFSVisit(v *adj.Node, visited map[*adj.Node]bool) {
	fmt.Println(v.Data)
	visited[v] = true
	for _, temp := range v.Neighbors {
		if _, ok := visited[temp]; !ok {
			DFSVisit(temp, visited)
		}
	}
}

func main() {
	vertices := adj.BuildGraph()
	DFS(vertices)
	//DFSVisit(vertices[3], make(map[*adj.Node]bool))
}
