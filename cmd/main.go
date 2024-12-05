package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"task-manager-cli/storage"
	"task-manager-cli/task"
	"time"
)

func main() {

	fmt.Println("task manager cli")
	add := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := add.String("t", "", "title of task")
	addDescription := add.String("d", "", "description of task")

	list := flag.NewFlagSet("list", flag.ExitOnError)

	delete := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := delete.Int("id", 0, "Task ID to delete")

	store := storage.NewStorage("task.json")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])
		if *addTitle == "" {
			log.Fatal("title is required")
		}
		tasks, err := store.LoadTasks()
		if err != nil {
			log.Fatal(err)
		}
		newTask := task.Task{
			ID:          len(tasks.Tasks) + 1,
			Title:       *addTitle,
			Description: *addDescription,
			Status:      "pending",
			CreatedAt:   time.Now(),
		}
		tasks.Tasks = append(tasks.Tasks, newTask)
		if err := store.SaveTasks(tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task added successfully!")

	case "list":
		list.Parse(os.Args[2:])
		tasks, err := store.LoadTasks()
		if err != nil {
			log.Fatal(err)
		}
		for _, t := range tasks.Tasks {
			fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreated: %s\n\n",
				t.ID, t.Title, t.Description, t.Status, t.CreatedAt.Format(time.RFC822))
		}

	case "delete":
		delete.Parse(os.Args[2:])
		if *deleteId == 0 {
			log.Fatal("id is required")
		}

		tasks, err := store.LoadTasks()
		if err != nil {
			log.Fatal(err)
		}
		var newTasks []task.Task
		for _, t := range tasks.Tasks {
			if t.ID != *deleteId {
				newTasks = append(newTasks, t)
			}
		}

		tasks.Tasks = newTasks
		if err := store.SaveTasks(tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task deleted successfully!")
	}

}
