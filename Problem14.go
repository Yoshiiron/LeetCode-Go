package main

import (
	"strings"
)

/*func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, str := range strs[1:] {
		if len(prefix) > 0 && (len(str) < len(prefix) || str[:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}
	return prefix
}
*/

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, str := range strs[1:] {
		if !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}
	return prefix
}
