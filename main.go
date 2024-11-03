package main

import (
	"console-todo/pkg"
	"fmt"
	"log"
	"strconv"
)

var (
	addTaskCmd    = "Добавить задачу"
	showTasksCmd  = "Показать все задачи"
	removeTaskCmd = "Удалить задачу"
	unknownCmd    = "Неизвестная команда"
)

func main() {
	storage := pkg.Storage{}
	for {
		fmt.Println("\nМеню:")
		fmt.Printf("\n1.%s:", addTaskCmd)
		fmt.Printf("\n2.%s:", showTasksCmd)
		fmt.Printf("\n3.%s:", removeTaskCmd)

		choise := pkg.GetUserInput("\nВыберите действие: ")

		switch choise {
		case showTasksCmd:
			showAllTasks(storage)
		case addTaskCmd:
			addTask(storage)
		case removeTaskCmd:
			removeTask(storage)
		default:
			fmt.Print(unknownCmd)
			return
		}
	}
}

func showAllTasks(s pkg.Storage) {

	tasks, err := s.AllTasks()
	if err != nil {
		log.Fatalf("error while get all tasks, %s: ", err.Error())
	}
	if len(tasks) == 0 {
		fmt.Printf("\nСписок задач пуст\n")
	}
	for _, task := range tasks {
		fmt.Printf("%d. %s: %s\n", task.ID, task.Title, task.Description)
	}
}

func addTask(s pkg.Storage) {
	task := pkg.Task{
		Done: false,
	}
	title := pkg.GetUserInput("\nНаименование:")
	task.Title = title
	description := pkg.GetUserInput("\nОписание:")
	task.Description = description

	s.AddTask(task)
}

func removeTask(s pkg.Storage) {

	id := pkg.GetUserInput("\nВведите номер задачи:")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Введите число")
	}
	s.RemoveTask((idNum))
}
