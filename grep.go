package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	args := flag.Args()
	pat := args[0]		
	f, err := os.OpenFile(args[1], os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	do_grep(pat, f)
}
func do_grep(pat string, f *os.File) {
	scanner := bufio.NewScanner(f)
	str := ""
	for scanner.Scan() {
		s := scanner.Text()
		r := regexp.MustCompile(pat)
		if r.MatchString(s) == true {
			str += s
		}
	}
	fmt.Println(str)
}
