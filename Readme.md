Task Manager CLI

Task Manager CLI is a simple command-line tool for managing tasks. It supports adding, listing, and deleting tasks. Tasks are stored in a JSON file for persistence.

Features

	1.	Add Task: Create a new task with a title, description, and status.
	2.	List Tasks: View all tasks with their details.
	3.	Delete Task: Remove a task by its ID.

Project Structure

task-manager-cli/
├── cmd/
│   └── main.go         # Main entry point for the CLI
├── storage/
│   └── storage.go      # Handles reading and writing tasks to a JSON file
├── task/
│   └── task.go         # Defines the Task and TaskList structures
├── go.mod              # Go module file

Prerequisites

	•	Go 1.23.3 or later installed on your system.
	•	Basic understanding of the command line.

Installation

	1.	Clone the repository:


cd task-manager-cli


	2.	Initialize dependencies:

go mod tidy


	3.	Build the project:

go build -o task-manager-cli cmd/main.go

This will generate a binary executable named task-manager-cli.

Usage

Run the built CLI tool with the following subcommands:

1. Add a Task

./task-manager-cli add -t "Task Title" -d "Task Description"

	•	Options:
	•	-t: Title of the task (required).
	•	-d: Description of the task (optional).
	•	Example:

./task-manager-cli add -t "Learn Go" -d "Complete the CLI project"


	•	Output:

Task added successfully!



2. List Tasks

./task-manager-cli list

	•	Example Output:

ID: 1
Title: Learn Go
Description: Complete the CLI project
Status: pending
Created: 06 Dec 24 1:00 AM



3. Delete a Task

./task-manager-cli delete -id <task_id>

	•	Options:
	•	-id: ID of the task to delete (required).
	•	Example:

./task-manager-cli delete -id 1


	•	Output:

Task deleted successfully!

Configuration

By default, tasks are stored in a file named task.json. You can change the file path by modifying the following line in main.go:

store := storage.NewStorage("task.json")

Development Guide

	1.	Add a New Feature:
	•	Modify main.go to add a new subcommand or extend existing functionality.
	2.	Error Handling:
	•	Ensure all errors are properly logged and handled.
	3.	Testing:
	•	Manually run all commands to verify correctness.

Example Workflow

	1.	Add a task:

./task-manager-cli add -t "Buy groceries" -d "Milk, eggs, bread"


	2.	List all tasks:

./task-manager-cli list


	3.	Delete a task:

./task-manager-cli delete -id 1

Notes

	•	The ID of a task is auto-incremented and starts from 1.
	•	Tasks are persisted in the JSON file and survive application restarts.

