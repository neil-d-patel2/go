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

Go provides usual arithmetic and logical operators, but when applied to strings, it concatentates the values, so sep + os.Args[i] represents the concat of sep and os.Args[i] os.Args[i] 


traditional for loop (no parentheses in go)  for initialization; condition; post

for initialization; condition; post {
    //statements
}


braces are mandatory, parentheses are not, braces must be on the same line as the post statement in go 


 

another from of gos for loop iterates over a range of values from a data type like a string or a slice, we can see this in echo2.go 

The solution to having a for loop but not using the index variable is using a 
blank identifier _, because go doesnt allow unused local variables

s := "" // can only be used in a function, can do this for a package lvl variable


Fetching a url:



for many apps, access to informatin from the internet is as important as access to the local file system , go provides a lot of packages grouped under net that make it easy to send and receive informatino through the Internet make low level network connections and set up servers, this is where gos concurrency features become really useful






