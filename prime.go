package main

import "fmt"

var primes = make(chan int)

func sieveOfEratosthenes(primes chan<- int, N int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes <- i
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	close(primes)
	return
}

func main() {
	go sieveOfEratosthenes(primes, 10000)
	for {
		v, ok := <-primes
		if !ok {
			return
		}
		fmt.Println("this is prime:", v)
	}
}
