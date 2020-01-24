package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"bufio"
	"flag"
	"encoding/binary"
)


const (
	defaultName	= "nil.bfs"				// Default disk file name
	prompt		= "> "					// Prefix for input prompt
	maxData		= 256					// Max bytes in a file
	maxKids		= 256					// Max children of a directory
)

var (
	chatty		bool					// Verbose debug output
	root		File					// Top level file of the file system
	user		string					// User of the file system
	endian		= binary.LittleEndian	// Disk binary endian-ness
)


// Basic disk file system
func main() {
	flag.BoolVar(&chatty, "D", false, "Verbose debug output")
	flag.StringVar(&user, "u", "none", "File system user")
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

			fatal("err: could not read input -", err)
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
			// TODO


		case "rm":
			if nfields < 2 {
				fmt.Println("fail: rm requires a 'name' argument")
				continue repl
			}

			// TODO

		case "emit":
			if nfields < 2 {
				fmt.Println("fail: emit requires a 'name' argument")
				continue repl
			}

			// TODO

		case "save":
			fname := defaultName
			if nfields > 1 {
				fname = fields[1]
			}

			f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0640)
			if err != nil {
				fatal("err: could not open file for write -", err)
			}
			defer f.Close()

			// TODO - write out differently, this is broken
			err = binary.Write(f, endian, root)
			if err != nil {
				fatal("err: could not write file -", err)
			}

		case "load":
			fname := defaultName
			if nfields > 1 {
				fname = fields[1]
			}

			f, err := os.Open(fname)
			if err != nil {
				fatal("err: could not open file for read -", err)
			}
			defer f.Close()

			// TODO - rewrite to match `save` format
			err = binary.Read(f, endian, &root)
			if err != nil {
				fatal("err: could not read file -", err)
			}

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
	help		- print this description
	exit		- terminate prompt loop
	cd path		- change current directory to 'path'
	ls [path]	- list directory contents, current by default
	rm name		- delete a file called 'name'
	emit name	- print the raw file data called 'name'
	save [name]	- write cache to disk called 'name', 'nil.bfs' by default
	load [name]	- load a disk into memory called 'name', 'nil.bfs' by default
`
