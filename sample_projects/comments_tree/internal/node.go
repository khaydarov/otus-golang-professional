package internal

import "strconv"

type NodeValue int

type Node struct {
	Val     NodeValue
	Replies []*Node
}

func (n *Node) RepliesCount() int {
	var helper func(node *Node) int
	helper = func(node *Node) int {
		if len(node.Replies) == 0 {
			return 0
		}

		var result int
		for _, reply := range node.Replies {
			result += helper(reply)
		}

		return result + len(node.Replies)
	}

	return helper(n)
}

func Serialize(node *Node) string {
	if node == nil {
		return ""
	}

	var helper func(node *Node) string
	helper = func(node *Node) string {
		result := strconv.Itoa(int(node.Val))

		if len(node.Replies) == 0 {
			return result + "#"
		}

		for _, v := range node.Replies {
			result += helper(v)
		}

		return result
	}

	return helper(node)
}

func Deserialize(s string) *Node {
	//var helper func(s string) *Node
	//helper = func(s string) *Node {
	//	if s[0] == '#' {
	//		return nil
	//	}
	//
	//	newNode := &Node{Val: NodeValue()}
	//}

	return nil
}
