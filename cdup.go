package main

import (
	"strings"
)

func (cs *CommandsStruct) CDUP(args string) (string, error) {
	
	response, err := writeAndreadOnMemory(cs.connection, []byte("CDUP "+ "\r\n"))
	if err != nil{
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	return ParseFTPCode(code)
}