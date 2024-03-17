package main

import (
	"fmt"
)

func (cs *CommandsStruct) CD(command string){
	response := wr(cs.connection, []byte("CWD " + command[3:] + "\r\n"))
	if starts_with(string(response), "250") {
		fmt.Println("Todo en talla")
	}
}