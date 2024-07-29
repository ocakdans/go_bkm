package consoleread

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConsoleRead() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your number is: ", num)
	}
}
func ReadFromConsole() {
	var score float64
	_, err := fmt.Print("Enter your score: ")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Scanf("%f", &score)
	switch {
	case score <= 10:
		fmt.Println("f")
	case score <= 60:
		fmt.Println("d")
	case score <= 70:
		fmt.Println("c")
	case score <= 80:
		fmt.Println("b")
	case score <= 90:
		fmt.Println("a")
	case score <= 100:
		fmt.Println("perfect")
	}
}
