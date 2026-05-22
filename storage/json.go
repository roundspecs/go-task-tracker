package storage

import (
	"encoding/json"
	"fmt"
	"go-task-tracker/task"
	"os"
	"path/filepath"
)

type dataWrapper struct {
	NextID int         `json:"next_id"`
	Tasks  []task.Task `json:"tasks"`
}

type JSONStorage struct {
	filename string
}

func NewJSONStorage(filename string) (*JSONStorage, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not find home directory %w", err)
	}
	storagePath := filepath.Join(homePath, filename)

	return &JSONStorage{filename: storagePath}, nil
}

func (s *JSONStorage) Save(tasks []task.Task, nextID int) error {
	data := dataWrapper{
		NextID: nextID,
		Tasks:  tasks,
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("serializing data: %w", err)
	}

	err = os.WriteFile(s.filename, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("writing task file: %w", err)
	}
	return nil
}

func (s *JSONStorage) Load() (tasks []task.Task, nextID int, err error) {
	jsonBytes, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, 1, nil
		}
		return nil, 0, fmt.Errorf("reading task file: %w", err)
	}

	var data dataWrapper
	if err = json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, 0, fmt.Errorf("parsing task file: %w", err)
	}

	return data.Tasks, data.NextID, nil
}
