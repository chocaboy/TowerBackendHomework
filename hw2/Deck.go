package main

import "fmt"

type Node struct {
	number int
	prev   *Node
	next   *Node
}
type temp struct {
	head  *Node
	left  *Node
	right *Node
}

func (temp *temp) AddFront(number int) {
	newNode := &Node{number: number}
	if temp.left == nil {
		temp.left = newNode
		temp.right = newNode
	} else {
		newNode.next = temp.left
		temp.left.prev = newNode
		temp.left = newNode
	}
}

func (temp *temp) AddBack(number int) {
	newNode := &Node{number: number}
	if temp.right == nil {
		temp.left = newNode
		temp.right = newNode
	} else {
		newNode.prev = temp.right
		temp.right.next = newNode
		temp.right = newNode
	}
}

func (temp *temp) PopFront() (int, bool) {
	if temp.left == nil {
		fmt.Println("No elements")
		return 0, false
	}
	pop_num := temp.left.number
	if temp.left != temp.right {
		temp.left = temp.left.next
		temp.left.prev = nil
	} else {
		temp.right = nil
		temp.left = nil
	}
	fmt.Println(pop_num, "was popped")
	return pop_num, true
}

func (temp *temp) PopBack() (int, bool) {
	if temp.right == nil {
		fmt.Println("No elements")
		return 0, false
	}
	pop_num := temp.right.number
	if temp.left != temp.right {
		temp.right = temp.right.prev
		temp.right.next = nil
	} else {
		temp.left = nil
		temp.right = nil
	}
	fmt.Println(pop_num, "was popped")
	return pop_num, true
}

func (temp *temp) IsExist(number int) bool {
	now := temp.left
	flag := false
	for now != nil {
		if now.number == number {
			flag = true
			break
		}
		now = now.next
	}
	if flag {
		fmt.Println(number, "is exist")
		return true
	} else {
		fmt.Println(number, "is not exist")
		return false
	}
}

func (temp *temp) PrintuyEmae() {
	now := temp.left
	for now != nil {
		fmt.Print(now.number, " ")
		now = now.next
	}
	fmt.Println()
}

func main() {
	temp := temp{}
	fmt.Println(temp.PopBack())
	temp.AddBack(5)
	temp.PrintuyEmae()
	fmt.Println(temp.PopBack())
	temp.PrintuyEmae()
	temp.AddBack(5)
	temp.AddBack(10)
	temp.AddBack(15)
	temp.PrintuyEmae()
	fmt.Println(temp.PopBack())
	fmt.Println(temp.PopFront())
}
