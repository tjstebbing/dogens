package dogens

import (
	"fmt"
	"net"
)

func DomainToAddr(domain string) (string, error) {
	txts, err := net.LookupTXT(domain)
	if err != nil {
		return "", err
	}
	fmt.Println(txts)
	return "", nil
}
