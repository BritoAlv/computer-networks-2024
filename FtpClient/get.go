package main

import (
	"os"
	"strconv"
	"strings"
)

func (cs *CommandsStruct) GET(arg string) (string, error) {
	useBinary := true
	if strings.HasPrefix(arg, binary_flag) {
		useBinary = true
		arg = strings.TrimSpace(arg[len(binary_flag):])
	}
	if strings.HasPrefix(arg, ascii_flag) {
		useBinary = false
		arg = strings.TrimSpace(arg[len(ascii_flag):])
	}
	return command_get(cs, strings.TrimSpace(arg), useBinary)
}

func command_get(cs *CommandsStruct, pathname string, useBinary bool) (string, error) {
	defer cs.release_connection()
	parts := strings.Split(pathname, "/")
	filename := parts[len(parts)-1]
	file, _ := os.Create(filename)
	err := cs.check_connection()
	if err != nil {
		return "", err
	}
	if useBinary {
		_, err := cs.TYPE("I")
		if err != nil {
			os.Remove(filename)
			return "", err
		}
	}

	sizeStr, err := cs.SIZE(pathname)
	if err != nil {
		os.Remove(filename)
		return "", err
	}
	
	_, err = writeAndreadOnMemory(cs, "RETR "+pathname)
	if err != nil {
		os.Remove(filename)
		return "", err
	}

	sizeInt, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		os.Remove(filename)
		return "", err
	}

	err = readOnFile(cs.connectionData, file, sizeInt)
	if err != nil {
		os.Remove(filename)
		return "", err
	}
	// this line made the code work !! .
	result, err := readOnMemoryDefault(cs)
	if err != nil {
		return "", err
	}
	_, err = cs.TYPE("A")
	if err != nil {
		return "", err
	}
	return string(result), nil
}