package message
	
import (
	"time"
)

type Message struct { 
	MsgId int
	Value interface{}
	CreatedTime time.Time
	Generator string  //message maker
	EXP time.Duration //expiration date
	MsgType string    
	SubNum int        
}

func NewMsg() Message{
	m := Message{
		MsgId:1,
		CreatedTime:time.Now(),
		Generator:"jinlong",
		EXP:time.Hour,
		MsgType:"",
		SubNum:1,
	}
	return m
}

