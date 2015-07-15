package control

import (
	
	//"error"
	"fmt"
	"github.com/jin06/Caesar/msgqueue"
	"github.com/jin06/Caesar/db"
	"strconv"
)

type MqAgentStat struct {
	MqIdArr []int
	MqNameArr []string
}

func (mastat *MqAgentStat) StartMq(arr []int,res *string) error {
	//fmt.Println("dd")
	already := ""
	nonexistent := ""
	startsuccess := ""
	for _, v := range arr {
		_, ok := msgqueue.DefaultMM[v]
		if ok {
			fmt.Println("Mqid=",v," already start")
			already += strconv.Itoa(v) + " "
		}else {
			mq2, err := db.GetMqDataById(v)
			if err != nil {
				fmt.Println("Mqid=", v, " is not non-existent.")
				nonexistent += strconv.Itoa(v) + " "
			}else {
				msgqueue.AddMQ(mq2)
				startsuccess += strconv.Itoa(v) + " "
				fmt.Println("Mqid:",strconv.Itoa(v) ,"start up.")
			}
		}
	}
	if startsuccess != "" {
		*res += startsuccess + "start success!  "
	}
	if already != "" {
		*res += already + "already started!!  "
	}
	if nonexistent != "" {
		*res += nonexistent + "not exist!  "
	}
	return nil
}

func (mastat *MqAgentStat) StopMq(arr []int,res *string) error {
	//fmt.Println("dd")
	alreadystoppe := ""
	nowstop := ""
	//nonexistent := ""
	for _, v := range arr {
		_, ok := msgqueue.DefaultMM[v]
		if ok {
			fmt.Println("Mqid=",v," is running, now will stop.")
			nowstop += strconv.Itoa(v) + " "
			msgqueue.DeleteMQbyId(v)
		}else {
			fmt.Println("Mqid=", v, " is not running.")
			alreadystoppe += strconv.Itoa(v) + " "
		}
	}
	if nowstop != "" {
		*res += nowstop + " stopped!  "
	}
	if alreadystoppe != "" {
		*res += alreadystoppe + " have stopped or nonexistent!!  "
	}
	return nil
}
