package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Catalog https://www.codewars.com/kata/59d9d8cb27ee005972000045/train/go
func Catalog(s string, article string) string {
	result := []string{}

	searcher := fmt.Sprintf("<prod><name>(.*?%s.*?)</name><prx>(\\d+(?:\\.\\d+)?)</prx><qty>(\\d+)</qty></prod>", article)
	re := regexp.MustCompile(searcher)

	for _, match := range re.FindAllStringSubmatch(s, -1) {
		result = append(result, fmt.Sprintf("%s > prx: $%s qty: %s", match[1], match[2], match[3]))
	}

	if len(result) == 0 {
		return "Nothing"
	}
	return strings.Join(result, "\n")
}

func main() {
}
