package client

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	
	"github.com/jin06/Caesar/command"
	
	"github.com/tsuru/config"
	
)

var rpcClient *rpc.Client

func StartService(localAddr, serverAddr, username, password string) {
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

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func Readconfig(local *string, server *string) {
	err := config.ReadConfigFile("../client/config/client.yml")
	handleError(err)
	
	*local, err = config.GetString("localaddress")
	handleError(err)
	
	*server, err = config.GetString("serveraddress")
	
	handleError(err)
	fmt.Println(*server, *local)
}
