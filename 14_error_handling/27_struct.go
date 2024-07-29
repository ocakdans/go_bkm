package errorhandling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d --- %s", b.parameter, b.message)
}

func TahminEt2(number int) (string, error) {
	if number < 1 || number > 100 {
		return "", &borderException{number, "out for border"}
	}
	return "Bildiniz", nil
}
