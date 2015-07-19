package object

import (
	//"net"
	//"net/rpc"
	//"math/rand"
	//"time"
	//"fmt"
	"strconv"
	
	"github.com/jin06/Caesar/db"
	"github.com/jin06/Caesar/log"
	"github.com/jin06/Caesar/msgqueue"
	//"github.com/jin06/Caesar/log"
	
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
func (users *Users) Login (user *User, res *string)  error {
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
					
					//fmt.Printf("%s login.\n", user.Name)
					logs := "A new client login. Client's name is " + user.Name
					log.Log("info", logs, nil)
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
	simres.LogInfo = ""
	//simres.Res = "You don't have queue."
	if users.isLogined(user) {
		if user.Key != users.UM[user.Name].Key{
			simres.LogInfo = "User has already login!!! You are not login."
			simres.Res = ""
		}else {
			mqArr, err := db.GetAllMqFromDB(user.Name)
			if err != nil {
				return err
			}else {
				i := 0
				runing := "Runing message queue :"
				for _, mq := range mqArr {
					if mq == nil {break}
					simres.Res += "ID: " + strconv.Itoa(mq.MQid) + "  NAME: " + mq.MQname +" \n"
					if _,ok := msgqueue.DefaultMM[mq.MQid]; ok {
						runing += strconv.Itoa(mq.MQid) + " "
					}
					i++
				}
				status := "You have " + strconv.Itoa(i) + " message queue:\n"
				if runing == "Runing message queue :" {
					runing = "No message queue is runing."
				}
				simres.Res =status + simres.Res + runing + "\n"
				
			}
		}
	}else {
		simres.LogInfo = "You are offline, please login."
		simres.Res = ""
	}
	return nil
}

func (users *Users) Users (user *User, simres *SimResult) error {
	simres.LogInfo = ""
	//simres.Res = "You don't have queue."
	if users.isLogined(user) {
		if user.Key != users.UM[user.Name].Key{
			simres.LogInfo = "User has already login!!! You are not login."
			simres.Res = ""
		}else {
			usersRes, err := db.GetAllUsersFromDB()
			if err != nil {
				return err
			}else {
				i := 0
				//status := "Users :"
				for _, user := range usersRes {
					if user == nil {break}
					simres.Res += "ID: " + strconv.Itoa(user.Id) + "  NAME: " + user.Name +" \n"
					i++
				}
				simres.Res += "Users num is :" + strconv.Itoa(i) + "\n"
				
				//simres.Res = simres.Res + "\n"
			}
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
			//fmt.Printf("%s exit.\n", user.Name)
			logs := "Client " + user.Name + " logoff."
			log.Log("info", logs, nil)
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

