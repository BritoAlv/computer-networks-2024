package main

import (
	"errors"
	"net"
	"sync"
)

type CommandsStruct struct {
	connectionConfig  *net.Conn
	connectionData    *net.Conn
	muRead            sync.Mutex
	muCheckConnection sync.Mutex
	connDataUsed      bool
	queueResponses    Queue
}

func (cs *CommandsStruct) check_connectionPort( connData *net.Conn ) error {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()

	if cs.connectionData == nil || !cs.connDataUsed{
		cs.connectionData = connData
		cs.connDataUsed = false
	} else {
		(*connData).Close()
		return errors.New("already in use")
	}
	return nil
}

func (cs *CommandsStruct) check_connection() error {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()

	if cs.connectionData == nil {
		err := cs.PASV()
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

func (cs *CommandsStruct) release_connection() {
	cs.muCheckConnection.Lock()
	defer cs.muCheckConnection.Unlock()
	(*cs.connectionData).Close()
	cs.connectionData = nil
	cs.connDataUsed = false
}