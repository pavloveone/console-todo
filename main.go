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
	updateTaskCmd = "Обновить задачу"
	unknownCmd    = "Неизвестная команда"
)

func main() {
	for {
		fmt.Println("\nМеню:")
		fmt.Printf("\n1.%s:", addTaskCmd)
		fmt.Printf("\n2.%s:", showTasksCmd)
		fmt.Printf("\n3.%s:", removeTaskCmd)
		fmt.Printf("\n4.%s:\n", updateTaskCmd)

		choise := pkg.GetUserInput("\nВыберите действие: ")

		switch choise {
		case showTasksCmd:
			showAllTasks()
		case addTaskCmd:
			createTask()
		case removeTaskCmd:
			remvoveTask()
		default:
			fmt.Print(unknownCmd)
			return
		}
	}
}

func showAllTasks() {
	storage := pkg.Storage{}
	tasks, err := storage.AllTasks()
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

func createTask() {
	storage := pkg.Storage{}
	task := pkg.Task{
		Done: false,
	}
	title := pkg.GetUserInput("\nНаименование:")
	task.Title = title
	description := pkg.GetUserInput("\nОписание:")
	task.Description = description

	storage.AddTask(task)
}

func remvoveTask() {
	storage := pkg.Storage{}

	id := pkg.GetUserInput("\nВведите номер задачи:")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Введите число")
	}
	storage.RemoveTask((idNum))
}
