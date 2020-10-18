package main

import "fmt"

func main2() {
	var r1, c1 int
	var r2, c2 int
	var a [5][5]int
	var b [5][5]int
	var ans [5][5]int
	var c int
	var s int
	fmt.Println("Enter no. of rows and columns for first array :")
	fmt.Scanln(&r1, &c1)

	fmt.Println(r1, c1)
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scanln(&a[i][j])
		}
	}
	fmt.Println("Enter no. of rows and columns for second array :")
	fmt.Scanln(&r2, &c2)
	fmt.Println(r2, c2)

	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scanln(&b[i][j])
		}
	}
	fmt.Println("Enter 0 for addition, 1 for subtraction, 2 for multiplication")
	fmt.Scanln(&c)
	switch c {
	case 0:
		if r1 == r2 && c1 == c2 {
			for i := 0; i < r1; i++ {
				for j := 0; j < c1; j++ {
					ans[i][j] = a[i][j] + b[i][j]
				}
			}
		} else {
			fmt.Println("The entered matrices have different dimensions")
		}
		fmt.Println("New matrix after addition is ")
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Printf("%d ", ans[i][j])
			}
			fmt.Printf("\n")

		}

	case 1:
		if r1 == r2 && c1 == c2 {
			for i := 0; i < r1; i++ {
				for j := 0; j < c1; j++ {
					ans[i][j] = a[i][j] - b[i][j]
				}
			}
		} else {
			fmt.Println("The entered matrices have different dimensions")
		}
		fmt.Println("New matrix after subtraction is ")
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Printf("%d ", ans[i][j])
			}
			fmt.Printf("\n")
		}
	case 2:
		fmt.Println("Enter the scalar to be multiplied")
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

		fmt.Println("The values of matrices after multiplying scalar values are")
		for i := 0; i < r1; i++ {
			for j := 0; j < c1; j++ {
				fmt.Printf("%d ", a[i][j])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n \n")
		for i := 0; i < r2; i++ {
			for j := 0; j < c2; j++ {
				fmt.Printf("%d ", b[i][j])
			}
			fmt.Printf("\n")
		}
	}

}
