package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}
func greet(fname, lname string) (s1, s2 string) {
	s1 = fmt.Sprint(fname, lname)
	s2 = fmt.Sprint(lname, fname)
	return
}
