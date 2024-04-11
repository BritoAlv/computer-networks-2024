package main

import (
	"strings"
)

func (cs *CommandsStruct) RMD(input string) (string, error) {
	// response, err := writeAndreadOnMemory(cs.connection, []byte("RMD "+input+"\r\n"))
	// if err != nil {
	// 	return "", err
	// }

	response, err := RMD_Recursive(cs, input)

	if err != nil {
		return "", err
	}

	return ParseFTPCode(string(response)[0:3])
}

func RMD_Recursive(cs *CommandsStruct, directory string) (string, error) {

	passiveConnection, _ := cs.PASV()

	_, _ = writeAndreadOnMemory(cs.connection, []byte("LIST "+directory+"\r\n"))

	passiveData, _ := readOnMemory(passiveConnection)

	_, _ = readOnMemory(cs.connection)

	files, folders := parseFilesAndFolder(strings.Split(string(passiveData), "\n"))

	for i := 0; i < len(files); i++ {
		writeAndreadOnMemory(cs.connection, []byte("DELE "+directory+"/"+files[i]+"\r\n"))
	}

	for i := 0; i < len(folders); i++ {
		RMD_Recursive(cs, directory+"/"+folders[i])
	}

	response, excep := writeAndreadOnMemory(cs.connection, []byte("RMD "+directory+"\r\n"))

	return string(response), excep
}

func parseFilesAndFolder(directoryDetails []string) (files []string, folders []string) {
	for i := 0; i < len(directoryDetails); i++ {
		directoryDetail := directoryDetails[i]
		if directoryDetail == "" {
			continue
		}

		split := strings.Split(directoryDetail, " ")
		directory := split[len(split)-1]
		directory = directory[:len(directory)-1]

		if directoryDetails[i][0] == 'd' {
			folders = append(folders, directory)
		} else {
			files = append(files, directory)
		}
	}

	return files, folders
}
