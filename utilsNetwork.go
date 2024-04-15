package main

import (
	"errors"
	"io"
	"net"
	"os"
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

func writeonMemoryDefault(connData *net.Conn, dataS string) (int, error) {
	data := []byte(dataS)
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

func writeonMemoryPassive(connData *net.Conn, data []byte ) (int, error) {
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

func readOnMemoryDefault(connData *net.Conn) (string, error) {
	tmp := make([]byte, max_size)
	(*connData).SetReadDeadline(time.Now().Add(timeout * time.Second))
	n, err := (*connData).Read(tmp)
	if err != nil {
		return "", err
	}
	response := string(tmp[:n])
	err = CheckResponseNumber(response[:3])
	if err != nil {
		return response, err
	}
	return string(tmp[:n]), nil
}

func writeAndreadOnMemory(connData *net.Conn, data string) (string, error) {
	data += "\r\n"
	_, err := writeonMemoryDefault(connData, data)
	if err != nil {
		return "", err
	}
	return readOnMemoryDefault(connData)
}

func readOnFile(connData *net.Conn, file *os.File, size int64) error {
	offset := int64(0)
	tmp := make([]byte, max_size)
	for offset != size {
		(*connData).SetReadDeadline(time.Now().Add(timeout * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			continue
		}
		file.WriteAt(tmp[:n], offset)
		offset += int64(n)
	}
	return nil
}

func CheckResponseNumber(code string) (error) {
	message, exists := FTPErrorMessages[code]
	if !exists {
		return errors.New("La response " + code + " n'est pas reconnue.")
	}
	if message[1] == "" {
		return nil
	}
	return errors.New(message[1])
}