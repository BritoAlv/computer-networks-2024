package main

import "strings"

func (cs *FtpSession) PUT_U(arg string) (string, error) {
	return cs.PUT(unique_flag + " " + strings.TrimSpace(arg))
}