package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan int)
	chB := make(chan int)

	go func() {
		for {
			select {
			case va:=<-chA:
				fmt.Println("channel A got new value",va)
			case vb:=<-chB:
				fmt.Println("channel B got new value",vb)

			default:
				time.Sleep(time.Second*2)
				fmt.Println("nothing happened")
			}
		}
	}()

	go func() {
		chA<-1
		chA<-2
	}()




		chB<-3
		chB<-4

}
