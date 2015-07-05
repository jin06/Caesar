package main 

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"github.com/jin06/Caesar/message"
	//"encoding/gob"
	//"bytes"
	"encoding/json"
	//"client"
	"flag"
)

var (
	localAddr string = "127.0.0.1:2213"
	serverAddr string = "127.0.0.1:1212"
	username = ""
	password = ""
)

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

//user login
func login() {
	
}

//user register
func register() {
	
}

//Resolve and handle the flag  
func flagResolve() {
	localFlag := flag.String("local", "", "local addr")
	serverFlag := flag.String("server", "", "server addr and port")
	userFlag := flag.String("user", "guest", "guest client")
	passwordFlag := flag.String("password", "123456", "password")
	flag.Parse()
	if *localFlag != "" {
		localAddr = *localFlag
	}
	if *serverFlag != "" {
		serverAddr = *serverFlag
	}
	if *userFlag != "" {
		username = *userFlag
	}
	if *userFlag != "" {
		password = *passwordFlag
	}
	
}

func cmdService(c *net.TCPConn) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _, err := r.ReadLine()
		handleError(err)
		c.Write(line)
	}
}

func send(v interface{}, conn *net.TCPConn) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}else {
		conn.Write(b)
		return nil 
	}
}

func main() {
	
	//flag resolve
	flagResolve()
	
	//change the localAddr and serverAddr type to tcp4 address
	laddr, err := net.ResolveTCPAddr("tcp4", localAddr)
	handleError(err)
	raddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	handleError(err)
	
	//print startup infomation
	fmt.Println("Client startup at",localAddr,".", "Connecting to Server:", serverAddr)
	
	//create tcpConnection 
	tcpConn, err := net.DialTCP("tcp4", laddr, raddr)
	if err != nil {
		fmt.Println("Connection failed!:", err)
	}else {
		fmt.Println("Success!")
		//cmdService(tcpConn)
	}
	
	ad := message.AD{
		Name:"tudou",
		Content:"sdfljsdklfjsdjfkdjfkj",
	}
	
    //fmt.Printf("%+v", ad)
	send(ad, tcpConn)
}

