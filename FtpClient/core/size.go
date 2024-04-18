package core

import (
	"strconv"
	"strings"
)

func (cs *FtpSession) SIZE(arg string) (string, error) {
	result, err := writeAndreadOnMemory(cs, "SIZE " + arg)
	if err != nil {
		return strconv.Itoa(no_size), nil
	}
	sizeStr := string(result)
	sizeStr = strings.Split(sizeStr, " ")[1]
	sizeStr = strings.Split(sizeStr, "\r")[0]
	return sizeStr, nil
}