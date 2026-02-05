package main

import (
	"fmt"
	limiter "ratelimiter/fixed_window_counter"
	"time"
)

func main() {

	client := time.NewTicker(4 * time.Second)
	client1 := time.NewTicker(5 * time.Second)

	done := make(chan bool)
	u := limiter.NewUser()
	go func() {
		for {
			select {
				case <-done:
					return
				case <-client.C:
					if u.Check(){
						fmt.Printf("request accepted %v\n",u.CurWindowSize)
					}else{
						fmt.Printf("request denied %v\n",u.CurWindowSize)
					}
				case <-client1.C:
					if u.Check(){
						fmt.Printf("request accepted %v\n",u.CurWindowSize)
					}else{
						fmt.Printf("request denied %v\n",u.CurWindowSize)
					}
			}
		}
	}()

	time.Sleep(1 * time.Minute)
	client.Stop()
	done <- true
	fmt.Println("client stopped")
}
