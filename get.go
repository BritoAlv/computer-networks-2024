package main

import (
	"errors"
	"os"
	"strings"
)

func (cs *CommandsStruct) GET(arg string) (string, error) {
	if arg[len(arg)-1] == 'A' {
		return command_get(cs, strings.TrimSpace(arg[:len(arg)-1]), false)
	} else if arg[len(arg)-1] == 'B'{
		return command_get(cs, strings.TrimSpace(arg[:len(arg)-1]), true)
	} else{
		return "", errors.New("wrong arguments")
	}
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
