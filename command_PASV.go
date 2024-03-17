package main

import (
	"fmt"
	"net"
)

func command_PASV(connConfig *net.Conn) *net.Conn {
	connData, err := open_conection(wr(connConfig, []byte("PASV \r\n")))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &connData
}