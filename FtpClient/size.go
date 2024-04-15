package main

import (
	"strings"
)

func (cs *CommandsStruct) SIZE(arg string) (string, error) {
	result, err := writeAndreadOnMemory(cs.connectionConfig, "SIZE " + arg)
	if err != nil {
		return "0", err
	}
	sizeStr := string(result)
	sizeStr = strings.Split(sizeStr, " ")[1]
	sizeStr = strings.Split(sizeStr, "\r\n")[0]
	return sizeStr, nil
}