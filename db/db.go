package db

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"

	"github.com/jin06/Caesar/log"
	. "github.com/jin06/Caesar/msgqueue"
	//"fmt"
	"errors"
)

var (
	user   string = "beijing"
	pwd    string = "beijing"
	dbname string = "caesar"
)

var (
	DB = mysql.New("tcp", "", "127.0.0.1:3306", user, pwd, dbname)
)

func InitDBService() {
	err := DB.Connect()
	if err != nil {
		log.Log("err", err.Error(), nil)
	}
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

func CreateMQ(msgq *MsQueue) error {
	
	err := DB.Connect()
	
	stmt, err := DB.Prepare("insert into msgqueue values (?, ?, ?, ?, ?)")
	
	handleErr(err)
	stmt.Bind(msgq)
	handleErr(err)
	//err = getData(msgq) 
	_, err = stmt.Run() //msgq.MQid, msgq.MQname, msgq.MQType, msgq.Owner, msgq.Persistence
	handleErr(err)
	return nil
}
