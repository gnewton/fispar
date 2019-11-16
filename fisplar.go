package fisplar

import (
	"errors"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

const DefaultSeparator = string(os.PathSeparator)

type Fisplar struct {
	Separator              *string
	Depth                  int
	Width                  int
	ErrorOnTooShortStrings bool
	length                 int
	inited                 bool
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

	f.inited = true

	return nil
}

func (f *Fisplar) Split(str string) (string, error) {
	if !f.inited {
		return "", errors.New("fisplar not inited")
	}

	if len(str) == 0 {
		return "", errors.New("String is empty")
	}

	if utf8.RuneCountInString(str) != len(str) {
		log.Println("==========")
		return splitRunes(f, str)
	}

	return splitString(f, str)
}

// TODO
// non-runes (i.e. string) implementation faster and less allocs.
//
func splitString(f *Fisplar, s string) (string, error) {
	if f.length > len(s) {
		if !f.ErrorOnTooShortStrings {
			return s, nil
		} else {
			return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(s)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
		}
	}

	var out string

	depthCount := 0
	var i int
	for i = 0; i < f.length; i += f.Width {
		//log.Println(s, i, f.Width, out)
		out = out + s[i:i+f.Width]

		if depthCount == f.Depth {
			break
		}
		out = out + *f.Separator
		depthCount++

	}

	if len(s) > f.length {
		out = out + s[f.length:]
	}

	//return splitRunes(f, s)
	return out, nil
}

func splitRunes(f *Fisplar, s string) (string, error) {

	r := []rune(s)
	if f.ErrorOnTooShortStrings && f.length > len(r) {
		return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(r)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
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
