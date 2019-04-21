package main

import (
	"fmt"
	"os"
	"time"
)

var (
	filename string
	message  string
)

func main() {
	// arg check. If no arg, error exit
	tasks := len(os.Args) - 1
	if tasks < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s task1 task2...", os.Args[0])
		os.Exit(1)
	}

	// print tasks
	for i := 1; i <= tasks; i++ {
		fmt.Printf("%s\n", os.Args[i])

	}
	// make filename
	time := time.Now().Format("2006-01-02")
	filename = "./store/" + fmt.Sprintf("%v", time) + ".txt"

	// open file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// make message
	for i := 1; i <= tasks; i++ {
		message = message + fmt.Sprintf("%s\n", os.Args[i])
	}

	//write file
	_, err = file.WriteString(message)
	if err != nil {
		panic(err)
	}

}
