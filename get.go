package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func (cs *CommandsStruct) GET(path string) (string, error) {
	if path[len(path)-1] == 'A' {
		return command_get(cs, strings.TrimSpace(path[:len(path)-1]), false)
	} else if path[len(path)-1] == 'B' {
		return command_get(cs, strings.TrimSpace(path[:len(path)-1]), true)
	} else {
		return "", errors.New("wrong arguments")
	}
}

func command_get(cs *CommandsStruct, pathname string, useBinary bool) (string, error) {
	parts := strings.Split(pathname, "/")
	filename := parts[len(parts)-1]
	file, _ := os.Create(filename)
	connData, err := cs.PASV()
	if err != nil {
		return "", err
	}
	if useBinary {
		_, err := writeAndreadOnMemory(cs.connection, "TYPE I")
		if err != nil {
			return "", err
		}
	}
	size, err := writeAndreadOnMemory(cs.connection, "SIZE " + pathname)
	if err != nil {
		return "", nil
	}
	
	_, err = writeAndreadOnMemory(cs.connection, "RETR "+ pathname )
	if err != nil {
		return "", err
	}

	// convert size
	sizeStr := string(size)

	sizeStr = strings.Split(sizeStr, " ")[1]
	sizeStr = strings.Split(sizeStr, "\r\n")[0]

	sizeint, _ := strconv.ParseInt(sizeStr, 10, 64)

	err = readOnFile(connData, file, sizeint)
	if err != nil {
		os.Remove(file.Name())
		return "", err
	}
	// this line made the code work !! .
	(*connData).Close()
	result, err := readOnMemoryDefault(cs.connection)
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, "TYPE A")
	if err != nil {
		return "", err
	}
	return string(result), nil
}