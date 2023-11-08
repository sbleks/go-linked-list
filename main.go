package main

import (
	"bufio"
	"fmt"
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

func printNodes() {
	if head == nil {
		fmt.Printf("Your list is empty\n")
		return
	}

	cur := head

	for cur != nil {
		fmt.Printf("%v->", cur)
		cur = cur.next
	}
	fmt.Printf("\n")
}

func printMenu() {
	fmt.Printf("Your options are:\n")
	fmt.Printf("\t1. Add a Node to the list.\n")
	fmt.Printf("\t2. Remove a Node from the list.\n")
	fmt.Printf("\t3. Insert a Node into the list.\n")
	fmt.Printf("\t4. Print your list.\n")
	fmt.Printf("\t5. Quit.\n")
}

var head *Node

func main() {

	option := -1

	for option != 5 {
		printMenu()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		selectedOption, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("The option you entered isn't valid, %s\n", err)
		}
		option = selectedOption
		if option > 0 && option <= 4 {
			switch option {
			case 1:
				// Add node to list
				fmt.Printf("What value should your node hold?\n")
				scanner.Scan()
				data := scanner.Text()
				dataInt, err := strconv.Atoi(data)
				if err != nil {
					fmt.Printf("The data you passed in isn't a number: %s\n", err)
				}
				newNode := addNode(dataInt)
				fmt.Printf("%v\n", newNode)
			case 2:
				// Remove a node from list
			case 3:
				// Insert a node into the list
			case 4:
				// Print list
				printNodes()
			case 5:
				break
			}
		}
	}

}
