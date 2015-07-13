package main

import (
	"net"
	//"github.com/jin06/Caesar/command"
	"github.com/jin06/Caesar/msgqueue"
	"github.com/jin06/Caesar/sflag"
	"github.com/jin06/Caesar/log"
	"github.com/jin06/Caesar/control"
	
	"github.com/tsuru/config"
)

var (
	listenAddr string          // = "127.0.0.1:1212"
	msgAddr    string          // = "127.0.0.1:1213"
	mqAgent    msgqueue.MqAgent = msgqueue.MqAgent{}
)

func handleErr(err error) {
	if err != nil {
		//fmt.Print(err)
		log.Log("err", err.Error(), nil)
	}
} 
//Load configuration file.
func loadConf() {
	log.Log("info", "Load config file...", nil)
	err := config.ReadConfigFile("../config/server.yaml")
	if err != nil {
		//fmt.Print(err)
		log.Log("err", err.Error(), nil)
	}else {
	log.Log("info", "Completely read!", nil)
	}
	   
	listenAddr, err = config.GetString("rpcAddress")
	handleErr(err)
	msgAddr, err = config.GetString("msgAddress")
	handleErr(err)
}

func main() {
	log.Log("","Welcom to use Caesar. Caesar is a high performance message queue.", nil)
	loadConf() 
	sflag.FlagResolve(&listenAddr)  
	
	//resolve TCPAddress
	log.Log("info", "Now start server...", log.Fields{"Listen Address" : listenAddr})  
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)
	handleErr(err)

	//Create listener
	ln, err := net.ListenTCP("tcp4", tcpAddr)
	control.Init(ln)
//	rpcServer := rpc.NewServer() 
//	users := new(command.Users)
//	rpcServer.Register(users)
//	log.Log("info", "Server start success, and now accept request from client.", nil)
//	rpcServer.Accept(ln)
} 
