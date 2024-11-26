package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	id       int
	name     string
	complete bool
}

func main() {
	//list commands available
	fmt.Println("Commands: list, create, delete, quit")
	//declare for user input
	//split for todo tasks
	taskList := make([]task, 0)
	//scanner for terminal input
	scanner := bufio.NewScanner(os.Stdin)
	//loop for user input
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {
		case "create":
			taskList = append(taskList, task{
				id:       (len(taskList) + 1),
				name:     strings.Join(split[1:], " "),
				complete: false,
			})
		case "delete":
			for i, task := range taskList {
				if task.name == strings.Join(split[1:], " ") {
					taskList = append(taskList[:i], taskList[i+1:]...)
				}
			}
		case "list":
			fmt.Println("Task List:")
			for _, t := range taskList {
				fmt.Printf("ID: %d, Name: %s, Completed: %t\n", t.id, t.name, t.complete)
			}
		case "complete":
			for i, task := range taskList {
				if task.name == strings.Join(split[1:], " ") {
					taskList[i].complete = true
				}
			}
		}
	}
}
