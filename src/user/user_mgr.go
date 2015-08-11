package user

import(
	
)

type UserManager struct{
	users map[string]*User
	persist* Persist
}

func NewUserManager() *UserManager {
	this := new(UserManager)
	this.persist = NewPersist("user.txt")
	this.users = this.persist.Load()
	return this
}

func (this *UserManager) GetUser(id string, password string) (*User){
	if user, ok := this.users[id]; ok {
		if user.Password == password{
			return user
		}
	}
	return nil
}


func (this *UserManager) AddUser(id string, password string) (*User){
	if _, ok := this.users[id]; ok {
    	return nil
	}
	user := NewUser(id + ";"+password)
	this.users[user.Id]=user
	return user
}

func (this *UserManager) Save(){
	for _, user := range this.users{
		this.persist.Save(user)
	}
}

func (this *UserManager) Clear(){
	this.users = make(map[string]*User)
}