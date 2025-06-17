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

func bytes_count(filepath string) int64 {
	file, err := os.Open(filepath)
	err_log_exit(err)
	text, err := file.Stat()
	err_log_exit(err)
	file.Close()
	return text.Size()
}

func lines_count(filepath string) int64 {
	file, err := os.Open(filepath)
	err_log_exit(err)
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var lines_c int64 = 0
	for scanner.Scan() {
		lines_c++
	}

	err_log_exit(scanner.Err())
	file.Close()
	return lines_c
}

func words_count(filepath string) int64 {
	file, err := os.Open(filepath)
	err_log_exit(err)
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var words_c int64 = 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words_c++
		// fmt.Println(scanner.Text())
	}

	err_log_exit(scanner.Err())
	file.Close()
	return words_c
}

func runes_count(filepath string) int64 {
	file, err := os.Open(filepath)
	err_log_exit(err)
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var chars_c int64 = 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		chars_c++
		// fmt.Println(scanner.Text())
	}
	err_log_exit(scanner.Err())
	file.Close()
	return chars_c
}
func main() {
	var arg_slice = os.Args[1:len(os.Args)]
	// if len(arg_slice) < 2 {
	// 	err_log_exit(fmt.Errorf("Too litle or too many args. try 'go run . -c text.txt'"))
	// }

	// fmt.Println(arg_slice)

	var filepath string
	var cmline_option string

	if strings.Contains(arg_slice[0], "-") {
		cmline_option = arg_slice[0]
		filepath = arg_slice[1]
	} else if len(arg_slice) > 1 && strings.Contains(arg_slice[1], "-") {
		cmline_option = arg_slice[1]
		filepath = arg_slice[0]
	} else {
		filepath = arg_slice[0]
	}

	switch cmline_option {
	case "":
		fmt.Println(lines_count(filepath), words_count(filepath), bytes_count(filepath), filepath)
	case "-c":
		fmt.Println(bytes_count(filepath), "file:", filepath)
	case "-l":
		fmt.Println(lines_count(filepath), filepath)
	case "-w":
		fmt.Println(words_count(filepath), filepath)
	case "-m":
		fmt.Println(runes_count(filepath), filepath)
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes in file")
		fmt.Println("-l: Lines in file")
		fmt.Println("-w: Words in file")
		fmt.Println("-m: Characters in file")

	}

	os.Exit(0)
}
