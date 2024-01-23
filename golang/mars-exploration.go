// https://www.hackerrank.com/challenges/mars-exploration/

package main

func marsExploration(s string) int32 {
	// Write your code here
	err := int32(0)

	for i := 0; i < len(s); i += 1 {
		if string(s[i]) != "S" {
			err += 1
		}
		if string(s[i+1]) != "O" {
			err += 1
		}
		if string(s[i+2]) != "S" {
			err += 1
		}

		i += 2
	}

	return err
}
