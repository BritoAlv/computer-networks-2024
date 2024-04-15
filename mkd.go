package main

import "strings"

func (cs *CommandsStruct) MKD(input string) (string, error) {

	directories := strings.Split(input, "/")
	currentDirectory := directories[0]

	for i := 0; i < len(directories); i++ {
		if i > 0 {
			currentDirectory += "/" + directories[i]
		}
		_, err := writeAndreadOnMemory(cs.connectionConfig, "MKD "+currentDirectory)
		if err != nil {
			return "", err
		}
	}
	return "Seems Like Done", nil
}
