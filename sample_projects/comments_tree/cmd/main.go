package main

import "fmt"

func main() {

	// todo: setup cobra
	// todo: separate

	tree := BuildFromNode(testTree())
	fmt.Println(tree.GetRepliesCount(1))
	fmt.Println(Serialize(tree.GetHead()))
}

func testTree() *Node {
	one := Node{Val: 1}
	two := Node{Val: 2}
	three := Node{Val: 3}
	four := Node{Val: 4}

	//"1.2.3.#.4.#"
	one.Replies = append(one.Replies, &two)
	one.Replies = append(one.Replies, &three)

	two.Replies = append(two.Replies, &four)

	return &one
}
