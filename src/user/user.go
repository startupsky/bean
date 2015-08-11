package user

import(
	"strings"
)
	
type User struct {
	Id          string
	Password      string
}

func NewUser(line string) *User {
	this := new(User)
	linearay := strings.Split(line, ";")
	this.Id = linearay[0]
	this.Password = linearay[1]
	return this
}