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

func LinkedListTest() {
	list := SinglyLinkedList[int]{}
	list.InsertBack(1)
	list.InsertBack(2)
	list.InsertFront(0)
	list.Display() // 0 1 2
	list.Delete(1)
	list.Display() // 0 2
}

func TreeTest() {
	tree := BinaryTree[int]{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Print("In-order: ")
	tree.InOrder() // 2 3 4 5 6 7 8

	fmt.Print("Pre-order: ")
	tree.PreOrder() // 5 3 2 4 7 6 8

	fmt.Print("Post-order: ")
	tree.PostOrder() // 2 4 3 6 8 7 5

	fmt.Println("Search 4:", tree.Search(4))
	fmt.Println("Search 9:", tree.Search(9))
}

func main() {
	StackTest()
	QueueTest()
    LinkedListTest()
    TreeTest()
}
