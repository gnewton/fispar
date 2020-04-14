package fisplar

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	for _, ft := range Tests {

		t.Run(ft.name, func(t *testing.T) {
			// init
			err := ft.Fisplar.Init()

			// Got error and expecting one (i.e. OK)
			if err != nil && ft.expectInitError {
				return
			}

			// Got error but not expecting one
			if err != nil && !ft.expectInitError {
				expectLog(t, ft.expectSplitError || ft.expectInitError)
				t.Fatal("Init() error: Expecting no error, go error. ", err, ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}

			// Got no error but expecting one
			if err == nil && ft.expectInitError {
				t.Fatal("Init() error: Expecting error, got no error", err, ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}

			// Got no error and not expecting one: OK, continue to split

			// Split
			v, err := ft.Fisplar.Split(ft.input)
			//t.Log("Result: ", ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)

			// Got error and expected on (i.e. OK)
			if err != nil && ft.expectSplitError {
				return
			}
			// Not expected Split error
			if err != nil {
				// Due to unknown internal error
				if !ft.expectSplitError {
					t.Fatal("Split: Not expected error, got error.", err, ft.Width, ft.Depth, v, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
				}
			}

			// Due to unknown error input != output
			if v != ft.output && !ft.expectSplitError {
				t.Fatalf("**Not expected error, got error [input != output].  width=%d depth=%d result=%s input=%s expectedResult=%s  expectSplitE=%t expecteInitE=%t", ft.Width, ft.Depth, v, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}

		})
	}
}

func expectLog(t *testing.T, b bool) {
	if b {
		t.Log("Expect errors")
	} else {
		t.Log("Expect NO errors")
	}
}
