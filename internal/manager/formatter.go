package manager

import "strings"

func formatNetList(input []string) string {
	return strings.Join(input, ",")
}
