package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logch = make(chan logEntry, 50)
var donech = make(chan struct{})

// struct{} has zero memory,which is strong-typed
func main() {
	go logger()
	/*defer func() {
		close(logch)
	}()*/
	logch <- logEntry{time.Now(), logInfo, "App is starting"}

	logch <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	donech <- struct{}{}
	//first{} is no field,second{} is initializer
}

func logger() {
	/*or entry := range logch {
		fmt.Printf("%v -[%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}*/
	for {
		//select:until a channel receive message,it will go to listen
		select {
		case entry := <-logch:
			fmt.Printf("%v -[%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-donech:
			break
		}
	}
}
