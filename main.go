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
	conn, err := net.Dial("tcp", "127.0.0.1"+":"+"21")
	if err != nil {
		fmt.Println("Connection can't be established: ")
		fmt.Println("	" + err.Error())
		return
	}
	X.connection = &conn
	response, err := readOnMemory(&conn)
	if err != nil {
		fmt.Println("There was a problem getting the response")
		fmt.Println("	" + err.Error())
		return
	}
	fmt.Println(string(response))

	// Login step, this can be abstracted away, but i find weird
	//  to take a function to loop over an input 
	for {
		fmt.Println("Introduce a user name, or type ANONYMOUS.")
		reader := bufio.NewReader(os.Stdin)
		userName, _ := reader.ReadString('\n')
		userName = strings.TrimSpace(userName)
		if (userName == "ANONYMOUS"){
			X.ANONYMOUS(" "); 
			break; 
		}		
		fmt.Println("Introduce your password")
		reader = bufio.NewReader(os.Stdin)
		passWord, _ := reader.ReadString('\n')
		passWord = strings.TrimSpace(passWord)
		X.USER(userName)
		res, err := X.PASS(passWord)
		if starts_with(res,"Wrong") || err != nil {
			fmt.Println("Login error")
		}else{
			fmt.Println(res)
			break; 
		}
	}

	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.Split(command, "\n")[0]
		parts := strings.Split(command, " ")
		if len(parts) == 0 {
			fmt.Println("Que haces Calavera")
			continue
		}
		command_name := parts[0]
		if command_name == "exit" {
			break
		}

		method := reflect.ValueOf(&X).MethodByName(strings.ToUpper(command_name))

		if !method.IsValid() {
			fmt.Println("Calavera pon algo que entienda... como help")
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
