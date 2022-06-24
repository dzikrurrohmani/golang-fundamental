package main

import "fmt"

type node struct {
	left  *node
	right *node
	data  int
}

func (n *node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &node{
				data:  data,
				// left:  nil,
				// right: nil,
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{
				data:  data,
				// left:  nil,
				// right: nil,
			}
		} else {
			n.right.insert(data)
		}
	}
}

type binaryTree struct {
	root *node
}

func (bt *binaryTree) insert(data int) {
	if bt.root == nil {
		bt.root = &node{
			data:  data,
			// left:  nil,
			// right: nil,
		}
	} else {
		bt.root.insert(data)
	}
}

func print(node *node, space int, ch string) {
	if node == nil {
		return
	}

	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%s:%v\n", ch, node.data)
	print(node.left, space+2, "L")
	print(node.right, space+2, "R")
}

func main() {
	tree := &binaryTree{}
	tree.insert(100)
	tree.insert(20)
	tree.insert(120)
	tree.insert(50)
	tree.insert(15)
	tree.insert(85)
	print(tree.root, 0, "R")
}
