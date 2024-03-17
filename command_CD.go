package main

import (
	"fmt"
	"net"
)

func command_CD(connConfig *net.Conn, command string){
	response := wr(connConfig, []byte("CWD " + command[3:] + "\r\n"))
	if starts_with(string(response), "250") {
		fmt.Println("Todo en talla")
	}
}