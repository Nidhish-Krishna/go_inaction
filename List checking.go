package main

import "fmt"

func intersection(a []int, b []int) []int {
	l1 := len(a)
	l2 := len(b)
	var ans []int
	ans = make([]int, 0)
	var i, j int
	ans = make([]int, 0)
	for i = 0; i < l1; i++ {
		for j = 0; j < l2; j++ {
			if a[i] == b[j] {
				ans = append(ans, a[i])
			}
		}
	}
	return ans
}
func main1() {

	var slc1 = []int{1, 2, 3, 4}
	var slc2 = []int{1, 2, 3, 4}
	var slc []int

	fmt.Println("Program to check subsets of list")
	fmt.Println("assuming the arrays in sorted order and none of the entries are repeated in a single array")

	fmt.Println(slc1)
	fmt.Println(slc2)

	slc = intersection(slc1, slc2)
	fmt.Println(slc)

	if len(slc) == len(slc1) && len(slc) == len(slc2) {
		fmt.Println("Both the arrays are contained inside each other")
	} else if len(slc) < len(slc2) {
		fmt.Println("List 1 is contained inside list 2")

	} else if len(slc) < len(slc1) {
		fmt.Println("List 2 is contained inside List 1")

	} else if len(slc) == 0 {
		fmt.Println("None of the predicted")

	} else {
		fmt.Println("None of the predicted")
	}

}
