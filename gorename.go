package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const DEFAULT_LENGTH = 10

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Renames a file to a random 10-digit number")
		fmt.Println("Usage: gorename <pattern>")
		os.Exit(1)
	}

	run()

}
func run() {
	// Verify that the file exists
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file pattern (e.g., *.jpg) or a directory (e.g. . ) ")
	}
	pattern := os.Args[1]

	var files []string
	var err error

	fmt.Println("Reading Files...")

	if pattern == "." {
		files, err = readAllFiles(".")
	} else {
		files, err = filepath.Glob(pattern)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files successfully read: ")
	fmt.Println(files)

	err = renameFiles(files)
	if err != nil {
		log.Fatal(err)
	}
	// Generate a random 10-digit number
}
func renameFiles(files []string) error {

	fmt.Println("Renaming files...")

	for _, file := range files {
		ext := filepath.Ext(file)
		fmt.Println("renaming: ", file, " to")
		randomDigits := generateRandomDigits()
		newName := randomDigits + ext
		fmt.Print(newName)
		err := os.Rename(file, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateRandomDigits() string {
	length := DEFAULT_LENGTH
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	fmt.Println("Generating Random Digits...")
	digits := make([]byte, length)
	for i := 0; i < length; i++ {
		digits[i] = byte(rng.Intn(10) + '0')
	}
	fmt.Println("Random Digigts: ", string(digits))
	return string(digits)
}

func readAllFiles(dir string) ([]string, error) {
	var files []string

	absPath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Reading all files from: ", absPath)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}
