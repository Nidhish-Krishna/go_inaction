package main

import "fmt"

func mainma() {
	m1 := make(map[int]int)
	var k int
	var v int
	var key int
	var val bool
	fmt.Println("Enter the key and value to be entered in map1")
	for i := 0; i < 2; i++ {
		fmt.Scanln(&k, &v)
		m1[k] = v
	}
	fmt.Println("map 1:", m1)
	k = 0
	v = 0
	m2 := make(map[int]int)
	fmt.Println("Enter the key and value to be entered in map2")
	for i := 0; i < 3; i++ {
		fmt.Scanln(&k, &v)
		m2[k] = v
	}
	fmt.Println("map 2:", m2)
	fmt.Println("To check if a key is in a map")
	fmt.Println("Enter the key to be checked:")
	fmt.Scanln(&key)

	_, val = m1[key]

	if val == true {
		fmt.Printf("\n")
		fmt.Printf("The key %d is present in map1 and its value is %d", key, m1[key])
	} else {
		fmt.Println("The key is not present in the maps")
	}

	fmt.Println("")
	fmt.Println("map 1 before merging:", m1)
	fmt.Println("map 2               :", m2)
	fmt.Println("Merging m2 with m1")
	for id2, value2 := range m2 {
		m1[id2] = value2
	}

	fmt.Println("map 1 after merging:", m1)

}
