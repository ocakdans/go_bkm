package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	RandomnessWithoutSeed()
}

func RandomnessWithSeed() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func RandomnessWithoutSeed() {
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}
