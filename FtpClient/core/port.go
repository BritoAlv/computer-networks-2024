package core

import "strings"

func (cs *FtpSession) PORT(args string) (string, error) {
	resp, err := writeAndreadOnMemory(cs, "PORT "+strings.TrimSpace(args))
	if err != nil {
		return "", err
	}
	conn, err := open_conection("(" + args + ")")
	if err != nil {
		return "", err
	}
	err = cs.check_connectionPort(&conn)
	if err != nil {
		return "", err
	}
	return resp, nil
}
