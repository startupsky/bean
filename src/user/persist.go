package user

import(
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

type Persist struct{
	file string
}

func NewPersist(file string) *Persist {
	this := new(Persist)
	this.file = file
	return this
}

func (this *Persist) Save(u *User)  {
	fout,err := os.Create(this.file)
    defer fout.Close()
	
    if err != nil {
        fmt.Println(this.file,err)
        return
    }
	
    fout.WriteString(u.Id + ";" + u.Password)
	fout.WriteString("\r\n")
}

func (this *Persist) Load() (users map[string]*User) {
	
	content, err := ioutil.ReadFile(this.file)
	if err != nil {
        fmt.Println(this.file,err)
		users = make(map[string]*User)
        return
    }
	
 	array := strings.Split(string(content), "\r\n")
	users = make(map[string]*User)
	
	for _, line := range array {
		if len(line) > 0{
			user := NewUser(line)
	        users[user.Id] = user
		}	
    }
	return
}