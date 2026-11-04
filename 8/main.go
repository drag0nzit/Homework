package main

import (
	"fmt"
)

type Car struct {
	Number  string
	Brand   string
	Model   string
	Year    int
	OwnerID int
}

type Client struct {
	ID   int
	Name string
}

type Repair struct {
	ID        int
	CarNumber string
	Works     []Service
	Status    string
}

type Service interface {
	GetName() string
	GetPrice() float64
}

type OilChange struct{}

func (o OilChange) GetName() string {
	return "Замена масла"
}

func (o OilChange) GetPrice() float64 {
	return 1500
}

type EngineRepair struct{}

func (e EngineRepair) GetName() string {
	return "Ремонт двигателя"
}

func (e EngineRepair) GetPrice() float64 {
	return 15000
}

type TireService struct{}

func (t TireService) GetName() string {
	return "Шиномонтаж"
}

func (t TireService) GetPrice() float64 {
	return 2500
}

type Diagnostics struct{}

func (d Diagnostics) GetName() string {
	return "Диагностика"
}

func (d Diagnostics) GetPrice() float64 {
	return 1000
}

type AutoService struct {
	Cars    []Car
	Clients []Client
	Repairs []Repair
}

func (a *AutoService) AddCar(car Car) {
	a.Cars = append(a.Cars, car)
}

func (a *AutoService) AddClient(client Client) {
	a.Clients = append(a.Clients, client)
}

func (a *AutoService) CreateRepair(id int, number string) {
	for _, c := range a.Cars {
		if c.Number == number {
			a.Repairs = append(a.Repairs, Repair{
				ID:        id,
				CarNumber: number,
				Status:    "Новая",
			})

			fmt.Println("Заявка создана")
			return
		}
	}

	fmt.Println("Автомобиль не найден")
}

func (a *AutoService) AddWork(repairID int, work Service) {
	for i := range a.Repairs {
		if a.Repairs[i].ID == repairID &&
			a.Repairs[i].Status != "Завершена" {
			a.Repairs[i].Works = append(
				a.Repairs[i].Works, work,
			)
			return
		}
	}
}

func (a *AutoService) ChangeStatus(id int, status string) {
	for i := range a.Repairs {
		if a.Repairs[i].ID == id {
			a.Repairs[i].Status = status
			return
		}
	}
}

func (a *AutoService) RepairPrice(id int) float64 {
	var total float64

	for _, r := range a.Repairs {
		if r.ID == id {
			for _, w := range r.Works {
				total += w.GetPrice()
			}
		}
	}

	return total
}

func (a *AutoService) FindCars(ownerID int) {
	for _, c := range a.Cars {
		if c.OwnerID == ownerID {
			fmt.Println(c.Brand, c.Model, c.Number)
		}
	}
}

func (a *AutoService) ClientExpenses(ownerID int) float64 {
	var total float64

	for _, r := range a.Repairs {
		for _, c := range a.Cars {
			if c.Number == r.CarNumber &&
				c.OwnerID == ownerID &&
				r.Status == "Завершена" {
				for _, w := range r.Works {
					total += w.GetPrice()
				}
			}
		}
	}

	return total
}

func main() {
	service := AutoService{}

	service.AddClient(Client{1, "Алексей"})

	service.AddCar(Car{
		Number:  "A123BC",
		Brand:   "Toyota",
		Model:   "Camry",
		Year:    2020,
		OwnerID: 1,
	})

	service.CreateRepair(1, "A123BC")

	service.AddWork(1, OilChange{})
	service.AddWork(1, Diagnostics{})
	service.AddWork(1, TireService{})

	fmt.Println("Стоимость ремонта:",
		service.RepairPrice(1))

	service.ChangeStatus(1, "Завершена")

	fmt.Println("Автомобили клиента:")
	service.FindCars(1)

	fmt.Println("Всего потрачено:",
		service.ClientExpenses(1))
}
