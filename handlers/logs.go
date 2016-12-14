package handlers

import "time"

type Log struct {
	time    time.Time
	context string
	message string
	level   string
}

func (l Log) Write(level string, message string, context string) error {
	var err error

	return nil
}
