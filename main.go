package main

import (
	"github.com/yerkortiz/sieve/sieve"
)

func main() {
	//primes := sieve.Eratosthenes(2<<10)
	//for i:=0; i <= 2<<10; i++ {
	//    if primes[i] {
	//        fmt.Printf("%d\n", i)
	//    }
	//}
    
    sieve.Sieve(100)
}
