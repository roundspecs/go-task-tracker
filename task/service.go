package task

import (
	"fmt"
	"time"
)

type Storage interface {
	Save(tasks []Task, nextID int) error
	Load() (tasks []Task, nextID int, err error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Add(description string) (Task, error) {
	tasks, nextID, err := s.storage.Load()
	if err != nil {
		return Task{}, fmt.Errorf("loading data: %w", err)
	}

	task := Task{
		ID:          nextID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	err = s.storage.Save(tasks, nextID+1)
	if err != nil {
		return Task{}, fmt.Errorf("saving data: %w", err)
	}

	return task, nil
}
