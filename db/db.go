package db

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"

	"github.com/jin06/Caesar/log"
	. "github.com/jin06/Caesar/msgqueue"
	"github.com/jin06/Caesar/message"
	"github.com/tsuru/config"
	//"fmt"
	"errors" 
	//"time"
)

var (
	user   string //= "beijing"
	pwd    string  // = "beijing"
	dbname string //= "caesar"
	//User string = "beijing"
	dbaddress string 
)

var (
	DB mysql.Conn//= mysql.New("tcp", "",dbaddress , user, pwd, dbname)
)

func InitDBService() {
	err := DB.Connect()
	if err != nil {
		log.Log("err", err.Error(), nil)
	}
}

func init() {
	err := config.ReadConfigFile("../config/db.yaml")
	if err != nil {
		//fmt.Print(err)
		log.Log("err", err.Error(), nil)
	}else {
		log.Log("info", "DB config read!", nil)
	}   
	user, err = config.GetString("username")
	handleErr(err)
	pwd, err = config.GetString("password")
	handleErr(err)
	dbname, err = config.GetString("dbname")
	handleErr(err)
	dbaddress, err = config.GetString("address")
	handleErr(err)
	//log.Log("info", "", log.Fields{"username":user,"password":pwd,"dbname":dbname,"address":dbaddress})
	DB = mysql.New("tcp", "", dbaddress , user, pwd, dbname)
}

func handleErr(err error) {
	if err != nil {
		log.Log("err", err.Error(), nil)
	}
}

//query user table, check user
func VerifyUser(usr, pwd string) (int, error) {
	return verifyUser(usr, pwd)
}

//confirm user's infomation
func verifyUser(usr, pwd string) (int, error) {
	err := DB.Connect()
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from user where user_name='%s'", usr)
	//fmt.Println(rows[0][3])
	switch l := len(rows); l {
	case 0:
		return 0, nil
	case 1:
		if rows[0].Str(3) == pwd {
			return 1, nil
		} else {
			return 2, nil
		}
	default:
		return 3, errors.New("DB error: there are two users have same name.")
	}
	//return 4, errors.New("Unknow error.")
}

func UpdateRegisterTime(s string) error {
	return nil
}
//
//func CreateMQ(msgq *MsQueue) error {
//	
//	err := DB.Connect()
//	
//	stmt, err := DB.Prepare("insert into msgqueue values (?, ?, ?, ?, ?)")
//	
//	handleErr(err)
//	stmt.Bind(msgq)
//	handleErr(err)
//	//err = getData(msgq) 
//	_, err = stmt.Run() //msgq.MQid, msgq.MQname, msgq.MQType, msgq.Owner, msgq.Persistence
//	handleErr(err)
//	return nil
//}
func CreateMqtoDB(mqTable *Mq_Table) error {
	err1 := DB.Connect()
	defer DB.Close()
	if err1 != nil {
		log.Log("err", err1.Error(), nil)
		return err1
	}
	stmt, err := DB.Prepare("insert into msgqueue values (?, ?, ?, ?, ?)")
	handleErr(err)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	stmt.Bind(mqTable)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//err = getData(msgq) 
	stmt.Bind(mqTable.Id, mqTable.Name, mqTable.Type, mqTable.User_Name, mqTable.Bool_Persist)
	_, err = stmt.Run() //msgq.MQid, msgq.MQname, msgq.MQType, msgq.Owner, msgq.Persistence
	handleErr(err)
	return nil
}

func CreateUsertoDB(userTable *User_Table) error {
	err1 := DB.Connect()
	defer DB.Close()
	if err1 != nil {
		log.Log("err", err1.Error(), nil)
		return err1
	}
	stmt, err := DB.Prepare("insert into user values (?, ?, ?, ?, ?, ?,?,?)")
	handleErr(err)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	stmt.Bind(userTable)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//err = getData(msgq) 
	stmt.Bind(userTable.Id, userTable.Group_id, userTable.Name, userTable.Password, userTable.Register_time,userTable.Sign, userTable.Last_login_time, userTable.Other)
	log.Log("info", string(userTable.Id), nil)
	_, err = stmt.Run() //msgq.MQid, msgq.MQname, msgq.MQType, msgq.Owner, msgq.Persistence
	handleErr(err)
	return nil
}

func DeleteMqById(mqid int) error {
	_, err := GetMqDataById(mqid)
	if err != nil {
		return err
	}
	
	err1 := DB.Connect()
	defer DB.Close()
	if err1 != nil {
		log.Log("err", err1.Error(), nil)
		return err1
	}
	stmt, err := DB.Prepare(`delete from msgqueue where msgqueue_id=?`)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//stmt.Bind(mqid)
	//_, _, err = stmt.Exec(mqid)
	_, _, err = stmt.Exec(mqid)  
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err   
	}
