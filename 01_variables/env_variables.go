package variables

import (
	"fmt"
	"os"
)

func EnvVariables() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	uName := os.Getenv("USERNAME")
	hPath := os.Getenv("TERM_PROGRAM_VERSION")

	fmt.Println(uName)
	fmt.Println(hPath)

}
