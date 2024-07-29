package errorhandling

import (
	"fmt"
	"os"
)

func ErrorHandling() {
	f, err := os.Open("demo1.txt")
	// if file is found, error is gonna be nil
	// errors.New("Bu bir hata") yazarsan da kendi hatanı dönmüş olursun. ve de bunu yazdırabilirsin.

	if err != nil {
		// Type Assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı with false of the ok")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
