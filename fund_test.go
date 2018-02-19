package main

import (
	"fmt"
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFunc(b *testing.B) {
	if b.N < WORKERS {
		return
	}
	fmt.Println("Total Iterations : ", b.N)
	server := NewFundServer(b.N)

	dollarsPerFounder := b.N / WORKERS
	//fmt.Println("dollarsPerFounder : ", dollarsPerFounder)

	var wg sync.WaitGroup

	pizzaTime := false

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < dollarsPerFounder; j++ {
				server.Transact(func(fund *Fund) {
					if fund.Balance() <= 10 {
						pizzaTime = true
						return
					}
					fund.Withdraw(1)
				})

				if pizzaTime {
					break
				}
			}
		}()
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 10 {
		b.Error("Balance wasn't zero : ", balance)
	}
}
