package db

import (
	"github.com/jin06/Caesar/message"
)

func PushMsgToDB(msg *message.Message) error{
	return CreateMsgtoDB(msg)
}

func PopMsgFromDBByFlag(flag int) (*message.Message, error){
	return GetMsgByFlag(flag)
}
