package parse

import (
	"log"
	"os"
	"path/filepath"
)

// Reads a markdown file from the given path and returns its raw bytes and filename.
func GetContent(path string) ([]byte, string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot read markdown: %v", err)
	}

	filename := filepath.Base(path)

	return content, filename
}