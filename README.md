# Basicfs

A basic and naive disk file system.

## Build

	go build

## Examples

Compile, run, print help, and exit:

	; go build && ./basicfs
	> help
	Valid commands:
		help		­ print this description
		exit		­ terminate prompt loop
		cd path		­ change current directory to 'path'
		ls [path]	­ list directory contents, current by default
		rm name		­ delete a file called 'name'
		emit name	- print the raw file data called 'name'
		save [name]	­ write cache to disk called 'name', 'nil.bfs' by default
		load [name]	­ load a disk into memory called 'name', 'nil.bfs' by default
	> quit
	Goodbye ☺
	;

