// https://www.hackerrank.com/challenges/circular-array-rotation/

package main

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	for i := 0; i < int(k); i += 1 {
		a = append([]int32{a[len(a)-1]}, a[:len(a)-1]...)
	}

	resp := []int32{}
	for _, v := range queries {
		resp = append(resp, a[v])
	}

	return resp
}
