package main

import(
	"fmt"
)

func counter(out chan<- int) {
	for i:=0; i<100; i++{
		out <- i
	}
	close(out)
}

func squaer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v*v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in{
		fmt.Println(v)
	}
	// close(in)
}


func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squaer(squares, naturals)
	printer(squares)

}