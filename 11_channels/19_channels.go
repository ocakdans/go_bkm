package channels

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
