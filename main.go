package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 2, "Number of smaller lists to create")
	flag.Parse()

	if n <= 0 {
		fmt.Println("The number of splits should be greater than 0.")
		os.Exit(1)
	}

	// Open stdin as a bufio.Reader
	reader := bufio.NewReader(os.Stdin)

	// Initialize counters and file writers
	totalLines := 0
	splits := make([]int, n)
	writers := make([]*bufio.Writer, n)
	files := make([]*os.File, n)

	defer func() {
		// Ensure all files are closed properly
		for i := 0; i < n; i++ {
			if writers[i] != nil {
				err := writers[i].Flush()
				check(err)
				err = files[i].Close()
				check(err)
			}
		}
	}()

	// Process input line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // End of input
		}
		totalLines++
		idx := totalLines % n
		splits[idx]++

		// Initialize file writer for the current split if not already initialized
		if writers[idx] == nil {
			filename := fmt.Sprintf("split_part%d.txt", idx+1)
			file, err := os.Create(filename)
			check(err)
			files[idx] = file
			writers[idx] = bufio.NewWriter(file)
		}

		// Write line to the corresponding split file
		_, err = writers[idx].WriteString(line)
		check(err)
	}

	// Output summary of splits
	fmt.Printf("Input lines: %d\n", totalLines)
	for i := 0; i < n; i++ {
		fmt.Printf("Split %d: %d lines\n", i+1, splits[i])
	}
	fmt.Println("Splitting completed.")
}
