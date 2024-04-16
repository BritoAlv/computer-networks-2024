package main

import "strings"

func (cs *CommandsStruct) PORT(args string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "PORT " + strings.TrimSpace(args)) 
}