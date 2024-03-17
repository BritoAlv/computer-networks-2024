package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func command_STORE(connConfig *net.Conn, command string){
	// split command in args.
	args := strings.Split(command, " ")
	if len(args) < 2 || args[1] != "A" || args[1] != "B" {
		fmt.Println("Provide Arguments: get filename binary/ascii")
		fmt.Println("put file.go A")
		fmt.Println("put file.mp4 B")
	}
	if(args[1] == "A"){
		command_store(connConfig, args[0], false)
	}
	if(args[1] == "B"){
		command_store(connConfig, args[0], true)
	}
}

func command_store(connConfig *net.Conn, filename string, useBinary bool) {
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