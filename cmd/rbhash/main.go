package main

import (
	"flag"
	"fmt"
	"os"
	"rapidbinary/internal/hasher"
	"rapidbinary/internal/shared"
)

func main() {
	flag.Usage = func() {
		shared.PrintHeader()
		fmt.Println("Usage: rbhash [file_path]")
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	filePath := args[0]
	shared.PrintHeader()

	fmt.Printf("Hashing: %s\n", filePath)

	hash, err := hasher.CalculateSHA256(filePath)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n%sSHA256:%s %s\n", shared.Bold, shared.Reset, hash)
}
