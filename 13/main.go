package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID      int
	Name    string
	Friends []int
}

type SocialNetwork struct {
	Users map[int]User
}

func (s *SocialNetwork) AddUser(id int, name string) {
	if _, exists := s.Users[id]; exists {
		fmt.Println("Пользователь уже существует")
		return
	}

	s.Users[id] = User{
		ID:   id,
		Name: name,
	}
}

func (s *SocialNetwork) DeleteUser(id int) {
	delete(s.Users, id)

	for userID, user := range s.Users {
		user.Friends = removeFriend(user.Friends, id)
		s.Users[userID] = user
	}
}

func removeFriend(friends []int, id int) []int {
	for i, friend := range friends {
		if friend == id {
			return append(friends[:i], friends[i+1:]...)
		}
	}
	return friends
}

func (s *SocialNetwork) AddFriend(userID, friendID int) {
	user, exists1 := s.Users[userID]
	friend, exists2 := s.Users[friendID]

	if !exists1 || !exists2 {
		fmt.Println("Пользователь не найден")
		return
	}

	if userID == friendID {
		fmt.Println("Нельзя добавить самого себя")
		return
	}

	for _, id := range user.Friends {
		if id == friendID {
			fmt.Println("Друг уже добавлен")
			return
		}
	}

	user.Friends = append(user.Friends, friendID)
	friend.Friends = append(friend.Friends, userID)

	s.Users[userID] = user
	s.Users[friendID] = friend
}

func (s *SocialNetwork) DeleteFriend(userID, friendID int) {
	user, exists1 := s.Users[userID]
	friend, exists2 := s.Users[friendID]

	if !exists1 || !exists2 {
		return
	}

	user.Friends = removeFriend(user.Friends, friendID)
	friend.Friends = removeFriend(friend.Friends, userID)

	s.Users[userID] = user
	s.Users[friendID] = friend
}

func (s *SocialNetwork) CommonFriends(id1, id2 int) {
	u1, exists1 := s.Users[id1]
	u2, exists2 := s.Users[id2]

	if !exists1 || !exists2 {
		return
	}

	for _, friend1 := range u1.Friends {
		for _, friend2 := range u2.Friends {
			if friend1 == friend2 {
				fmt.Println(s.Users[friend1].Name)
			}
		}
	}
}

func (s *SocialNetwork) MostPopular() {
	var best User

	for _, user := range s.Users {
		if len(user.Friends) > len(best.Friends) ||
			best.ID == 0 {
			best = user
		}
	}

	if best.ID != 0 {
		fmt.Println("Больше всего друзей у:", best.Name)
		fmt.Println("Количество друзей:", len(best.Friends))
	}
}

func (s *SocialNetwork) Recommend(id int) {
	user, exists := s.Users[id]

	if !exists {
		return
	}

	friends := make(map[int]bool)

	for _, friendID := range user.Friends {
		friends[friendID] = true
	}

	recommendations := make(map[int]bool)

	for _, friendID := range user.Friends {
		friend := s.Users[friendID]

		for _, secondID := range friend.Friends {
			if secondID != id && !friends[secondID] {
				recommendations[secondID] = true
			}
		}
	}

	ids := make([]int, 0, len(recommendations))

	for id := range recommendations {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	fmt.Println("Рекомендации:")

	for _, id := range ids {
		fmt.Println(s.Users[id].Name)
	}
}

func main() {
	network := SocialNetwork{
		Users: make(map[int]User),
	}

	network.AddUser(1, "Алексей")
	network.AddUser(2, "Иван")
	network.AddUser(3, "Пётр")
	network.AddUser(4, "Дмитрий")
	network.AddUser(5, "Сергей")

	network.AddFriend(1, 2)
	network.AddFriend(1, 3)
	network.AddFriend(2, 4)
	network.AddFriend(3, 4)
	network.AddFriend(4, 5)

	fmt.Println("Общие друзья:")
	network.CommonFriends(2, 3)

	fmt.Println("\nСамый популярный пользователь:")
	network.MostPopular()

	fmt.Println("\nРекомендации для Алексея:")
	network.Recommend(1)

	fmt.Println("\nУдаление друга:")
	network.DeleteFriend(1, 3)
}
