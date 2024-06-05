package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filepath := os.Args[1]

	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file : %v\n", err)
		return
	}
	fmt.Printf("%s", file)

	// file, err := os.Open("quest8.txt")

	// if err != nil {
	// 	fmt.Printf("Error is: %v", err.Error())
	// }

	// arr := make([]byte, 32)

	// file.Read(arr)

	// fmt.Println(string(arr))

	// file.Close()
}
