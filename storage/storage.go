package storage

import (
	"encoding/json"
	"os"
	"task-manager-cli/task"
)

type Storage struct {
	filePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{filePath: filePath}
}

func (s *Storage) SaveTasks(tasks task.TaskList) error {
	// marshal means converts Go objects → JSON
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}

func (s *Storage) LoadTasks() (task.TaskList, error) {
	var tasks task.TaskList
	//read file here
	data, err := os.ReadFile(s.filePath)
	// check error
	if err != nil {
		// check if file doesn't exist
		if os.IsNotExist(err) {
			return task.TaskList{Tasks: []task.Task{}}, nil
		}

		return tasks, err
	}
	// unmarshal means converts JSON → Go objects
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
