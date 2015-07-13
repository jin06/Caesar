package sflag

import (
	"flag"
	"fmt"
	"os"
	"regexp"  
	
	"github.com/tsuru/config"
	"github.com/jin06/Caesar/log"
)
 
var (
	listenFlag = flag.String("address", "", "local addr")
	versionFlag = flag.Bool("version", false, "Caesar Server version")
	runtimeFlag = flag.String("runtime", "", "start mode")
)

var (
	//IP and port regexp.   "201.201.122.132:1234" is a correct address.
	ipPattern string = "(2[5][0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\:([1-9][0-9]+$)"
)

func FlagResolve(listenAddr *string) {
	flag.Parse()
	
	if *versionFlag != false {
		err := config.ReadConfigFile("../config/server.yml")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		version, _ := config.GetString("version")
		update, _ := config.GetList("update")
		instruction, _ := config.GetString("instruction")
		fmt.Printf("CaeserClient version: %s\n", version)
		fmt.Printf("New speciality contrast to old version: \n")
		for k, v := range update {
			fmt.Printf("%d-- %s\n",k+1,v)
		}
		fmt.Printf("       %s\n", instruction)
		os.Exit(0)
	}
	if *runtimeFlag == "min" {
		log.Log("warn", "Start up with least CPU's resousces", log.Fields{"Occupy CPU Num": 1})
	}
	
	if *listenFlag != "" {
		fp, err := regexp.MatchString(ipPattern, *listenFlag)
		handleError(err)
		if !fp{  
			//fmt.Printf("\"%s\" is not a valid address, please check it and try again!\n", *listenFlag)
			log.Log("fatal", "-address is not a valid address, please check it and try again!", map[string]interface{}{"address":*listenFlag})
		os.Exit(0)
		}
		*listenAddr = *listenFlag
		fmt.Println("--Notice: you have set a new listen address", *listenAddr)
	}else {
		//fmt.Println("--Didn't set the listen address. Server will start at default address.")
		log.Log("info", "Didn't set the listen address. Server will start at default address.", log.Fields{"Default Address": *listenAddr})
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

