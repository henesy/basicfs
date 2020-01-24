package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"bufio"
	"flag"
)


const (
	defaultName	= "nil.bfs"	// Default disk file name
	prompt		= "> "		// Prefix for input prompt
)

var (
	chatty		bool		// Verbose debug output
)


// Basic disk filesystem
func main() {
	flag.BoolVar(&chatty, "D", false, "Verbose debug output")
	flag.Parse()

	in := bufio.NewReader(os.Stdin)

	// Main input loop
	repl:
	for {
		fmt.Print(prompt)
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break repl
			}

			fatal("err: could not read input - ", err)
		}

		fields := strings.Fields(string(line))
		nfields := len(fields)

		if nfields < 1 {
			fmt.Println("fail: specify a command such as `help`")
			continue repl
		}

		switch fields[0] {
		case "help":
			fmt.Print(helpStr)

		case "cd":
			if nfields < 2 {
				fmt.Println("fail: cd requires a 'path' argument")
				continue repl
			}

		case "ls":


		case "rm":
			if nfields < 2 {
				fmt.Println("fail: rm requires a 'name' argument")
				continue repl
			}

		case "save":


		case "load":


		case "quit":
			fallthrough
		case "exit":
			break repl
		}
	}

	fmt.Println("Goodbye ☺")
}


// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}


var helpStr string = `Valid commands:
	help		­ print this description
	exit		­ terminate prompt loop
	cd path		­ change current directory to 'path'
	ls [path]	­ list directory contents, current by default
	rm name		- delete a file called 'name'
	save [name]	­ write cache to disk called 'name', 'nil.bfs' by default
	load [name]	- load a disk into memory called 'name', 'nil.bfs' by default
`
