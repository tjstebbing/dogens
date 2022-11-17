package main

import (
	"fmt"
	"log"
	"os"

	dogens "github.com/tjstebbing/dogens/pkg"
)

func main() {

	if len(os.Args) > 1 {
		addr, err := dogens.DomainToAddr(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Dogecoin addr for %s: %s\n", os.Args[1], addr)
	} else {
		log.Fatal("Please provide a domain")
	}

}
