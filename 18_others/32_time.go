package others

import (
	"fmt"
	"time"
)

func Times() {
	t := time.Date(2024, time.May, 01, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Time is %s\n", t)

	now := time.Now()
	fmt.Printf("Now, time is %s", now)
}
