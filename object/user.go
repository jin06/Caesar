package object

import (
	//"net"
	//"net/rpc"
	//"math/rand"
	//"time"
	"fmt"
	
	"github.com/jin06/Caesar/db"
	"github.com/jin06/Caesar/log"
	
)

//Users map save onlin users
type Users struct {
	 UM map[string]User
	 Number int
}

//user online
type User struct {
	Id int
	Name string
	Password string
	//RpcClient *rpc.Client  
	Role string   //user's role, admin or general user
	Group string  //user belong to group
	Key int
}

func NewUsers() *Users {
	users := new(Users)
	users.UM = make(map[string]User)
	return users
}

//user login control
func (users *Users) Login (user *User, res *string) error {
	//log.Log("warn","1234455566666666" , nil)
	if users.isLogined(user) {
		*res = "User has logined."
		return nil
	}else {
		i, err := db.VerifyUser(user.Name, user.Password)
		if err != nil {
			log.Log("err", err.Error(), nil)
			*res = "Server error."
			return nil
		}else {
			switch i {
				case 0://no user
					*res = "Wrong username."
					return nil
				case 1:
					users.UM[user.Name] = *user
					err = db.UpdateRegisterTime(user.Name)
					if err != nil {
						log.Log("err", err.Error(), nil )
					}
					
					//fmt.Println(user.Key)
					*res = "Login success."
					fmt.Printf("%s login.\n", user.Name)
					return  nil
				case 2:	//wrong password
					*res = "Wrong password."
					return nil
				default :
					*res = "Server error."
					return nil
			}
		}
	}
}

func (users *Users) MyMQ (user *User, simres *SimResult) error {
	if users.isLogined(user) {
		if user.Key != users.UM[user.Name].Key{
			simres.LogInfo = "User has already login!!! You are not login."
			simres.Res = ""
			
		}else {
			simres.LogInfo = ""
			simres.Res = "result...."
		}
	}else {
		simres.LogInfo = "You are offline, please login."
		simres.Res = ""
	}
	return nil
}

func (users *Users) LogOff (user *User, res *string) error {
	if users.isLogined(user) && user.Key == users.UM[user.Name].Key{
			*res = ""
			delete(users.UM, user.Name)
			fmt.Printf("%s exit.\n", user.Name)
	}else {
		*res = ""
	}
	return nil
}

//check user if logined
func (users *Users) isLogined(user *User) bool {
	//log.Log("err", "111111",nil)
	_, ok := users.UM[user.Name]
	if ok {
		return true
	} else {
		return false
	}
}

