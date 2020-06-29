package main

import (
	"reflect"
	"testing"
)

func Test_sieveOfEratosthenes(t *testing.T) {
	primesTest := make(chan int)
	const max = 12
	expected := []int{2, 3, 5, 7, 11}
	var received []int

	go sieveOfEratosthenes(primesTest, max)

	for {
		v, ok := <-primesTest
		if ok {
			break
		}
		received = append(received, v)
	}
	if reflect.DeepEqual(received, expected) {
		t.Errorf("Expected %d got %d", expected, received)
	}
	return
}

func Benchmark_sieveOfEratosthenes(b *testing.B) {
	const max = 10000

	for i := 0; i < b.N; i++ {
		primesTest := make(chan int)
		go sieveOfEratosthenes(primesTest, max)
	}
}
