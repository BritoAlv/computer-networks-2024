package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
)

func main() {
	var X CommandsStruct

	ftptouse := Local

	conn, err := net.Dial("tcp", strings.TrimSpace(ftptouse.ip)+":"+strings.TrimSpace(ftptouse.port))
	if err != nil {
		fmt.Println("Connection can't be established: ")
		fmt.Println("	" + err.Error())
		return
	}
	X.connection = &conn
	response, err := readOnMemoryDefault(&conn)
	if err != nil {
		fmt.Println("There was a problem getting the response")
		fmt.Println("	" + err.Error())
		return
	}
	X.USER(ftptouse.user)
	X.PASS(ftptouse.password)

	fmt.Println(string(response))
	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.Split(command, "\n")[0]
		parts := strings.Split(command, " ")
		if len(parts) == 0 {
			fmt.Println("Wrongg")
			continue
		}
		command_name := parts[0]
		if command_name == "exit" {
			break
		}

		method := reflect.ValueOf(&X).MethodByName(strings.ToUpper(command_name))

		if !method.IsValid() {
			fmt.Println("help maybe ?")
			continue
		}
		resultCommand := method.Call([]reflect.Value{reflect.ValueOf(strings.TrimSpace(command[len(command_name):]))})

		resultString, _ := resultCommand[0].Interface().(string)
		
		if resultCommand[1].IsNil() {
			fmt.Println("Command Says : \n" + resultString)
		} else {
			resultError, _ := resultCommand[1].Interface().(error)
			fmt.Println("Error : " + resultError.Error())
		}
	}
	defer conn.Close()
}