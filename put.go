package main

import (
	"io"
	"os"
	"strings"
)

func (cs *CommandsStruct) PUT(command string) (string, error) {
	// split command in args.
	args := strings.Split(command, " ")
	if len(args) < 2 || (args[1] != "A" && args[1] != "B") {
		return "Provide Arguments: put filename binary/ascii" + "put file.go A" + "put file.mp4 B", nil
	}
	if args[1] == "A" {
		return command_store(cs, strings.TrimSpace(args[0]), false)
	}
	if args[1] == "B" {
		return command_store(cs, strings.TrimSpace(args[0]), true)
	}
	return "", nil
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
	_, err = writeAndreadOnMemory(cs.connection, []byte("TYPE A\r\n"))
	if err != nil {
		return "", err
	}
	return string(result), nil
}