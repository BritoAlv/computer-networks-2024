package main

import (
	"fmt"
	"net"
)

func command_LIST(connConfig *net.Conn, command string) {
	connData := *command_PASV(connConfig)
	write(connConfig, []byte("LIST \r\n"))
	data, _ := read(&connData)
	fmt.Println(string(data))
	defer connData.Close()
}