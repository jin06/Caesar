package client

import (
	"strings"
	//"regexp"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/jin06/Caesar/db"
	"github.com/jin06/Caesar/log"
)

func CmdPara(cmd string) bool{
	
	if strings.HasPrefix(cmd, "start") {
		arr := []int{}
		cmd := strings.TrimPrefix(cmd, "start")
		sArr := strings.Fields(cmd)
		res := ""
		if len(sArr) >0 {
			for _, v := range sArr {
				i, err := strconv.Atoi(v)
				arr = append(arr, i)
				if err != nil {
					fmt.Println(err)
					return true
				}
				//fmt.Println(arr[k])
			}
			rpcClient.Call("MqAgentStat.StartMq", arr, &res)
			fmt.Println(res)
		}else {
			fmt.Println("No parameters.----For example start 1234")
		}
		return true
	}
	
	if strings.HasPrefix(cmd, "stop") {
		arr := []int{}
		cmd := strings.TrimPrefix(cmd, "stop")
		sArr := strings.Fields(cmd)
		res := ""
		if len(sArr) >0 {
			for _, v := range sArr {
				i, err := strconv.Atoi(v)
				arr = append(arr, i)
				if err != nil {
					fmt.Println(err)
					return true
				}
				//fmt.Println(arr[k])
			}
			err := rpcClient.Call("MqAgentStat.StopMq", arr, &res)
			handleError(err)
			fmt.Println(res)
		}else {
			fmt.Println("No parameters.----For example stop 1234")
		}
		return true
	}
	
	if strings.HasPrefix(cmd, "create mq"){
		if checkLogined() {
			return true
		}
		
		res := ""
		mqTable := new(db.Mq_Table)
		mqTable.Id = genMqIdKey()
		mqTable.User_Name = Me.Name
		
		fmt.Println("Please input mq name:")
		line, _, err := r.ReadLine()
		handleError(err)
		mqTable.Name = string(line)
		fmt.Println("Please input mq type(int type):")
		line, _, err = r.ReadLine()
		handleError(err)
		mqTable.Type, err = strconv.Atoi(string(line))
		handleError(err)
		if err != nil {
			return true
		}
		fmt.Println("Is persistent(message save to db, 0 is false, 1 is true)?")
		line, _, err = r.ReadLine()
		handleError(err)
		mqTable.Bool_Persist, err = strconv.Atoi(string(line)) 
		handleError(err)
		if err != nil {
			return true
		}
		//fmt.Println(mqTable.Id,mqTable.User_Name, mqTable.Name, mqTable.Type,mqTable.Bool_Persist)
		
		err = rpcClient.Call("Mq_Table.CreateMQ", &mqTable, &res)
		handleError(err)
		fmt.Println(res)
		return true
	}
	
		if strings.HasPrefix(cmd, "delete mq") {
		arr := []int{}
		cmd := strings.TrimPrefix(cmd, "delete mq")
		sArr := strings.Fields(cmd)
		res := ""
		if len(sArr) >0 {
			for _, v := range sArr {
				i, err := strconv.Atoi(v)  
				arr = append(arr, i)
				if err != nil {
					fmt.Println(err)
					return true
				}
				//fmt.Println(arr[k])
			}
			err := rpcClient.Call("Mq_Table.DeleteMq", arr, &res)
			handleError(err)
			fmt.Println(res)
		}else {
			fmt.Println("No parameters.----For example delete mq 1234")
		}
		return true
	}
	
	return false
}

func genMqIdKey() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Int31n(100000))
}

func checkLogined() bool{
	if StatusLine == 0{
		log.Log("info", "You are offline, please login.", nil)
		return true
	}else {
		return false
	}
}

