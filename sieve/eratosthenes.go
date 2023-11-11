package sieve

func Eratosthenes(n int) []bool{
    primes := make([]bool, n + 1, n +1)
    
    for i := 2; i <= n; i++ {
        primes[i] = true
    }

    for i := 2; i <= n; i++ {
        if primes[i] == true && i*i <= n {
            for j := i*i; j <= n; j = j + i {
                primes[j] = false
            }
        }
    }

    return primes
}
