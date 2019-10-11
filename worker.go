package main

import "fmt"

const (
	workerNum int = 3
)

var (
	i   int
	val = make(chan int, 50)
)

func doWorkers() {
	for idx := 0; idx < workerNum; idx++ {
		go func() {
			for {
				fmt.Println("adding : ", i, " to chan")
				val <- i
				i++
				if i > 1000000000 {
					i = 0
				}
			}
		}()
	}
}

func getI() int {
	return <-val
}
