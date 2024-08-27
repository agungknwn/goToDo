package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	crud "github.com/waksun0x00/goToDo/internal"
)

var taskIndex int
var input string
var task_input, newTask string

func main() {
	fmt.Println("Golang To do list Console")

	// var discard string

	fmt.Println("1. Add Task")
	fmt.Println("2. Delete Task")
	fmt.Println("3. Update Task")
	fmt.Println("4. Mark done Task")
	fmt.Println("5. Read Task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Options (1,2,3,4,5,6)")
		// _, err := fmt.Scanf("%d", &choice)

		// choice error handler
		// if err != nil {
		// 	fmt.Println("Invalid Input")

		// 	fmt.Scanln(&discard)
		// 	continue
		// }

		scanner.Scan()
		input = scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid Input")
			continue
		}

		if choice < 1 || choice > 6 {
			fmt.Println("invalid Choice")
			continue
		} else if choice == 6 {
			fmt.Println("Bye !!!")
			break
		}

		// crud framework switchcase
		switch choice {
		case 1:
			fmt.Println("Task Description : ")
			scanner.Scan()
			task_input = scanner.Text()
			crud.Add_task(task_input)
		case 2:
			fmt.Println("Task Index to be removed : ")
			indexScan()
			crud.Delete_task(taskIndex)
		case 3:
			fmt.Println("Task Index to be edited : ")
			indexScan()
			fmt.Println("please enter new task : ")
			scanner.Scan()
			newTask = scanner.Text()
			// fmt.Scanln(&newTask)
			crud.Update_task(taskIndex, newTask)
		case 4:
			fmt.Println("Task index to be mark done : ")
			// fmt.Scanln(&taskIndex)
			indexScan()
			crud.Mark_done(taskIndex)
		case 5:
			crud.Read_task()
		}
	}

}

func indexScan() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input = scanner.Text()

	taskIndex, _ = strconv.Atoi(input)
}
