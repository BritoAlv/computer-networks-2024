package main

import "strings"

func (cs *CommandsStruct) ACCT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "ACCT " + strings.TrimSpace(args))
}