package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func write(connConfig *net.Conn , data string) error {
	_, err := (*connConfig).Write([]byte(data + "\r\n"))
	time.Sleep( 1* time.Second)
    return err
}

func read(connData *net.Conn) ([]byte, error) {
	tmp := make([]byte, 1024)
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
		if n < 1024{
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

func main(){
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
}