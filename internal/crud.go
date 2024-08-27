package crud

import (
	"fmt"
)

type tasks struct {
	desc      string
	completed bool
}

// todolist slices with tasks type struct
var todoList []tasks

func Add_task(task string) {
	newTask := tasks{task, false}
	todoList = append(todoList, newTask)

	fmt.Println("\nNew Task added")

}

func Delete_task(index int) {
	if index >= 1 && index <= len(todoList) {
		todoList = append(todoList[:index-1], todoList[index:]...)
		fmt.Println("\nTask removed")
	} else {
		fmt.Println("invalid Index")
	}
}

func Read_task() {
	fmt.Println("\nYour to do list is here :")
	for i, task := range todoList {
		status := "on Going"
		if task.completed {
			status = "done"
		}

		fmt.Printf("%d. %v [%v]\n", i+1, task.desc, status)
	}
}

func Update_task(index int, newTask string) {
	if index >= 1 && index <= len(todoList) {
		todoList[index-1].desc = newTask
		fmt.Println("\nTask succesfully updated")
	} else {
		fmt.Println("invalid Index")
	}
}

func Mark_done(index int) {
	if index >= 1 && index <= len(todoList) {
		todoList[index-1].completed = true
		fmt.Println("\nTask marked as done")
	} else {
		fmt.Println("invalid Index")
	}
}
