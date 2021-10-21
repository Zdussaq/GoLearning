//Removed adjacent duplicate entities in array.
package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 6, 2}
	fmt.Println(arr)
	arr = removeDuplicates(arr)
	fmt.Println(arr)
}

func removeDuplicates(arr []int) []int {
	i := 0
	for _, v := range arr {
		if arr[i] != v {
			i++
			arr[i] = v
		}
	}
	return arr[:i+1]

}
