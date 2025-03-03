package main

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)

	for _, str := range strs {
		strArr := strings.Split(str, "")
		slices.Sort(strArr)
		key := strings.Join(strArr, "")
		m[key] = append(m[key], str)
	}

	result := make([][]string, 0)
	for _, sl := range m {
		result = append(result, sl)
	}

	return result

}
