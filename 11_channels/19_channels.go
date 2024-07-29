package channels

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 15}
	c := make(chan int)
	go sum(s, c)
	x := <-c
	fmt.Println(x)
}

func GoRoutine(ciftSayiCn chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum = sum + i
	}

	ciftSayiCn <- sum // atama yapıyoruz...channel sum ı taşıyacak
}

func GoRoutine2(tekSayiCn chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum = sum + i
	}

	tekSayiCn <- sum // atama yapıyoruz...
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
