package main

import (
	"fmt"
	"sort"
	"strings"
)

type Task struct {
	ID          int
	Name        string
	Description string
	Priority    int
	Status      string
}

type TaskManager struct {
	Tasks map[int]Task
}

func (tm *TaskManager) AddTask(task Task) {
	if task.ID <= 0 || task.Name == "" ||
		task.Priority < 1 || task.Priority > 5 {
		fmt.Println("Некорректные данные задачи")
		return
	}

	if _, exists := tm.Tasks[task.ID]; exists {
		fmt.Println("Задача с таким ID уже существует")
		return
	}

	task.Status = "Новая"
	tm.Tasks[task.ID] = task
}

func (tm *TaskManager) ChangeTask(id int, name, description string, priority int) {
	task, exists := tm.Tasks[id]

	if !exists {
		return
	}

	if priority < 1 || priority > 5 {
		return
	}

	task.Name = name
	task.Description = description
	task.Priority = priority

	tm.Tasks[id] = task
}

func (tm *TaskManager) DeleteTask(id int) {
	delete(tm.Tasks, id)
}

func (tm *TaskManager) ChangeStatus(id int, status string) {
	task, exists := tm.Tasks[id]

	if !exists {
		return
	}

	if status != "Новая" &&
		status != "В работе" &&
		status != "Выполнена" {
		fmt.Println("Неизвестный статус")
		return
	}

	task.Status = status
	tm.Tasks[id] = task
}

func (tm *TaskManager) Search(name string) {
	for _, task := range tm.Tasks {
		if strings.Contains(
			strings.ToLower(task.Name),
			strings.ToLower(name),
		) {
			fmt.Println(task.ID, task.Name, task.Status)
		}
	}
}

func (tm *TaskManager) Filter(status string) {
	for _, task := range tm.Tasks {
		if task.Status == status {
			fmt.Println(task.ID, task.Name)
		}
	}
}

func (tm *TaskManager) SortTasks() []Task {
	var result []Task

	for _, task := range tm.Tasks {
		result = append(result, task)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Priority > result[j].Priority
	})

	return result
}

func (tm *TaskManager) Statistics() {
	newTasks := 0
	inProgress := 0
	completed := 0

	for _, task := range tm.Tasks {
		switch task.Status {
		case "Новая":
			newTasks++
		case "В работе":
			inProgress++
		case "Выполнена":
			completed++
		}
	}

	fmt.Println("Новых задач:", newTasks)
	fmt.Println("В работе:", inProgress)
	fmt.Println("Выполнено:", completed)
}

func main() {
	manager := TaskManager{
		Tasks: make(map[int]Task),
	}

	manager.AddTask(Task{1, "Изучить Go", "Основы языка", 3, ""})
	manager.AddTask(Task{2, "Сделать проект", "Написать программу", 5, ""})
	manager.AddTask(Task{3, "Прочитать книгу", "Изучить документацию", 2, ""})

	manager.ChangeStatus(1, "В работе")
	manager.ChangeStatus(2, "Выполнена")

	fmt.Println("Поиск задач:")
	manager.Search("Go")

	fmt.Println("\nЗадачи в работе:")
	manager.Filter("В работе")

	fmt.Println("\nСортировка по приоритету:")
	for _, task := range manager.SortTasks() {
		fmt.Println(task.ID, task.Name, task.Priority)
	}

	fmt.Println("\nСтатистика:")
	manager.Statistics()
}
