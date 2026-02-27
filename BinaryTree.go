package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

type TreeNode[T Ordered] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T Ordered] struct {
	Root *TreeNode[T]
}

func (t *BinaryTree[T]) Insert(value T) {
    newNode := &TreeNode[T]{Value: value}
    if t.Root == nil {
        t.Root = newNode
        return
    }
    insertNode(t.Root, newNode)
}

func insertNode[T Ordered](node, newNode *TreeNode[T]) {
    if newNode.Value < node.Value {
        if node.Left == nil {
            node.Left = newNode
        } else {
            insertNode(node.Left, newNode)
        }
    } else {
        if node.Right == nil {
            node.Right = newNode
        } else {
            insertNode(node.Right, newNode)
        }
    }
}

func (t *BinaryTree[T]) InOrder() {
    inOrder(t.Root)
    fmt.Println()
}

func inOrder[T Ordered](node *TreeNode[T]) {
    if node != nil {
        inOrder(node.Left)
        fmt.Print(node.Value, " ")
        inOrder(node.Right)
    }
}

func (t *BinaryTree[T]) PreOrder() {
	preOrder(t.Root)
	fmt.Println()
}

func preOrder[T Ordered](node *TreeNode[T]) {
	if node != nil {
		fmt.Print(node.Value, " ")
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func (t *BinaryTree[T]) PostOrder() {
	postOrder(t.Root)
	fmt.Println()
}

func postOrder[T Ordered](node *TreeNode[T]) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Print(node.Value, " ")
	}
}

func (t *BinaryTree[T]) Search(value T) bool {
	return search(t.Root, value)
}

func search[T Ordered](node *TreeNode[T], value T) bool {
    if node == nil {
        return  false
    }
    if value == node.Value {
        return true
    } else if value < node.Value {
        return search(node.Left, value)
    } else {
        return search(node.Right, value)
    }
    return false
}