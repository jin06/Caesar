package Command

import (
	//"net"
	"fmt"
)

type Command struct {
	Cmd string
}

func (c *Command) DisCmd(cmd string) {
	switch(cmd){
		case "":
		default:
		fmt.Println("Command not found!", cmd) 
	} 
}

