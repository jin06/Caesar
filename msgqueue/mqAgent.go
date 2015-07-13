package msgqueue

import (
	//"net"
	//"fmt"
	//"encoding/json"
	"github.com/jin06/Caesar/message"
	//"net"
)

//var MqMap map[string]MsQueue 
type MqAgent map[int]*MsQueue

var DefaultMM = make(MqAgent)

func (mqagent MqAgent) AddtoMq(msg message.Message) {
	mqagent[msg.MQid].AddMsg(msg) 
}

func NewMqAgent() *MqAgent {
	mqAgent := MqAgent{} 
	//mqAgent.MQ = NewMQ()
	return &mqAgent
}
func AddMQ(mq *MsQueue){
	DefaultMM.AddMQ(mq)
}

func (mqagent MqAgent)AddMQ(mq *MsQueue) {
	mqagent[mq.MQid] = mq
}

func GetMQ(mqid int) *MsQueue{
	return DefaultMM[mqid]
}

func DeleteMQ(mq *MsQueue){
	DefaultMM.DeleteMQ(mq)
}

func (mqagent MqAgent)DeleteMQ(mq *MsQueue) {
	delete(mqagent, mq.MQid)
}


