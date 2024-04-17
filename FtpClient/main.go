package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"sync"
)

var X FtpSession
var wg sync.WaitGroup

func execute_command(command string, X *FtpSession) {
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

	fmt.Println("Command Says : " + result)
	if !resultCommand[1].IsNil() {
		resultError, _ := resultCommand[1].Interface().(error)
		fmt.Println("Also : " + resultError.Error())
	}
}

func readPart(commandCh chan string) {
	defer wg.Done()
	for {
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		commandCh <- command
		fmt.Println("Command sent")
	}
}

func writePart(commandCh chan string) {
	defer wg.Done()
	for {
		command := <-commandCh
		go execute_command(command, &X)
	}
}

func main() {
	ftptouse := Local
	conn, err := net.Dial("tcp", strings.TrimSpace(ftptouse.ip)+":"+strings.TrimSpace(ftptouse.port))
	if err != nil {
		fmt.Println("Connection can't be established: ")
		fmt.Println("	" + err.Error())
		return
	}
	X.connectionConfig = &conn
	X.connectionData = nil
	X.connDataUsed = false
	X.queueResponses = *NewQueue()
	defer conn.Close()
	response, err := readOnMemoryDefault(&X)
	if err != nil {
		fmt.Println("There was a problem getting the response")
		fmt.Println("	" + err.Error())
		return
	}
	fmt.Println(string(response))
	commandCh := make(chan string)
	
	wg.Add(2)
	go readPart(commandCh)
	go writePart(commandCh)
	// Wait for the goroutines to finish
	wg.Wait()
}