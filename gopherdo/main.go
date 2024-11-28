package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct for tasks
type task struct {
	id       int
	name     string
	complete bool
	category string
	urgency  string
	dueDate  time.Time
}

// function will create task based off user input string
func create(taskList []task, input string, currentID int) []task {
	//defaults
	defaultname := ""
	defaultcomplete := false
	defaultcategory := "general"
	defaulturgency := "non-urgent"
	defaultdueDate := time.Time{}

	//split input on , to get separate commands
	cmdList := strings.Split(input, ",")
	//for all in cmd list split command on space
	for _, cmd := range cmdList {
		fmt.Println("cmd:", cmd)
		split := strings.Split(cmd, " ")
		fmt.Println("split", split)
		fmt.Println("split1:", split[1])
		//switch statement for first word/the command
		switch split[0] {
		case "create":
			defaultname = strings.Join(split[1:], " ")
		case "complete":
			//true if entered true
			defaultcomplete = (strings.Join(split[1:], " ") == "true")
		case "category":
			defaultcategory = strings.Join(split[1:], " ")
		case "urgency":
			defaultcategory = strings.Join(split[1:], " ")
		case "due":
			//parse user input for date
			var err error
			defaultdueDate, err = time.Parse(time.DateOnly, strings.Join(split[1:], " "))
			if err != nil {
			}
		}
	}
	//assign values. any values not found will be assigned default
	taskList = append(taskList, task{
		id:       currentID,
		name:     defaultname,
		complete: defaultcomplete,
		category: defaultcategory,
		urgency:  defaulturgency,
		dueDate:  defaultdueDate,
	})
	return taskList

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
			taskList = create(taskList, line, len(taskList)+1)
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
