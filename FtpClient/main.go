package main

import (
	"FTPClient/api"
	"sync"
)

func main() {
	var vg sync.WaitGroup
	vg.Add(1)
	go api.Run_web_gui()
	vg.Wait()
}
