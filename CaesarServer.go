package main 

import (
	//"net/rpc"
	"net"
	"fmt"
	"github.com/jin06/Caesar/msgqueue"
	//"encoding/gob" 
//	/"message"
	//"bytes"
	//"encoding/json"
)

const (
	listenAddr string = "127.0.0.1:1212" 
	msgAddr string = "127.0.0.1:1213"
)	

var (
	mqAgent msgqueue.MqAgent = msgqueue.MqAgent{}
)

func handleErr(err error) {
	if err != nil {
		fmt.Print(err)	
	}	
}

func main() {
	//resolve TCPAddress 
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)
	handleErr(err)
	
	//Create listener 
	ln, err := net.ListenTCP("tcp4", tcpAddr)
	handleErr(err)
	
	//Create TCP connection
	
	tcpConn , err := ln.AcceptTCP()
	
	handleErr(err)   
	
	 
//	var	value interface{} = message.AD{}
//	b := make([]byte,1024)
//	msgqueue.RecAndRes(tcpConn, b, &value)
//	fmt.Printf("%+v", value)	
	
		 mqAgent.Start(tcpConn)	
	
}

