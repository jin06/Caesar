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
		fmt.Printf("Create message queue name by '%s', type is '%d', persistent is '%d', yes or no?",mqTable.Name,mqTable.Type,mqTable.Bool_Persist)
		line, _, err = r.ReadLine()
		handleError(err)
		if string(line) == "yes" || string(line) == "y" {
			err = rpcClient.Call("Mq_Table.CreateMQ", &mqTable, &res)
			handleError(err)
			fmt.Println(res)
		}else {
			fmt.Println("Create faile, cancel it.")
		}
		
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
			fmt.Printf("Delete message queue, yes or no?")
			line, _, err := r.ReadLine()
			handleError(err)
			if string(line) == "yes" || string(line) == "y" {
				err := rpcClient.Call("Mq_Table.DeleteMq", arr, &res)
				handleError(err)
				fmt.Println(res)
			}else {
				fmt.Println("Delete faile, cancel it.")
			}
		}else {
			fmt.Println("No parameters.----For example delete mq 1234")
		}
		return true
	}
		
	if strings.HasPrefix(cmd, "create user"){
		if checkLogined() {
			return true
		}
		if Me.Name != "admin" {
			log.Log("err", "You are not administrator, no create right!", nil)
			return true
		}
		
		res := ""
		userTable := new(db.User_Table)
		userTable.Id = genUserIdKey()
		userTable.Register_time = time.Now()
		userTable.Sign = 1
		userTable.Group_id = 2
		//userTable.User_Name = Me.Name
		
		fmt.Println("Please input user name:")
		line, _, err := r.ReadLine()
		handleError(err)
		userTable.Name = string(line)
		fmt.Println("Please input password:")
		line, _, err = r.ReadLine()
		handleError(err)
		userTable.Password = string(line)
		//handleError(err)
//		if err != nil {
//			return true
//		}
		
		//fmt.Println(mqTable.Id,mqTable.User_Name, mqTable.Name, mqTable.Type,mqTable.Bool_Persist)
		fmt.Printf("Create user name by '%s', password by '%s', yes or no?",userTable.Name,userTable.Password)
		line, _, err = r.ReadLine()
		handleError(err)
		if string(line) == "yes" || string(line) == "y" {
			err = rpcClient.Call("User_Table.CreateUser", &userTable, &res)
			handleError(err)
			fmt.Println(res)
		}else {
			fmt.Println("Create faile, cancel it.")
		}
		return true
	}	
	if strings.HasPrefix(cmd, "delete user") {
		if Me.Name != "admin" {
			log.Log("err", "You are not administrator, no delete user right!", nil)
			return true
		}
		arr := []int{}
		cmd := strings.TrimPrefix(cmd, "delete user")
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
			fmt.Printf("Delete users, yes or no?")
			line, _, err := r.ReadLine()
			handleError(err)
			if string(line) == "yes" || string(line) == "y" {
				err := rpcClient.Call("User_Table.DeleteUser", arr, &res)
				handleError(err)
				fmt.Println(res)
			}else {
				fmt.Println("Delete faile, cancel it.")
			}
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

func genUserIdKey() int {
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

