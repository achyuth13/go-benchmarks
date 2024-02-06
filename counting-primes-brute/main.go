package main

import (
	"fmt"
	"math"
	"time"
)

var MAX_INT = 100000000
var totalPrimeNumbers int32 = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	totalPrimeNumbers++
}

//From 3 to the Max int, we call the function checkPrime to see if it's a Prime or Not
//How are we doing so?
//If the number is 1, we ignore
//Else, from 3 to the Square root of the number, we check if the number is divisible by any number
//If it is, we return, and add the counter

func main() {
	start := time.Now()

	for i := 3; i < MAX_INT; i++ {
		checkPrime(i)
	}
	fmt.Println("Checking till ", MAX_INT, "found ", totalPrimeNumbers+1, "prime numbers. Took ", time.Since(start))
}
