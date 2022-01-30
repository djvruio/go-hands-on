package main

import "fmt"

func main() {
	fmt.Println("Hello World 02!")
}

/*
go mod init hello_world02
Option 1:
go install
compile the program and install(copy) the binary to location C:\Users\rafae\go\bin
Option 2:
go build
creates the binary inside the location from which go build is called
Option 3:
go run main.go
or
go run --work main.go
compiles the file to a temporary location and runs the file from that location
*/
