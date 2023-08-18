package main

import (
	"flag"
	"fmt"
	"time"
)

// main() В данном коде основная функция  main
// использует пакет  flag  для определения строковых флагов
// командной строки. Далее вызывается функция  StartScan rootPath)
// которая выполняет дальнейшую логику.
func main() {
	starttime := time.Now()
	var rootPath string
	var fullVisibility string
	// Define string flags
	flag.StringVar(&rootPath, "ROOT", "", "Dir for scan")
	flag.StringVar(&fullVisibility, "fullPath", "", "will the paths be complete")

	// Parse the command-line arguments
	flag.Parse()

	// Access the string flag values
	fmt.Println("Dir for scan: ", rootPath)

	errMessage := StartScan(rootPath, fullVisibility)
	if errMessage != nil {
		fmt.Println(errMessage)
	}

	workTime := time.Since(starttime)
	fmt.Println("The program has worked: ", workTime, " seconds.")
}

//go run . --ROOT=/home/anton/go
