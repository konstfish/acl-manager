package manager

import (
	"net"
	"strings"
)

func formatToCSV(input []string) string {
	return strings.Join(input, ",")
}

func validateCSV(input string) bool {
	split := strings.Split(input, ",")
	ip := net.ParseIP(strings.Split(split[0], "/")[0])
	if ip == nil {
		return false
	}
	return true
}
