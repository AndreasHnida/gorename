package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const DEFAULT_LENGTH = 10

// global variables for flags
var prefix string
var length int
var sequence bool
var verbose bool

func main() {
	// define flags
	// flag for prefix
	prefixPtr := flag.String("prefix", "", "a string")
	// flag for length
	lengthPtr := flag.Int("length", DEFAULT_LENGTH, "an int")
	// bool flag for sequence
	sequencePtr := flag.Bool("sequence", false, "a bool")
	verbosePtr := flag.Bool("v", false, "a bool")

	flag.Parse()

	prefix = *prefixPtr
	length = *lengthPtr
	sequence = *sequencePtr
	verbose = *verbosePtr

	if *verbosePtr {
		fmt.Println("prefix: ", prefix)
		fmt.Println("length: ", length)
		fmt.Println("sequence: ", sequence)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Renames a file to a random number with a given length and optional prefix")
		fmt.Println("Usage: gorename [-prefix <prefix>] [-length <length>] <pattern>")
		os.Exit(1)
	}

	pattern := flag.Arg(0)

	run(pattern)
}
func run(pattern string) {
	// Verify that the file exists
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

	if !sequence {
		err = renameFilesWithRandomNumbers(files)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = renameFilesInSequence(files)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Generate a random 10-digit number
}
func renameFilesWithRandomNumbers(files []string) error {

	fmt.Println("Renaming files...")

	for _, file := range files {
		ext := filepath.Ext(file)
		fmt.Println("renaming: ", file, " to")
		randomDigits := generateRandomDigits(length)
		newName := prefix + randomDigits + ext
		fmt.Println(newName)
		err := os.Rename(file, newName)
		if err != nil {
			return err
		}
	}
	return nil
}
func renameFilesInSequence(files []string) error {

	fmt.Println("Renaming files...")
	seqNum := 1
	for _, file := range files {
		ext := filepath.Ext(file)
		fmt.Println("renaming: ", file, " to")
		// generate a sequencial number with 3 Digits
		sequencialNumber := fmt.Sprintf("%03d", seqNum)
		newName := prefix + sequencialNumber + ext
		fmt.Println(newName)
		err := os.Rename(file, newName)
		if err != nil {
			return err
		}
		seqNum++
	}
	return nil
}

func generateRandomDigits(length int) string {
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
