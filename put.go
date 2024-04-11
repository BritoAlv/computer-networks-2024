package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

func (cs *CommandsStruct) PUT(arg string) (string, error) {
	if arg[len(arg)-1] == 'A' {
		return command_store(cs, strings.TrimSpace(arg[:len(arg)-1]), false)
	} else if arg[len(arg)-1] == 'B' {
		return command_store(cs, strings.TrimSpace(arg[:len(arg)-1]), true)
	} else {
		return "", errors.New("wrong arguments")
	}
}

func command_store(cs *CommandsStruct, filename string, useBinary bool) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if useBinary {
		_, err := writeAndreadOnMemory(cs.connection, []byte("TYPE I\r\n"))
		if err != nil {
			return "", err
		}
	}
	conn_data, err := cs.PASV()
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, []byte("STOR "+filename+"\r\n"))
	if err != nil {
		return "", err
	}
	buffer := make([]byte, max_size)
	for {
		bytesRead, err := file.Read(buffer) // Read the file in chunks
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		_, err = writeonMemory(conn_data, buffer[:bytesRead])
		if err != nil {
			return "", err
		}
	}
	// this line made the code work !! .
	(*conn_data).Close()
	result, err := readOnMemory(cs.connection)
	if err != nil {
		return "", err
	}
	/*
	some folders do not have allowed access.
	*/
	_, err = writeAndreadOnMemory(cs.connection, []byte("TYPE A\r\n"))
	if err != nil {
		return "", err
	}
	return string(result), nil
}
