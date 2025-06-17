package main

//NOTE:
//https://pkg.go.dev/bufio#example-Scanner-Words
//https://codingchallenges.fyi/challenges/challenge-wc/
//https://go.dev/blog/error-handling-and-go
//https://pkg.go.dev/os

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func err_log_exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var arg_slice = os.Args[1:len(os.Args)]
	if len(arg_slice) != 2 {
		err_log_exit(fmt.Errorf("Too litle or too many args. try 'go run . -c text.txt'"))
	}
	// fmt.Println(arg_slice)

	var filepath string
	var cmline_option string
	if strings.Contains(arg_slice[0], "-") {
		cmline_option = arg_slice[0]
		filepath = arg_slice[1]
	} else {
		cmline_option = arg_slice[1]
		filepath = arg_slice[0]
	}

	file, err := os.Open(filepath)
	err_log_exit(err)

	switch cmline_option {
	case "-c":
		text, err := file.Stat()
		err_log_exit(err)
		fmt.Println(text.Size(), "file:", text.Name())
	case "-l":
		var scanner *bufio.Scanner = bufio.NewScanner(file)
		var lines_c int = 0
		for scanner.Scan() {
			lines_c++
			// fmt.Println(scanner.Text())
		}

		fmt.Println(lines_c, filepath)
		err_log_exit(scanner.Err())
	case "-w":
		var scanner *bufio.Scanner = bufio.NewScanner(file)
		var words_c int = 0
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words_c++
			// fmt.Println(scanner.Text())
		}

		fmt.Println(words_c, filepath)
		err_log_exit(scanner.Err())
	case "-m":
		var scanner *bufio.Scanner = bufio.NewScanner(file)
		var chars_c int = 0
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			chars_c++
			// fmt.Println(scanner.Text())
		}

		fmt.Println(chars_c, filepath)
		err_log_exit(scanner.Err())
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes and name of file")
		fmt.Println("-l: Lines of text and name of file")
	}

	file.Close()
	os.Exit(0)
}
