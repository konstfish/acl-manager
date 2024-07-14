package manager

import "strings"

func formatToCSV(input []string) string {
	return strings.Join(input, ",")
}
