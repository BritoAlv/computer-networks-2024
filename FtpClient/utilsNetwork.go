package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func open_conection(connDataConfig string) (net.Conn, error) {
	connDataIP, connDataPort := parse_get_connection_ftp(connDataConfig)
	connData, err := net.Dial("tcp", connDataIP+":"+connDataPort)
	if err != nil {
		return nil, err
	}
	return connData, nil
}

func writeonMemoryDefault(cs *CommandsStruct, dataS string) (int, error) {
	data := []byte(dataS)
	pr := 0
	for {
		(*cs.connectionConfig).SetWriteDeadline(time.Now().Add(timeout * time.Second))
		size_to_write := min(max_size, len(data)-pr)
		n, err := (*cs.connectionConfig).Write(data[pr : pr+size_to_write])
		if err != nil {
			return pr, nil
		}
		if n != size_to_write {
			return n, errors.New("not everything was written")
		}
		pr += n
		if pr == len(data) {
			break
		}
	}
	return pr, nil
}

func writeonMemoryPassive(connData *net.Conn, data []byte) (int, error) {
	pr := 0
	for {
		(*connData).SetWriteDeadline(time.Now().Add(timeout * time.Second))
		size_to_write := min(max_size, len(data)-pr)
		n, err := (*connData).Write(data[pr : pr+size_to_write])
		if err != nil {
			return pr, nil
		}
		if n != size_to_write {
			return n, errors.New("not everything was written")
		}
		pr += n
		if pr == len(data) {
			break
		}
	}
	return pr, nil
}

func readOnMemoryPassive(connData *net.Conn) (string, error) {
	tmp := make([]byte, max_size)
	data := make([]byte, 0)
	for {
		(*connData).SetReadDeadline(time.Now().Add(timeout * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		data = append(data, tmp[:n]...)
	}
	return string(data), nil
}

func getResponse(cs *CommandsStruct) (string, error) {
	result := cs.queueResponses.Dequeue()
	respCode := result[:3]
	err := CheckResponseNumber(respCode)
	if err != nil {
		return "", err
	}
	return result, nil
}

func readOnMemoryDefault(cs *CommandsStruct) (string, error) {
	cs.muRead.Lock()
	defer cs.muRead.Unlock()
	if cs.queueResponses.list.Len() == 0 {
		tmp := make([]byte, max_size)
		data := make([]byte, 0)
		(*cs.connectionConfig).SetReadDeadline(time.Now().Add(timeout * time.Second))
		n, err := (*cs.connectionConfig).Read(tmp)
		if err != nil {
			return "", err
		}
		data = append(data, tmp[:n]...)
		dataStr := string(data)
		lines := SplitString(dataStr, '\n')
		for lines[len(lines)-1][3] != ' ' {
			(*cs.connectionConfig).SetReadDeadline(time.Now().Add(timeout * time.Second))
			n, err := (*cs.connectionConfig).Read(tmp)
			if err != nil {
				return "", err
			}
			data = append(data, tmp[:n]...)
			dataStr := string(data)
			lines = SplitString(dataStr, '\n')
		}
		start := 0
		for index, line := range lines {
			if len(line) > 3 && line[3] == ' ' {
				cs.queueResponses.Enqueue( strings.Join(lines[start : index+1], "\n"))
				start = index+1
			}
		}
	}
	return getResponse(cs)
}

func writeAndreadOnMemory(cs *CommandsStruct, data string) (string, error) {
	data += "\r\n"
	_, err := writeonMemoryDefault(cs, data)
	if err != nil {
		return "", err
	}
	return readOnMemoryDefault(cs)
}

func readOnFile(connData *net.Conn, file *os.File, size int64) error {
	offset := int64(0)
	tmp := make([]byte, max_size)
	if size == no_size {
		for {
			(*connData).SetReadDeadline(time.Now().Add(timeout * time.Second))
			n, err := (*connData).Read(tmp)
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			file.WriteAt(tmp[:n], offset)
			offset += int64(n)
		}
	} else {
		for offset != size {
			(*connData).SetReadDeadline(time.Now().Add(timeout * time.Second))
			n, err := (*connData).Read(tmp)
			if err != nil {
				continue
			}
			file.WriteAt(tmp[:n], offset)
			offset += int64(n)
		}
	}
	return nil
}

func CheckResponseNumber(code string) error {
	message, exists := FTPErrorMessages[code]
	if !exists {
		return errors.New("La response " + code + " n'est pas reconnue.")
	}
	if message[1] == "" {
		return nil
	}
	return errors.New(message[1])
}