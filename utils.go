package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const max_size = 1024

func starts_with(command, prefix string) bool {
	// Function to check if a string starts with a prefix
	return len(command) >= len(prefix) && command[:len(prefix)] == prefix
}

func is_valid_ip(ip string) bool {
	// Function to check if a string is a valid IP
	// This is a very simple check, it only checks if the string has 3 dots and 4 numbers
	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		return false
	}
	for _, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			return false
		}
		if value < 0 || value > 255 {
			return false
		}
	}
	return true
}

func open_conection(connDataConfig string) (net.Conn, error) {
	connDataIP, connDataPort := parse_get_connection_ftp(connDataConfig)
	connData, err := net.Dial("tcp", connDataIP+":"+connDataPort)
	if err != nil {
		return nil, err
	}
	return connData, nil
}

func parse_get_connection_ftp(input string) (string, string) {
	// This method is crying for refactoring.
	// Parse the input string to get the IP and Port
	// Example: "227 Entering Passive Mode (127,0,0,1,195,149)"
	// The port is calculated as (195*256+149)
	// Output: "127.0.0.1", "195*256+149".
	split1 := strings.Split(input, "(")
	split2 := strings.Split(split1[1], ",")
	split3 := strings.Split(split2[5], ")")
	ip := split2[0] + "." + split2[1] + "." + split2[2] + "." + split2[3]
	first_part_port, _ := strconv.ParseInt(split2[4], 10, 32)
	second_part_port, _ := strconv.ParseInt(split3[0], 10, 32)

	port := strconv.FormatInt((first_part_port*256 + second_part_port), 10)
	return ip, port
}

func writeonMemory(connData *net.Conn, data []byte) (int, error) {
	// I wll read from server in chunks of size 1024 KBs.
	pr := 0
	for {
		(*connData).SetWriteDeadline(time.Now().Add(2 * time.Second))
		size_to_write := min(1024, len(data)-pr)
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

func readOnMemory(connData *net.Conn) ([]byte, error) {
	// I wll read from server in chunks of size 1024 KBs.
	tmp := make([]byte, 1024)
	data := make([]byte, 0)
	for {
		(*connData).SetReadDeadline(time.Now().Add(2 * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			if err != io.EOF {
				r := make([]byte, 0)
				return r, err
			}
			break
		}
		data = append(data, tmp[:n]...)
		if n < max_size {
			break
		}
	}
	return data, nil
}

func writeAndreadOnMemory(connData *net.Conn, data []byte) ([]byte, error) {
	_, err := writeonMemory(connData, data)
	if err != nil {
		return nil, err
	}
	return readOnMemory(connData)
}

func readOnFile(connData *net.Conn, file *os.File) error {
	// I wll read from server in chunks of size 1024 KBs.
	tmp := make([]byte, 1024)
	for {
		(*connData).SetReadDeadline(time.Now().Add(2 * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		file.Write(tmp[:n])
		if n < max_size {
			break
		}
	}
	return nil
}


var FTPErrorMessages = map[string][]string{
	// GO says that errors should not be capitalized 
	"200": {"comando se ha ejecutado con éxito.", ""},
	"421": {"", "servicio FTP no está disponible."},
}

func ParseFTPCode(code string) (string, error) {
	message, exists := FTPErrorMessages[code]
	if !exists {
		return "",errors.New("codigo de error desconocido")
	}
	msg := message[0]
	if( message[1] == ""){
		return msg, nil
	}
	return msg, errors.New(message[1])
}