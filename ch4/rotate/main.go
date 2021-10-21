//exercise 4.4
package main

import "fmt"

func rotate(arr []int) {
	front := arr[0]
	copy(arr[:], arr[1:])
	arr[len(arr)-1] = front
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	rotate(arr)
	fmt.Println(arr)
}
