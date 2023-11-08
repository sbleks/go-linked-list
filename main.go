package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

func addNode(data int) *Node {
	var newNode = &Node{}
	if head == nil {
		newNode.data = data
		head = newNode
	} else {
		newNode.data = data
		newNode.next = head
		head = newNode
	}

	return newNode
}

func removeNode(data int) bool {
	// if you have not added any nodes
	if head == nil {
		fmt.Printf("You have no nodes so you can't delete anything\n")
		return false
	}

	curr := head
	prev := head

	for curr != nil {
		if curr.data == data {

			// if the node you want to remove is the head
			if curr == head {
				head = curr.next
			} else {
				prev.next = curr.next
			}

			return true

		}

		prev = curr
		curr = curr.next
	}

	return false
}

func insertNode(data, nodeVal int) bool {
	var newNode = &Node{data: data}

	if head == nil {
		fmt.Printf("You have no nodes so you can't delete anything\n")
		return false
	}

	curr := head
	for curr != nil {
		if curr.data == nodeVal {

			// if you want to insert
			if curr == head {
				newNode.next = curr.next
				head = newNode
			} else {
				newNode.next = curr.next
				curr.next = newNode
			}

			return true
		}
		curr = curr.next
	}
	return false
}

func printNodes() {
	if head == nil {
		fmt.Printf("Your list is empty\n")
		return
	}

	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.data)
		cur = cur.next
	}
	fmt.Printf("\n")
}

func seedNodes() {
	for i := 0; i < 15; i++ {
		num := rand.Intn(100)
		addNode(num)
	}
}

func printMenu() {
	fmt.Printf("Your options are:\n")
	fmt.Printf("\t1. Add a Node to the list.\n")
	fmt.Printf("\t2. Remove a Node from the list.\n")
	fmt.Printf("\t3. Insert a Node into the list.\n")
	fmt.Printf("\t4. Print your list.\n")
	fmt.Printf("\t5. Seed list.\n")
	fmt.Printf("\t6. Quit.\n")
}

func scanNextNum(scanner *bufio.Scanner) int {
	scanner.Scan()
	data := scanner.Text()
	dataInt, err := strconv.Atoi(data)

	if err != nil {
		fmt.Printf("The data you passed in isn't a number: %s\n", err)
	}

	return dataInt
}

var head *Node

func main() {

	option := -1

	for option != 6 {
		printMenu()
		scanner := bufio.NewScanner(os.Stdin)
		selectedOption := scanNextNum(scanner)
		option = selectedOption
		if option > 0 && option <= 6 {
			switch option {
			case 1:
				// Add node to list
				fmt.Printf("What value should your node hold?\n")
				selectedOption := scanNextNum(scanner)
				newNode := addNode(selectedOption)
				fmt.Printf("%v\n", newNode)
				break
			case 2:
				// Remove a node from list
				fmt.Printf("What node should we remove?\n")
				printNodes()
				selectedOption := scanNextNum(scanner)
				if !removeNode(selectedOption) {
					fmt.Printf("The node you selected was not found in the list, please try again\n")
				} else {
					break
				}
			case 3:
				// Insert a node into the list
				fmt.Printf("Where should we add a node?\n")
				printNodes()
				nodePosition := scanNextNum(scanner)
				fmt.Printf("What value should your node hold?\n")
				data := scanNextNum(scanner)
				if !insertNode(data, nodePosition) {
					fmt.Printf("The node you selected was not found in the list, please try again\n")
				} else {
					break
				}
			case 4:
				// Print list
				printNodes()
				break
			case 5:
				seedNodes()
				break
			case 6:
				break
			}
		}
	}

}
