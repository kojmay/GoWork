package main

import(
	"sync"
	"time"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	var gs [5]struct {
		id int "id"
		result int "result"
	}

	for i:=0; i<len(gs); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id +1) * 100

			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}

	println("main ...")
	wg.Wait()
	fmt.Println("%v\n", gs)
	println("main exit.")

 }