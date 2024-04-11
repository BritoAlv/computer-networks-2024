package main

import (
	"strings"
)

func (cs *CommandsStruct) RNTO(newName string) (string, error) {
	// Lets assume that RNFR was emitted. 
	command := "RNTO " + newName + "\r\n"
	response, err := writeAndreadOnMemory(cs.connection, []byte(command))
	if err != nil {
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	return ParseFTPCode(code)
}