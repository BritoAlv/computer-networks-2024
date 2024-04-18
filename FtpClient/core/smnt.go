package core

import "strings"

func (cs *FtpSession) SMNT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "SMNT " + strings.TrimSpace(args))
}