package main

func (cs *CommandsStruct) LS(command  string) (string, error) {
	// first try yo establish a PASSIVE Connection Data.
	connData, err := cs.PASV()
	if err != nil{
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, []byte("LIST \r\n"))
	if err != nil{
		return "", err
	}
	data, err := readOnMemory(connData)
	if err != nil{
		return "", err
	}
	defer (*connData).Close()
	return string(data), nil
}