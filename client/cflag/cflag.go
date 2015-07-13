
/*
	Package cflag provides method of resolving the CaesarClient.go flag.
*/
package cflag

import (
	"flag"
	"fmt"
	"os"
	"regexp"  
	
	"github.com/tsuru/config"
	"github.com/jin06/Caesar/log"
)

var (
	localFlag = flag.String("local", "", "local addr")
	serverFlag = flag.String("server", "", "server addr and port")
	userFlag = flag.String("user", "", "guest client")
	passwordFlag = flag.String("password", "", "password")
	helpFlag = flag.Bool("help",false,"Print usage")
	versionFlag = flag.Bool( "version", false, "Print version information and quit")
)

var (
	//IP and port regexp.   "201.201.122.132:1234" is a correct address.
	ipPattern string = "(2[5][0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\:([1-9][0-9]+$)"
)

func FlagResolve(localAddr *string, serverAddr *string, username *string, password *string) {
	flag.Parse()
	
	if *helpFlag != false {
		//log.Log("info", "", nil)
		fmt.Println("Usage:...........!!!!!")
		os.Exit(0)
	}
	if *versionFlag != false {
		err := config.ReadConfigFile("../client/config/version.yml")
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
	if *localFlag != "" {
		*localAddr = *localFlag
		log.Log("info", "you set a new addres", log.Fields{"address" : *localFlag})
		//fmt.Println("--Notice: you have set a new address", *localAddr)
	}else {
		//fmt.Println("--Didn't set the start port. Caesar will start at default port.")
		log.Log("info", "Didn't set the start port. Caesar will start at default port.",  log.Fields{"default address" : *localAddr})
	}
	if *serverFlag != "" {
		fp, err := regexp.MatchString(ipPattern, *serverFlag)
		handleError(err)
		if !fp{  
			//fmt.Printf("\"%s\" is not a valid address, please check it and try again!\n", *serverFlag)
			warnMsg := *serverFlag + "is not a valid address, please check it and try again!"
			log.Log("warn", warnMsg, nil)
			os.Exit(0)
		}
		*serverAddr = *serverFlag
		log.Log("info", "You have set a new server address", log.Fields{"new address" : *serverAddr})
		//fmt.Println("--Notice: you have set a new server address", *serverAddr)
	}else {
		log.Log("info", "Didn't set the server address.Caesar will connect the default address.", log.Fields{"new address" : *serverAddr})
		//fmt.Println("--Didn't set the server address. Caesar will connect the default address.")
	}
	if *userFlag != "" && *passwordFlag != ""{
		*username = *userFlag 
		*password = *passwordFlag
		fmt.Println(*username, *password)
	}else {
		//fmt.Println("--Anonymous login, can do nothing! Please login with exgist user or register a new user.")
		log.Log("info", "Anonymous login, can do nothing! Please login with exgist user or register a new user.", nil)
	}
}

func handleError(err error) {
	if err != nil {
		//fmt.Print(err)
		log.Log("err", err.Error(), nil)
	}
}