package hello

import (
	"fmt"
	col "github.com/colors"
	"os"
	"time"
)

func printOrange(message string) {
	fmt.Println(col.Orange + message + col.X)
}

func printBlue(message string) {
	fmt.Println(col.Blue + message + col.X)
}

func printViolet(message string) {
	fmt.Println(col.Violet + message + col.X)
}

func PrintMulti(message string) {
	fmt.Println(col.Multi(message) + col.X)
}

func PrintForever(message string) {
	for {
		printOrange(message)
		time.Sleep(700)
		printBlue(message)
		time.Sleep(700)
		printViolet(message)
		time.Sleep(700)
	}

func GetArgs() (string,string){
	who="Andrew"
	option=""
	if len(os.Args)=1{
		if strings.HasPrefix(os.Args[1]){}
	}
}
