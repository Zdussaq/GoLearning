package main

import "fmt"

func main() {
	fmt.Println(max(3, 5, 3, 2, -3, 1000, 3))
	fmt.Println(min(3, 5, 3, 2, -3, 1000, 3))
}

func max(vars ...int) (int, error) {

	if len(vars) < 1 {
		return -1, fmt.Errorf("must pass at least one var.")
	}

	max := vars[0]

	for _, v := range vars {
		if v > max {
			max = v
		}
	}

	return max, nil

}

func min(vars ...int) (int, error) {

	if len(vars) < 1 {
		return -1, fmt.Errorf("must pass at least one var.")
	}

	min := vars[0]

	for _, v := range vars {
		if v < min {
			min = v
		}
	}

	return min, nil

}
