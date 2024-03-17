package main

import (
	"fmt"
)

func (cs *CommandsStruct) CD(args string){
	response := wr(cs.connection, []byte("CWD " + args + "\r\n"))
	if starts_with(string(response), "250") {
		fmt.Println("Todo en talla")
	}
}