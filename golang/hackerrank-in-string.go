// https://www.hackerrank.com/challenges/hackerrank-in-a-string/
package main

import "fmt"

func hackerrankInString(s string) string {
	// Write your code here
	hackerrank := []string{"h", "a", "c", "k", "e", "r", "r", "a", "n", "k"}
	i := 0
	for _, v := range s {
		if string(v) == hackerrank[i] {
			i += 1
		}
		if i == len(hackerrank) {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(hackerrankInString("hackerrank"))
	fmt.Println(hackerrankInString("hackeronek"))
	fmt.Println(hackerrankInString("hereiamstackerrank"))
}
