package msgqueue

import (
	"net" 
	"fmt"
	"encoding/json"
	"github.com/jin06/Caesar/message"
	//"net"
)

type MqAgent struct {
	MQ *MsQueue
	TCPConn *net.TCPConn	
}

func NewMqAgent() *MqAgent{
	mqAgent := MqAgent{}
	mqAgent.MQ = NewMQ()
	return &mqAgent
}

func Server(tcpConn *net.TCPConn) {
	for {
		b := make([]byte, 100)
		tcpConn.Read(b)
		fmt.Println(string(b))
	}
}


//receive and resole value to v type
func RecAndRes(conn *net.TCPConn, b []byte, v *interface{}){
	len, err := conn.Read(b)
	if err != nil {
		fmt.Println(err)
		//return nil 
		
	}else {
		err = json.Unmarshal(b[:len], v)	
		//return v
		//fmt.Printf("%+v", v)	
	}
}

func CreateMessage(v interface{}) message.Message{
	msg := message.NewMsg()
	msg.Value = v
	return msg
}

func (mqAgent *MqAgent) Start(tcpConn *net.TCPConn) {
	mqAgent.MQ = NewMQ()
	mqAgent.TCPConn = tcpConn
	b := make([]byte, 1024)
	var value interface{} = message.AD{}
	RecAndRes(tcpConn, b, &value)
	//fmt.Printf("%+v", value)
	msg := CreateMessage(value)
	mqAgent.MQ.List.PushBack(msg)
	fmt.Printf("%+v", mqAgent.MQ.List.Front().Value)
	//fmt.Printf("%+v", msg)
//	fmt.Println()
//	fmt.Println(msg.Value)
}

func (mqAgent *MqAgent) AddConn() {
	
}



