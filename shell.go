package main

import (
	"fmt"
	"net"
)


func MainShell(){
	fmt.Println("Calavera FTP Client")
	var connConfig net.Conn
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
		connConfig = conn
		fmt.Println("Established connection, Greetings should come from the server: ")
		break
	}
	fmt.Println("SERVER SAYS: ")
	response, _ := read(&connConfig)
	fmt.Println(string(response))
	fmt.Println("As you can read I need User and Password")
	fmt.Print("User: ")
	user := read_input()
	fmt.Print("Password: ")
	password := read_input()
	write(&connConfig, []byte("USER " + user + "\r\n"))
	_, err :=  read(&connConfig)
	if err != nil {
		fmt.Println("Something weird happent so I'm going to close everything")
		return		
	}
	write(&connConfig, []byte("PASS " + password + "\r\n"))
	_, err =  read(&connConfig)
	if err != nil {
		fmt.Println("Something weird happent so I'm going to close everything")
		return		
	}
	
	fmt.Println("You can start")
	for {
		fmt.Print(">> ")
		command := read_input()
		if(starts_with(command, "exit")){
			fmt.Println("Goodbye! Calavera")
			break
		} else if (starts_with(command, "cd")){
			command_CD(&connConfig, command)			
		} else if (starts_with(command, "ls")){
			command_LIST(&connConfig, command)
		} else if (starts_with(command, "get")){
			command_GET(&connConfig, command)
		} else if(starts_with(command, "put")){
			 command_STORE(&connConfig, command)
		} else if(starts_with(command, "help")){
			fmt.Println("cd", "ls", "get", "put", "help")
		} else{
			fmt.Println("Calavera pon algo que entienda... como help")
		}
	}
	defer connConfig.Close()
}