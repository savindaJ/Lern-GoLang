package main

import "fmt"

func Array() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array: ", arr)
}

func LinkedList() {
	linkedList := []int{1, 2, 3, 4, 5}
	linkedList = append(linkedList, 6)
	fmt.Println("Linked List: ", linkedList)
}

func Stack() {
	stack := []int{1, 2, 3, 4, 5}
	stack = append(stack, 6)
	stack = append(stack, 7)

	stack = stack[:len(stack)-1]	
	fmt.Println("Stack: ", stack)
}

func Queue() {
	queue := []int{1, 2, 3, 4, 5}
	fmt.Println("Queue: ", queue)
}

func Tree() {
	tree := []int{1, 2, 3, 4, 5}
	fmt.Println("Tree: ", tree)
}

func Graph() {
	graph := []int{1, 2, 3, 4, 5}
	fmt.Println("Graph: ", graph)
}

func HashTable() {
	hashTable := []int{1, 2, 3, 4, 5}
	fmt.Println("Hash Table: ", hashTable)
}

func Heap() {
	heap := []int{1, 2, 3, 4, 5}
	fmt.Println("Heap: ", heap)
}		

func Trie() {
	trie := []int{1, 2, 3, 4, 5}
	fmt.Println("Trie: ", trie)
}

func BinarySearchTree() {
	binarySearchTree := []int{1, 2, 3, 4, 5}
	fmt.Println("Binary Search Tree: ", binarySearchTree)
}

func RedBlackTree() {
	redBlackTree := []int{1, 2, 3, 4, 5}
	fmt.Println("Red Black Tree: ", redBlackTree)
}

func AVLTree() {
	avlTree := []int{1, 2, 3, 4, 5}
	fmt.Println("AVL Tree: ", avlTree)
}

func BTree() {
	bTree := []int{1, 2, 3, 4, 5}
	fmt.Println("B Tree: ", bTree)
}

func BPlusTree() {
	bPlusTree := []int{1, 2, 3, 4, 5}
	fmt.Println("B Plus Tree: ", bPlusTree)
}

func BStarTree() {
	bTree := []int{1, 2, 3, 4, 5}
	fmt.Println("B Star Tree: ", bTree)
}

func Slice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice: ", slice)

	// use make to create a slice
	slice = make([]int, 10)
	fmt.Println("Slice: ", slice)

	// use make to create a slice with a specific length
	slice = make([]int, 10, 20)
	fmt.Println("Slice: ", slice)

	// use make to create a slice with a specific length and capacity
	slice = make([]int, 10, 20)
	fmt.Println("Slice: ", slice)

	// slice expansion
	slice = append(slice, 11)
	fmt.Println("Slice: ", slice)
	s := slice[0:5]
	fmt.Println("Slice: ", s)

	fmt.Println(slice[0])
	fmt.Println("Slice capacity: ", cap(slice))

}
// func main() {
	// Slice()
	// Array()
	// LinkedList()
	// Stack()
	// Queue()
	// Tree()
	// Graph()
	// HashTable()
	// Heap()
	// Trie()
	// BinarySearchTree()
	// RedBlackTree()
	// AVLTree()
	// BTree()
	// BPlusTree()
	// BStarTree()
// }