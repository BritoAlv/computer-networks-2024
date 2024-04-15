package main

import "strings"

func (cs *CommandsStruct) PUT_U(arg string) (string, error) {
	return cs.GET(unique_flag + " " + strings.TrimSpace(arg))
}