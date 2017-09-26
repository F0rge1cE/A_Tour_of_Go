package main

import (
	"golang.org/x/tour/wc"	
	"strings"
)

func WordCount(s string) map[string]int {
	
	var count int
	var flag bool
	var res map[string]int
		
	res = make(map[string]int)
	str := strings.Fields(s)
	
	for _,v := range str {
		
		count, flag = res[v]
		
		if flag==true {
			res[v] = count + 1
		}else {
			res[v] = 1
		}
		
	}
	
	return res
}

func main() {
	wc.Test(WordCount)
}
