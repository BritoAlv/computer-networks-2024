package main

import (
	"fmt"
)

func (cs *CommandsStruct) LS(command  string) {
	connData := *cs.PASV("")
	write(cs.connection, []byte("LIST \r\n"))
	data, _ := read(&connData)
	fmt.Println(string(data))
	defer connData.Close()
}