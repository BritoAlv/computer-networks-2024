package main

import (
	"os"
	"strings"
)



func (cs *CommandsStruct) GET(arg string) (string, error) {
	// split command in args.
	args := strings.Split(arg, " ")
	if len(args) < 2 || (args[1] != "A" && args[1] != "B") {
		return "Bad Arguments" + "Provide Arguments: get filename binary/ascii" + "get file.go A" + "get file.mp4 B", nil
	}
	if args[1] == "A" {
		return command_get(cs, strings.TrimSpace(args[0]), false)
	}
	if args[1] == "B" {
		return command_get(cs, strings.TrimSpace(args[0]), true)
	}
	return "", nil
}

func command_get(cs *CommandsStruct, s string, useBinary bool) (string, error) {
	file, _ := os.Create(s)
	if useBinary {
		_, err := writeAndreadOnMemory(cs.connection, []byte("TYPE I\r\n"))
		if err != nil {
			return "", err
		}
	}
	connData, err := cs.PASV()
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, []byte("RETR "+s+"\r\n"))
	if err != nil {
		return "", err
	}
	err = readOnFile(connData, file)
	if err != nil {
		os.Remove(file.Name())
		return "", err
	}
	// this line made the code work !! .
	(*connData).Close()
	result, err := readOnMemory(cs.connection)
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, []byte("TYPE A\r\n"))
	if err != nil {
		return "", err
	}
	return string(result), nil
}
