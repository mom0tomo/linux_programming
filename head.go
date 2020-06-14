package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: head n [file file...]\n")
		return
	}
	nlines, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	if len(args) == 2 {
		f, err := os.OpenFile(args[1], os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		do_head(f, nlines)
	} else {
		for i := 1; i <len(args); i++ {
			f, err := os.OpenFile(args[i], os.O_RDONLY, 0644)
			defer f.Close()
			if err != nil {
				log.Fatal(err)
				return
			}
			if err := do_head(f, nlines); err != nil {
				log.Fatal(err)
			}
		}
	}
	return
}

func do_head(f *os.File, nlines int) error {
	if nlines <= 0 {
		return nil
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	fmt.Println(scanner.Text())
		nlines--
		if nlines == 0 {
			return nil
		}
	}
	return nil
}
