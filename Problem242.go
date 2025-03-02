package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hms, hmt := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hms[s[i]] += 1
		hmt[t[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if hms[s[i]] != hmt[s[i]] {
			return false
		}
	}

	return true
}
