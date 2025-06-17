package main

//NOTE:"bufio" scanner
//https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go

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
		var line_counter int = 0
		for scanner.Scan() {
			line_counter++
			// fmt.Println(scanner.Text())
		}

		fmt.Println(line_counter, filepath)
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	case "-w":
		var scanner *bufio.Scanner = bufio.NewScanner(file)
		var line_counter int = 0
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			line_counter++
			// fmt.Println(scanner.Text())
		}

		fmt.Println(line_counter, filepath)
		err_log_exit(scanner.Err())
	default:
		fmt.Println("Non existing arg")
		fmt.Println("-c: Bytes and name of file")
		fmt.Println("-l: Lines of text and name of file")
	}

	file.Close()
	os.Exit(0)
}
