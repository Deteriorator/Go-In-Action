/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-14 11:23
 */

package chapter1

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //<co id="createwg" />

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i) //<co id="receivevalue" />
	}
	wg.Done()
}

func main() {
	c := make(chan int) //<co id="createchan" />
	go printer(c)       //<co id="goroutine" />
	wg.Add(1)           //<co id="addwg" />

	//Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i //<co id="sendonchan" />
	}

	close(c)  //<co id="closechan" />
	wg.Wait() //<co id="wgwait" />
}
