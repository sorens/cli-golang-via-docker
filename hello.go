package main

import "os"
import "fmt"

func main() {
	name := "world"
	if (len(os.Args) > 1) {
		name = os.Args[1]
	}
	fmt.Printf("hello, %s\n", name)
}