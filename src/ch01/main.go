package main


import "fmt"

func main(){
	cmd := parseCmd();

	if cmd.versionFlag {
		fmt.Println("Version 0.0.1")
	}else if cmd.helpFlag {
		printUsage()
	}else {
		fmt.Println("now started....")
	}
}
