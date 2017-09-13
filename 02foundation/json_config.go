package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage json_config json_filename")
	}

	filename := os.Args[1]

	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Error decoding configuration: ", err)
	}
	fmt.Printf("Configuration path %s is enabled: %v\n", conf.Path, conf.Enabled)

}
