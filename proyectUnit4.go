package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
	taskType    string
}

type TasksList struct {
	tasks []Task
}

// Add new tasks to the TasksList
func (t *TasksList) addNewTask(task Task) {
	t.tasks = append(t.tasks, task)
}

// Update Task
func (t *TasksList) updateTask(index int, task Task) {
	t.tasks[index] = task
}

// Delete Task
func (t *TasksList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

// To set completed Tasks
func (t *TasksList) setCompleted(index int) {
	t.tasks[index].completed = true
}

func main() {

	list := TasksList{}

	//Instance to introduce data for a structure
	read := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Select one of the following options:\n",
			"1. Add a new task \n",
			"2. Set a task as completed \n",
			"3. Edit task \n",
			"4. Delete task \n",
			"5. Exit ")
		fmt.Print("Enter one option :")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task
			fmt.Print("Enter task name: ")
			t.name, _ = read.ReadString('\n')
			fmt.Print("Enter task description: ")
			t.description, _ = read.ReadString('\n')
			fmt.Print("Enter task type: ")
			t.taskType, _ = read.ReadString('\n')
			list.addNewTask(t)
			fmt.Println("Task added to list correctly")
		case 2:
			var index int
			fmt.Print("Enter task index: ")
			fmt.Scanln(&index)
			list.setCompleted(index)
			fmt.Println("Task updated to completed ")
		case 3:
			var index int
			var t Task
			fmt.Print("Enter task index: ")
			fmt.Scanln(&index)
			fmt.Print("Enter task name: ")
			t.name, _ = read.ReadString('\n')
			fmt.Print("Enter task description: ")
			t.description, _ = read.ReadString('\n')
			fmt.Print("Enter task type: ")
			t.taskType, _ = read.ReadString('\n')
			list.updateTask(index, t)
			fmt.Println("Task updated to completed ")
		case 4:
			var index int
			fmt.Print("Enter task index: ")
			fmt.Scanln(&index)
			list.deleteTask(index)
			fmt.Println("Task deleted from list correctly")
		case 5:
			fmt.Print("Exit program")
			return
		default:
			fmt.Println("Unrecognized option!")
		}
		//Make a list for the Tasks
		fmt.Println("List of tasks:")
		fmt.Println("====================================================================")
		for i, task := range list.tasks {
			fmt.Printf("%d. Description: %s - %s - Complete : %t \n", i, task.name, task.description, task.completed)
		}
	}

}
