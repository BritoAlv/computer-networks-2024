package main

import (
	"net"
)

func (cs *CommandsStruct) PASV() (*net.Conn, error) {
	data, err := writeAndreadOnMemory(cs.connection, "PASV ")
	if err != nil{
		return nil, err
	}
	connData, err := open_conection(data)
	if err != nil {
		return nil, err
	}
	return &connData, nil
}