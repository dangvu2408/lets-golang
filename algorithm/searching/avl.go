package main

import (
	"fmt"
	"math"
)

type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

func newNode(val int) *Node {
	return &Node{value: val, height: 1}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func updateHeight(n *Node) {
	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
}

// Cân bằng cây
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}
	if key < node.value {
		node.left = insert(node.left, key)
	} else if key > node.value {
		node.right = insert(node.right, key)
	} else {
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	if balance > 1 && key < node.left.value {
		return rightRotate(node)
	}
	if balance < -1 && key > node.right.value {
		return leftRotate(node)
	}
	if balance > 1 && key > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	if balance < -1 && key < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key == node.value {
		return true
	} else if key < node.value {
		return search(node.left, key)
	} else {
		return search(node.right, key)
	}
}

func main() {
	var root *Node

	arr := []int{10, 20, 5, 4, 15, 25, 7}
	for _, val := range arr {
		root = insert(root, val)
	}

	fmt.Println("Tìm 15:", search(root, 15))
	fmt.Println("Tìm 8:", search(root, 8))
}
