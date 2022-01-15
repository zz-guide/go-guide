package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (user *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", user, user)
}

func (user *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", user, user)
}

func main() {
	m := Manager{User{1, "Tom"}, "Administrator"}

	fmt.Println(m.ToString())

	fmt.Println(m.User.ToString())
}
