package mystruct

import "fmt"

type User struct {
	Login string	
	Tokens float64	
}

var user = User{}

func GetStruct() *User {
	return &user
}

func PrintInfo() {
	fmt.Println(user)
}

func (u *User) MineTokens() {
	u.Tokens += 0.1
}

func (u *User) SetLogin(login string) {
	u.Login = login
}