package main

import (
	
	"encoding/json"
	"fmt"
	"github.com/jin06/Caesar/client"
	"github.com/jin06/Caesar/client/cflag"
	//"github.com/jin06/Caesar/client/control"
	"net"
	"net/rpc"
)

var (
	localAddr  string = ""
	serverAddr string = ""
	username   string = ""
	password   string = ""
	rpcClient  *rpc.Client
)

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func send(v interface{}, conn *net.TCPConn) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	} else {
		conn.Write(b)
		return nil
	}
}

func main() {
	client.Readconfig(&localAddr, &serverAddr)
	//flag resolve    
	cflag.FlagResolve(&localAddr, &serverAddr, &username, &password)
	//fmt.Println(localAddr,serverAddr,username,password)

	//start rpc service in port: localAddr, and the server address is serverAddr 
	client.StartService(localAddr, serverAddr, username, password)
}
