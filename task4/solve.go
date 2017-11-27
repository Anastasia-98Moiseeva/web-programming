package main

import "unicode"
import "strings"

func RemoveEven(x []int) []int {
    res := make([]int, 0)
    for _, val := range x {
        if val % 2 == 1 {
            res = append(res, val)
        }
    }
    return res
}

func PowerGenerator(x int) func() int {
    y := 1
    return func() int {
        y = y*x
        return y
    }
}

func DifferentWordsCount(st string) int {
    s := make(map[string]bool)
    w := ""
    res := 0
    str = st + " "
    for _, l := range (str) {
        if unicode.IsLetter(l) {
            w = w + string(unicode.ToLower(l))
        } else if w != "" {
            if !s[w] {
                res = res + 1
            }
            s[w] = true
            w = ""
        }
    }
    return res
}