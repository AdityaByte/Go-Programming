package main

import (
	"bytes"
	"fmt"
)

/* Bytes Buffer :
A bytes.bufffer is a mutable slice of byte that allows efficient appending and reading operation.*/

func main() {
	fmt.Println("Bytes Buffer in golang")
	buf := new(bytes.Buffer) // Here the new keyword just create a memory space of the default size and doesn't initialize it.
	buf.Write([]byte("Aditya"))
	buf.WriteString(" Pawar")
	fmt.Println("Length of buffer:", buf.Len())
	fmt.Println(buf.String())
}
