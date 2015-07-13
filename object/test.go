package object

import (
	//"fmt"
	//"github.com/jin06/Caesar/msgqueue"
)
  
type Test struct {
	Id int
	Name string
}

type User1 struct {
	Name string
	Password string
}

func (u *Test) Login(user *User1,s *string) error{
	*s =  user.Name + user.Password
	return nil
}