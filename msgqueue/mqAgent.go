package msgqueue

import (
	//"net"
	"fmt"
	//"encoding/json"
	"github.com/jin06/Caesar/message"
	//"github.com/jin06/Caesar/db"
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

func RuningMq() {
	for _, v := range DefaultMM {
		fmt.Println(v.MQid, v.MQname, v.Owner)
	}
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

//func DeleteMQ(mq *MsQueue) {
//	DefaultMM.DeleteMQ(mq)
//}

func (mqagent MqAgent)DeleteMQbyId(mqid int) {
	delete(mqagent, mqid)
}

func DeleteMQbyId(mqid int) {
	DefaultMM.DeleteMQbyId(mqid)
}

func (mqagent *MqAgent)Test(cmd *[]string, res *string) error{
	*res = "mqagent rpc method test"
	return nil
}


