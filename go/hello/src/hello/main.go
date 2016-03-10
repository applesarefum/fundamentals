package main

import (
	"hello"
	"os"
)

func main() {
	name := "Andrew"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	hello.PrintForever("Hey " + name + ", whasssuuuuuuup?")
}
