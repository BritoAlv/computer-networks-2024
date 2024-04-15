package main

import (
	"strconv"
	"strings"
)

func (cs *CommandsStruct) SIZE(arg string) (int64, error) {
	result, err := writeAndreadOnMemory(cs.connectionConfig, "SIZE " + arg)
	if err != nil {
		return 0, err
	}
	sizeStr := string(result)
	sizeStr = strings.Split(sizeStr, " ")[1]
	sizeStr = strings.Split(sizeStr, "\r\n")[0]

	sizeint, _ := strconv.ParseInt(sizeStr, 10, 64) 
	return sizeint, nil
}