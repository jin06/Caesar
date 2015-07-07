package main 

import (
	"flag"
	"net"
	"fmt"
	"net/rpc"
	
	"github.com/jin06/Caesar/msgqueue"
	"github.com/jin06/Caesar/command" 
	
)

var (
	listenAddr string = "127.0.0.1:1212" 
	msgAddr string = "127.0.0.1:1213"
	mqAgent msgqueue.MqAgent = msgqueue.MqAgent{}
)	

func handleErr(err error) {
	if err != nil {
		fmt.Print(err)	
	}	
}

func flagResolve() {
	addrFlag := flag.String("Address", "", "listen Addr and port")
	if *addrFlag != "" {
		listenAddr = *addrFlag	
	} 
}

func main() {
	
	flagResolve()
	
	//resolve TCPAddress 
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)
	handleErr(err)
	
	//Create listener 
	ln, err := net.ListenTCP("tcp4", tcpAddr)
	handleErr(err)
	
	//Create TCP connection
	
//	tcpConn , err := ln.AcceptTCP()	
//	handleErr(err)  
	
	rpcServer := rpc.NewServer()
	users := new(command.Users)
	rpcServer.Register(users)
	rpcServer.Accept(ln)
	
	//command.Service(tcpConn) 
//	var	value interface{} = message.AD{}
//	b := make([]byte,1024)
//	msgqueue.RecAndRes(tcpConn, b, &value)
//	fmt.Printf("%+v", value)	
	// mqAgent.Start(tcpConn)	
}

