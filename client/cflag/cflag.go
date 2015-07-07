
/*
	Package cflag provides method of resolving the CaesarClient.go flag.
*/
package cflag

import (
	"flag"
	"fmt"
	"os"
	"regexp"  
)

var (
	localFlag = flag.String("local", "", "local addr")
	serverFlag = flag.String("server", "", "server addr and port")
	userFlag = flag.String("user", "guest", "guest client")
	passwordFlag = flag.String("password", "123456", "password")
	helpFlag = flag.Bool("help",false,"Print usage")
	flVersion = flag.Bool( "version", false, "Print version information and quit")
)

var (
	//IP and port regexp.   "201.201.122.132:1234" is a correct address.
	ipPattern string = "(2[5][0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\.(25[0-5]|2[0-4]\\d|1\\d{2}|\\d{1,2})\\:([1-9][0-9]+$)"
)

func FlagResolve(localAddr *string, serverAddr *string, username *string, password *string) {
	flag.Parse()
	
	if *helpFlag != false {
		fmt.Println("Usage:...........")
		os.Exit(0)
	}
	if *flVersion != false {
		fmt.Println("Caesar1.0")
		os.Exit(0)
	}
	if *localFlag != "" {
		*localAddr = *localFlag
	}
	if *serverFlag != "" {
		fp, err := regexp.MatchString(ipPattern, *serverFlag)
		handleError(err)
		if !fp{  
			fmt.Printf("\"%s\" is not a valid address, please check it and try again!\n", *serverFlag)
		os.Exit(0)
		}
		*serverAddr = *serverFlag
	}
	if *userFlag != "" {
		*username = *userFlag
	}
	if *passwordFlag != "" {
		*password = *passwordFlag
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}