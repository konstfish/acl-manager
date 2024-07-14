package manager

import "strings"

func praseFromNetList(input string) []string {
	var results []string

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 && !strings.HasPrefix(trimmedLine, "#") {
			results = append(results, trimmedLine)
		}
	}

	return results
}

func parseFromCSV(input string) []string {
	lines := strings.Split(input, ",")

	return lines
}
