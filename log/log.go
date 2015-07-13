package log

import (
	"github.com/Sirupsen/logrus"
	"time"
)

var (
	timeFormat string = "2006-01-02 15:04:05"
)

type Fields map[string]interface{}

func Log(level string, msg string, fields map[string]interface{}) {
	switch(level){
		case "info" :
			if fields == nil {  
				logrus.Infof("%s--%s",time.Now().Format(timeFormat), msg)
				//logrus.Infof("%s shijian", logrus.Entry.Time.Format("2006-01-02 15:04:05"))
			}else {
				logrus.WithFields(fields).Infof("%s--%s",time.Now().Format(timeFormat), msg)
			}
		case "warn" :
			if fields == nil {
				logrus.Warnf("%s--%s",time.Now().Format(timeFormat), msg)
			}else {
				logrus.WithFields(fields).Warnf("%s--%s",time.Now().Format(timeFormat), msg)
			}
		case "fatal" :
			if fields == nil {
				logrus.Fatalf("%s--%s",time.Now().Format(timeFormat), msg)   
			}else {
				 logrus.WithFields(fields).Fatalf("%s--%s",time.Now().Format(timeFormat), msg) 
			}
		case "err"  :
			if fields == nil {
				logrus.Errorf("%s--%s",time.Now().Format(timeFormat), msg)
			}else {
				logrus.WithFields(fields).Errorf("%s--%s",time.Now().Format(timeFormat), msg)
			}
		case "debug" :
			if fields == nil {
				logrus.Debugf("%s--%s",time.Now().Format(timeFormat), msg)
			}else {
				logrus.WithFields(fields).Debugf("%s--%s",time.Now().Format(timeFormat), msg)
			}
		default :
			logrus.Printf("%s--%s",time.Now().Format(timeFormat), msg)
	}
}