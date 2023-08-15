package main

import (
	"os"
	"fmt"
	"log"
	"io"
)

type UselessReader struct {}
type UselessWriter struct {}

func (this *UselessReader) Read (destination []byte) (int, error) {
	fmt.Print("shoot> ")
	return os.Stdin.Read(destination) // returns both size of bytes and err
}

func (this *UselessWriter) Write (source []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(source) // returns both size of bytes and err
}

func UselessCopy(reader *UselessReader, writer *UselessWriter) {
	var (
		size int
		err error
	)
	input := make([]byte, 1024)

	if size, err = reader.Read(input); err != nil {
		log.Fatalln("Congrats; You just fucked up reading! ", err)
	}
	fmt.Println("Successfully read", size, "bytes.\n")

	if size, err = writer.Write(input[:size]); err != nil {
		// slicing the input slice prevents the writer to write empty bytes of it 
		//(because input len is 1024 bytes) 
		log.Fatalln("Congrats; You just fucked up writing! ", err)
	}
	fmt.Println("Successfully wrote", size, "bytes.\n")
}
func main() {
	var (
		reader UselessReader
		writer UselessWriter
	)
	fmt.Println("By My Useless Reader & Writer:")
	UselessCopy(&reader, &writer)

	fmt.Println("By io.Copy:")

	if written, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Congrats for fucking up while copying!", err)
	} else {
		fmt.Println("Successfully copied", written, "bytes.")
	}
}