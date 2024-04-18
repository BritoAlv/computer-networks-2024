package core

import "strings"

func (cs *FtpSession) MODE(typeC string) (string, error) {
	return writeAndreadOnMemory(cs, "MODE " + strings.TrimSpace(typeC))
}