//13. Roman to Integer

package main

import "fmt"

func romanToInt(s string) (res int) {
	mp := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		val, _ := mp[s[i]]
		nextVal := 0

		if i+1 < len(s) {
			nextVal = mp[s[i+1]]
		}

		if val < nextVal {
			res -= val
			continue
		}
		res += val
	}
	return
}

func main() {
	fmt.Println(romanToInt("IV"))
}
