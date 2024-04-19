package main

import (
	"FTPClient/api"
	"FTPClient/console"
	"sync"
)

func main() {
	var vg sync.WaitGroup
	vg.Add(1)
	go api.Run_web_gui()
	go console.StartConsole()
	vg.Wait()
}