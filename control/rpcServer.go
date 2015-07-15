package control

import (
	"net/rpc"
	"net"
	"github.com/jin06/Caesar/log"
	"github.com/jin06/Caesar/object"
	"github.com/jin06/Caesar/db"
	//"github.com/jin06/Caesar/db"
	//"github.com/jin06/Caesar/msgqueue"
)

//default rpc server, handle ruquest form client
var (
	 DefaultServer *rpc.Server = rpc.NewServer()	
)

//method need to publish
var (
	 test = new(object.Test)
	 onlineUsers = object.NewUsers()  
	 mqAgentStat = new(MqAgentStat)
	 mqTable = new(db.Mq_Table) 
)

//publice DefaultServer's methods
func Register() {
	DefaultServer.Register(test)
	DefaultServer.Register(onlineUsers)  
	DefaultServer.Register(mqAgentStat)  
	DefaultServer.Register(mqTable)
}

//init the defaultServer, publice receiver's method and accept the listener from client.
func Init(ln *net.TCPListener) {
	//DefaultServer = rpc.NewServer()
	//test := new(Test)
	//DefaultServer.Register(test)
	//users := new(command.Users)
	//rpcServer.Register(users) 
	Register()
	log.Log("info", "Server start success, and now accept connection request from client.", nil)
	DefaultServer.Accept(ln)
	//log.Log("info", "run here now.", nil)
}



