package main

import (
	"fmt"
	"strconv"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) Delete(index int) {
	t := *todos

	if index < len(t) && index >= 0 {
		gone := t[index].Title
		*todos = append(t[:index], t[index+1:]...)
		fmt.Printf("\n\n")
		fmt.Println("Deleted your task to: " + gone)

	} else {
		fmt.Printf("\n\n")
		fmt.Println("Not Valid Index")
	}

}

func (todos *Todos) ChangeCompletion(index int) {
	t := *todos

	if index < len(t) && index >= 0 {
		completed := t[index].Completed

		if !completed {
			time := time.Now()
			t[index].CompletedAt = &time
		}

		t[index].Completed = !completed
		fmt.Printf("\n\n")
		fmt.Println("You have now completed your task to: " + t[index].Title)
	} else {
		fmt.Printf("\n\n")
		fmt.Println("Not Valid Index")
	}

}

func (todos *Todos) EditTitle(index int, title string) {
	t := *todos

	if index < len(t) && index >= 0 {

		t[index].Title = title
		fmt.Printf("\n\n")
		fmt.Println("Title changed to: " + title)
	} else {
		fmt.Printf("\n\n")
		fmt.Println("Not Valid Index")
	}

}

func (todos *Todos) ShowAllTasks() {
	t := *todos
	fmt.Printf("\n\n")
	fmt.Println("---------------------------Your Current List---------------------------")
	for i := 0; i < len(t); i++ {
		fmt.Printf("\n")
		fmt.Println("Task " + strconv.Itoa(i+1) + ":")
		fmt.Println("Title: " + t[i].Title)
		fmt.Println("Created At: " + t[i].CreatedAt.Format(time.RFC850))
		if t[i].Completed {
			fmt.Println("Completed? Yes")
			fmt.Println("Completed At: " + t[i].CompletedAt.Format(time.RFC850))
		} else {
			fmt.Println("Completed? No")
		}
	}

}
