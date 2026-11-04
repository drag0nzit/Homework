package main

import (
	"fmt"
	"sort"
)

type Dish struct {
	Name      string
	Category  string
	Price     float64
	Available bool
	Ordered   int
}

type Order struct {
	TableNumber int
	Dishes      []Dish
}

type Discount interface {
	Calculate(total float64) float64
}

type NoDiscount struct{}

func (n NoDiscount) Calculate(total float64) float64 {
	return total
}

type RegularDiscount struct{}

func (r RegularDiscount) Calculate(total float64) float64 {
	return total * 0.9
}

type HolidayDiscount struct{}

func (h HolidayDiscount) Calculate(total float64) float64 {
	return total * 0.8
}

type Restaurant struct {
	Menu   []Dish
	Orders []Order
}

func (r *Restaurant) AddDish(dish Dish) {
	r.Menu = append(r.Menu, dish)
}

func (r *Restaurant) FindDish(name string) *Dish {
	for i := range r.Menu {
		if r.Menu[i].Name == name {
			return &r.Menu[i]
		}
	}
	return nil
}

func (r *Restaurant) DeleteDish(name string) {
	for i, dish := range r.Menu {
		if dish.Name == name {
			r.Menu = append(r.Menu[:i], r.Menu[i+1:]...)
			return
		}
	}
}

func (r *Restaurant) ChangePrice(name string, price float64) {
	dish := r.FindDish(name)

	if dish != nil && price >= 0 {
		dish.Price = price
	}
}

func (r *Restaurant) CreateOrder(table int, names []string) {
	order := Order{TableNumber: table}

	for _, name := range names {
		dish := r.FindDish(name)

		if dish == nil || !dish.Available {
			fmt.Println("Блюдо недоступно:", name)
			return
		}

		order.Dishes = append(order.Dishes, *dish)
		dish.Ordered++
	}

	r.Orders = append(r.Orders, order)
}

func (o Order) Calculate(discount Discount) float64 {
	var total float64

	for _, dish := range o.Dishes {
		total += dish.Price
	}

	return discount.Calculate(total)
}

func (r *Restaurant) PopularDish() {
	if len(r.Menu) == 0 {
		return
	}

	var popular Dish

	for _, dish := range r.Menu {
		if dish.Ordered > popular.Ordered {
			popular = dish
		}
	}

	fmt.Println("Самое популярное блюдо:", popular.Name)
	fmt.Println("Количество заказов:", popular.Ordered)
}

func (r *Restaurant) SortMenu() {
	sort.Slice(r.Menu, func(i, j int) bool {
		return r.Menu[i].Price < r.Menu[j].Price
	})
}

func main() {
	restaurant := Restaurant{}

	restaurant.AddDish(Dish{"Пицца", "Основное", 700, true, 0})
	restaurant.AddDish(Dish{"Бургер", "Основное", 450, true, 0})
	restaurant.AddDish(Dish{"Сок", "Напитки", 150, true, 0})

	restaurant.CreateOrder(1, []string{"Пицца", "Сок"})
	restaurant.CreateOrder(2, []string{"Бургер", "Сок"})

	fmt.Println("Заказ первого столика:")

	if len(restaurant.Orders) > 0 {
		fmt.Println("Без скидки:",
			restaurant.Orders[0].Calculate(NoDiscount{}))

		fmt.Println("Постоянный клиент:",
			restaurant.Orders[0].Calculate(RegularDiscount{}))

		fmt.Println("Праздничная скидка:",
			restaurant.Orders[0].Calculate(HolidayDiscount{}))
	}

	fmt.Println("\nПопулярное блюдо:")
	restaurant.PopularDish()

	fmt.Println("\nСортировка меню:")
	restaurant.SortMenu()

	for _, dish := range restaurant.Menu {
		fmt.Println(dish.Name, dish.Price)
	}
}
