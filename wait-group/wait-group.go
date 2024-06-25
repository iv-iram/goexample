package wait_group

import (
    "fmt"
    "sync"
    "time"
)

//This WaitGroup is used to wait for all the goroutines launched here to finish. 
//Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func Wait_group() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()

}