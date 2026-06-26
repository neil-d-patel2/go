
Go is a compiled lanuage, toolchain converts .go files to machine code. 

This is accessed from go on the command line, simplest command is run which compiles the source code links to libs and then creates exectable and runs it

To create an executable, use go build

Go has over 100 stl libraries, fmt is one of them that allows u to print (contains Println)

Package main defines a standalone executable program, not a lib

Within package main the main function serves as a typical main function 

We need to tell compiler what packages we need, so we ONLY import what you need, or else you cannot compile 

Go is very strong on code formatting. Newlines following certain tokens are converted into semicolons, so newlines matter to properly parse go code

go's gofmt tool properly rewrites code into the standard format. 

1.2 CL args:

Most programs get some input, like an external source such as a file, network connection, etc 

The os package provides functions and other valeus fro dealiong with the os in a platfrom independent fashion 

command line arguments are availablke to a variable named args that is part of the os package 

variable os.Args is a slice of strings, and they are a dynamically sized sequence s

the first element of os.Args[0] is the name of the command itself, lets implement the unix echo command which prints its command line arguments on a single line

