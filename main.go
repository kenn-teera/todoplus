package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kenn-teera/todoplus/color"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task

func addTask(taskName string) {
	newtask := Task{TaskName: taskName, completed: false}
	tasks = append(tasks, newtask)

	fmt.Println("Task Added")
}

func listTasks() {
	for i, task := range tasks {
		var status = "[N]"
		if task.completed {
			status = "[D]"
		}
		fmt.Println("----------------------")
		if status == "[D]" {
			fmt.Printf("%d. %s %s\n", i+1, color.Green(task.TaskName), color.Green(status))
		} else {
			fmt.Printf("%d. %s %s\n", i+1, task.TaskName, color.Red(status))
		}
		fmt.Println("----------------------")
	}
}

func markCompleted(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}
	tasks[index].completed = true
	fmt.Println("Task marked as completed")
}

func editTask(index int, newString string) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}
	tasks[index].TaskName = newString
	fmt.Println("Task edited successfully")
}

func deleteTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task deleted")
}

func main() {

	var indexInput int
	var taskInput string
	var newTask string

	fmt.Println("Welcome to TODOPlus")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your choice: ")
		var choice int
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter task index to mark as completed: ")
			scanner.Scan()
			indexInput, _ := strconv.Atoi(scanner.Text())
			markCompleted(indexInput - 1)
		case 4:
			fmt.Print("Enter task index to edit: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter new task name: ")
			scanner.Scan()
			newTask = scanner.Text()
			editTask(indexInput-1, newTask)
		case 5:
			fmt.Print("Enter task index to delete: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput - 1)
		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 0:
			fmt.Println("0. Help")
			fmt.Println("1. Add Task")
			fmt.Println("2. List Tasks")
			fmt.Println("3. Mark Task Completed")
			fmt.Println("4. Edit Task")
			fmt.Println("5. Delete Task")
			fmt.Println("6. Exit")
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
