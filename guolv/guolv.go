package guolv

import (
	"strings"
)

func ContainsAny(b []byte, list []string) bool {
	for _, s := range list {
		if strings.Contains(string(b), s) {
			return true
		}
	}
	return false
}
