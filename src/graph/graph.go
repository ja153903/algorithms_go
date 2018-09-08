package graph

type GraphNode struct {
	Val int
	Children []*GraphNode
	Visited bool
}

func (gn *GraphNode) AddChild(child *GraphNode) {
	gn.Children = append(gn.Children, child)
}

type Graph struct {
	Start *GraphNode
	Nodes []*GraphNode
	Bidirectional bool
}

func (g *Graph) AddGraphNode(gn *GraphNode) {
	g.Nodes = append(g.Nodes, gn)
}

func (g *Graph) SearchForNode(val int) *GraphNode {
	if len(g.Nodes) == 0 {
		return nil
	}

	for _, gn := range g.Nodes {
		if gn.Val == val {
			return gn
		}
	}

	return nil
}
