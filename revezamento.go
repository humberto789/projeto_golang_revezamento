package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func runner(order chan int, numberRunner int, wg *sync.WaitGroup) {
	for {
		value := <-order
		if value == numberRunner {
			fmt.Println("Runner ", numberRunner, " starting to running...")

			randomInt := (rand.Intn(9) + 1)
			time.Sleep(time.Duration(randomInt) * time.Second)

			fmt.Println("Runner ", numberRunner, " finished an running")

			order <- (value + 1)

			break
		}

		order <- value
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	runnersQuantity := 4

	wg.Add(runnersQuantity)

	order := make(chan int, 1)
	order <- 1

	for i := 1; i <= runnersQuantity; i++ {
		go runner(order, i, &wg)
	}

	wg.Wait()

	fmt.Println("Running is finished")

}
