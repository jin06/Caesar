package msgqueue

import (
	"container/list"

	"github.com/jin06/Caesar/message"
	//"message"
)

var DefaultMQ *MsQueue

type MsQueue struct {
	MQid        int        //message queue id
	MQname      string     //message queue name
	List        *list.List //message push in and pull from list
	MQType      int        //message queue type
	Owner       string     //message queue owner, who own this message queue
	Persistence bool       //if true, it means that the message save in database
}

func NewMQ() *MsQueue {
	mq := MsQueue{}
	mq.List = list.New()
	//mq.MQType = ONLINE
	return &mq
}

func NewElement() list.Element {
	element := list.Element{}
	return element
}

func (mq *MsQueue) AddMsg(msg message.Message) {
	mq.List.PushBack(msg)
}

func (mq *MsQueue) PopMsg() *message.Message{
	e := mq.List.Front()
	if e == nil {
		return nil 
	}
	msg := mq.List.Remove(e).(message.Message)
	return &msg
}
