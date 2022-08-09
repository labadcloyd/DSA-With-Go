package data_structures

import (
	"fmt"
)

// Queue is data structure that follows the First in First out rule
type Queue struct {
	Items []int
}

func (q *Queue) add(num int) {
	// adding an item to the end
	q.Items = append(q.Items, num)
}
func (q *Queue) pop() int {
	firstNum := q.Items[0]
	// removing the first item
	q.Items = q.Items[1:]
	return firstNum
}

// * Function for using the queue
func RunQueue(c *string) {
	choice := ""
	close := false
	queue := Queue{}

	// setting up items
	queue.Items = make([]int, 3, 10)
	for i := 0; i < 3; i++ {
		queue.Items[i] = i + 1
	}
	printItems(queue.Items)

	for !close {
		printChoice(&choice)

		if choice == "0" {
			*c = "99999"
			close = true
		}
		if choice == "1" {
			num := 0
			fmt.Println("Please enter the number you want to add: ")
			_, err := fmt.Scan(&num)
			for err != nil {
				fmt.Println("Please enter a number:")
				_, err = fmt.Scan(&num)
			}
			queue.add(num)
			printItems(queue.Items)
		}
		if choice == "2" {
			num := queue.pop()
			fmt.Println("Popped number: ", num)
			printItems(queue.Items)
		}
	}
}

func printChoice(choice *string) {
	fmt.Println("\nWhat do you want to do with your queue? (Enter The number)")
	fmt.Println("1. Add")
	fmt.Println("2. Pop")
	fmt.Println("0. Exit")
	fmt.Printf(": ")
	fmt.Scan(choice)
	fmt.Printf("\n\n")
}

func printItems(items []int) {
	fmt.Println("Current Queue:")
	for i := 0; i < len(items); i++ {
		fmt.Printf("-[%v]%v ", i, items[i])
	}
}
