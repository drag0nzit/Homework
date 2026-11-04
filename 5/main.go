package main

import (
	"fmt"
	"sort"
)

type Student struct {
	ID     int
	Name   string
	Age    int
	Grades map[string][]int
}

type Teacher struct {
	ID   int
	Name string
}

type Subject struct {
	Name string
}

type Journal struct {
	Students []Student
	Teachers []Teacher
	Subjects []Subject
}

func (j *Journal) AddStudent(id int, name string, age int) {
	j.Students = append(j.Students, Student{
		ID:     id,
		Name:   name,
		Age:    age,
		Grades: make(map[string][]int),
	})
}

func (j *Journal) AddSubject(name string) {
	j.Subjects = append(j.Subjects, Subject{name})
}

func (j *Journal) FindStudent(id int) *Student {
	for i := range j.Students {
		if j.Students[i].ID == id {
			return &j.Students[i]
		}
	}
	return nil
}

func (j *Journal) AddGrade(id int, subject string, grade int) {
	s := j.FindStudent(id)

	if s == nil || grade < 2 || grade > 5 {
		return
	}

	s.Grades[subject] = append(s.Grades[subject], grade)
}

func (j *Journal) ChangeGrade(id int, subject string, index, grade int) {
	s := j.FindStudent(id)

	if s == nil || grade < 2 || grade > 5 {
		return
	}

	if index >= 0 && index < len(s.Grades[subject]) {
		s.Grades[subject][index] = grade
	}
}

func (j *Journal) DeleteGrade(id int, subject string, index int) {
	s := j.FindStudent(id)

	if s == nil {
		return
	}

	grades := s.Grades[subject]

	if index >= 0 && index < len(grades) {
		s.Grades[subject] = append(
			grades[:index],
			grades[index+1:]...,
		)
	}
}

func AverageStudent(s Student) float64 {
	var sum, count int

	for _, grades := range s.Grades {
		for _, grade := range grades {
			sum += grade
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func (j *Journal) AverageSubject(subject string) float64 {
	var sum, count int

	for _, s := range j.Students {
		for _, grade := range s.Grades[subject] {
			sum += grade
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func (j *Journal) BestAndWorst() {
	if len(j.Students) == 0 {
		return
	}

	best, worst := j.Students[0], j.Students[0]

	for _, s := range j.Students {
		if AverageStudent(s) > AverageStudent(best) {
			best = s
		}

		if AverageStudent(s) < AverageStudent(worst) {
			worst = s
		}
	}

	fmt.Println("Лучший студент:", best.Name,
		AverageStudent(best))

	fmt.Println("Худший студент:", worst.Name,
		AverageStudent(worst))
}

func (j *Journal) SortStudents() {
	sort.Slice(j.Students, func(i, k int) bool {
		return AverageStudent(j.Students[i]) >
			AverageStudent(j.Students[k])
	})
}

func (j *Journal) ShowBadStudents() {
	for _, s := range j.Students {
		bad := false

		for _, grades := range s.Grades {
			for _, grade := range grades {
				if grade < 3 {
					bad = true
				}
			}
		}

		if bad {
			fmt.Println("Неудовлетворительная оценка:", s.Name)
		}
	}
}

func (j *Journal) ShowStudents() {
	for _, s := range j.Students {
		fmt.Println("\nСтудент:", s.Name)
		fmt.Println("Средний балл:", AverageStudent(s))

		for subject, grades := range s.Grades {
			fmt.Println(subject, grades)
		}
	}
}

func main() {
	journal := Journal{}

	journal.AddStudent(1, "Алексей", 20)
	journal.AddStudent(2, "Иван", 19)
	journal.AddStudent(3, "Пётр", 21)

	journal.AddSubject("Математика")
	journal.AddSubject("Программирование")

	journal.AddGrade(1, "Математика", 5)
	journal.AddGrade(1, "Программирование", 4)

	journal.AddGrade(2, "Математика", 3)
	journal.AddGrade(2, "Программирование", 2)

	journal.AddGrade(3, "Математика", 4)
	journal.AddGrade(3, "Программирование", 5)

	fmt.Println("Электронный журнал:")
	journal.ShowStudents()

	fmt.Println("\nСредний балл по математике:",
		journal.AverageSubject("Математика"))

	fmt.Println("\nЛучший и худший:")
	journal.BestAndWorst()

	fmt.Println("\nСтуденты с плохими оценками:")
	journal.ShowBadStudents()

	fmt.Println("\nСортировка:")
	journal.SortStudents()
	journal.ShowStudents()
}