//	num := res.AffectedRows()    
//	fmt.Println(num)  
	return nil
}

func DeleteUserById(userId int) error {
	
	err1 := DB.Connect()
	defer DB.Close()
	if err1 != nil {
		log.Log("err", err1.Error(), nil)
		return err1
	}
	stmt, err := DB.Prepare(`delete from user where user_id=?`)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//stmt.Bind(mqid)
	//_, _, err = stmt.Exec(mqid)
	_, _, err = stmt.Exec(userId)  
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err   
	}
	
	return nil
}

func DeleteMqByUserId(userId int) error {
	err := DB.Connect()
	defer DB.Close()
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	stmt, err := DB.Prepare(`delete from msgqueue where user_id=?`)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	_, _, err = stmt.Exec(userId)  
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err   
	}
	return nil
}

//get message queue data from db
func GetMqFromDB(mqid int) (*MsQueue, error) {
	err := DB.Connect()
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from msgqueue where msgqueue_id='%d'", mqid)
	//fmt.Println(rows[0][3])
	if l := len(rows); l != 1 {
		return nil, errors.New("DB error, more than one message queue with same id.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		return NewMsgQue(rows[0].Int(0),rows[0].Str(1), rows[0].Int(2), rows[0].Str(3) , rows[0].Bool(4)), nil
	}
	//return 4, errors.New("Unknow error.")
}

func GetAllMqFromDB(mqOwner string) ([]*MsQueue, error) {
	err := DB.Connect()
	mqArr := make([]*MsQueue, 20, 40)
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from msgqueue where msgqueue_user='%s'", mqOwner)
	
	if l := len(rows); l == 0 {
		return nil, errors.New("You don't have mq.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		for k,row := range rows {
				mqArr[k] = NewMsgQue(row.Int(0),row.Str(1), row.Int(2), row.Str(3) , row.Bool(4))
		}
		return mqArr, nil
	}
}

func GetAllUsersFromDB() ([]*User_Table, error) {
	err := DB.Connect()
	usersArr := make([]*User_Table, 20, 40)
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from user")
	
	if l := len(rows); l == 0 {
		return nil, errors.New("No user.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		for k,row := range rows {
				usersArr[k] = &User_Table{
					Id:row.Int(0),
					Name:row.Str(2),
					}
		}
		return usersArr, nil
	}
}

func GetMqDataById(mqid int) (*MsQueue, error) {
	err := DB.Connect()
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from msgqueue where msgqueue_id='%d'", mqid)
	
	if l := len(rows); l == 0 {
		return nil, errors.New("You don't have mq.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		mq := NewMsgQue(rows[0].Int(0),rows[0].Str(1), rows[0].Int(2), rows[0].Str(3) , rows[0].Bool(4))
		return mq, nil
	}
}

func CreateMsgtoDB(msg *message.Message) error {
	err1 := DB.Connect()
	defer DB.Close()
	if err1 != nil {
		log.Log("err", err1.Error(), nil)
		return err1
	}
	stmt, err := DB.Prepare("insert into message values (?, ?, ?, ?, ?,?,?,?,?)")
	handleErr(err)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//stmt.Bind(mqTable)
	if err != nil {
		log.Log("err", err.Error(), nil)
		return err
	}
	//err = getData(msgq) 
	stmt.Bind(msg.MsgId, msg.Value, msg.CreatedTime, msg.Generator, msg.EXP, msg.MsgType, msg.SubNum, msg.MQid, nil)
	_, err = stmt.Run() //msgq.MQid, msgq.MQname, msgq.MQType, msgq.Owner, msgq.Persistence
	handleErr(err)
	return nil
}


func GetMsgFlag(msgId int) (int, error) {
	err := DB.Connect()
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from message where message_id='%d'", msgId)
	
	if l := len(rows); l == 0 {
		return 0, errors.New("message not exist.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		flag := rows[0].Int(8)
		return flag, nil
	}
}

func GetMsgByFlag(flag int) (*message.Message, error) {
	err := DB.Connect()
	defer DB.Close()
	handleErr(err)
	rows, _, err := DB.Query("select * from message where flag='%d'", flag)
	
	if l := len(rows); l == 0 {
		return nil, errors.New("message not exist.")
	}else {
		//return NewMsgQue(id int,name string, mqType int, owner string, per bool), nil
		msg := message.Message{
			MsgId:rows[0].Int(0),
			Value:rows[0].Str(1), 
			//CreatedTime:rows[0].Time(2,t), 
			Generator:rows[0].Str(3) ,
			//EXP:time.Hour ,
			MsgType:rows[0].Str(5),
			SubNum:rows[0].Int(6),
			MQid:rows[0].Int(7),
		}
		log.Log("info", "msg",nil)
		return &msg, nil
	}
}

