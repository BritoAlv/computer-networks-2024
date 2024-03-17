package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func command_GET(connConfig *net.Conn, command string){
	// split command in args.
	args := strings.Split(command, " ")
	if len(args) < 2 || args[1] != "A" || args[1] != "B" {
		fmt.Println("Provide Arguments: get filename binary/ascii")
		fmt.Println("get file.go A")
		fmt.Println("get file.mp4 B")
	}
	if(args[1] == "A"){
		command_get(connConfig, args[0], false)
	}
	if(args[1] == "B"){
		command_get(connConfig, args[0], true)
	}
}

func command_get(connConfig *net.Conn, s string, useBinary bool) {
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