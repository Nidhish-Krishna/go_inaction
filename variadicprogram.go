package main

import "fmt"

func prt(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf(" name %d : %s\n", i+1, s[i])

	}
}

func main() {
	n := 0
	fmt.Println("Enter the number of strings")
	fmt.Scanln(&n)
	st := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&st[i])
	}

	fmt.Println("Calling variadic function to print a set of names...")
	prt(st...)
}
