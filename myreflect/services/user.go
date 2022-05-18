package services

import "fmt"

type User struct {
	use  string
	Name string
	Id   int64
}

func (u *User) Say(content string) string {
	return fmt.Sprintf("My Name is %s, Id is %d, Say %s", u.Name, u.Id, content)
}
func (u User) Run() {
	return
}
func (u *User) Add(num int) int {
	return int(u.Id) + num
}

func (u *User) SetUse(use string) {
	u.use = use
}

func (u *User) GetUse() string {
	return u.use
}
