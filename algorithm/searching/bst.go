package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}
	if value < root.value {
		root.left = insert(root.left, value)
	} else if value > root.value {
		root.right = insert(root.right, value)
	}
	return root
}

func search(root *Node, key int) *Node {
	if root == nil || root.value == key {
		return root
	}
	if key < root.value {
		return search(root.left, key)
	}
	return search(root.right, key)
}

func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Printf("%d ", root.value)
		inorder(root.right)
	}
}

func main() {
	var root *Node
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	for _, v := range arr {
		root = insert(root, v)
	}

	fmt.Print("Inorder Traversal: ")
	inorder(root)
	fmt.Println()

	key := 6
	found := search(root, key)
	if found != nil {
		fmt.Printf("Đã tìm thấy %d trong BST.\n", key)
	} else {
		fmt.Printf("Không tìm thấy %d trong BST.\n", key)
	}
}
