package main

import (
	"fmt"
	"os"
)

func readFile() {
	file, err := os.Open("./open.txt")
	// file, err := os.Open("C:\\Users\\user\\my_go_tutorial\\open.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // registered here, runs when readFile() exits

	// do work with file...
	fmt.Println("doing work...")
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Content:", string(buf[:n]))
} // file.Close() runs here automatically

func main() {
	readFile()

	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
}
