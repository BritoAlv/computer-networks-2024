package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

func write(connConfig *net.Conn, data string) error {
	_, err := (*connConfig).Write([]byte(data + "\r\n"))
	time.Sleep(1 * time.Second)
	return err
}

func read(connData *net.Conn) ([]byte, error) {
	max_size := 1024
	tmp := make([]byte, max_size)
	data := make([]byte, 0)
	length := 0
	for {
		n, err := (*connData).Read(tmp)
		if err != nil {
			if err != io.EOF {
				r := make([]byte, 0)
				return r, err
			}
			break
		}
		data = append(data, tmp[:n]...)
		length += n
		if n < max_size {
			break
		}
	}
	return data, nil
}

func wr(connConfig *net.Conn, data string) string {
	write(connConfig, data)
	ans, _ := read(connConfig)
	return string(ans)
}

func open_conection(connDataConfig string) (net.Conn, error) {
	connDataIP, connDataPort := parse_get_connection_ftp(connDataConfig)
	connData, err := net.Dial("tcp", connDataIP+":"+connDataPort)
	if err != nil {
		fmt.Println(err)
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

func command_PASV(connConfig *net.Conn) *net.Conn {
	connData, err := open_conection(wr(connConfig, "PASV"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &connData
}

func command_LIST(connConfig *net.Conn) {
	connData := *command_PASV(connConfig)
	write(connConfig, "LIST")
	data, _ := read(&connData)
	fmt.Println(string(data))
	defer connData.Close()
}

func command_GET(connConfig *net.Conn, s string) {
	connData := *command_PASV(connConfig)
	write(connConfig, "RETR " + s + "\r\n")
	data, err := read(&connData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	connConfig, err := net.Dial("tcp", "localhost:21")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connConfig.Close()
	fmt.Println(wr(&connConfig, "Habla"))
	fmt.Println(wr(&connConfig, "USER brito"))
	fmt.Println(wr(&connConfig, "PASS password"))
	fmt.Println(wr(&connConfig, "CWD /upload"))
	command_LIST(&connConfig)
	command_GET(&connConfig, "demo.txt")
}