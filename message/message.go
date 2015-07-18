
/*Package message provides message struct.
*/
package message
	
import (
	"time"
//	"sync"
//	"fmt"
//	"strconv"
	//"net/http"
	"github.com/ant0ine/go-json-rest/rest"
)

type Message struct {
	MsgId int
	//Value interface{}
	Value string
	CreatedTime time.Time
	Generator string  //message maker
	EXP time.Duration //expiration date
	MsgType string
	SubNum int
	MQid int
	//sync.RWMutex
	
}

func NewMsg() *Message{
	m := Message{
		MsgId:1,
		CreatedTime:time.Now(),
		Generator:"",   
		EXP:time.Hour,
		MsgType:"",
		SubNum:1,
	}
	return &m  
}

func (msg *Message)TestMsg(w rest.ResponseWriter, r *rest.Request){
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}

//func (msg *Message) GetMsg(w rest.ResponseWriter, r *rest.Request) {
//    mqid := r.PathParam("mqid")
//    msg.RLock()
//    
//    msg.RUnlock()
//    if user == nil {
//        rest.NotFound(w, r)
//        return
//    }
//    w.WriteJson(msg)
//}



//func genMsgId() int {
//	rand.Seed(time.Now().UnixNano())
//	return int(rand.Int31n(100000))
//}

