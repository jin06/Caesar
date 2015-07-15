package db

import (
	"time"
	//"github.com/ziutek/mymysql/native"
	//"fmt"
	"github.com/jin06/Caesar/log"
	"strconv"
	//"github.com/ziutek/mymysql/mysql"
	//"github.com/ziutek/mymysql/native"
)
  
func (lead *Mq_Table) CreateMQ(mqtable *Mq_Table, res *string) error {
	//fmt.Println("here")
	//log.Log("info", mqtable.Name, nil)
	err := CreateMqtoDB(mqtable)
	if err != nil {
		*res = "Mq create failed!"
	}else {
		*res = mqtable.Name + " create success."
		log.Log("info", mqtable.Name + " created.", nil)
	}
	return nil
}

func (lead *Mq_Table) DeleteMq(mqId int, res *string) error {
	//fmt.Println("here")
	//log.Log("info", mqtable.Name, nil)
	err := DeleteMqById(mqId)
	if err != nil {
		*res = "Delete failed!"
	}else {
		*res = strconv.Itoa(mqId) + " delete success."
		log.Log("info", strconv.Itoa(mqId) + " delete.", nil)
	}
	return nil
}
  
type User_Table struct {
	Id int
	Group_id int
	Name string
	Password string
	Register_time  time.Time
	Sign int
	Last_login_time time.Time 
	Other string
}

type Mq_Table struct {
	Id        int
	Name      string
	Type      int 
	User_Name   string
	Bool_Persist  int
}


