package main

import "fmt"

func Sum(x int, y int) []byte {
	v := fmt.Sprintf("%v", x+y)

	return []byte(v)
}
