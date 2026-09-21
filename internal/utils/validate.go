package utils

import "strings"

const maxIDLength = 64

// ValidID reports whether s is a usable identifier: non-empty once trimmed, within a sane
// length, and free of control characters. It deliberately doesn't constrain character set
// beyond that, since imported Parent IDs, ID Numbers, and Employee Numbers may use formats
// this service doesn't own.
func ValidID(s string) bool {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" || len(trimmed) > maxIDLength {
		return false
	}
	for _, r := range trimmed {
		if r < 0x20 || r == 0x7f {
			return false
		}
	}
	return true
}
