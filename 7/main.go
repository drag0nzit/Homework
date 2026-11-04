package main

import (
	"fmt"
	"sort"
)

type Order struct {
	Number     int
	Address    string
	Weight     float64
	GoodsPrice float64
	Delivery   Delivery
}

type Delivery interface {
	Calculate(weight float64) float64
	GetInfo() string
}

type Courier struct{}

func (c Courier) Calculate(weight float64) float64 {
	return 300 + weight*50
}

func (c Courier) GetInfo() string {
	return "Курьерская доставка"
}

type Post struct{}

func (p Post) Calculate(weight float64) float64 {
	return 150 + weight*30
}

func (p Post) GetInfo() string {
	return "Почтовая доставка"
}

type Pickup struct{}

func (p Pickup) Calculate(weight float64) float64 {
	return 50 + weight*10
}

func (p Pickup) GetInfo() string {
	return "Доставка в пункт выдачи"
}

func (o Order) TotalPrice() float64 {
	return o.GoodsPrice + o.Delivery.Calculate(o.Weight)
}

type Shop struct {
	Orders []Order
}

func (s *Shop) AddOrder(order Order) {
	s.Orders = append(s.Orders, order)
}

func (s *Shop) ShowOrders() {
	for _, o := range s.Orders {
		fmt.Println("Заказ:", o.Number)
		fmt.Println("Адрес:", o.Address)
		fmt.Println("Способ:", o.Delivery.GetInfo())
		fmt.Println("Стоимость товаров:", o.GoodsPrice)
		fmt.Println("Стоимость доставки:",
			o.Delivery.Calculate(o.Weight))
		fmt.Println("Итог:", o.TotalPrice())
		fmt.Println()
	}
}

func (s *Shop) SortOrders() {
	sort.Slice(s.Orders, func(i, j int) bool {
		return s.Orders[i].TotalPrice() <
			s.Orders[j].TotalPrice()
	})
}

func (s *Shop) MostExpensive() {
	if len(s.Orders) == 0 {
		return
	}

	max := s.Orders[0]

	for _, o := range s.Orders {
		if o.TotalPrice() > max.TotalPrice() {
			max = o
		}
	}

	fmt.Println("Самый дорогой заказ:", max.Number)
	fmt.Println("Итоговая стоимость:", max.TotalPrice())
}

func main() {
	shop := Shop{}

	shop.AddOrder(Order{
		Number:     1,
		Address:    "Улица Ленина, 10",
		Weight:     2,
		GoodsPrice: 2000,
		Delivery:   Courier{},
	})

	shop.AddOrder(Order{
		Number:     2,
		Address:    "Улица Мира, 5",
		Weight:     5,
		GoodsPrice: 3000,
		Delivery:   Post{},
	})

	shop.AddOrder(Order{
		Number:     3,
		Address:    "Улица Садовая, 15",
		Weight:     1,
		GoodsPrice: 1000,
		Delivery:   Pickup{},
	})

	fmt.Println("Все заказы:")
	shop.ShowOrders()

	shop.SortOrders()

	fmt.Println("Самый дорогой заказ:")
	shop.MostExpensive()
}
