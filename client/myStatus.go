package client
//status is the client's status parameters. For instance: client's login messsage
import (

)

var DefautStatus *Status = newStatus()



type Status struct {
	Login bool
	LoginName string
}

func newStatus() *Status{
	return &Status{false, ""}	
}

