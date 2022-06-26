package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	val      int
	nextNode *Node
}

func reverse(list Node) *Node {
	var current *Node = &list
	var prev *Node = nil

	for current != nil {
		var next *Node = current.nextNode
		current.nextNode = prev
		prev, current = current, next
	}
	// list = *prev
	return prev
}

func printList(current *Node) {
	if current == nil {
		return
	}
	fmt.Println(current.val)
	printList(current.nextNode)
}

func main() {
	var head Node = Node{}
	fmt.Println("Вводите элементы связанного списка, нажимая enter.\nКогда закончите - нажмите enter на пустой строке")
	scanner := bufio.NewScanner(os.Stdin)

	var prev *Node = &head
	var next *Node = prev

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else {
			number, err := strconv.Atoi((scanner.Text()))
			if err == nil {
				next.val = number
				prev.nextNode = *&next
				prev, next = next, &Node{}
			}
		}
	}

	head = *reverse(head)

	fmt.Println("развернутый список:")
	printList(&head)
}
