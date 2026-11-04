package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Project struct {
	Name string
	Path string
}

var projects = []Project{
	{"Интернет-магазин", "1"},
	{"Банковская система", "2"},
	{"Гостиница", "3"},
	{"Игровая система персонажей", "4"},
	{"Электронный журнал", "5"},
	{"Библиотека", "6"},
	{"Доставка", "7"},
	{"Автосервис", "8"},
	{"Анализатор текста", "9"},
	{"Управление задачами", "10"},
	{"Склад", "11"},
	{"Кинотеатр", "12"},
	{"Социальная сеть", "13"},
	{"Ресторан", "14"},
	{"Учебный центр", "15"},
}

func runProject(project Project) {
	cmd := exec.Command("go", "run", "./"+project.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите задачу:")

		for i, project := range projects {
			fmt.Printf("%d. %s\n", i+1, project.Name)
		}

		fmt.Println("0. Запустить все")
		fmt.Println("-1. Выход")
		fmt.Print("Ваш выбор: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		if choice == -1 {
			break
		}

		if choice == 0 {
			for _, project := range projects {
				runProject(project)
			}
		} else if choice >= 1 && choice <= len(projects) {
			runProject(projects[choice-1])
		}
	}
}
