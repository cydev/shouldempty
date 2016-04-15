package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if len(b) == 0 {
		os.Exit(0)
	}
	message := "Should be empty, but got:"
	if len(os.Args) > 1 {
		message = fmt.Sprintf("%s:", os.Args[1])
	}
	fmt.Fprintln(os.Stdout, message)
	os.Stdout.Write(b)
	os.Exit(-1)
}
