package loops

import "fmt"

func Workshop4() {

	// 220 ve 284 arkadaş sayıları mı
	// 10 ve 65 girilmiş ise arkadaş sayı değildir.
	a := 220
	b := 284
	var aDividers, bDividers int = 0, 0

	for i := 1; i < a; i++ {
		if a%i == 0 {
			aDividers = i + aDividers
		}
	}

	for i := 1; i < b; i++ {
		if b%i == 0 {
			bDividers = i + bDividers
		}
	}

	if a == bDividers && b == aDividers {
		fmt.Println("Arkadaş sayılardır")
	} else {
		fmt.Println("Arkadaş sayı değiller")
	}

}
