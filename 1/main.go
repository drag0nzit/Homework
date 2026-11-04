package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Stock    int
}

type User struct {
	ID      int
	Name    string
	Balance float64
}

type CartItem struct {
	ProductID int
	Quantity  int
}

type Payment interface {
	Pay(amount float64) bool
}

type CardPayment struct{}

func (p CardPayment) Pay(amount float64) bool {
	fmt.Println("Оплата банковской картой:", amount)
	return true
}

type CashPayment struct{}

func (p CashPayment) Pay(amount float64) bool {
	fmt.Println("Оплата наличными:", amount)
	return true
}

type OnlinePayment struct{}

func (p OnlinePayment) Pay(amount float64) bool {
	fmt.Println("Оплата через интернет:", amount)
	return true
}

type Shop struct {
	Products []Product
	Users    []User
	Cart     []CartItem
}

func (s *Shop) ShowProducts() {
	for _, p := range s.Products {
		fmt.Printf("%d | %s | %s | %.2f руб. | Остаток: %d\n",
			p.ID, p.Name, p.Category, p.Price, p.Stock)
	}
}

func (s *Shop) FindByID(id int) *Product {
	for i := range s.Products {
		if s.Products[i].ID == id {
			return &s.Products[i]
		}
	}
	return nil
}

func (s *Shop) FindByName(name string) *Product {
	for i := range s.Products {
		if strings.EqualFold(s.Products[i].Name, name) {
			return &s.Products[i]
		}
	}
	return nil
}

func (s *Shop) AddProduct(p Product) {
	s.Products = append(s.Products, p)
}

func (s *Shop) ChangePrice(id int, price float64) {
	p := s.FindByID(id)
	if p != nil && price >= 0 {
		p.Price = price
	}
}

func (s *Shop) AddStock(id, amount int) {
	p := s.FindByID(id)
	if p != nil && amount > 0 {
		p.Stock += amount
	}
}

func (s *Shop) DeleteProduct(id int) {
	for i, p := range s.Products {
		if p.ID == id {
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			return
		}
	}
}

func (s *Shop) AddToCart(id, quantity int) {
	p := s.FindByID(id)

	if p == nil || quantity <= 0 || quantity > p.Stock {
		fmt.Println("Невозможно добавить товар")
		return
	}

	for i := range s.Cart {
		if s.Cart[i].ProductID == id {
			if s.Cart[i].Quantity+quantity <= p.Stock {
				s.Cart[i].Quantity += quantity
			}
			return
		}
	}

	s.Cart = append(s.Cart, CartItem{id, quantity})
}

func (s *Shop) RemoveFromCart(id int) {
	for i, item := range s.Cart {
		if item.ProductID == id {
			s.Cart = append(s.Cart[:i], s.Cart[i+1:]...)
			return
		}
	}
}

func (s *Shop) ChangeCartQuantity(id, quantity int) {
	p := s.FindByID(id)

	if p == nil || quantity <= 0 || quantity > p.Stock {
		return
	}

	for i := range s.Cart {
		if s.Cart[i].ProductID == id {
			s.Cart[i].Quantity = quantity
			return
		}
	}
}

func (s *Shop) Checkout(userID int, payment Payment) {
	var user *User

	for i := range s.Users {
		if s.Users[i].ID == userID {
			user = &s.Users[i]
			break
		}
	}

	if user == nil || len(s.Cart) == 0 {
		fmt.Println("Пользователь или корзина не найдены")
		return
	}

	var total float64

	for _, item := range s.Cart {
		p := s.FindByID(item.ProductID)

		if p == nil || p.Stock < item.Quantity {
			fmt.Println("Недостаточно товара на складе")
			return
		}

		total += p.Price * float64(item.Quantity)
	}

	if user.Balance < total {
		fmt.Println("Недостаточно средств")
		return
	}

	if !payment.Pay(total) {
		fmt.Println("Оплата не прошла")
		return
	}

	for _, item := range s.Cart {
		p := s.FindByID(item.ProductID)
		p.Stock -= item.Quantity
	}

	user.Balance -= total
	s.Cart = nil

	fmt.Println("Покупка успешно оформлена!")
	fmt.Println("Итого:", total)
}

func (s *Shop) SortByPrice() {
	sort.Slice(s.Products, func(i, j int) bool {
		return s.Products[i].Price < s.Products[j].Price
	})
}

func (s *Shop) SortByName() {
	sort.Slice(s.Products, func(i, j int) bool {
		return s.Products[i].Name < s.Products[j].Name
	})
}

func (s *Shop) Statistics() {
	if len(s.Products) == 0 {
		fmt.Println("Каталог пуст")
		return
	}

	min := s.Products[0]
	max := s.Products[0]
	var total float64

	for _, p := range s.Products {
		if p.Price < min.Price {
			min = p
		}

		if p.Price > max.Price {
			max = p
		}

		total += p.Price * float64(p.Stock)
	}

	fmt.Println("Самый дешёвый:", min.Name, min.Price)
	fmt.Println("Самый дорогой:", max.Name, max.Price)
	fmt.Println("Стоимость всего склада:", total)
}

func main() {
	shop := Shop{
		Products: []Product{
			{1, "Телефон", "Электроника", 25000, 10},
			{2, "Наушники", "Аксессуары", 5000, 20},
			{3, "Ноутбук", "Электроника", 60000, 5},
		},
		Users: []User{
			{1, "Алексей", 100000},
		},
	}

	fmt.Println("Каталог товаров:")
	shop.ShowProducts()

	shop.AddToCart(1, 1)
	shop.AddToCart(2, 2)

	fmt.Println("\nОформление заказа:")
	shop.Checkout(1, CardPayment{})

	fmt.Println("\nСтатистика:")
	shop.Statistics()

	fmt.Println("\nОставшиеся товары:")
	shop.ShowProducts()
}
