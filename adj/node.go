package adj

import "fmt"

//Node represents a vertex in a graph.
//This graph representation uses adjacency lists.
type Node struct {
	Data      int
	Neighbors []*Node
}

//Add appends neighbor as an adjacent vertex to n.
func (n *Node) Add(neighbor *Node) {
	n.Neighbors = append(n.Neighbors, neighbor)
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.Data)
}

//Pretty prints the Node data and its neighbors' data.
func (n Node) Pretty() string {
	str := fmt.Sprintf("%d : { ", n.Data)
	for _, neighbor := range n.Neighbors {
		str += fmt.Sprintf("%d ", neighbor.Data)
	}
	str += "}"
	return str
}

//BuildGraph reads from STDIN to build a graph represented by adjacency lists.
//First line of input contains the number of vertices, v, and the number of edges, e.
//Then there are e lines of input, representing each edge.
func BuildGraph() []*Node {

	var v, e int
	fmt.Scanf("%d", &v)
	fmt.Scanf("%d", &e)

	vertices := make([]*Node, v)

	for i := 0; i < len(vertices); i++ {
		vertices[i] = &Node{i, make([]*Node, 0)}
	}

	for i := 0; i < e; i++ {
		var v1, v2 int
		fmt.Scanf("%d", &v1)
		fmt.Scanf("%d", &v2)
		vertices[v1].Add(vertices[v2])
		vertices[v2].Add(vertices[v1])
	}
	return vertices
}
