package main

import "strings"

func (cs *CommandsStruct) CD(args string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("CWD " + args + "\r\n"))
	if err != nil{
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	return ParseFTPCode(code)
}