package main

import (
	"fmt"
	"flag"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for i := 0; i < len(args); i++ {
		fmt.Printf("arg %d=%s\n", i, args[i])
	}
}

