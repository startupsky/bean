package user

import (
	"testing"
)

func Test_InitUserManagerIsEmpty(t *testing.T) {
	user_mgr := NewUserManager()
	user_mgr.Clear()
	if len(user_mgr.users) > 0{
		t.Error("User count should be 0 after clear")
	}
}

func Test_Add(t *testing.T) {
	user_mgr := NewUserManager()
	user_mgr.Clear()
	user_mgr.AddUser("tom", "tom1")
	if len(user_mgr.users) != 1{
		t.Error("User count should be 1 after add tom")
	}
}

func Test_GetTom(t *testing.T) {
	user_mgr := NewUserManager()
	user_mgr.Clear()
	user_mgr.AddUser("tom", "tom1")
	user := user_mgr.GetUser("tom", "tom1")
	if user == nil{
		t.Error("Should get tom when id and password correct")
	}
}

func Test_GetTomFailWhenPasswordWrong(t *testing.T) {
	user_mgr := NewUserManager()
	user_mgr.Clear()
	user_mgr.AddUser("tom", "tom1")
	user := user_mgr.GetUser("tom", "tom")
	if user != nil{
		t.Error("Should get nil when id and password wrong")
	}
}

func Test_AddUserFailWhenIdDuplicate(t *testing.T) {
	user_mgr := NewUserManager()
	user_mgr.Clear()
	user := user_mgr.AddUser("tom", "tom1")
	if user == nil{
		t.Error("Should add tom 1st time")
	}
	user = user_mgr.AddUser("tom", "tom")
	if user != nil{
		t.Error("Should add tom fail 2nd time")
	}
}