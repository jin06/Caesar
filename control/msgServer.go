package control

import (
	"github.com/ant0ine/go-json-rest/rest"
    "log"
   // "fmt"
    "net/http"
    //"time"
   // "sync"
    "github.com/tsuru/config"
    mylog "github.com/jin06/Caesar/log"
    "github.com/jin06/Caesar/msgqueue"
    "github.com/jin06/Caesar/message"
)

var (
	ListenPort string
	
	//msg = message.Message{}
	
)

func init() {
	err := config.ReadConfigFile("../config/msgserver.yaml")
	if err != nil {
		//fmt.Print(err)
		mylog.Log("err", err.Error(), nil)
	}else {
		mylog.Log("info", "Message server config read!", nil)
	}   
	ListenPort, err = config.GetString("listenport")   
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		mylog.Log("err", err.Error(), nil)
	}
}

func MsgServerStart() {
	
	mqmsg := msgqueue.MqMsg{
		Msg:message.NewMsg(),
	}
	persistent := PerMqAgent{
		Msg:message.NewMsg(),
	}

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    
    router, err := rest.MakeRouter(        
        //rest.Get("/test", msg.TestMsg),
        rest.Get("/testmq/:mqid", msgqueue.DefaultMM.TestMq),
        rest.Post("/send_msg/:mqid", mqmsg.PostMsg),
        rest.Get("/receive_msg/:mqid", mqmsg.GetMsg),
        rest.Post("/send/:mqid", persistent.PostMsg),
        rest.Get("/receive/:mqid", persistent.GetMsg),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":"+ListenPort, api.MakeHandler()))
}






