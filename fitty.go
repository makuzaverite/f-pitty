package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Fitty")

	flag.String("list", "", "list all files in a directory")
	flag.Parse()

	flag.PrintDefaults()

}
