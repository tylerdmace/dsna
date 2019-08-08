package bst

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root node
}

func NewBST(v int) *tree {
	n := new(node)
	n.value = v
	t := new(tree)
	t.root = *n
	return t
}

func (b *tree) Insert(v int) {
	checkAndInsert(&b.root, v)
}

func (b *tree) Delete(v int) node {
	panic("Not Implemented")
}

func (b *tree) Min() node {
	panic("Not Implemented")
}

func (b *tree) Max() node {
	panic("Not Implemented")
}

func (b *tree) Find(v int) node {
	panic("Not Implemented")
}

func (b *tree) Print() {
	printChildren(&b.root)
}

func checkAndInsert(n *node, v int) {
	if v < n.value {
		if n.left == nil {
			nn := new(node)
			nn.value = v
			n.left = nn
		} else {
			checkAndInsert(n.left, v)
		}
	} else if v > n.value {
		if n.right == nil {
			nn := new(node)
			nn.value = v
			n.right = nn
		} else {
			checkAndInsert(n.right, v)
		}
	}
}

func printChildren(n *node) {
	fmt.Printf("Node value: %v\r\n", n.value)

	if n.left != nil {
		printChildren(n.left)
	}

	if n.right != nil {
		printChildren(n.right)
	}
}
