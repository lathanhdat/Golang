package main

import "fmt"

/*yields an additional untyped boolean result reporting whether the communication succeeded.
* The value of ok is true if the value received was delivered by a
*successful send operation to the channel, or false if it is a zero value generated
*because the channel is closed and empty.
 */
func main() {
	channel := make(chan int)
	go func() {
		channel <- 42
		close(channel)
	}()
	v, ok := <-channel
	if !ok {
		fmt.Println("!OK", v, ok)
	} else {
		fmt.Println("OK", v, ok)
	}
}
