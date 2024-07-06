package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "path/filepath"
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
        fmt.Println("Usage: go run split_large.go -n <number_of_splits> <filename>")
        os.Exit(1)
    }

    filename := flag.Args()[0]

    // Step 1: Count the total number of lines
    file, err := os.Open(filename)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalLines := 0
    for scanner.Scan() {
        totalLines++
    }
    check(scanner.Err())

    if n <= 0 || n > totalLines {
        fmt.Println("Invalid number of splits. It should be between 1 and the number of lines in the file.")
        os.Exit(1)
    }

    // Step 2: Calculate lines per file
    linesPerFile := totalLines / n
    remainder := totalLines % n

    // Step 3: Reopen the file and split it into smaller files
    file.Seek(0, 0) // Reset the file pointer to the beginning
    scanner = bufio.NewScanner(file)

    base := filepath.Base(filename)
    extension := filepath.Ext(base)
    baseName := base[:len(base)-len(extension)]

    for i := 0; i < n; i++ {
        partFileName := fmt.Sprintf("%s_part%d%s", baseName, i+1, extension)
        partFile, err := os.Create(partFileName)
        check(err)
        defer partFile.Close()

        writer := bufio.NewWriter(partFile)
        linesToWrite := linesPerFile
        if remainder > 0 {
            linesToWrite++
            remainder--
        }

        for j := 0; j < linesToWrite; j++ {
            if !scanner.Scan() {
                break
            }
            _, err := writer.WriteString(scanner.Text() + "\n")
            check(err)
        }
        writer.Flush()
    }

    fmt.Println("File splitting completed.")
}
