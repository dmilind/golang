#### Go Compiler

In a nutshell go compiler is a program which processes the source code and provides object code and executable binary code. Source code is the computer program written in golang, which is human readable. Once source code is compiled, compiler creates an object code (a file with .o extension). An object code is a machine code called as machine language which is understood by machine with specific architecture. Some compiler creates assembly language which is also human readable.

Ig golang, object code can be created using go tools. `go tool compile <program.go>` generates object code. This code is in relocatable format. This relocatable information is used by linker to assemble other object files and provide one executable file. Object code also contains other information such as program symbols (variable names and functions).   

Go compiler can provide an archive file also which is with .a extension. `go tool compile --pack` is responsible for providing archive file. This file is a collection or all other files after linking. This is a binary file. Contains of this archive file can be seen by executing `ar t <program>.a` command.

`mdhoke＠codingFun[master ?] ➤ ar t example01.a
__.PKGDEF
_go_.o
`
As you can see archive file contains object code.Since my program is not using other files thats why single object file is encapsulated.

#### Go Garbage Collector

In program , memory always gets allocated to the objects specified in the program. If memory is not freed up then there would be a performance delay with CPU and RAM. So it is necessary to revoke all the memory consumed by the objects which are not referenced in the program execution. This method of revoking memory , is called as garbage collection.
In go , garbage collection, runs with mutator thread. That means when program is in execution, multiple GX threads runs to free up the memory. To understand the GC concept find example-gc.go from the repo. 
