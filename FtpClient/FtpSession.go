package main

import (
	"errors"
	"net"
	"sync"
)

type FtpSession struct {
	connectionConfig  *net.Conn
	connectionData    *net.Conn
	muRead            sync.Mutex
	muCheckConnection sync.Mutex
	connDataUsed      bool
	queueResponses    Queue
}

func SessionBuilder(exam FTPExample) (*FtpSession, error) {
	conn, err := net.Dial("tcp", exam.ip+":"+exam.port)
	if err != nil {
		return nil, err
	}
	cs := &FtpSession{
		connectionConfig:  &conn,
		connectionData:    nil,
		muRead:            sync.Mutex{},
		muCheckConnection: sync.Mutex{},
		connDataUsed:      false,
		queueResponses:    *NewQueue(),
	}
	_, err = readOnMemoryDefault(cs)
	if err != nil {
		return nil, err
	}
	_, err = cs.USER(exam.user)
	if err != nil {
		return nil, err
	}
	_ , err = cs.PASS(exam.password)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func SessionFinish(cs *FtpSession) {
	(*cs.connectionConfig).Close()
}

func (cs *FtpSession) check_connectionPort(connData *net.Conn) error {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()

	if cs.connectionData == nil || !cs.connDataUsed {
		cs.connectionData = connData
		cs.connDataUsed = false
	} else {
		(*connData).Close()
		return errors.New("already in use")
	}
	return nil
}

func (cs *FtpSession) check_connection() error {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()

	if cs.connectionData == nil {
		_, err := cs.PASV("")
		if err != nil {
			return err
		}
		cs.connDataUsed = true
		return nil
	} else if !cs.connDataUsed {
		cs.connDataUsed = true
		return nil
	} else {
		return errors.New("connectionData is being used by another thread")
	}
}

func (cs *FtpSession) release_connection() {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()
	if (cs.connectionData) != nil {
		(*cs.connectionData).Close()
	}
	cs.connectionData = nil
	cs.connDataUsed = false
}


