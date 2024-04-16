package main

import (
	"io"
	"os"
	"strings"
)



func (cs *CommandsStruct) PUT(arg string) (string, error) {
	useUnique := false
	useBinary := true
	useAppend := false
	if strings.HasPrefix(arg, append_flag) {
		useAppend = true
		arg = strings.TrimSpace(arg[len(append_flag):])
	}
	if strings.HasPrefix(arg, unique_flag) {
		useUnique = true
		arg = strings.TrimSpace(arg[len(unique_flag):])
	}
	if strings.HasPrefix(arg, binary_flag) {
		useBinary = true
		arg = strings.TrimSpace(arg[len(binary_flag):])
	}
	if strings.HasPrefix(arg, ascii_flag) {
		useBinary = false
		arg = strings.TrimSpace(arg[len(ascii_flag):])
	}
	return command_store(cs, strings.TrimSpace(arg), useUnique, useBinary, useAppend)
}

func command_store(cs *CommandsStruct, filename string, useUnique bool, useBinary bool, useAppend bool) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if useBinary {
		_, err := cs.TYPE("I")
		if err != nil {
			return "", err
		}
	}
	conn_data, err := cs.PASV()
	if err != nil {
		return "", err
	}
	command := "STOR "
	if useAppend {
		command = "APPE "
	}
	if useUnique {
		command = "STOU "
	}
	_, err = writeAndreadOnMemory(cs.connectionConfig, command+filename)
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
		_, err = writeonMemoryPassive(conn_data, buffer[:bytesRead])
		if err != nil {
			return "", err
		}
	}
	(*conn_data).Close()
	result, err := readOnMemoryDefault(cs.connectionConfig)
	if err != nil {
		return "", err
	}
	_, err = cs.TYPE("A")
	if err != nil {
		return "", err
	}
	return string(result), nil
}