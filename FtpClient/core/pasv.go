package core

import (
	"errors"
	"strings"
)

func (cs *FtpSession) PASV(args string) (string, error) {
	data, err := writeAndreadOnMemory(cs, "PASV ")
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(data, "227") {
		connData, err := open_conection(data)
		if err != nil {
			return "",err
		}
		cs.connectionData = &connData
		result := data[:len(data)-1] + " Opened the connection for you" 
		return result, nil
	} else {
		return "", errors.New("PASV got wrong response : " + data)
	}
}
