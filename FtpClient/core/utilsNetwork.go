package core

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

func writeonMemoryDefault(cs *FtpSession, dataS string) (int, error) {
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

func getResponse(cs *FtpSession) (string, error) {
	result := cs.queueResponses.Dequeue()
	respCode := result[:3]
	err := CheckResponseNumber(respCode)
	if err != nil {
		return "", err
	}
	return result, nil
}

func readOnMemoryDefault(cs *FtpSession) (string, error) {
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
		response := string(data)
		lines := SplitString(response, '\n')
		for len(lines) > 0 {
			responseCode := lines[0][0:3]
			found := false
			for i := 0; i < len(lines); i++ {
				if strings.HasPrefix(lines[i], responseCode+" ") {
					found = true
					cs.queueResponses.Enqueue(strings.Join(lines[0:i+1], "\n"))
					if i+1 < len(lines) {
						lines = lines[i+1 :]
					} else {
						lines  = make([]string, 0)
					}
					break
				}
			}
			if !found {
				(*cs.connectionConfig).SetReadDeadline(time.Now().Add(timeout * time.Second))
				n, err := (*cs.connectionConfig).Read(tmp)
				if err != nil {
					return "", err
				}
				newLines := SplitString(string(tmp[:n]), '\n')
				lines = append(lines, newLines...)
			}
		}
	}
	return getResponse(cs)
}

func writeAndreadOnMemory(cs *FtpSession, data string) (string, error) {
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

