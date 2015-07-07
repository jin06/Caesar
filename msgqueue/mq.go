package msgqueue

import (
	"container/list"
	//"message" 
)

const ( 
	ONLINE = 1	
)

type MsQueue struct {
	List *list.List
	MQType int
}

func NewMQ() *MsQueue {
	mq := MsQueue{}
	mq.List = list.New()
	mq.MQType = ONLINE
	return &mq
}

func NewElement() list.Element {
	element := list.Element{}
	return element
}


