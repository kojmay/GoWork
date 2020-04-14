package main

import(
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {		// 任务b
			println("b")
		}()

		for i := 0; i<4; i++{ 	// 任务a
			println("a:", i)

			if i == 1 {
				runtime.Goshed()
			}
		}
	}()

	<-exit
}