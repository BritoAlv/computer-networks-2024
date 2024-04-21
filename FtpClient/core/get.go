package core

import (
	"os"
	"strconv"
	"strings"
)

func (cs *FtpSession) GET(arg string) (string, error) {
	useBinary := true
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
		return command_get(cs, parts[0], "/", useBinary)
	}
	return command_get(cs, strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), useBinary)
}

func command_get(cs *FtpSession, pathnameS string, pathnameD string , useBinary bool) (string, error) {
	filename := Get_filename_path(pathnameS)
	file, _ := os.Create(pathnameD + "/" + filename)
	err := cs.check_connection()
	if err != nil {
		cs.release_connection()
		return "", err
	}
	if useBinary {
		_, err := cs.TYPE("I")
		if err != nil {
			cs.release_connection()
			return "", err
		}
	}

	sizeStr, err := cs.SIZE(pathnameS)
	if err != nil {
		cs.release_connection()
		return "", err
	}
	
	_, err = writeAndreadOnMemory(cs, "RETR "+pathnameS)
	if err != nil {
		cs.release_connection()
		return "", err
	}

	sizeInt, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		cs.release_connection()
		return "", err
	}

	err = readOnFile(cs.connectionData, file, sizeInt)
	if err != nil {
		cs.release_connection()
		return "", err
	}
	// this line made the code work !! .
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