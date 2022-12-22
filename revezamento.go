package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func runner(order chan int, winner chan bool, group int, numberRunner int, wg *sync.WaitGroup) {
	for {
		value := <-order
		win := <-winner
		if value == numberRunner {
			fmt.Println("Runner", numberRunner, "| group", group, "took the baton...")

			randomInt := (rand.Intn(9) + 1)
			time.Sleep(time.Duration(randomInt) * time.Second)

			fmt.Println("Runner", numberRunner, "| group", group,"handed over the baton")

			order <- (value + 1)
			if win == false && value == 4 {
				winner <- true
				fmt.Println("Team",group, "are the winner of the race")
			}else {
				winner <- win
			}
			break
		}
		winner <- win
		order <- value
	}
	defer wg.Done()
}

func main() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	runnersQuantity := 4

	wg1.Add(runnersQuantity)
	wg2.Add(runnersQuantity)

	orderG1 := make(chan int, 1)
	orderG1 <- 1

	orderG2 := make(chan int, 1)
	orderG2 <- 1

	winner := make(chan bool, 1)
	winner <- false 

	fmt.Println("Race started:")

	for i := 1; i <= runnersQuantity; i++ {
		go runner(orderG1, winner, 1, i, &wg1)
		go runner(orderG2, winner, 2, i, &wg2)
	}


	wg1.Wait()
	wg2.Wait()

	fmt.Println("Race finished.")

}
