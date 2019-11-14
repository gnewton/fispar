package fisplar

import (
	"encoding/hex"
	"errors"
	"hash"
	"os"
	"strconv"
	//"strings"
	"sync"
)

const DefaultSeparator = string(os.PathSeparator)

type Fisplar struct {
	Depth               int
	Width               int
	length              int
	Separator           *string
	Hash                func() hash.Hash
	pool                *sync.Pool
	inited              bool
	errorOnShortStrings bool
}

func (f *Fisplar) Init() error {
	if f == nil {
		return errors.New("fisplay is nil")
	}
	if f.Depth < 1 {
		return errors.New("Depth must be > 0")
	}
	if f.Width < 1 {
		return errors.New("Width must be > 0")
	}
	f.length = f.Width * f.Depth

	if f.Separator == nil {
		f.Separator = new(string)
		*(f.Separator) = DefaultSeparator
	}

	if f.Hash != nil {
		f.initPool()
	}
	f.inited = true

	return nil
}

func (f *Fisplar) Split(s string) (string, error) {
	if !f.inited {
		return "", errors.New("fisplar not inited")
	}

	if len(s) == 0 {
		return "", errors.New("String is empty")
	}
	if f.Hash != nil {
		s = f.hashString(s)
	}

	r := []rune(s)
	if f.errorOnShortStrings && f.length > len(r) {
		return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(s)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
	}

	//var ns []string

	var z string
	n := 0
	for i := 0; i < len(r); i++ {
		{
			if n == f.Width && i <= f.length {
				n = 0
				z = z + *f.Separator
			}
			z = z + string(r[i])
			n++
		}
	}

	//if len(r) > f.length {
	//ns = append(ns, s[f.length:])
	//}

	//return strings.Join(ns, *f.Separator), nil
	return z, nil
}

func (f *Fisplar) hashString(s string) string {
	var h hash.Hash
	h = f.pool.Get().(hash.Hash)
	h.Reset()

	h.Write([]byte(s))
	hexString := hex.EncodeToString(h.Sum(nil))

	f.pool.Put(h)

	return hexString
}

func (f *Fisplar) initPool() {
	f.pool = &(sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			//hash := sha512.New
			return f.Hash()
		},
	})

}
