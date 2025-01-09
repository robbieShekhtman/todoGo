package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	todos := Todos{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n---- TODO Application ----")
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. Mark Task as Completed/Incomplete")
		fmt.Println("4. Edit Task Title")
		fmt.Println("5. Show All Tasks")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			todos.Add(title)
			fmt.Println("Task added successfully!")
		case "2":
			fmt.Print("Enter task index to delete(starting at 1): ")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
			} else {
				todos.Delete(index - 1)
			}
		case "3":
			fmt.Print("Enter task index to toggle completion(starting at 1): ")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
			} else {
				todos.ChangeCompletion(index - 1)
			}
		case "4":
			fmt.Print("Enter task index to edit(starting at 1): ")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Print("Enter new title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			todos.EditTitle(index-1, title)
		case "5":
			todos.ShowAllTasks()
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
