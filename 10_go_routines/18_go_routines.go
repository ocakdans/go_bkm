package goroutines

import (
	"fmt"
	"time"
)

//parallel programming değpildir!
// parallel için runtime.GOMAXPROCS(1) de yazman lazım.

func GoRoutine() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("çift sayi:", i)
		time.Sleep(1 * time.Second)
	}
}

func GoRoutine2() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("tek sayi:", i)
		time.Sleep(1 * time.Second)

	}
}
