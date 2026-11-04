package main

import (
	"fmt"
	"sort"
)

type Transaction struct {
	Type   string
	Amount float64
	Info   string
}

type Account struct {
	Number  int
	Balance float64
	Blocked bool
	History []Transaction
}

type Client struct {
	ID       int
	Name     string
	Accounts []Account
}

type Operation interface {
	Execute() bool
}

type Deposit struct {
	Account *Account
	Amount  float64
}

func (d Deposit) Execute() bool {
	if d.Account.Blocked || d.Amount <= 0 {
		return false
	}

	d.Account.Balance += d.Amount
	d.Account.History = append(d.Account.History,
		Transaction{"Пополнение", d.Amount, "Внесение средств"})
	return true
}

type Withdraw struct {
	Account *Account
	Amount  float64
}

func (w Withdraw) Execute() bool {
	if w.Account.Blocked || w.Amount <= 0 ||
		w.Amount > w.Account.Balance {
		return false
	}

	w.Account.Balance -= w.Amount
	w.Account.History = append(w.Account.History,
		Transaction{"Снятие", w.Amount, "Снятие средств"})
	return true
}

type Transfer struct {
	From   *Account
	To     *Account
	Amount float64
}

func (t Transfer) Execute() bool {
	if t.From.Blocked || t.To.Blocked ||
		t.Amount <= 0 || t.Amount > t.From.Balance {
		return false
	}

	t.From.Balance -= t.Amount
	t.To.Balance += t.Amount

	t.From.History = append(t.From.History,
		Transaction{"Перевод", t.Amount, "Отправка денег"})

	t.To.History = append(t.To.History,
		Transaction{"Перевод", t.Amount, "Получение денег"})

	return true
}

type Bank struct {
	Clients []Client
}

func (b *Bank) AddClient(id int, name string) {
	b.Clients = append(b.Clients, Client{
		ID:   id,
		Name: name,
	})
}

func (b *Bank) AddAccount(clientID, number int, balance float64) {
	if balance < 0 {
		return
	}

	for i := range b.Clients {
		if b.Clients[i].ID == clientID {
			b.Clients[i].Accounts = append(
				b.Clients[i].Accounts,
				Account{Number: number, Balance: balance},
			)
			return
		}
	}
}

func (b *Bank) FindClient(id int) *Client {
	for i := range b.Clients {
		if b.Clients[i].ID == id {
			return &b.Clients[i]
		}
	}
	return nil
}

func (b *Bank) FindAccount(number int) *Account {
	for i := range b.Clients {
		for j := range b.Clients[i].Accounts {
			if b.Clients[i].Accounts[j].Number == number {
				return &b.Clients[i].Accounts[j]
			}
		}
	}
	return nil
}

func (b *Bank) BlockAccount(number int) {
	account := b.FindAccount(number)

	if account != nil {
		account.Blocked = true
		fmt.Println("Счёт заблокирован")
	}
}

func (b *Bank) ShowHistory(number int) {
	account := b.FindAccount(number)

	if account == nil {
		return
	}

	for _, t := range account.History {
		fmt.Println(t.Type, t.Amount, t.Info)
	}
}

func (b *Bank) SortClients() {
	sort.Slice(b.Clients, func(i, j int) bool {
		var sum1, sum2 float64

		for _, a := range b.Clients[i].Accounts {
			sum1 += a.Balance
		}

		for _, a := range b.Clients[j].Accounts {
			sum2 += a.Balance
		}

		return sum1 > sum2
	})
}

func (b *Bank) ShowClients() {
	for _, c := range b.Clients {
		fmt.Println("Клиент:", c.Name, "ID:", c.ID)

		for _, a := range c.Accounts {
			fmt.Println("Счёт:", a.Number,
				"Баланс:", a.Balance,
				"Заблокирован:", a.Blocked)
		}
	}
}

func main() {
	bank := Bank{}

	bank.AddClient(1, "Алексей")
	bank.AddClient(2, "Иван")

	bank.AddAccount(1, 1001, 5000)
	bank.AddAccount(2, 1002, 3000)

	account1 := bank.FindAccount(1001)
	account2 := bank.FindAccount(1002)

	Deposit{account1, 2000}.Execute()

	Withdraw{account1, 1000}.Execute()

	Transfer{account1, account2, 1500}.Execute()

	fmt.Println("История счёта 1001:")
	bank.ShowHistory(1001)

	bank.SortClients()

	fmt.Println("\nИнформация о клиентах:")
	bank.ShowClients()

	bank.BlockAccount(1002)
}
