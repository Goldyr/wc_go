package main

//NOTE:
//https://pkg.go.dev/bufio#example-Scanner-Words
//https://codingchallenges.fyi/challenges/challenge-wc/
//https://go.dev/blog/error-handling-and-go
//https://pkg.go.dev/os
//https://freshman.tech/snippets/go/read-console-input/

//WARN:
//Problems: Scanning a file the file empties the scanner making
//You have re read the file every instance
//I didnt have this problem because I was only doing one scan at the time eg: words only, lines only
//Suspect: File pointer
//Gotta use Seek(0,0) to reset file after using them

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func err_log_exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func bytes_count(file io.Reader) int {
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var bytes_c int = 0
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes_c++
	}

	err_log_exit(scanner.Err())
	return bytes_c
}

func lines_count(file io.Reader) int64 {
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var lines_c int64 = 0
	for scanner.Scan() {
		lines_c++
	}

	err_log_exit(scanner.Err())
	return lines_c
}

func words_count(file io.Reader) int64 {
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var words_c int64 = 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words_c++
		// fmt.Println(scanner.Text())
	}

	err_log_exit(scanner.Err())
	return words_c
}

func runes_count(file io.Reader) int64 {
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	var chars_c int64 = 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		chars_c++
		// fmt.Println(scanner.Text())
	}
	err_log_exit(scanner.Err())
	return chars_c
}
func main() {
	var arg_slice = os.Args[1:len(os.Args)]
	var filepath string
	var cmline_option string

	switch len(arg_slice) {
	case 0:
		{
			//No args no filepath
			//NOTE: filepath and cmline_option are defaulted by go to ""
		}
	case 1:
		{
			if strings.Contains(arg_slice[0], "-") {
				//Args no filepath
				cmline_option = arg_slice[0]
			} else {
				//No args filepath
				filepath = arg_slice[0]
			}
		}
	case 2:
		{
			//Args and filepath
			if strings.Contains(arg_slice[0], "-") {
				//Arg in 0
				cmline_option = arg_slice[0]
				filepath = arg_slice[1]
			}
			if strings.Contains(arg_slice[1], "-") {
				//Arg in 1
				cmline_option = arg_slice[1]
				filepath = arg_slice[0]
			}
		}
	default:
		{
			err_log_exit(fmt.Errorf("Multiple args not supported"))
		}

	}

	if filepath == "" {
		reader := bufio.NewReader(os.Stdin)

		// text, _ := reader.ReadString('\n')
		//WARN:Reader
		// reader := bufio.NewReader(os.Stdin)
		// text, _ := reader.ReadString('\n')
		// if text != "" {
		// 	fmt.Println(text)
		// } else {
		// 	fmt.Println("erro")
		// }
		// fmt.Println(lines_count(reader), words_count(reader), "TODO")

		copy := reader
		copy2 := reader
		bytes := bytes_count(reader)
		words := words_count(copy)
		lines := lines_count(copy2)
		fmt.Println(lines, words, bytes, filepath)
		os.Exit(0)
	}

	file, err := os.Open(filepath)
	err_log_exit(err)

	switch cmline_option {
	case "":
		bytes := bytes_count(file)
		file.Seek(0, 0)
		words := words_count(file)
		file.Seek(0, 0)
		lines := lines_count(file)
		fmt.Println(lines, words, bytes, filepath)
		// fmt.Println(lines_count(file), words_count(file), bytes_count(file), filepath)
	case "-c":
		fmt.Println(bytes_count(file), "file:", filepath)
	case "-l":
		fmt.Println(lines_count(file), filepath)
	case "-w":
		fmt.Println(words_count(file), filepath)
	case "-m":
		fmt.Println(runes_count(file), filepath)
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes in file")
		fmt.Println("-l: Lines in file")
		fmt.Println("-w: Words in file")
		fmt.Println("-m: Characters in file")
	}

	file.Close()
	os.Exit(0)
}
