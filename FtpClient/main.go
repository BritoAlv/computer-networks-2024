package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
)

func execute_command(command string, X *CommandsStruct) {
	command = strings.Split(command, "\n")[0]
	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		fmt.Println("no command written")
		return 
	}
	command_name := parts[0]
	method := reflect.ValueOf(X).MethodByName(strings.ToUpper(command_name))

	if !method.IsValid() {
		fmt.Println("command not implemented by the client")
		return
	}
	resultCommand := method.Call([]reflect.Value{reflect.ValueOf(strings.TrimSpace(command[len(command_name):]))})

	result, _ := resultCommand[0].Interface().(string)
	
	if !resultCommand[1].IsNil() {
		resultError, _ := resultCommand[1].Interface().(error)
		fmt.Println(resultError.Error())
	}
	fmt.Println("Command Says : \n" + result)
}

func main() {
	var X CommandsStruct
	ftptouse := Scene
	conn, err := net.Dial("tcp", strings.TrimSpace(ftptouse.ip)+":"+strings.TrimSpace(ftptouse.port))
	if err != nil {
		fmt.Println("Connection can't be established: ")
		fmt.Println("	" + err.Error())
		return
	}
	X.connectionConfig = &conn
	defer conn.Close()
	response, err := readOnMemoryDefault(&conn)
	if err != nil {
		fmt.Println("There was a problem getting the response")
		fmt.Println("	" + err.Error())
		return
	}
	fmt.Println(string(response))
	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		execute_command(command, &X)
	}
}