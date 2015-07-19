package control

import (
	"sync"
	"strconv"
	//"log"
	"fmt"
	
	"github.com/jin06/Caesar/message"
	"github.com/jin06/Caesar/msgqueue"
	"github.com/jin06/Caesar/db"
	"github.com/jin06/Caesar/log"
	
	"github.com/ant0ine/go-json-rest/rest"
)

var Counter = 10000
var Flag = 21

type PerMqAgent struct {
	sync.RWMutex
	Msg *message.Message
}

func (perAgent PerMqAgent)PostMsg(w rest.ResponseWriter, r *rest.Request) {
	//msg := message.NewMsg()
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := msgqueue.DefaultMM[mqid]
	if ok {
		perAgent.Msg.MQid = mqid
		perAgent.Msg.Generator = mq.Owner
		perAgent.Msg.MsgId = Counter
		Counter++
		r.DecodeJsonPayload(perAgent.Msg)
		
			//fmt.Println(mqmsg.Msg.Value)
		
		//fmt.Println("post msg")
		mq.Lock()
		//mq.AddMsg(*mqmsg.Msg)
		err = db.PushMsgToDB(perAgent.Msg)
		mq.Unlock()
		if err !=nil {
			w.WriteJson(map[string]string{"1011": "server receive, but not save to db"})
			return 
		}
		//w.WriteJson(perAgent.Msg)
		w.WriteJson(map[string]string{"1016": "post success"})
		//flag, err := db.GetMsgFlag(perAgent.Msg.MsgId)
//		if flag == 0 {
//			log.Log("err", err.Error(), nil)
//		}else {
//			Flag = flag
//		}
		fmt.Println(Flag)
	}else {
		w.WriteJson(map[string]string{"1010": "mq not running"})
	}
}

func (perAgent PerMqAgent)GetMsg(w rest.ResponseWriter, r *rest.Request) {
	//msg := message.NewMsg()
	mqid, err := strconv.Atoi(r.PathParam("mqid"))
	if err != nil {
		log.Log("info", err.Error(), nil)
	}
	mq, ok := msgqueue.DefaultMM[mqid]
	if ok {
		//fmt.Println("post msg")
		mq.Lock()
		msg, _ := db.PopMsgFromDBByFlag(Flag)
		mq.Unlock()
		if msg == nil {
			w.WriteJson(map[string]string{"1011": "no message in db"})
		}else {
			w.WriteJson(msg)
			Flag++
			//w.WriteJson(map[string]string{"1001": "get success"})
		}
	}else {
		w.WriteJson(map[string]string{"1010": "mq not running"})
	}
}
