package main

import (
	"fmt"
	"net"
)

func (cs *CommandsStruct) PASV(command string)  *net.Conn {
	connData, err := open_conection(wr(cs.connection, []byte("PASV \r\n")))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &connData
}