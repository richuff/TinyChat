package test

import "fmt"

var ch chan int

func main() {
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		default:
			fmt.Println("last")
			return
		}
	}
}
