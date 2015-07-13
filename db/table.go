package db

import (
	"time"
)
  
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
