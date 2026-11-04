package main

import (
	"fmt"
	"sort"
)

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Quantity int
}

type Warehouse struct {
	ID       int
	Name     string
	Products []Product
}

type Storage struct {
	Warehouses []Warehouse
}

func (w *Warehouse) FindProduct(id int) *Product {
	for i := range w.Products {
		if w.Products[i].ID == id {
			return &w.Products[i]
		}
	}
	return nil
}

func (w *Warehouse) AddProduct(id int, amount int) {
	p := w.FindProduct(id)

	if p != nil && amount > 0 {
		p.Quantity += amount
		fmt.Println("Товар добавлен")
	}
}

func (w *Warehouse) RemoveProduct(id int, amount int) {
	p := w.FindProduct(id)

	if p == nil || amount <= 0 {
		return
	}

	if p.Quantity < amount {
		fmt.Println("Недостаточно товара")
		return
	}

	p.Quantity -= amount
	fmt.Println("Товар списан")
}

func (w *Warehouse) AddNewProduct(p Product) {
	if p.Quantity < 0 || p.Price < 0 || w.FindProduct(p.ID) != nil {
		return
	}

	w.Products = append(w.Products, p)
}

func (s *Storage) Transfer(fromID, toID, productID, amount int) {
	var from, to *Warehouse

	for i := range s.Warehouses {
		if s.Warehouses[i].ID == fromID {
			from = &s.Warehouses[i]
		}

		if s.Warehouses[i].ID == toID {
			to = &s.Warehouses[i]
		}
	}

	if from == nil || to == nil || fromID == toID || amount <= 0 {
		fmt.Println("Некорректные склады или количество")
		return
	}

	source := from.FindProduct(productID)

	if source == nil || source.Quantity < amount {
		fmt.Println("Недостаточно товара для перемещения")
		return
	}

	destination := to.FindProduct(productID)

	if destination == nil {
		copyProduct := *source
		copyProduct.Quantity = amount
		to.Products = append(to.Products, copyProduct)
	} else {
		destination.Quantity += amount
	}

	source.Quantity -= amount
	fmt.Println("Товар перемещён")
}

func (w *Warehouse) SearchProduct(name string) {
	for _, p := range w.Products {
		if p.Name == name {
			fmt.Println(p.ID, p.Name, p.Quantity, p.Price)
		}
	}
}

func (w *Warehouse) ShowLowStock(minimum int) {
	for _, p := range w.Products {
		if p.Quantity < minimum {
			fmt.Println(p.Name, "Осталось:", p.Quantity)
		}
	}
}

func (w *Warehouse) MostValuableCategory() {
	totals := make(map[string]float64)

	for _, p := range w.Products {
		totals[p.Category] += p.Price * float64(p.Quantity)
	}

	var category string
	var max float64

	for name, value := range totals {
		if value > max {
			max = value
			category = name
		}
	}

	if category != "" {
		fmt.Println("Самая дорогая категория:", category)
		fmt.Println("Общая стоимость:", max)
	}
}

func (w *Warehouse) SortProducts() {
	sort.Slice(w.Products, func(i, j int) bool {
		if w.Products[i].Category == w.Products[j].Category {
			return w.Products[i].Price < w.Products[j].Price
		}

		return w.Products[i].Category < w.Products[j].Category
	})
}

func main() {
	storage := Storage{
		Warehouses: []Warehouse{
			{ID: 1, Name: "Главный склад"},
			{ID: 2, Name: "Дополнительный склад"},
		},
	}

	storage.Warehouses[0].AddNewProduct(
		Product{1, "Ноутбук", "Электроника", 60000, 10},
	)

	storage.Warehouses[0].AddNewProduct(
		Product{2, "Мышь", "Аксессуары", 1500, 50},
	)

	storage.Warehouses[0].AddNewProduct(
		Product{3, "Монитор", "Электроника", 20000, 5},
	)

	fmt.Println("Приход товара:")
	storage.Warehouses[0].AddProduct(1, 5)

	fmt.Println("\nСписание товара:")
	storage.Warehouses[0].RemoveProduct(2, 10)

	fmt.Println("\nПеремещение товара:")
	storage.Transfer(1, 2, 1, 3)

	fmt.Println("\nТовары ниже минимума:")
	storage.Warehouses[0].ShowLowStock(10)

	fmt.Println("\nКатегория с максимальной стоимостью:")
	storage.Warehouses[0].MostValuableCategory()

	fmt.Println("\nСортировка:")
	storage.Warehouses[0].SortProducts()

	for _, p := range storage.Warehouses[0].Products {
		fmt.Println(p.Name, p.Category, p.Price, p.Quantity)
	}
}
