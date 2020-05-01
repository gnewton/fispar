package lib

import (
	//"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"strings"
)

// Hashes [filename + contents], returns string representation of hash
func filehash(filename string, h hash.Hash) (string, error) {
	h.Reset()
	log.Println(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// filename + contents
	readers := []io.Reader{strings.NewReader(fstat.Name()), f}

	for i, _ := range readers {
		r := readers[i]
		if _, err := io.Copy(h, r); err != nil {
			log.Println(err)
			return "", err
		}
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
