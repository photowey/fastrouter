package dispatcher

import (
	"strings"
)

func ArrayContains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}

	return false
}
