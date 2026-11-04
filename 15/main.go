package main

import (
	"fmt"
	"sort"
	"strings"
)

type Student struct {
	ID      int
	Name    string
	Age     int
	Courses []int
}

type Teacher struct {
	ID     int
	Name   string
	Groups []int
}

type Course struct {
	ID       int
	Name     string
	Duration int
	Price    float64
}

type Group struct {
	ID        int
	Name      string
	CourseID  int
	Students  []int
	TeacherID int
}

type CourseStatistics struct {
	CourseName string
	Students   int
	AverageAge float64
	TotalPrice float64
}

type Center struct {
	Students map[int]Student
	Teachers map[int]Teacher
	Courses  map[int]Course
	Groups   map[int]Group
}

func NewCenter() *Center {
	return &Center{
		Students: make(map[int]Student),
		Teachers: make(map[int]Teacher),
		Courses:  make(map[int]Course),
		Groups:   make(map[int]Group),
	}
}

func (c *Center) AddStudent(student Student) {
	c.Students[student.ID] = student
}

func (c *Center) AddTeacher(teacher Teacher) {
	c.Teachers[teacher.ID] = teacher
}

func (c *Center) AddCourse(course Course) {
	c.Courses[course.ID] = course
}

func (c *Center) AddGroup(group Group) {
	c.Groups[group.ID] = group
}

func (c *Center) AddStudentToGroup(studentID, groupID int) bool {
	student, studentExists := c.Students[studentID]
	group, groupExists := c.Groups[groupID]

	if !studentExists || !groupExists {
		return false
	}

	if _, exists := c.Courses[group.CourseID]; !exists {
		return false
	}

	for _, id := range group.Students {
		if id == studentID {
			return false
		}
	}

	group.Students = append(group.Students, studentID)
	student.Courses = append(student.Courses, group.CourseID)

	c.Groups[groupID] = group
	c.Students[studentID] = student

	return true
}

func (c *Center) AssignTeacher(teacherID, groupID int) bool {
	teacher, teacherExists := c.Teachers[teacherID]
	group, groupExists := c.Groups[groupID]

	if !teacherExists || !groupExists {
		return false
	}

	if group.TeacherID != 0 && group.TeacherID != teacherID {
		oldTeacher := c.Teachers[group.TeacherID]
		oldTeacher.Groups = removeID(oldTeacher.Groups, groupID)
		c.Teachers[group.TeacherID] = oldTeacher
	}

	for _, id := range teacher.Groups {
		if id == groupID {
			group.TeacherID = teacherID
			c.Groups[groupID] = group
			return true
		}
	}

	teacher.Groups = append(teacher.Groups, groupID)
	group.TeacherID = teacherID

	c.Teachers[teacherID] = teacher
	c.Groups[groupID] = group

	return true
}

func (c *Center) SearchStudent(name string) []Student {
	var result []Student

	for _, student := range c.Students {
		if strings.Contains(
			strings.ToLower(student.Name),
			strings.ToLower(name),
		) {
			result = append(result, student)
		}
	}

	return result
}

func (c *Center) SearchGroup(name string) []Group {
	var result []Group

	for _, group := range c.Groups {
		if strings.Contains(
			strings.ToLower(group.Name),
			strings.ToLower(name),
		) {
			result = append(result, group)
		}
	}

	return result
}

func (c *Center) SortByName() []Student {
	result := make([]Student, 0, len(c.Students))

	for _, student := range c.Students {
		result = append(result, student)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}

func (c *Center) SortByAge() []Student {
	result := make([]Student, 0, len(c.Students))

	for _, student := range c.Students {
		result = append(result, student)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Age < result[j].Age
	})

	return result
}

func (c *Center) CountStudents() int {
	return len(c.Students)
}

