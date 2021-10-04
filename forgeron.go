package main

import "fmt"

func ContainsItem(tab []string, s string) int {
	var res int = 0
	for _, a := range tab {
		if a == s {
			res ++
		}
	}
	return res
}
