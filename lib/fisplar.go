package lib

import (
	"errors"
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
	runeLength             int
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

	//Contains utf8 runes?
	if utf8.RuneCountInString(str) != len(str) {
		return splitRunes(f, str)
	} else {
		// No UTF in string: byte=char
		return splitString(f, str)
	}
}

func splitString(f *Fisplar, s string) (string, error) {

	if f.length > len(s) {
		if f.ErrorOnTooShortStrings {
			return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(s)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
		}
	}

	var newString string

	var i int
	for i = 0; i < len(s); i++ {
		newString = newString + string(s[i])
		if (i+1)%f.Width == 0 && i < f.length {
			newString = newString + *f.Separator
		}
	}
	return newString, nil
}

func splitString_orig(f *Fisplar, s string) (string, error) {
	if f.length > len(s) {
		return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(s)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))

	}

	var newString string

	depthCount := 0
	var i int
	for i = 0; i < f.length; i += f.Width {
		newString = newString + s[i:i+f.Width]

		if depthCount == f.Depth {
			break
		}
		newString = newString + *f.Separator
		depthCount++

	}

	if len(s) > f.length {
		newString = newString + s[f.length:]
	}
	return newString, nil
}

func splitRunes(f *Fisplar, s string) (string, error) {

	r := []rune(s)
	if f.ErrorOnTooShortStrings && f.length > len(r) {
		return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(r)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
	}

	var newString string

	for i := 0; i < len(r); i++ {
		{
			newString = newString + string(r[i])
			if (i+1)%f.Width == 0 && i < f.length {
				newString = newString + *f.Separator
			}
		}
	}

	// if len(r) > f.length {
	// 	newString = newString + string(r[f.length:])
	// }

	//return strings.Join(ns, *f.Separator), nil
	return newString, nil
}

func splitRunes_orig(f *Fisplar, s string) (string, error) {

	r := []rune(s)
	if f.ErrorOnTooShortStrings && f.length > len(r) {
		return "", errors.New("String too short [" + s + "] len=" + strconv.Itoa(len(r)) + " depth=" + strconv.Itoa(f.Depth) + " width=" + strconv.Itoa(f.Width))
	}

	var newString string
	widthCounter := 0
	for i := 0; i < len(r); i++ {
		{
			if widthCounter == f.Width && i <= f.length {
				widthCounter = 0
				newString = newString + *f.Separator
			}
			newString = newString + string(r[i])
			widthCounter++
		}
	}

	//if len(r) > f.length {
	//ns = append(ns, s[f.length:])
	//}

	//return strings.Join(ns, *f.Separator), nil
	return newString, nil
}

type Mover struct {
}
