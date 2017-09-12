package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to")
var spanish bool
var help bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
	flag.BoolVar(&help, "h", false, "Print this help")
	flag.BoolVar(&help, "help", false, "Print this help")
}

func main() {

	//go run .\02foundation\flag_cli.go -s -name=Mundo
	//go run .\02foundation\flag_cli.go -h

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if spanish == true {
		fmt.Printf("Hola %s\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}

}
