package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	toDoList := make([]string, 0)
	fmt.Println("Welcome to the TO DO List CLI app!")
	for {
		fmt.Println("\nEnter your command(create, read, update, delete):")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if !(userInput == "stop") {
			switch userInput {
			case "create":
				fmt.Println("Enter task name:")
				toDoList = create(toDoList, reader)
			case "read":
				read(toDoList)
			case "update":
				toDoList = update(toDoList, reader)
			case "delete":
				toDoList = delete(toDoList, reader)
			default:
				fmt.Println("Invalid command! Please try again!")
			}
		} else {
			read(toDoList)
			break
		}
	}

}

func create(list []string, reader *bufio.Reader) []string {
	userTask, _ := reader.ReadString('\n')
	userTask = strings.TrimSpace(userTask)
	list = append(list, userTask)
	return list
}

func read(list []string) {
	for i, v := range list {
		fmt.Printf("%v. %v\n", i+1, v)
	}
}

func update(list []string, reader *bufio.Reader) []string {
	for {	
	fmt.Println("Enter task name to update:")
	userTaskToUpdate, _ := reader.ReadString('\n')
	userTaskToUpdate = strings.TrimSpace(userTaskToUpdate)
	isFound := false
	for i, v := range list {
		if strings.TrimSpace(v) == userTaskToUpdate {
			fmt.Println("Enter new task name")
			userTaskUpdated, _ := reader.ReadString('\n')
			userTaskUpdated = strings.TrimSpace(userTaskUpdated)
			list[i] = userTaskUpdated
			isFound = true
			break
		} 
	}
	if !isFound {
		fmt.Println("There is no such task. Try again!")
		continue
	}
	break
	}

	return list
}

func delete(list []string, reader *bufio.Reader) []string {
	for {
		fmt.Println("Enter task name to remove:")
		taskToRemove, _ := reader.ReadString('\n')
		taskToRemove = strings.TrimSpace(taskToRemove)
		isFound := false
		for i, v := range list {
			if strings.TrimSpace(v) == taskToRemove {
				list = append(list[:i], list[i+1:]...)
				fmt.Printf("Removed task #%v with name %v successfully!\n", i+1, v)
				isFound = true
				break
			}
		}
		if !isFound {
			fmt.Println("There is no such task. Try again!")
			continue
		}
		break
	}
	return list
}
