package main

import (
	"fmt"
	"os"
	limiter "ratelimiter/fixed_window_counter"
	"time"
)

func main() {
	arg := os.Args[1]
	switch arg {

	case "1":

		client := time.NewTicker(4 * time.Second)

		done := make(chan bool)
		u := limiter.NewUser()
		go func() {
			for {
				select {
				case <-done:
					return
				case <-client.C:
					if u.Check() {
						fmt.Printf("request accepted %v\n", u.CurCount)
					} else {
						fmt.Printf("request denied %v\n", u.CurCount)
					}
				}
			}
		}()

		time.Sleep(1 * time.Minute)
		client.Stop()
		done <- true
		fmt.Println("client stopped")

	case "2":
		client := time.NewTicker(20 * time.Second)
		done := make(chan bool)
		bursty:=make(chan int,5)
		go func(){
			time.Sleep(time.Minute-time.Second)
			for i:=0;i<5;i++{
				bursty<-i
			}
		}()
		u := limiter.NewUser()
		go func() {
			for {
				select {
				case <-done:
					return
				case <-bursty:
					if u.Check() {
						fmt.Printf("request accepted %v\n", u.CurCount)
					} else {
						fmt.Printf("request denied %v\n", u.CurCount)
					}
				case <-client.C:
					if u.Check() {
						fmt.Printf("request accepted %v\n", u.CurCount)
					} else {
						fmt.Printf("request denied %v\n", u.CurCount)
					}
				}
			}
		}()

		time.Sleep(1 * time.Minute)
		client.Stop()
		done <- true
		fmt.Println("client stopped")
	}
}
