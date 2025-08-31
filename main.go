package main

import "fmt"

type OldLog struct{}

func (l *OldLog) Print(message string) {
	fmt.Println("OldLog:", message)
}

type NewLog interface {
	Log(message string)
}

type LogAdapter struct {
	oldLog *OldLog
}

func (l *LogAdapter) Log(message string) {
	l.oldLog.Print(message)
}

func main() {

	old := &OldLog{}

	old.Print("Old Log")

	var logger NewLog = &LogAdapter{oldLog: old}

	logger.Log("New Log")
}
