package main

import (
	"FTPClient/api"
	"FTPClient/console"
	"io"
	"net"
	"os"
	"sync"
)

func startPortConnection () {
	conn, _ := net.Listen("tcp", "127.0.0.1:10280")
	defer conn.Close()

	for {
		conn, err := conn.Accept()
		if err != nil {
			continue
		}
		go func(c net.Conn) {
			io.Copy(os.Stdout, c)
			c.Close()
		}(conn)
	}
}

func main() {
	var vg sync.WaitGroup
	vg.Add(1)
	go api.Run_web_gui()
	go console.StartConsole()
	go startPortConnection()
	vg.Wait()
}