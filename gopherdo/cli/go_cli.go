package go_cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// struct for tasks
type task struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Complete     bool      `json:"complete"`
	Category     string    `json:"category"`
	Urgency      string    `json:"urgency"`
	DueDate      time.Time `json:"dueDate"`
	CreationDate time.Time `json:"creationDate"`
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
		ID:           currentID,
		Name:         defaultname,
		Complete:     defaultcomplete,
		Category:     defaultcategory,
		Urgency:      defaulturgency,
		DueDate:      defaultdueDate,
		CreationDate: time.Now(),
	})
	return tasks
}

func delete(tasks []task, searchID int) []task {
	for i, task := range tasks {
		if task.ID == searchID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks
		}
	}
	return tasks
}

func list(tasks []task) {
	fmt.Println("Task List:")
	for _, t := range tasks {
		fmt.Printf("ID: %d, Name: %s, Completed: %t, Category: %s, Urgency: %s, Due date: %v\n", t.ID, t.Name, t.Complete, t.Category, t.Urgency, t.DueDate)
	}
}

func SaveTasks(tasks []task) error {
	file, err := os.Create("saved_tasks.json")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error encoding tasks: %v", err)
	}

	return nil
}

func complete(tasks []task, searchID int) []task {
	for i, task := range tasks {
		if task.ID == searchID {
			tasks[i].Complete = true
			return tasks
		}
	}
	return tasks
}

func LoadTasks() ([]task, error) {
	// Check if file exists
	if _, err := os.Stat("saved_tasks.json"); os.IsNotExist(err) {
		return make([]task, 0), nil // Return empty slice if file doesn't exist
	}

	// Open the file
	file, err := os.Open("saved_tasks.json")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a slice to store the tasks
	var tasks []task

	// Create a decoder and decode into tasks slice
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("error decoding tasks: %v", err)
	}

	return tasks, nil
}

func main() {
	//list commands available
	fmt.Println("Commands: list, create, delete, complete, quit, save")

	//split for todo tasks
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		tasks = make([]task, 0) // Create empty slice if loading fails
	}

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
		case "quit":
			SaveTasks(tasks)
			os.Exit(0)
		case "save":
			SaveTasks(tasks)
		}
	}
}
