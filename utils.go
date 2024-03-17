package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func starts_with(command, prefix string) bool{
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

func read_input() string {
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	return input[:len(input)-1]
}

const max_size = 1024

func write(connConfig *net.Conn, data []byte) error {
	_, err := (*connConfig).Write(data)
	time.Sleep(1 * time.Second)
	return err
}

func read(connData *net.Conn) ([]byte, error) {
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

func wr(connConfig *net.Conn, data []byte) string {
	write(connConfig, data)
	time.Sleep(2*time.Second)
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