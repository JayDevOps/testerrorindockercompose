package main

import (
	"fmt"
	"errors"
	"os"
	"bufio"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		errors.New("Invalid File Name")
	}
	fmt.Println("Checking for any errors")

	y, err1 := checkDockerComposeIsError(os.Args[1])
	if err1 != nil {
		fmt.Printf("%d: ",y)
		fmt.Println(fmt.Errorf("%v", err1))
	} else {
		fmt.Println("No Error")
	}
}


func checkDockerComposeIsError(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
			return 0, err
	}
	defer f.Close()

	// splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	var isError bool = false
	// https://golang.org/pkg/bufio/#scanner.scan
	for scanner.Scan() {
			fmt.Println(scanner.Text())
			if strings.Contains(scanner.Text(), "Err") {
					isError = true
			}
			line++
	}

	if err := scanner.Err(); err != nil {
			return 125, err
			// handle the error
	}
	
	if isError == true {
		return line, errors.New("One or more errors in processing pipeline task")
	} else {
		return 0, nil
	}
}
