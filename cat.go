package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "%s: file name not given\n", os.Args[0])
		return
	}
	for i := 0; i < len(args); i++ {
		if err := do_cat(args[i]); err != nil {
			log.Fatal(err)
		}
	}
}

func do_cat(path string) (output error) {
	fd, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	buf := make([]byte, 64)
	for {
		bytesread, err := fd.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		fmt.Print(string(buf[:bytesread]))
	}
	return nil
}
