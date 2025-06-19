package main

import (
	"bufio"
	"bytes"
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

func bytes_count(file []byte) int {
	reader := bytes.NewReader(file)
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var bytes_c int = 0
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes_c++
	}

	err_log_exit(scanner.Err())
	return bytes_c
}

func lines_count(file []byte) int {
	reader := bytes.NewReader(file)
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var lines_c int = 0
	for scanner.Scan() {
		lines_c++
	}

	err_log_exit(scanner.Err())
	return lines_c
}

func words_count(file []byte) int {
	reader := bytes.NewReader(file)
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var words_c int = 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words_c++
	}

	err_log_exit(scanner.Err())
	return words_c
}

func runes_count(file []byte) int {
	reader := bytes.NewReader(file)
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var chars_c int = 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		chars_c++
	}
	err_log_exit(scanner.Err())
	return chars_c
}

func give_me_bytes(filepath string) []byte {
	var raw_bytes []byte
	var err error
	if filepath == "" {
		reader := bufio.NewReader(os.Stdin)
		raw_bytes, err = io.ReadAll(reader)
	} else {
		file, err := os.Open(filepath)
		err_log_exit(err)

		reader := bufio.NewReader(file)
		raw_bytes, err = io.ReadAll(reader)
		file.Close()
	}
	err_log_exit(err)
	return raw_bytes
}

func main() {
	var arg_slice = os.Args[1:len(os.Args)]
	var filepath string
	var cmline_option string

	switch len(arg_slice) {
	case 0:
		{
			//No arg no filepath
			//NOTE: filepath and cmline_option are defaulted by go to ""
		}
	case 1:
		{
			if strings.Contains(arg_slice[0], "-") {
				//Arg no filepath
				cmline_option = arg_slice[0]
			} else {
				//No arg filepath
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

	//NOTE: Either brings the bytes or quits
	var raw_bytes []byte = give_me_bytes(filepath)

	if filepath == "" {
		filepath = "STDIN"
	}

	switch cmline_option {
	case "":
		bytes := bytes_count(raw_bytes)
		words := words_count(raw_bytes)
		lines := lines_count(raw_bytes)
		fmt.Println(lines, words, bytes, filepath)
	case "-c":
		fmt.Println(bytes_count(raw_bytes), filepath)
	case "-l":
		fmt.Println(lines_count(raw_bytes), filepath)
	case "-w":
		fmt.Println(words_count(raw_bytes), filepath)
	case "-m":
		fmt.Println(runes_count(raw_bytes), filepath)
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes in file")
		fmt.Println("-l: Lines in file")
		fmt.Println("-w: Words in file")
		fmt.Println("-m: Characters in file")
	}

	os.Exit(0)
}
