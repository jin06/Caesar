package client

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	//"os/exec"
	"math/rand"
	"time"
	//"strings"

	"github.com/jin06/Caesar/command"
	"github.com/jin06/Caesar/log"
	"github.com/jin06/Caesar/object"
	//	"github.com/jin06/Caesar/msgqueue"

	"github.com/tsuru/config"
)

var (
	rpcClient *rpc.Client
	r = bufio.NewReader(os.Stdin)
	StatusLine = 0
)
var Me = &object.User{
	Id:0,  
	Role:"regular", 
	Group:"unknow",
	Key:genKey(),
}


func StartService(localAddr, serverAddr, username, password string) {

	//Create a new rpc client for handle the request to server and respons from server.
	var err error
	rpcClient, err = rpc.Dial("tcp4", serverAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer rpcClient.Close()
	fmt.Println("Start success! Client startup at", localAddr, ".", "Connected server address:", serverAddr)
	//defer logOff()
	//Scan the input command.
	
	for {
		fmt.Print(">")
		line, _, err := r.ReadLine()
		handleError(err)
		if !CmdPara(string(line)){
			disCmd(string(line))
		}
	}
}

//resolve and excutive the command "cmd"
func disCmd(cmd string) bool {
	switch cmd {
	case "exit":
		fmt.Println("Caesar exit! Bye!")
		logOff()
		os.Exit(0)
		return true
	case "info":
		if DefautStatus.Login == true {
			log.Log("info", "You are online.", log.Fields{"Login Name": DefautStatus.LoginName})
		} else {
			log.Log("info", "You are offline.", nil)
		}
		return true
	case "test":
		var s string
		user := &command.User{"vvvvvvdfdf", "1111111"}
		rpcClient.Call("Test.Login", user, &s)
		fmt.Println(s)
		return true
	case "login":
		if DefautStatus.Login == true {
			log.Log("info", "You are already online.", log.Fields{"Login Name": DefautStatus.LoginName})
		} else {
			//input name and password
			fmt.Println("Please input name:")
			line, _, err := r.ReadLine()
			handleError(err)
			username := string(line)
			fmt.Println("Please input password:")
			line, _, err = r.ReadLine()
			handleError(err)
			password := string(line)

			var res string
			//user := &object.User{0, username, password, "regular", "unknow",genKey()}
			Me.Name = username
			Me.Password = password
			
			//fmt.Println(Me)
			rpcClient.Call("Users.Login", Me, &res)
			if res == "Login success." {StatusLine = 1}
			//fmt.Println(Me.Id)
			//fmt.Println(Me.Key)
			log.Log("info", res, nil)
		}
		return true
	case "myqueue":
		var simRes object.SimResult
		rpcClient.Call("Users.MyMQ", Me, &simRes)
		if simRes.LogInfo != "" {
			fmt.Print(simRes.LogInfo)
		}else {
			fmt.Printf(simRes.Res)
		}
		return true	
	case "newqueue":
		fmt.Println("Please input name:")
		return true

	default:
		fmt.Printf("Command \"%s\" not found!\n", cmd)
		return false
	}
}

func handleError(err error) {
	if err != nil {
		log.Log("err", err.Error(), nil)
	}
}

//Read the config file and set the config.
func Readconfig(local *string, server *string) {
	err := config.ReadConfigFile("../client/config/client.yml")
	handleError(err)

	*local, err = config.GetString("localaddress")
	handleError(err)

	*server, err = config.GetString("serveraddress")

	handleError(err)

}

//gennerate user's key
func genKey() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Int31n(100000000))
}

//when you exit, send a command to server that you are exit
func logOff() {
	//fmt.Println("run here")
	var res string
	rpcClient.Call("Users.LogOff", Me, &res)
}
