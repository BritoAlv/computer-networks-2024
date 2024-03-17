package main

import (
	"fmt"
	"net"
	"reflect"
	"strings"
)

func MainShell() {
	var X CommandsStruct
	fmt.Println("Calavera FTP Client")
	
	// SETUP CONNECTION.
	for {
		fmt.Print("IP of the FTP server you wish connect to: ")

		ip := read_input()

		if !is_valid_ip(ip) {
			fmt.Println("I don't understand that IP")
			continue
		}

		conn, err := net.Dial("tcp", ip+":"+"21")
		if err != nil {
			fmt.Println("Connection can't be established")
			fmt.Println("	" + err.Error())
			continue
		}

		X.connection = &conn
		fmt.Println("Established connection, Greetings should come from the server: ")

		fmt.Print("SERVER SAYS: ")
		response, err := read(X.connection)
		if err != nil{
			fmt.Println("Something weird happent with the connection so I'm going to restart")
			continue
		}
		fmt.Println(string(response))
		break
	}

	// LOGIN.
	for {
		fmt.Println("As you can tell I need User and Password")
		fmt.Print("User: ")
		user := read_input()
		err := X.USER(user)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print("Password: ")
		password := read_input()
		err = X.PASS(password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	// AFTER LOGGED IN.
	for {
		fmt.Print(">> ")
		command := read_input()
		// use reflection to execute a command.
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
		method.Call([]reflect.Value{reflect.ValueOf(strings.TrimSpace(command[len(command_name):]))})
	}
	defer (*(X.connection)).Close()
}