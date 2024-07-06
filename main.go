package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//"strconv"
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

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run split.go -n <number_of_splits> <filename>")
		os.Exit(1)
	}

	filename := flag.Args()[0]

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	totalLines := len(lines)
	if n <= 0 || n > totalLines {
		fmt.Println("Invalid number of splits. It should be between 1 and the number of lines in the file.")
		os.Exit(1)
	}

	linesPerFile := totalLines / n
	remainder := totalLines % n

	base := filepath.Base(filename)
	extension := filepath.Ext(base)
	baseName := base[:len(base)-len(extension)]

	start := 0
	for i := 0; i < n; i++ {
		end := start + linesPerFile
		if remainder > 0 {
			end++
			remainder--
		}

		partFileName := fmt.Sprintf("%s_part%d%s", baseName, i+1, extension)
		partFile, err := os.Create(partFileName)
		check(err)
		defer partFile.Close()

		writer := bufio.NewWriter(partFile)
		for j := start; j < end; j++ {
			_, err := writer.WriteString(lines[j] + "\n")
			check(err)
		}
		writer.Flush()
		start = end
	}

	fmt.Println("File splitting completed.")
}
