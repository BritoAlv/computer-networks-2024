package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func command_PASV(connConfig *net.Conn) *net.Conn {
	connData, err := open_conection(wr(connConfig, []byte("PASV \r\n")))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &connData
}

func command_LIST(connConfig *net.Conn) {
	connData := *command_PASV(connConfig)
	write(connConfig, []byte("LIST \r\n"))
	data, _ := read(&connData)
	fmt.Println(string(data))
	defer connData.Close()
}

func command_GET(connConfig *net.Conn, s string, useBinary bool) {
	file, _ := os.Create("R" + s)
	if useBinary {
		fmt.Println(wr(connConfig, []byte("TYPE I\r\n")))
	}
	buffer := make([]byte, max_size)
	
	connData := *command_PASV(connConfig)
	write(connConfig, []byte("RETR " + s + "\r\n"))

	for {
        bytesRead, err := connData.Read(buffer) // Read the file in chunks

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
		file.Write(buffer[:bytesRead])
    }
	// this line made the code work !! .
	connData.Close()
	
	result, _ := read(connConfig)
	fmt.Println(string(result))
	
	fmt.Println(wr(connConfig, []byte("TYPE A\r\n")))
}

func command_STORE(connConfig *net.Conn, filename string, useBinary bool) {
	file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
	defer file.Close()
	buffer := make([]byte, max_size)
	if useBinary {
		fmt.Println(wr(connConfig, []byte("TYPE I\r\n")))
	}
    conn_data := command_PASV(connConfig)
	fmt.Println(wr(connConfig, []byte("STOR " + filename + "\r\n")))

	for {
        bytesRead, err := file.Read(buffer) // Read the file in chunks

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
		(*conn_data).Write(buffer[:bytesRead])
    }
	
	// this line made the code work !! .
	(*conn_data).Close()

	result, _ := read(connConfig)
	fmt.Println(string(result))
	
	fmt.Println(wr(connConfig, []byte("TYPE A\r\n")))
	
}

func main() {
	connConfig, err := net.Dial("tcp", "localhost:21")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connConfig.Close()
	fmt.Println(wr(&connConfig, []byte("Habla\r\n")))
	fmt.Println(wr(&connConfig, []byte("USER brito\r\n")))
	fmt.Println(wr(&connConfig, []byte("PASS password\r\n")))
	fmt.Println(wr(&connConfig, []byte("CWD /upload\r\n")))
	command_STORE(&connConfig, "a.mp4", true)
	command_GET(&connConfig, "a.mp4", true)
	command_STORE(&connConfig, "client.go", false)
	command_LIST(&connConfig)
}