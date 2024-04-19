package core

import (
	"errors"
	"io"
	"os"
	"strings"
)

func (cs *FtpSession) PUT(arg string) (string, error) {
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
	arg = strings.TrimSpace(arg)
	parts := strings.Split(arg, Separator)
	if len(parts) == 1 || parts[1] == "" {
		return "", errors.New("where put the file ")
	}
	return command_store(cs, strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), useUnique, useBinary, useAppend)
}

func command_store(cs *FtpSession, fileS string, fileD string, useUnique bool, useBinary bool, useAppend bool) (string, error) {
	file, err := os.Open(fileS)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if useBinary {
		_, err := cs.TYPE("I")
		if err != nil {
			cs.release_connection()
			return "", err
		}
	}
	err = cs.check_connection()
	if err != nil {
		cs.release_connection()
		return "", err
	}
	command := "STOR "
	if useAppend {
		command = "APPE "
	}
	if useUnique {
		command = "STOU "
	}
	_, err = writeAndreadOnMemory(cs, command+fileD)
	if err != nil {
		cs.release_connection()
		return "", err
	}
	buffer := make([]byte, max_size)
	for {
		bytesRead, err := file.Read(buffer) // Read the file in chunks
		if err != nil {
			if err != io.EOF {
				cs.release_connection()
				return "", err
			}
			break
		}
		_, err = writeonMemoryPassive(cs.connectionData, buffer[:bytesRead])
		if err != nil {
			cs.release_connection()
			return "", err
		}
	}
	cs.release_connection()
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