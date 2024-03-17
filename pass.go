package main

import (
	"fmt"
)

func (cs *CommandsStruct) PASS(args string) error  {
	response := wr(cs.connection, []byte("PASS " + args + "\r\n"))
	if starts_with(response, "230") {
		fmt.Println("Todo en talla")
		return nil
	} else {
		return fmt.Errorf("error: %s", response)		
	}
}