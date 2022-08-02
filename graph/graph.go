package graph

type GraphNode struct {
	Value any
}

type Graph struct {
	Nodes []*GraphNode
	Edges map[GraphNode][]*GraphNode
}

// NewGraph atua como um alias para AddNode(nil, node)
func NewGraph(node *GraphNode) *Graph {
	return AddNode(nil, node)
}

// AddNode adiciona um novo nó no grafo
func AddNode(graph *Graph, node *GraphNode) *Graph {
	if graph == nil {
		return &Graph{
			Nodes: []*GraphNode{node},
		}
	}

	graph.Nodes = append(graph.Nodes, node)

	return graph
}

// AddEdge adiciona uma nova conexão dupla entre dois nós do grafo
func AddEdge(graph *Graph, node1, node2 *GraphNode) *Graph {
	if graph == nil {
		graph = &Graph{
			Nodes: []*GraphNode{node1, node2},
		}
	}

	if graph.Edges == nil {
		graph.Edges = make(map[GraphNode][]*GraphNode)
	}

	graph.Edges[*node1] = append(graph.Edges[*node1], node2)
	graph.Edges[*node2] = append(graph.Edges[*node2], node1)

	return graph
}
