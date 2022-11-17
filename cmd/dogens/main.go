package main

import (
	"log"

	dogens "github.com/tjstebbing/dogens/pkg"
)

func main() {

	addr, err := dogens.DomainToAddr("dogecoin.com")
	if err != nil {
		log.Println(err)
		log.Fatal("No address associated with this domain")
	}
	log.Println(addr)

}
