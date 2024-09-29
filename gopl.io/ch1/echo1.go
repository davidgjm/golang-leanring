// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//when i==1, sep is using its implicit initial value: empty string ""
	// hence, the result string does not start with a prefix ','
	//
	// When i>1, a reassigned sep ',' is inserted before the commnad line argument
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ","
	}
	fmt.Println(s)
}
