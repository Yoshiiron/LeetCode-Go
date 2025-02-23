package main

import (
	"strconv"
)

func Reverse(x string) (res string) {
	for _, run := range x {
		res = string(run) + res
	}
	return
}

func isPalindrome(x int) bool {
	reversPol := Reverse(strconv.Itoa(x))
	strX := strconv.Itoa(x)
	return reversPol == strX
}

/*
Best solution

func isPalindrome(x int) bool {

    str := strconv.Itoa(x)

    for i := 0; i < len(str) / 2; i++ {
        if(str[i] != str[len(str) - i - 1]) {
            return false
        }
    }

    return true
}
*/
