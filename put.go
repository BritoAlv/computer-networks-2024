package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func (cs *CommandsStruct) PUT (command string){
	// split command in args.
	args := strings.Split(command, " ")
	if len(args) < 2 || args[1] != "A" || args[1] != "B" {
		fmt.Println("Provide Arguments: get filename binary/ascii")
		fmt.Println("put file.go A")
		fmt.Println("put file.mp4 B")
	}
	if(args[1] == "A"){
		command_store(cs, args[0], false)
	}
	if(args[1] == "B"){
		command_store(cs, args[0], true)
	}
}

func command_store(cs *CommandsStruct, filename string, useBinary bool) {
	file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
	defer file.Close()
	buffer := make([]byte, max_size)
	if useBinary {
		fmt.Println(wr(cs.connection, []byte("TYPE I\r\n")))
	}
    conn_data := *cs.PASV("")
	fmt.Println(wr(cs.connection, []byte("STOR " + filename + "\r\n")))

	for {
        bytesRead, err := file.Read(buffer) // Read the file in chunks

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
		(conn_data).Write(buffer[:bytesRead])
    }
	
	// this line made the code work !! .
	(conn_data).Close()

	result, _ := read(cs.connection)
	fmt.Println(string(result))
	fmt.Println(wr(cs.connection, []byte("TYPE A\r\n")))
}