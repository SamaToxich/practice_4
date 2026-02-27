package main

import "fmt"

type SinglyNode[T comparable] struct {
    Value T
    Next *SinglyNode[T]
}

type SinglyLinkedList[T comparable] struct {
    Head *SinglyNode[T]
}

func (l *SinglyLinkedList[T]) InsertFront(value T) {
    node := &SinglyNode[T]{Value: value}
    node.Next = l.Head
    l.Head = node
}

func (l *SinglyLinkedList[T]) InsertBack(value T) {
    node := &SinglyNode[T]{Value: value}
    if l.Head == nil {
        l.Head = node
        return
    }
    current := l.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = node
}

func (l *SinglyLinkedList[T]) Delete(value T) bool {
    if l.Head == nil {
        return false
    }
    if l.Head.Value == value {
        l.Head = l.Head.Next
        return true
    }
    prev := l.Head
    current := l.Head.Next
    for current != nil {
        if current.Value == value {
            prev.Next = current.Next
            return true
        }
        prev = current
        current = current.Next
    }
    return false
}

func (l *SinglyLinkedList[T]) Display() {
    current := l.Head
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}