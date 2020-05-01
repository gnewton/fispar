package lib

import (
	"crypto/sha256"
	"log"
	//"math"
	"testing"
)

func TestFileMain(t *testing.T) {

	fis := Fisplar{
		//Depth: math.MaxInt32,
		//Depth: math.MaxInt32,
		Depth: 12,
		Width: 3,
	}
	err := fis.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	h := sha256.New()

	//hs, err := filehash("filehash.go", h)
	files := []string{"filehash.go", "foo.txt", "foo3.txt", "tmp/foo.txt", "foo2.txt"}
	//files := []string{"foo.txt", "tmp/foo.txt", "foo.txt", "tmp/foo.txt"}
	for _, filename := range files {
		hs, err := filehash(filename, h)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(filename, hs)
		hs, err = filehash(filename, h)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(filename, hs)

		v, err := fis.Split(hs)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(v)
	}
}
