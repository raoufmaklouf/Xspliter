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

	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	totalLines := len(lines)
	if totalLines == 0 {
		fmt.Println("No input provided.")
		os.Exit(1)
	}

	if n > totalLines {
		fmt.Printf("Number of splits %d is greater than the number of lines %d in the input.\n", n, totalLines)
		os.Exit(1)
	}

	// Calculate lines per file
	linesPerFile := totalLines / n
	remainder := totalLines % n

	// Get the base name for the output files
	base := "split_part"
	extension := ".txt"

	for i := 0; i < n; i++ {
		partFileName := fmt.Sprintf("%s%d%s", base, i+1, extension)
		partFile, err := os.Create(partFileName)
		check(err)
		defer partFile.Close()

		writer := bufio.NewWriter(partFile)
		startLine := i * linesPerFile
		endLine := startLine + linesPerFile
		if i == n-1 {
			// Add the remainder to the last file
			endLine += remainder
		}

		for _, line := range lines[startLine:endLine] {
			_, err := writer.WriteString(line + "\n")
			check(err)
		}
		writer.Flush()
	}

	fmt.Println("Splitting completed.")
}
