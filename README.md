# recreating wc in go
Recreated wc in go from the coreutils
Reads lines words characters and bytes from a file or input

## installing

## USAGE
Using the test.txt

Lines, words, characters
 `./go_wc test.txt`
7145 58164 342190 test.txt

Characters
`./go_wc test.txt -c`
342190 test.txt

Lines
`./go_wc test.txt -l`
7145 test.txt

Words
`./go_wc test.txt -w`
58164 test.txt

Runes | Characters (Porque -m? Ni idea)
`./go_wc test.txt -m`
339292 test.txt

Using standard input

`cat test.txt | ./go_wc`
7145 58164 342190 STDIN
`cat test.txt | ./go_wc -c`
342190 STDIN
`cat test.txt | ./go_wc -l`
7145 STDIN
`cat test.txt | ./go_wc -w`
58164 STDIN
`cat test.txt | ./go_wc -m`
339292 STDIN

## random toughts while doing this
Scanning a file the file empties the scanner making so I ended up re reading the file every time I did something with the text
I tried using Seek(0,0) to reset the file pointer on use
Then I found the same problem trying to read from stdin
The better aproach was to create this function that either reads from the stdin or the filepath given and returns the bytes
So changing the args to []byte and use it universally around all the functions that count the words, lines, runes, etc.

I did notice most of this functions looks alike maybe I can refactor them into something more reusable

Refactored into a single function that takes the bufio.SplitFunc, I couldnt find a good name for it, so its a little less readable but ill take that for reducing 4 functions into 1

## sources:
https://pkg.go.dev/bufio#example-Scanner-Words
https://codingchallenges.fyi/challenges/challenge-wc/
https://go.dev/blog/error-handling-and-go
https://pkg.go.dev/os
https://freshman.tech/snippets/go/read-console-input/
