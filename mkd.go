package main

import "strings"

func (cs *CommandsStruct) MKD(input string) (string, error) {

	directories := strings.Split(input, "/")
	var response []byte
	var err error
	currentDirectory := directories[0]

	for i := 0; i < len(directories); i++ {
		if i > 0 {
			currentDirectory += "/" + directories[i]
		}

		response, err = writeAndreadOnMemory(cs.connection, []byte("MKD "+currentDirectory+"\r\n"))
		if err != nil {
			return "", err
		}
	}

	return ParseFTPCode(string(response)[0:3])
}
