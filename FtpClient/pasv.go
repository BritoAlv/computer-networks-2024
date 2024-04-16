package main

import (
	"errors"
	"strings"
)

func (cs *CommandsStruct) PASV() error {
	data, err := writeAndreadOnMemory(cs, "PASV ")
	if err != nil {
		return err
	}
	if strings.HasPrefix(data, "227") {
		connData, err := open_conection(data)
		if err != nil {
			return err
		}
		cs.connectionData = &connData
		return nil
	} else {
		return errors.New("PASV got wrong response : " + data)
	}

}
