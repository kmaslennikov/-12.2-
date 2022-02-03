package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
  fileName:="messages.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := make([]byte, fileInfo.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", buf)

}
