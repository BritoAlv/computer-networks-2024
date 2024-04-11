package main

import (
	"fmt"
	"strings"
)

func (cs *CommandsStruct) HELP(args string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("HELP" +"\r\n"))
	
	fmt.Print(string(response))
	fmt.Print(err)
	
	if err != nil {
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	return ParseFTPCode(code)
}