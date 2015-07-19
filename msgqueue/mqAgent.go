package msgqueue

import (
	//"net"
	"fmt"
	"strconv"
	//"net/http"
	//"encoding/json"
	"github.com/jin06/Caesar/message"
	"github.com/jin06/Caesar/log"
	"github.com/ant0ine/go-json-rest/rest"
	
	//"github.com/jin06/Caesar/db"
	
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

func (mqagent MqAgent)GetMsg(w rest.ResponseWriter, r *rest.Request) {
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := mqagent[mqid]
	if ok {
		mq.Lock()
		msg := mq.PopMsg()
		mq.Unlock()
		w.WriteJson(msg.Value)
	}else { //msgqueue nonexist
		
	}
}

func (mqagent MqAgent)PostMsg(w rest.ResponseWriter, r *rest.Request) {
	msg := message.NewMsg()
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := mqagent[mqid]
	if ok {
		msg.MQid = mqid
		r.DecodeJsonPayload(msg.Value)
		mq.Lock()
		mq.AddMsg(*msg)
		mq.Unlock()
		w.WriteJson(msg.Value)
	}else {
		
	}
}

//test message queue , if mq is runing,reuturn true
func (mqagent MqAgent)TestMq(w rest.ResponseWriter, r *rest.Request){
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("err", err.Error(), nil)
	}
	_, ok := mqagent[mqid]
	if ok {
		w.WriteJson(map[string]string{"1020": "mq is running"})
	}else {
		w.WriteJson(map[string]string{"1010": "mq not running"})
	}
	
}


