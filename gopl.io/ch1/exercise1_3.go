package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func concat(args []string) {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = ","
	}
	fmt.Println(s)
	fmt.Printf("%ds elapsed\n", time.Since(start).Microseconds())
}

func join(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(args[1:], ","))
	fmt.Printf("%ds elapsed\n", time.Since(start).Microseconds())
}

func main() {
	// fmt.Println(strings.Join(os.Args[1:], ","))
	// fmt.Println(os.Args[1:])
	concat(os.Args[1:])
	join(os.Args[1:])
}
