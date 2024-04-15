package main

import "strings"

func (cs *CommandsStruct) APPEND(arg string) (string, error) {
	return cs.PUT(append_flag + " " + strings.TrimSpace(arg))
}