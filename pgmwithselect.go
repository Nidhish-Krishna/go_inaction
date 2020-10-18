/*
Date created		: 03/10/2020 - 9 AM
Date last modified	: 05/10/2020 - 7 PM
Problem Statement 	:
	1. The idea of a simple test taking platform is implemented in this file
	2. The program receives questions from one file and write the answers recorded in another file

	3. The entire program should run for say, x minutes or seconds
	4. Once a question is displayed, user has a maximum of 10 seconds/5seconds to answer a question
	5. If user does not enter an answer within 10/5 seconds, a null string is returned and written to the res.txt
	6. Once time is up, the program  should terminate
Name of question paper file		: Sample.txt
Name of the file to be created	: res.txt

Need 				:
	1. The functions like Scanln() and Scanf() are blocking statements that block input stream
	2. So we use goroutines and channel synchronization in golang to achieve step 5 and 6
	3. Also the program should run for say x seconds which is the duration of the test i.e
	the idea is that our main program execution should wait until user is completing the execution of program

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

/*
The readfile1() is for identifying how many characters are present in each line
The program will not know how many bytes of character is present in each line
For identifying how many characters are present in each line, we use readfile()

readfile returns an array that contains no. of characters present in eachline

This is important when we are trying to print the question by navigating file line by line later !

*/
func readfile1() []int {
	file, err := os.Open("Sample.txt") // open the question paper file
	a := make([]int, 0)                // create a slice of no length which is similar to list in Python
	if err != nil {                    // to detect if there is file open error
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // NewScanner() returns a *Scanner type that is used to navigate through file

	for scanner.Scan() { // Scan() method navigates file till EOF is reached, it takes a split parameter as input

		a = append(a, len(scanner.Text())) //  by default, the parameter is SplitLines
	}
	return a //return slice a that contains no. of characters in each line
}

/*
The getInput () is used to get input from the user without blocking the input stream
The functions gets a channel as an input
and if user enters an input from input stream
That input is converted to a string and is written to an input channel
*/
func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- result
	}
}

/*
This program is the main routine program
It reads a sample file which contains questions and writes to the res.txt file

*/

func readfile2(arr []int) {
	file, err := os.Open("Sample.txt") //open question paper file
	if err != nil {                    // read if error in file open occurs
		log.Fatal(err)
	}
	defer file.Close()
	offs := int64(0) //offset variable to navigate a file

	fi, er := os.Create("res.txt") // create answer file res.txt
	if er != nil {                 // to read if there is error in file creation
		log.Fatal(er)
	}
	input := make(chan string) // input channel used to get input from channel

	for i := 0; i < 16; i += 2 { // read ith line, (i+2)th line and so on
		for j := i; j <= i+1; j++ { // read ith line and (i+1)th line and print on the screen
			data := make([]byte, arr[j]+1)
			file.Read(data) // read i th line from file
			fmt.Println(string(data))
			offs = offs + int64(arr[j]+2) //
			file.Seek(offs, 0)            // navigate file upto (offs) bytes from starting position of file pointer
		}
		fmt.Println("Enter a response for question: ")

		go getInput(input) //routine for getting innput from user

		select {
		case i := <-input: // this is executed if an input is entered in the input stream

			fi.WriteString(i)
			fmt.Println(i)
		case <-time.After(5000 * time.Millisecond): // is no input is entered, this executes
			fi.WriteString(" ")
			fmt.Println("timedout")

		}

	}
	defer file.Close()
	defer fi.Close()

}

func main() {
	linecharcount := readfile1()

	fmt.Println("Instructions: ")
	fmt.Println("1. The test duration is 16 seconds")
	fmt.Println("2. You have to enter an answer within 5 seconds")
	fmt.Println("Enter password for starting test: ")
	passwd := ""
	fmt.Scanln(&passwd)
	if passwd == "0000" { //password authentication
		fmt.Println("Your test has started")
		go readfile2(linecharcount) // start routine for test
	} else {
		fmt.Println("The password you entered is incorrect")
	}

	time.Sleep(16 * time.Second) // sleep mainroutine till the duration of test

	fmt.Println("Test is completed")
}
