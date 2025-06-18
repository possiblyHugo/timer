package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second)

	fmt.Printf("%d\n", timer1)
	fmt.Printf("%d\n", timer1)

}
