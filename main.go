package main

import "fmt"

func StackTest() {
    stack := Stack[int]{}
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println("Стек тест:")
    for !stack.IsEmpty() {
        if val, ok := stack.Pop(); ok {
            fmt.Println(val)
        }
    }
}

func QueueTest() {
    queue := Queue[string]{}
    queue.Enqueue("a")
    queue.Enqueue("b")
    queue.Enqueue("c")

    fmt.Println("Очередь тест:")
    for !queue.IsEmpty() {
        if val, ok := queue.Dequeue(); ok {
            fmt.Println(val)
        }
    }
}

func main() {
    StackTest()
    QueueTest()
}