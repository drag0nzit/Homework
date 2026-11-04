package main

import (
	"fmt"
)

type Film struct {
	Title    string
	Duration int
	AgeLimit int
}

type Session struct {
	ID      int
	Film    Film
	Time    string
	Price   float64
	Seats   [][]bool
	Revenue float64
}

type Viewer struct {
	ID   int
	Name string
	Age  int
}

type Ticket struct {
	ViewerID  int
	SessionID int
	Row       int
	Seat      int
}

type Cinema struct {
	Sessions []Session
	Tickets  []Ticket
}

func (c *Cinema) AddSession(session Session) {
	c.Sessions = append(c.Sessions, session)
}

func (c *Cinema) FindSession(id int) *Session {
	for i := range c.Sessions {
		if c.Sessions[i].ID == id {
			return &c.Sessions[i]
		}
	}
	return nil
}

func (c *Cinema) BuyTicket(viewer Viewer, sessionID, row, seat int) {
	s := c.FindSession(sessionID)

	if s == nil {
		fmt.Println("Сеанс не найден")
		return
	}

	if viewer.Age < s.Film.AgeLimit {
		fmt.Println("Возрастное ограничение не позволяет купить билет")
		return
	}

	if row < 0 || row >= len(s.Seats) ||
		seat < 0 || seat >= len(s.Seats[row]) {
		fmt.Println("Неверное место")
		return
	}

	if s.Seats[row][seat] {
		fmt.Println("Место уже занято")
		return
	}

	s.Seats[row][seat] = true
	s.Revenue += s.Price

	c.Tickets = append(c.Tickets, Ticket{
		ViewerID:  viewer.ID,
		SessionID: sessionID,
		Row:       row,
		Seat:      seat,
	})

	fmt.Println("Билет куплен")
}

func (c *Cinema) CancelTicket(viewerID, sessionID, row, seat int) {
	s := c.FindSession(sessionID)

	if s == nil {
		return
	}

	for i, t := range c.Tickets {
		if t.ViewerID == viewerID &&
			t.SessionID == sessionID &&
			t.Row == row && t.Seat == seat {

			s.Seats[row][seat] = false
			s.Revenue -= s.Price

			c.Tickets = append(
				c.Tickets[:i],
				c.Tickets[i+1:]...,
			)

			fmt.Println("Билет отменён")
			return
		}
	}
}

func (c *Cinema) ShowFreeSeats(sessionID int) {
	s := c.FindSession(sessionID)

	if s == nil {
		return
	}

	for i, row := range s.Seats {
		for j, occupied := range row {
			if !occupied {
				fmt.Printf("Ряд %d, место %d\n", i+1, j+1)
			}
		}
	}
}

func (c *Cinema) ShowRevenue() {
	for _, s := range c.Sessions {
		fmt.Println(s.Film.Title, "Выручка:", s.Revenue)
	}
}

func (c *Cinema) MostProfitable() {
	if len(c.Sessions) == 0 {
		return
	}

	best := c.Sessions[0]

	for _, s := range c.Sessions {
		if s.Revenue > best.Revenue {
			best = s
		}
	}

	fmt.Println("Самый прибыльный сеанс:", best.Film.Title)
	fmt.Println("Выручка:", best.Revenue)
}

func main() {
	cinema := Cinema{}

	seats := make([][]bool, 5)

	for i := range seats {
		seats[i] = make([]bool, 5)
	}

	cinema.AddSession(Session{
		ID:    1,
		Film:  Film{"Интерстеллар", 169, 12},
		Time:  "18:00",
		Price: 500,
		Seats: seats,
	})

	viewer := Viewer{1, "Алексей", 20}

	cinema.BuyTicket(viewer, 1, 0, 0)
	cinema.BuyTicket(viewer, 1, 0, 1)

	fmt.Println("\nСвободные места:")
	cinema.ShowFreeSeats(1)

	fmt.Println("\nВыручка:")
	cinema.ShowRevenue()

	fmt.Println("\nСамый прибыльный сеанс:")
	cinema.MostProfitable()

	fmt.Println("\nОтмена билета:")
	cinema.CancelTicket(1, 1, 0, 0)
}
