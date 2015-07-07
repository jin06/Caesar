package command

import (
	"net"
	"fmt"
	"net/rpc"
	
)

type Command struct {
	Cmd string
}

func DisCmd(cmd string) {
	switch(cmd){
		case "":
		default:
		fmt.Println("Command not found!", cmd) 
	} 
}

func Service(conn *net.TCPConn) {
	b := make([]byte, 1024)
	for {
		lenth, _ := conn.Read(b) 
		DisCmd(string(b[:lenth]))
		
	}	
}

func RpcService(c *net.TCPConn) {
	rpcServer := rpc.NewServer()
	users := new(Users)
	rpcServer.Register(users)
	rpcServer.ServeConn(c)	
}

type Users struct {
	Users string
	
}

type User struct {
	Name string
	Password string
}

func (u *Users) Login(user *User,s *string) error{
	*s =  user.Name + user.Password
	return nil
}
