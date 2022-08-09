package main

import (
	"dsa/data_structures"
	"fmt"
)

func main() {
	c := ""
	close := false
	// 1 = queues
	// 2 = stacks
	// 3 = linked lists
	for !close {
		printChoice(&c)
		if c == "0" {
			close = true
		}
		if c == "1" {
			data_structures.RunQueue(&c)
		}
	}
}

func printChoice(c *string) {
	fmt.Println("What data structure would you like to try out?")
	fmt.Println("1. Queue")
	fmt.Println("2. Stack")
	fmt.Println("3. Linked list")
	fmt.Println("0. Exit")
	fmt.Printf("Please enter the number of your choice: ")
	fmt.Scan(c)
	fmt.Printf("\n\n\n")
}
