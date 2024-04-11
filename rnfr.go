package main

import (
	"strings"
)

func (cs *CommandsStruct) RNFR(oldName string) (string, error) {
	command := "RNFR " + oldName + "\r\n"
	response, err := writeAndreadOnMemory(cs.connection, []byte(command))
	if err != nil {
		return "There was something wrong", err
	}
	code := strings.Split(string(response), " ")[0]
	return ParseFTPCode(code)
}