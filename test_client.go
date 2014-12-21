package main

import (
	"fmt"
	"github.com/arukim/overmind/client"
	"strconv"
	"time"
)

func main() {
	var input string

	var res interface{}
	for {
		start := time.Now()
		for i := 0; i < 10000; i++ {
			key := strconv.Itoa(i)

			err := client.Get(key, &res, time.Millisecond*100, time.Minute*15,
				func() {
					// here we can print or check data
				})
			if err != nil {
				fmt.Printf("Catched error %v while processing request\n", err)
			}
		}
		elapsed := time.Since(start)
		fmt.Printf("Elapsed %s\n", elapsed)
	}
	fmt.Scanln(&input)
}
