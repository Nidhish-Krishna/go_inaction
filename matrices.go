package main

import "fmt"

func maind() {
	var a [5][5]int
	var b [5][5]int
	var ans [5][5]int

	var r1 int

	var c1 int
	var r2 int
	var c2 int
	var s int
	var c int
	fmt.Println("Enter rows, cols. and entires of first matrix :")
	fmt.Scanln(&r1, &c1)
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scanln(&a[i][j])
		}
	}
	/*
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Print(a[i][j], " ")
			}
			fmt.Printf("\n")
		}
	*/

	fmt.Println("Enter rows, cols. and entires of second matrix :")
	fmt.Scanln(&r2, &c2)
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scanln(&b[i][j])
		}
	}
	/*
		for i := 0; i < r2; i++ {
			for j := 0; j < c2; j++ {
				fmt.Print(b[i][j], " ")
			}
			fmt.Printf("\n")
		}
	*/
	fmt.Println("Enter 0 for addition, 1 for subtraction and 2 for scalar multiplication ")
	fmt.Scanln(&c)
	//fmt.Println(c)

	switch c {
	case 0:
		if r1 == r2 && c1 == c2 {
			for i := 0; i < r1; i++ {
				for j := 0; j < c1; j++ {
					ans[i][j] = a[i][j] + b[i][j]
				}
			}
		}
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Print(ans[i][j], " ")
			}
			fmt.Print("\n")
		}
	case 1:
		if r1 == r2 && c1 == c2 {
			for i := 0; i < r1; i++ {
				for j := 0; j < c1; j++ {
					ans[i][j] = a[i][j] - b[i][j]
				}
			}
		}
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Print(ans[i][j], " ")
			}
			fmt.Print("\n")
		}
	case 2:
		fmt.Println("Enter the value of scalar that has to be multiplied : ")
		fmt.Scanln(&s)
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				a[i][j] = a[i][j] * s
			}
		}
		for i := 0; i < r2; i++ {
			for j := 0; j < c2; j++ {
				b[i][j] = b[i][j] * s
			}
		}
		fmt.Println("The matrices after multiplication are as follows...")
		fmt.Println("Matrix A")
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Print(a[i][j], " ")
			}
			fmt.Printf("\n")
		}
		fmt.Println("Matrix B")
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Print(b[i][j], " ")
			}
			fmt.Printf("\n")
		}

	default:
		fmt.Println("Enter a valid choice")

	}

}
