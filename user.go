package main


func (cs *CommandsStruct) USER(args string) error  {
	wr(cs.connection, []byte("USER " + args + "\r\n"))
	return nil 
}