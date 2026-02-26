package main

type SinglyNode[T any] struct {
    Value T
    Next *SinglyNode[T]
}

type SinglyLinkedList[T any] struct {
    Head *SinglyNode[T]
}

func (l *SinglyLinkedList[T]) InsertFront(value T) {
    node := &SinglyNode[T]{Value:value}
    node.Next = l.Head
    l.Head = node
}

func (l *SinglyLinkedList[T]) InsertBack(value T) {
    node := &SinglyNode[T]{Value:value}
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

    return false
}


