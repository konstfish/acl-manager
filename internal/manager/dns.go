package manager

import (
	"log"
	"net"
)

func getDNSList(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Printf("Error looking up IPs for domain %s: %v\n", domain, err)
		return nil, err
	}

	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	return ipStrings, nil
}
