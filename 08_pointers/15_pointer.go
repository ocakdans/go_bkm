package pointers

import "fmt"

func Pointer(number *int) { // * ile bellekteki adresi ile çalışıcam demek.
	*number = *number + 1
	fmt.Println("number in the pointer: ", *number)
}

func Pointer2(numbers []int) {
	numbers[0] = 100
	fmt.Println("numbers[0] in the Pointer2", numbers[0])
}

func SayHello(message *string) {
	println(*message)
	*message = "Hi Go!"

}
