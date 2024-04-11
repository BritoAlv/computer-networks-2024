package main

import (
	"net"
)

func (cs *CommandsStruct) PASV() (*net.Conn, error) {
	data, err := writeAndreadOnMemory(cs.connection, []byte("PASV \r\n"))
	if err != nil{
		return nil, err
	}
	response := string(data)
	_, err = ParseFTPCode(response[:3])
	if err != nil {
		return nil, err
	} 
	connData, err := open_conection(string(data))
	if err != nil {
		return nil, err
	}
	return &connData, nil
}