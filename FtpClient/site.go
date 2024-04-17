package main

import "strings"

func (cs *CommandsStruct) SITE(args string) (string, error) {
	return writeAndreadOnMemory(cs, "SITE " + strings.TrimSpace(args))
}