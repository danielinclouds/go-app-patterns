package main

import (
	"errors"
	"time"
)

func RunWithTimeout() error {
	done := make(chan bool)
	go func() {
		// Do something
		// time.Sleep(5 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		return nil
	case <-time.After(4 * time.Second):
		return errors.New("Timeout")
	}
}

func main() {
	err := RunWithTimeout()
	if err != nil {
		panic(err)
	}
}
