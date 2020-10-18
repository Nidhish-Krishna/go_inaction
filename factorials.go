package main

import (
	"fmt"
	"time"
)

func fact(n int) {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * i
	}
	fmt.Println("Factorial of ", n, " is", ans)
}
func main1() {
	for i := 1; i <= 10; i++ {
		go fact(i)
	}
	time.Sleep(1 * time.Second)

}
