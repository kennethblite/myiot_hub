package main

import(
	"os/exec"
	"os"
	"fmt"
	"strings"
	"bufio"
	"encoding/hex"
)



var userstream DataStream

type DataStream struct{
	//s []Stats 
	username string
	dryer string
}
//gatttool -b 20:91:48:4D:94:BC --char-write-req -a 0x0013 -n 0100 --listen

func main(){
	cmd := exec.Command("gatttool", "-b", "20:91:48:4D:94:BC", "--char-write-req" ,"-a" ,"0x0013" ,"-n", "0100" ,"--listen")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	Reader := bufio.NewReader(stdout)
	for {
	Message, err := Reader.ReadString(10)
		if (err != nil){
			if err.Error() == "EOF" {
				fmt.Println("device Or Resource Busy")
				os.Exit(1)
			}
			fmt.Println(err)
		}	
	//fmt.Println(Message)
	go ProcessReadings(Message)
	}
}



func ProcessReadings(s string){	
	final_string := ""
	s = strings.Replace(s, "\n","",-1)
	s = strings.Replace(s, "Notification handle = 0x0012 value: ", "",-1)
	s = strings.Replace(s, " " , "",-1)
	s = strings.Replace(s, ";" , "",-1)
	final_string = s

	//fmt.Println(final_string)
	decoded, _  := hex.DecodeString(final_string)
	fmt.Println(string(decoded))
}
