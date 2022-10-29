package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fp, err := os.Create("3/plan.txt")
	if err != nil {
		fmt.Println(err)
	}
	content := "I’m not afraid of difficulties and insist on learning programming"
	fp.Write([]byte(content))
	fp.Close()
	file, _ := os.Open("3/plan.txt")
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}
	fmt.Println("File Content =", string(chunk))
	file.Close()
}
