package main

import "strings"

func (cs *CommandsStruct) CD(args string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "CWD " + strings.TrimSpace(args))
}