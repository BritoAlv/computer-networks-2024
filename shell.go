package main

import (
	"fmt"
	"net"
	"reflect"
	"strings"
)


func MainShell(){
	var X CommandsStruct 
	fmt.Println("Calavera FTP Client")
	for {
		fmt.Print("IP of the FTP server you wish connect to: ")
		
		ip := read_input()

		if !is_valid_ip(ip){
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
		break
	}
	fmt.Println("SERVER SAYS: ")
	response, _ := read(X.connection)
	fmt.Println(string(response))
	fmt.Println("As you can read I need User and Password")
	fmt.Print("User: ")
	user := read_input()
	fmt.Print("Password: ")
	password := read_input()
	write(X.connection, []byte("USER " + user + "\r\n"))
	_, err :=  read(X.connection)
	if err != nil {
		fmt.Println("Something weird happent so I'm going to close everything")
		return		
	}
	write(X.connection, []byte("PASS " + password + "\r\n"))
	_, err =  read(X.connection)
	if err != nil {
		fmt.Println("Something weird happent so I'm going to close everything")
		return		
	}
	
	fmt.Println("You can start")
	for {
		fmt.Print(">> ")
		command := read_input()
		// use reflection to execute a command.
		parts := strings.Split(command, " ")
		if len(parts) == 0{
				fmt.Println("Que haces Calavera")
				continue
		}
		command_name := parts[0]

		if(command_name == "exit"){
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