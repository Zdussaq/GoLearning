package main

import "fmt"

func main() {
	fmt.Println(join("e ", "eeere", "ef dd ", "eeeee"))
}

func join(str ...string) string {
	var retStr string
	for _, v := range str {
		retStr += v
	}
	return retStr
}
