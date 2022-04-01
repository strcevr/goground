package chans

import "fmt"

func SelectCases () {
	// Prints 321
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			fmt.Print(1)
		case <-c:
			fmt.Print(2)
			c = nil
		case c <- 1:
			fmt.Print(3)
		}
	}
	// With small probability prints 332
	a := make(chan int, 2)
	for range [3]struct{}{} {
		select {
		default:
			fmt.Print(1)
		case <-a:
			fmt.Print(2)
			a = nil
		case a <- 1:
			fmt.Print(3)
		}
	}
}