package pkg

import (
	"encoding/json"
	"log"
	"os"
)

type Storage struct {
	Tasks []Task
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (s *Storage) AllTasks() ([]Task, error) {
	fName := "todo.txt"
	byteArray, err := os.ReadFile(fName)
	if err != nil {
		log.Fatalf("Невозможно прочитать файл: %s", err)
	}
	if len(byteArray) == 0 {
		return []Task{}, nil
	}
	var deserializedTasks []Task
	err = json.Unmarshal(byteArray, &deserializedTasks)
	if err != nil {
		log.Fatalf("Ошибка при десериализации: %s", err)
	}
	return deserializedTasks, nil
}

func (s *Storage) AddTask(newTask Task) error {
	tasks, err := s.AllTasks()
	if err != nil {
		return err
	}
	if len(tasks) > 0 {
		newTask.ID = len(tasks) + 1
	} else {
		newTask.ID = 1
	}
	tasks = append(tasks, newTask)
	if err = saveTasks(tasks); err != nil {
		log.Fatalf("Error while creating file: %s", err.Error())
	}
	return nil
}

func (s *Storage) RemoveTask(taskId int) error {
	tasks, err := s.AllTasks()
	if err != nil {
		return err
	}
	newTasks := make([]Task, 0, len(tasks)-1)
	for _, task := range tasks {
		if task.ID == taskId {
			continue
		}
		newTasks = append(newTasks, task)
	}
	for i := 0; i < len(newTasks); i++ {
		newTasks[i].ID = i + 1
	}
	if err = saveTasks(newTasks); err != nil {
		log.Fatalf("Error while creating file: %s", err.Error())
	}
	return nil
}

func saveTasks(t []Task) error {
	fName := "todo.txt"

	byteArray, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Ошибка при сериализации: %s", err)
	}
	if err := os.WriteFile(fName, byteArray, 0777); err != nil {
		log.Fatalf("Error while save file: %s", err.Error())
	}
	return nil
}
