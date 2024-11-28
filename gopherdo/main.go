package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// struct for tasks
type task struct {
	id           int
	name         string
	complete     bool
	category     string
	urgency      string
	dueDate      time.Time
	creationDate time.Time
}

// function will create task based off user input string
func create(tasks []task, input string, currentID int) []task {
	//defaults
	defaultname := ""
	defaultcomplete := false
	defaultcategory := "general"
	defaulturgency := "non-urgent"
	defaultdueDate := time.Time{}
	//split input on , to get separate commands
	cmdList := strings.Split(input, ";")
	//for all in cmd list split command on space
	for _, cmd := range cmdList {
		split := strings.Split(cmd, " ")
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
			defaulturgency = strings.Join(split[1:], " ")
		case "due":
			//parse user input for date
			var err error
			defaultdueDate, err = time.Parse(time.DateOnly, strings.Join(split[1:], " "))
			if err != nil {
				fmt.Println("oops")
			}
		}
	}
	//assign values. any values not found will be assigned default
	tasks = append(tasks, task{
		id:           currentID,
		name:         defaultname,
		complete:     defaultcomplete,
		category:     defaultcategory,
		urgency:      defaulturgency,
		dueDate:      defaultdueDate,
		creationDate: time.Now(),
	})
	return tasks
}

func delete(tasks []task, searchID int) []task {
	for i, task := range tasks {
		if task.id == searchID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks
		}
	}
	return tasks
}

func list(tasks []task) {
	fmt.Println("Task List:")
	for _, t := range tasks {
		fmt.Printf("ID: %d, Name: %s, Completed: %t, Category: %s, Urgency: %s, Due date: %v\n", t.id, t.name, t.complete, t.category, t.urgency, t.dueDate)
	}
}

func complete(tasks []task, searchID int) []task {
	for i, task := range tasks {
		if task.id == searchID {
			tasks[i].complete = true
			return tasks
		}
	}
	return tasks
}

func main() {
	//list commands available
	fmt.Println("Commands: list, create, delete, complete, quit")
	//declare for user input
	//split for todo tasks
	tasks := make([]task, 0)
	//scanner for terminal input
	scanner := bufio.NewScanner(os.Stdin)
	//loop for user input
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {
		case "create":
			tasks = create(tasks, line, len(tasks)+1)
		case "delete":
			//convert number part of input to int so can be passed as id
			searchID, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			tasks = delete(tasks, searchID)
		case "list":
			list(tasks)
		case "complete":
			//convert number part of input to int so can be passed as id
			searchID, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			tasks = complete(tasks, searchID)
		}
	}
}
