package main

import "fmt"

func main() {
	fmt.Println(average(43, 56, 87, 12, 45, 57))
}
func average(sf ...float64) (total float64) {
	for _, v := range sf {
		total += v
	}
	total /= float64(len(sf))
	return
}
