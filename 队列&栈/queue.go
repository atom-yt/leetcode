package main

import (
	"fmt"
)

func checkValid(s string) bool {
	match := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	queue := []string{}
	rus := []rune(s)
	for _, rut := range rus {
		ru := string(rut)
		left, ok := match[ru]
		if ok {
			if len(queue) > 0 && queue[len(queue)-1] == left {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else {
	
			queue = append(queue, ru)
		}
	}
	return len(queue) == 0
}

func main() {
	fmt.Println(checkValid("{}"))
	fmt.Println(checkValid("{}}"))
}
