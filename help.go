package main

import (
	"strings"
)

func (cs *CommandsStruct) HELP(args string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("HELP"+"\r\n"))
	if err != nil {
		return "There was something wrong", err
	}
	// TODO: we should give something like what this gives but with our commands
	code := strings.Split(string(response), "-")[0]
	if code == "214" {
		return string(response), nil
	}
	return ParseFTPCode(code)
}