func (c *Center) Statistics() []CourseStatistics {
	var result []CourseStatistics

	for _, course := range c.Courses {
		var totalAge float64
		var studentCount int

		for _, student := range c.Students {
			for _, courseID := range student.Courses {
				if courseID == course.ID {
					totalAge += float64(student.Age)
					studentCount++
					break
				}
			}
		}

		var averageAge float64

		if studentCount > 0 {
			averageAge = totalAge / float64(studentCount)
		}

		result = append(result, CourseStatistics{
			CourseName: course.Name,
			Students:   studentCount,
			AverageAge: averageAge,
			TotalPrice: float64(studentCount) * course.Price,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CourseName < result[j].CourseName
	})

	return result
}

func removeID(ids []int, target int) []int {
	result := make([]int, 0, len(ids))

	for _, id := range ids {
		if id != target {
			result = append(result, id)
		}
	}

	return result
}

type Payment interface {
	Pay(amount float64) bool
}

type CashPayment struct{}

func (p CashPayment) Pay(amount float64) bool {
	fmt.Printf("Оплата наличными: %.2f руб.\n", amount)
	return true
}

type CardPayment struct{}

func (p CardPayment) Pay(amount float64) bool {
	fmt.Printf("Оплата картой: %.2f руб.\n", amount)
	return true
}

type OnlinePayment struct{}

func (p OnlinePayment) Pay(amount float64) bool {
	fmt.Printf("Онлайн-оплата: %.2f руб.\n", amount)
	return true
}

func (c *Center) PayCourse(studentID, courseID int, payment Payment) bool {
	_, studentExists := c.Students[studentID]
	course, courseExists := c.Courses[courseID]

	if !studentExists || !courseExists {
		fmt.Println("Студент или курс не найден")
		return false
	}

	return payment.Pay(course.Price)
}

func main() {
	center := NewCenter()

	center.AddStudent(Student{ID: 1, Name: "Иван Иванов", Age: 18})
	center.AddStudent(Student{ID: 2, Name: "Пётр Петров", Age: 20})
	center.AddStudent(Student{ID: 3, Name: "Анна Смирнова", Age: 19})

	center.AddTeacher(Teacher{ID: 1, Name: "Алексей Сергеев"})
	center.AddTeacher(Teacher{ID: 2, Name: "Мария Кузнецова"})

	center.AddCourse(Course{
		ID:       1,
		Name:     "Программирование",
		Duration: 120,
		Price:    15000,
	})

	center.AddCourse(Course{
		ID:       2,
		Name:     "Электроника",
		Duration: 90,
		Price:    12000,
	})

	center.AddGroup(Group{
		ID:       1,
		Name:     "Группа А",
		CourseID: 1,
	})

	center.AddGroup(Group{
		ID:       2,
		Name:     "Группа Б",
		CourseID: 2,
	})

	center.AddStudentToGroup(1, 1)
	center.AddStudentToGroup(2, 1)
	center.AddStudentToGroup(3, 2)

	center.AssignTeacher(1, 1)
	center.AssignTeacher(2, 2)

	fmt.Println("=== Студенты по имени ===")
	for _, student := range center.SortByName() {
		fmt.Println(student.Name)
	}

	fmt.Println("\n=== Студенты по возрасту ===")
	for _, student := range center.SortByAge() {
		fmt.Println(student.Name, "-", student.Age)
	}

	fmt.Println("\n=== Поиск студента ===")
	for _, student := range center.SearchStudent("Иван") {
		fmt.Println(student.ID, student.Name)
	}

	fmt.Println("\n=== Поиск группы ===")
	for _, group := range center.SearchGroup("А") {
		fmt.Println(group.ID, group.Name)
	}

	fmt.Println("\n=== Количество студентов ===")
	fmt.Println(center.CountStudents())

	fmt.Println("\n=== Статистика курсов ===")
	for _, stat := range center.Statistics() {
		fmt.Printf(
			"Курс: %s | Студентов: %d | Средний возраст: %.1f | Общая стоимость: %.2f руб.\n",
			stat.CourseName,
			stat.Students,
			stat.AverageAge,
			stat.TotalPrice,
		)
	}

	fmt.Println("\n=== Оплата курса ===")
	center.PayCourse(1, 1, CardPayment{})
	center.PayCourse(2, 1, CashPayment{})
	center.PayCourse(3, 2, OnlinePayment{})
}
