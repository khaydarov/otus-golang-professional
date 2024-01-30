package internal

type CommentTree struct {
	head *Node
	m    map[NodeValue]*Node
}

func (c *CommentTree) GetRepliesCount(v NodeValue) int {
	node, ok := c.m[v]
	if !ok {
		return 0
	}

	return node.RepliesCount()
}

func (c *CommentTree) GetHead() *Node {
	return c.head
}

func BuildFromNode(node *Node) *CommentTree {
	m := make(map[NodeValue]*Node)
	var helper func(node *Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}

		m[node.Val] = node
		for _, reply := range node.Replies {
			helper(reply)
		}
	}

	helper(node)

	return &CommentTree{
		node,
		m,
	}
}
