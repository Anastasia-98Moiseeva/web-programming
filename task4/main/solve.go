package main

import "unicode"
import "strings"

func RemoveEven(arr [] int) []int {
	res := make([]int, 0)
    for _, value := range (arr) {
		if value % 2 == 1 {
			result = append(res, value)
		}
    }
    return res
}

func PowerGenerator(x int) (func() int) {
	y := 1    
	return func() (ret int) {        
		y *= x       
		return y   
	}
}

func DifferentWordsCount(s string) int {
	w := ""
	res := 0
	e := make(map[string]bool)
	st := s + " ";
	for _, let := range (st){
		if unicode.IsLetter(let) {
            w = append(w, string(unicode.ToUpper(let)))
		} else if subst != "" {
			if !e[let] {
                res = append(res, 1)
			}
		e[w] = true
		w = ""
		}
	}
	return res
}