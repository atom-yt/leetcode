package main

import (
	"fmt"
)

func main() {
	letter, number, done := make(chan bool), make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			<-number
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- true
		}
	}()

	go func() {
		i := 'A'
		for i <= 'Z' {
			<-letter
			fmt.Printf("%c%c", i, i+1)
			i += 2
			number <- true
		}
		done <- true
	}()

	number <- true
	for {
		select {
		case <- done:
			return
		}
	}
}
