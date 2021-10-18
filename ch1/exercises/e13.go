//Exercise 1.3 is to time and find which print method is most effecient
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	getTime(echo1, "Echo1")
	getTime(echo2, "Echo2")
	getTime(echo3, "Echo3")

}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func getTime(f func(), functionName string) {
	start := time.Now()
	f()
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("%s took %s\n", functionName, elapsed)
}
