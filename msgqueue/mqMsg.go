package msgqueue

import (
	"sync"
	"strconv"
	"fmt"
	
	"github.com/jin06/Caesar/message"
	"github.com/jin06/Caesar/log"
	
	"github.com/ant0ine/go-json-rest/rest"
)

var Counter int = 1

type MqMsg struct {
	sync.RWMutex
	Msg *message.Message
}

func (mqmsg MqMsg)PostMsg(w rest.ResponseWriter, r *rest.Request) {
	//msg := message.NewMsg()
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := DefaultMM[mqid]
	if ok {
		mqmsg.Msg.MQid = mqid
		mqmsg.Msg.Generator = mq.Owner
		mqmsg.Msg.MsgId = Counter
		Counter++
		r.DecodeJsonPayload(mqmsg.Msg)
		
			fmt.Println(mqmsg.Msg.Value)
		
		//fmt.Println("post msg")
		mq.Lock()
		mq.AddMsg(*mqmsg.Msg)
		mq.Unlock()
		//w.WriteJson(mqmsg.Msg)
		w.WriteJson(map[string]string{"1016": "post success"})
	}else {
		w.WriteJson(map[string]string{"1010": "mq not running"})
	}
}

func (mqmsg MqMsg)GetMsg(w rest.ResponseWriter, r *rest.Request) {
	//msg := message.NewMsg()
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := DefaultMM[mqid]
	if ok {
		//fmt.Println("post msg")
		mq.Lock()
		msg := mq.PopMsg()
		mq.Unlock()
		if msg == nil {
			w.WriteJson(map[string]string{"1015": "no message in mq"})
		}else {
			//w.WriteJson(map[string]string{"1016": "post success"})
			w.WriteJson(msg)
		}
	}else {
		w.WriteJson(map[string]string{"1010": "mq not running"})
	}
}