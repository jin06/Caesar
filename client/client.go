package client

import (
	"net"
)

type Client struct {
	Username string
	UserId  int
	Conn *net.TCPConn
	Level string
	Subscription int
}

