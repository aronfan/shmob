package main

import (
	"flag"
	"fmt"
)

var (
	e = flag.String("e", "", "pipe mode, such as \"op=dump&key=1\"")
	h = flag.Bool("h", false, "display help")
)

func main() {
	flag.Parse()
	if *h {
		fmt.Println("share memory observer")
		return
	}
	if *e != "" {
		err := pipe(*e)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	} else {
		// display the console ui
		fmt.Println("display the console ui")
	}
}
