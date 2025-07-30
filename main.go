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
	fmt.Println("----------------------")
	fmt.Println("TodoPlus List")
	fmt.Println("----------------------")
	for i, task := range tasks {
		var status = "[N]"
		if task.completed {
			status = "[D]"
		}
		if status == "[D]" {
			fmt.Printf("%d. %s %s\n", i+1, color.Green(task.TaskName), color.Green(status))
		} else {
			fmt.Printf("%d. %s %s\n", i+1, task.TaskName, color.Red(status))
		}
	}
	fmt.Println("----------------------")
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
	fmt.Println("0: Exit")
	fmt.Println("1: Add")
	fmt.Println("2: List")
	fmt.Println("3: Completed")
	fmt.Println("4: Edit")
	fmt.Println("5: Delete")
	fmt.Println("6: Help")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		listTasks()
		fmt.Print("Enter your action (0-6): ")
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
			fmt.Println("0. Exit")
			fmt.Println("1. Add Task")
			fmt.Println("2. List Tasks")
			fmt.Println("3. Mark Task Completed")
			fmt.Println("4. Edit Task")
			fmt.Println("5. Delete Task")
			fmt.Println("6. Help")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		fmt.Print("\033[H\033[2J")
	}
}
