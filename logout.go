package main

import (
	"os"
	"strings"
)

func (cs *CommandsStruct) LOGOUT(args string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("QUIT"+ "\r\n"))
	if err != nil {
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	
	defer os.Exit(0)
	return ParseFTPCode(code)
}