package core

import "strings"

func (cs *FtpSession) CD(args string) (string, error) {
	return writeAndreadOnMemory(cs, "CWD " + strings.TrimSpace(args))
}