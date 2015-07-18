package control

import (
	"github.com/ant0ine/go-json-rest/rest"
    "log"
    "fmt"
    "net/http"
    //"time"
    "sync"
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
	users := Users{
        Store: map[string]*User{},
    }
	
	mqmsg := msgqueue.MqMsg{
		Msg:message.NewMsg(),
	}

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    
    router, err := rest.MakeRouter(
        rest.Get("/users", users.GetAllUsers),
        rest.Post("/users", users.PostUser),
        rest.Get("/users/:id", users.GetUser),
        rest.Put("/users/:id", users.PutUser),
        rest.Delete("/users/:id", users.DeleteUser),
        //rest.Get("/test", msg.TestMsg),
        rest.Get("/testmq/:mqid", msgqueue.DefaultMM.TestMq),
        rest.Post("/send_msg/:mqid", mqmsg.PostMsg),
        rest.Get("/receive_msg/:mqid", mqmsg.GetMsg),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":"+ListenPort, api.MakeHandler()))
}

type User struct {
    Id   string
    Name string
}

type Users struct {
    sync.RWMutex
    Store map[string]*User
}

func (u *Users) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
    u.RLock()
    users := make([]User, len(u.Store))
    i := 0
    for _, user := range u.Store {
        users[i] = *user
        i++
    }
    u.RUnlock()
    w.WriteJson(&users)
}

func (u *Users) GetUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.RLock()
    var user *User
    if u.Store[id] != nil {
        user = &User{}
        *user = *u.Store[id]
    }
    u.RUnlock()
    if user == nil {
        rest.NotFound(w, r)
        return
    }
    w.WriteJson(user)
}

func (u *Users) PostUser(w rest.ResponseWriter, r *rest.Request) {
    user := User{}
    err := r.DecodeJsonPayload(&user)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    u.Lock()
    id := fmt.Sprintf("%d", len(u.Store)) // stupid
    user.Id = id
    u.Store[id] = &user
    u.Unlock()
    w.WriteJson(&user)
}

func (u *Users) PutUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    if u.Store[id] == nil {
        rest.NotFound(w, r)
        u.Unlock()
        return
    }
    user := User{}
    err := r.DecodeJsonPayload(&user)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        u.Unlock()
        return
    }
    user.Id = id
    u.Store[id] = &user
    u.Unlock()
    w.WriteJson(&user)
}

func (u *Users) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    delete(u.Store, id)
    u.Unlock()
    w.WriteHeader(http.StatusOK)
}


