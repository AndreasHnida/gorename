package main

import (
    "fmt"
    "math/rand"
    "os"
    "path/filepath"
    "time"
)

func main() {
    if len(os.Args) < 2 {
      fmt.Println("Renames a file to a random 10-digit number")
        fmt.Println("Usage: gorename <filename>")
        os.Exit(1)
    }

    // Get the original file path from command line arguments
    originalFile := os.Args[1]

    // Verify that the file exists
    if _, err := os.Stat(originalFile); os.IsNotExist(err) {
        fmt.Printf("The file '%s' does not exist.\n", originalFile)
        os.Exit(1)
    }

    // Extract the directory and suffix from the original file
    dir := filepath.Dir(originalFile)
    ext := filepath.Ext(originalFile)

    // Generate a random 10-digit number as a string
    rand.Seed(time.Now().UnixNano())
    randomDigits := fmt.Sprintf("%010d", rand.Int63n(10000000000))

    // Construct new file name with the original extension
    newFileName := filepath.Join(dir, randomDigits+ext)

    // Rename the file
    err := os.Rename(originalFile, newFileName)
    if err != nil {
        fmt.Println("Error renaming the file:", err)
        os.Exit(1)
    }

    fmt.Println("File renamed to", newFileName)
}

