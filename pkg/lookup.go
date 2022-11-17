package dogens

import (
	"errors"
	"net"
	"strings"
)

func DomainToAddr(domain string) (string, error) {
	txts, err := net.LookupTXT(domain)
	if err != nil {
		return "", err
	}

	for _, record := range txts {
		if strings.HasPrefix(record, "dogecoin:") {
			_, addr, _ := strings.Cut(record, ":")
			return addr, nil
		}
	}

	return "", errors.New("No address found for this domain")
}
