package main

import (
	"fmt"
	"sort"
)

type Room struct {
	Number int
	Places int
	Price  float64
	Busy   bool
	Type   string
}

type Guest struct {
	ID   int
	Name string
	Age  int
}

type Booking struct {
	RoomNumber int
	GuestID    int
	Days       int
}

type Pricing interface {
	Calculate(days int) float64
}

type Standard struct {
	Price float64
}

func (s Standard) Calculate(days int) float64 {
	return s.Price * float64(days)
}

type Family struct {
	Price float64
}

func (f Family) Calculate(days int) float64 {
	return f.Price * float64(days) * 1.2
}

type Luxury struct {
	Price float64
}

func (l Luxury) Calculate(days int) float64 {
	return l.Price * float64(days) * 1.5
}

type Hotel struct {
	Rooms    []Room
	Guests   []Guest
	Bookings []Booking
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms = append(h.Rooms, room)
}

func (h *Hotel) AddGuest(guest Guest) {
	h.Guests = append(h.Guests, guest)
}

func (h *Hotel) FindRoom(number int) *Room {
	for i := range h.Rooms {
		if h.Rooms[i].Number == number {
			return &h.Rooms[i]
		}
	}
	return nil
}

func (h *Hotel) BookRoom(roomNumber, guestID, days int) {
	room := h.FindRoom(roomNumber)

	if room == nil || room.Busy || days <= 0 {
		fmt.Println("Невозможно забронировать комнату")
		return
	}

	guestExists := false

	for _, g := range h.Guests {
		if g.ID == guestID {
			guestExists = true
			break
		}
	}

	if !guestExists {
		fmt.Println("Гость не найден")
		return
	}

	room.Busy = true

	h.Bookings = append(h.Bookings, Booking{
		RoomNumber: roomNumber,
		GuestID:    guestID,
		Days:       days,
	})

	fmt.Println("Комната забронирована")
}

func (h *Hotel) CancelBooking(roomNumber int) {
	for i, b := range h.Bookings {
		if b.RoomNumber == roomNumber {
			room := h.FindRoom(roomNumber)
			if room != nil {
				room.Busy = false
			}

			h.Bookings = append(h.Bookings[:i], h.Bookings[i+1:]...)
			fmt.Println("Бронирование отменено")
			return
		}
	}
}

func (h *Hotel) EvictGuest(guestID int) {
	for i, b := range h.Bookings {
		if b.GuestID == guestID {
			room := h.FindRoom(b.RoomNumber)

			if room != nil {
				room.Busy = false
			}

			h.Bookings = append(h.Bookings[:i], h.Bookings[i+1:]...)
			fmt.Println("Гость выселен")
			return
		}
	}
}

func (h *Hotel) CalculatePrice(roomNumber int) {
	for _, b := range h.Bookings {
		if b.RoomNumber == roomNumber {
			room := h.FindRoom(roomNumber)

			if room == nil {
				return
			}

			var pricing Pricing

			switch room.Type {
			case "Стандарт":
				pricing = Standard{room.Price}
			case "Семейный":
				pricing = Family{room.Price}
			case "Люкс":
				pricing = Luxury{room.Price}
			default:
				fmt.Println("Неизвестный тип номера")
				return
			}

			fmt.Println("Стоимость проживания:",
				pricing.Calculate(b.Days), "руб.")
			return
		}
	}
}

func (h *Hotel) ShowFreeRooms() {
	for _, r := range h.Rooms {
		if !r.Busy {
			fmt.Println("Комната:", r.Number,
				"Мест:", r.Places,
				"Цена:", r.Price)
		}
	}
}

func (h *Hotel) SortRooms() {
	sort.Slice(h.Rooms, func(i, j int) bool {
		return h.Rooms[i].Price < h.Rooms[j].Price
	})
}

func (h *Hotel) ShowMinMax() {
	if len(h.Rooms) == 0 {
		return
	}

	min, max := h.Rooms[0], h.Rooms[0]

	for _, r := range h.Rooms {
		if r.Price < min.Price {
			min = r
		}
		if r.Price > max.Price {
			max = r
		}
	}

	fmt.Println("Самая дешёвая:", min.Number, min.Price)
	fmt.Println("Самая дорогая:", max.Number, max.Price)
}

func main() {
	hotel := Hotel{}

	hotel.AddRoom(Room{101, 2, 2000, false, "Стандарт"})
	hotel.AddRoom(Room{102, 4, 4000, false, "Семейный"})
	hotel.AddRoom(Room{103, 2, 7000, false, "Люкс"})

	hotel.AddGuest(Guest{1, "Алексей", 20})

	hotel.BookRoom(101, 1, 5)
	hotel.CalculatePrice(101)

	fmt.Println("\nСвободные комнаты:")
	hotel.ShowFreeRooms()

	fmt.Println("\nСтоимость номеров:")
	hotel.ShowMinMax()

	hotel.EvictGuest(1)
}
