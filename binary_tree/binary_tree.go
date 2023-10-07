package binary_tree

import (
	"cmp"
	"fmt"
)

type (
	Tree[E cmp.Ordered] struct {
		root *node[E]
		size int
	}

	node[E cmp.Ordered] struct {
		key   E
		left  *node[E]
		right *node[E]
	}
)

func New[E cmp.Ordered]() *Tree[E] {
	return &Tree[E]{}
}

func (tree *Tree[E]) Insert(value E) {
	if tree.root == nil {
		tree.root = &node[E]{value, nil, nil}
	}
	tree.size++
	tree.root.insert(&node[E]{value, nil, nil})
}

func (root *node[E]) insert(newNode *node[E]) {

	//if data exists, skip
	if root.key == newNode.key {
		return
	}

	// to right-subtree
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

func (tree *Tree[E]) Size() int {
	return tree.size
}

func (tree *Tree[E]) Search(value E) bool {
	tree.size--
	return searchElement[E](tree.root, value)
}

func searchElement[E cmp.Ordered](root *node[E], value E) bool {
	if root != nil {
		if value == root.key {
			return true
		} else if value > root.key {
			return searchElement(root.right, value)
		} else {
			return searchElement(root.left, value)
		}
	}
	return false
}

func (tree *Tree[E]) Show() {
	printNode(tree.root)
}

func printNode[E cmp.Ordered](root *node[E]) {
	if root != nil {
		printNode(root.left)
		fmt.Printf("key %+v\n", root.key)
		printNode(root.right)
	}
}

func (tree *Tree[E]) FindMin() {
	fmt.Println(minValue(tree.root))
}

func minValue[E cmp.Ordered](root *node[E]) E {
	if root != nil {
		if root.left == nil {
			return root.key
		}
		return minValue(root.left)
	}
	return root.key
}

// FindMax - print max element tree
func (tree *Tree[E]) FindMax() {
	fmt.Println(maxValue(tree.root))
}

func maxValue[E cmp.Ordered](root *node[E]) E {
	if root != nil {
		if root.right == nil {
			return root.key
		}
		return maxValue(root.right)
	}
	return root.key
}

// Delete element tree
func (tree *Tree[E]) Delete(value E) bool {
	if !tree.Search(value) || tree.root == nil {
		return false
	}

	if tree.root.key == value {
		tempRoot := &node[E]{}
		tempRoot.left = tree.root
		r := del(tree.root, tempRoot, value)
		tree.root = tempRoot.left
		return r
	}
	return del(tree.root.left, tree.root, value) || del(tree.root.right, tree.root, value)
}

func del[E cmp.Ordered](root *node[E], parent *node[E], value E) bool {
	switch {
	case root.key == value:
		if root.left != nil && root.right != nil {
			root.key = minValue(root.right)
			return del(root.right, root, root.key)
		}
		link(parent, root)
		return true
	case root.key > value:
		if root.left == nil {
			return false
		}
		return del(root.left, root, value)
	case root.key < value:
		if root.right == nil {
			return false
		}
		return del(root.right, root, value)
	}
	return false
}

func link[E cmp.Ordered](parent *node[E], root *node[E]) {
	if parent.left == root {
		if root.left != nil {
			parent.left = root.left
		} else {
			parent.left = root.right
		}
	} else if parent.right == root {
		if root.left != nil {
			parent.right = root.left
		} else {
			parent.right = root.right
		}
	}
}
