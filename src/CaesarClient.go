package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jin06/Caesar/client"
	"github.com/jin06/Caesar/client/cflag"
	"github.com/jin06/Caesar/command"
	"net"
	"net/rpc"
	"os"
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

func startService() {
    
	//Create a new rpc client for handle the request to server and respons from server.
	var err error
	rpcClient, err = rpc.Dial("tcp4", serverAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer rpcClient.Close()
	fmt.Println("Client startup at", localAddr, ".", "Connected to Server:", serverAddr)

	r := bufio.NewReader(os.Stdin)
	//Scan the input command.
	for {
		fmt.Print(">")
		line, _, err := r.ReadLine()
		handleError(err)
		disCmd(string(line))
	}
}

//resolve and excutive the command "cmd"
func disCmd(cmd string) bool {
	switch cmd {
	case "exit":
		fmt.Println("Exit!Bye!")
		os.Exit(0)
		return true
	case "login":
		var s string
		user := &command.User{"jinlong", "123"}
		rpcClient.Call("Users.Login", user, &s)
		fmt.Println(s)
		return true
	default:
		fmt.Println("Command not found!")
		return false
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
