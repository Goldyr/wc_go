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

// Takes bytes and counts them depeding on the buifio.SplitFunc given
// bufio.ScanWords
// bufio.ScanBytes
// bufio.ScanLines
// bufio.ScanRunes (chars)
func count_by_function(file []byte, split_function bufio.SplitFunc) int {
	reader := bytes.NewReader(file)
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var counter int = 0
	scanner.Split(split_function)
	for scanner.Scan() {
		counter++
	}
	err_log_exit(scanner.Err())
	return counter
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
		bytes := count_by_function(raw_bytes, bufio.ScanBytes)
		words := count_by_function(raw_bytes, bufio.ScanWords)
		lines := count_by_function(raw_bytes, bufio.ScanLines)
		fmt.Println(lines, words, bytes, filepath)
	case "-c":
		bytes := count_by_function(raw_bytes, bufio.ScanBytes)
		fmt.Println(bytes, filepath)
	case "-l":
		lines := count_by_function(raw_bytes, bufio.ScanLines)
		fmt.Println(lines, filepath)
	case "-w":
		words := count_by_function(raw_bytes, bufio.ScanWords)
		fmt.Println(words, filepath)
	case "-m":
		runes := count_by_function(raw_bytes, bufio.ScanRunes)
		fmt.Println(runes, filepath)
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes in file")
		fmt.Println("-l: Lines in file")
		fmt.Println("-w: Words in file")
		fmt.Println("-m: Characters in file")
	}

	os.Exit(0)
}
