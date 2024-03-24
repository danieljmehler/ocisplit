package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func SplitOCIDir(path string, dest string) bool {
	return true
}

func SplitOCIArchive(path string, dest string) bool {
	return true
}

func OCISplit(path string, dest string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if stat.IsDir() {
		return SplitOCIDir(path, dest)
	} else {
		return SplitOCIArchive(path, dest)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("ocisplit: ")

	var path string
	var dest string
	flag.StringVar(&path, "path", ".", "The path to the OCI arhive or decompressed OCI archive directory.")
	flag.StringVar(&dest, "dest", ".", "The path to the directory to which individual OCI archives will be saved.")
	flag.Parse()

	// Clean input parameters
	path = strings.TrimSuffix(path, "/")
	dest = strings.TrimSuffix(dest, "/")

	OCISplit(path, dest)
}
