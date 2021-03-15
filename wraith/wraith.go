package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/0x1a8510f2/wraith/comms"
	"github.com/0x1a8510f2/wraith/config"
	"github.com/0x1a8510f2/wraith/hooks"
)

// Useful globals
var startTime time.Time

// Exit handling
var exitTrigger chan struct{}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(
		c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	go func() {
		<-c
		if config.Config.Process.RespectExitSignals {
			exitTrigger <- struct{}{}
		}
	}()
}

func init() {
	startTime = time.Now()
	exitTrigger = make(chan struct{})
	setupCloseHandler()
}

func main() {
	// Run OnStart hooks
	hooks.RunOnStart()

	// Run OnExit hooks always on exit
	defer hooks.RunOnExit()

	// Keep track of active goroutines
	var wg sync.WaitGroup

	txQueue := make(comms.TxQueue)
	rxQueue := make(comms.RxQueue)
	wg.Add(1)
	go comms.Manage(txQueue, rxQueue, &wg)

	// Mainloop: Transmit, receive and process stuff
	for {
		select {
		case data := <-rxQueue:
			fmt.Printf("%v", data) // Debug
		case <-exitTrigger:
			return
		}
	}

}
