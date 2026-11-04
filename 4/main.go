package main

import (
	"fmt"
	"sort"
)

type Fighter interface {
	Attack() int
	TakeDamage(damage int)
	IsAlive() bool
	GetName() string
	GetLevel() int
	GetHealth() int
}

type Warrior struct {
	Name    string
	Health  int
	Level   int
	Power   int
	Defense int
}

type Mage struct {
	Name   string
	Health int
	Level  int
	Power  int
	Mana   int
}

type Archer struct {
	Name   string
	Health int
	Level  int
	Power  int
	Arrows int
}

func (w *Warrior) Attack() int {
	return w.Power + w.Level*2
}

func (m *Mage) Attack() int {
	if m.Mana < 10 {
		return 0
	}

	m.Mana -= 10
	return m.Power * 2
}

func (a *Archer) Attack() int {
	if a.Arrows <= 0 {
		return 0
	}

	a.Arrows--
	return a.Power + 5
}

func (w *Warrior) TakeDamage(damage int) {
	damage -= w.Defense
	if damage < 0 {
		damage = 0
	}

	w.Health -= damage
	if w.Health < 0 {
		w.Health = 0
	}
}

func (m *Mage) TakeDamage(damage int) {
	m.Health -= damage
	if m.Health < 0 {
		m.Health = 0
	}
}

func (a *Archer) TakeDamage(damage int) {
	a.Health -= damage
	if a.Health < 0 {
		a.Health = 0
	}
}

func (w *Warrior) IsAlive() bool {
	return w.Health > 0
}

func (m *Mage) IsAlive() bool {
	return m.Health > 0
}

func (a *Archer) IsAlive() bool {
	return a.Health > 0
}

func (w *Warrior) GetName() string { return w.Name }
func (m *Mage) GetName() string    { return m.Name }
func (a *Archer) GetName() string  { return a.Name }

func (w *Warrior) GetLevel() int { return w.Level }
func (m *Mage) GetLevel() int    { return m.Level }
func (a *Archer) GetLevel() int  { return a.Level }

func (w *Warrior) GetHealth() int { return w.Health }
func (m *Mage) GetHealth() int    { return m.Health }
func (a *Archer) GetHealth() int  { return a.Health }

func Heal(f Fighter, amount int) {
	switch p := f.(type) {
	case *Warrior:
		p.Health += amount
	case *Mage:
		p.Health += amount
	case *Archer:
		p.Health += amount
	}

	fmt.Println(f.GetName(), "восстановил здоровье")
}

func LevelUp(f Fighter) {
	switch p := f.(type) {
	case *Warrior:
		p.Level++
		p.Health += 10
	case *Mage:
		p.Level++
		p.Health += 10
		p.Mana += 10
	case *Archer:
		p.Level++
		p.Health += 10
		p.Arrows += 5
	}

	fmt.Println(f.GetName(), "повысил уровень")
}

func Battle(first, second Fighter) {
	attacker := first
	defender := second

	for first.IsAlive() && second.IsAlive() {
		damage := attacker.Attack()

		if damage == 0 {
			fmt.Println(attacker.GetName(), "не может атаковать")
			break
		}

		defender.TakeDamage(damage)

		fmt.Println(attacker.GetName(), "наносит", damage,
			"урона персонажу", defender.GetName())

		fmt.Println("Здоровье противника:", defender.GetHealth())

		attacker, defender = defender, attacker
	}

	if first.IsAlive() {
		fmt.Println("Победитель:", first.GetName())
	} else if second.IsAlive() {
		fmt.Println("Победитель:", second.GetName())
	} else {
		fmt.Println("Ничья!")
	}
}

func SortByLevel(players []Fighter) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].GetLevel() > players[j].GetLevel()
	})
}

func SortByHealth(players []Fighter) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].GetHealth() > players[j].GetHealth()
	})
}

func main() {
	warrior := &Warrior{"Воин", 100, 5, 15, 5}
	mage := &Mage{"Маг", 80, 5, 20, 50}
	archer := &Archer{"Лучник", 90, 4, 15, 30}

	players := []Fighter{warrior, mage, archer}

	fmt.Println("Начало боя:")
	Battle(warrior, mage)

	fmt.Println("\nЛечение:")
	Heal(archer, 20)

	fmt.Println("\nПовышение уровня:")
	LevelUp(archer)

	fmt.Println("\nСортировка по уровню:")
	SortByLevel(players)

	for _, p := range players {
		fmt.Println(p.GetName(), "Уровень:", p.GetLevel())
	}
}
