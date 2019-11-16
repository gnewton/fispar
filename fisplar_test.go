package fisplar

import (
	"testing"
)

type FisplarTest struct {
	Fisplar
	name             string
	input            string
	output           string
	expectSplitError bool
	expectInitError  bool
}

var tests []FisplarTest = []FisplarTest{
	//// Correct values
	FisplarTest{
		name: "simple test",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "simple test, support utf8, first rune",
		Fisplar: Fisplar{
			Depth: 2,
			Width: 2,
		},
		input:            "€234567",
		output:           "€2/34/567",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "simple test, support utf8, last rune",
		Fisplar: Fisplar{
			Depth: 2,
			Width: 2,
		},
		input:            "123456€",
		output:           "12/34/56€",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "simple test, support utf8, last rune, after separator",
		Fisplar: Fisplar{
			Depth: 3,
			Width: 2,
		},
		input:            "123456€",
		output:           "12/34/56/€",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Support utf8, single rune",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "€",
		output:           "€",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test single letter at end",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/56/7",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 2*4",
		Fisplar: Fisplar{Depth: 4,
			Width: 2,
		},
		input:            "12345678",
		output:           "12/34/56/78/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 4*2",
		Fisplar: Fisplar{
			Depth: 2,
			Width: 4,
		},
		input:            "12345678",
		output:           "1234/5678/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 8*1",
		Fisplar: Fisplar{Depth: 8,
			Width: 1,
		},
		input:            "12345678",
		output:           "1/2/3/4/5/6/7/8/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 7*1",
		Fisplar: Fisplar{Depth: 7,
			Width: 1,
		},
		input:            "12345678",
		output:           "1/2/3/4/5/6/7/8",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 1*8",
		Fisplar: Fisplar{
			Depth: 1,
			Width: 8,
		},
		input:            "12345678",
		output:           "12345678/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width, 9*8",
		Fisplar: Fisplar{Depth: 9,
			Width: 8,
		},
		input:            "12345678",
		output:           "12345678",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test every letter, full length (no trailing /)",
		Fisplar: Fisplar{Depth: 1,
			Width: 8,
		},
		input:            "12345678",
		output:           "12345678/",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test last letter, full length minus 1",
		Fisplar: Fisplar{Depth: 1,
			Width: 7,
		},
		input:            "12345678",
		output:           "1234567/8",
		expectSplitError: false,
		expectInitError:  false,
	},

	//// Init Fails
	FisplarTest{
		name: "Init Fail: depth is zero",
		Fisplar: Fisplar{Depth: 0,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  true,
	},

	FisplarTest{
		name: "Init Fail: width is zero",
		Fisplar: Fisplar{Depth: 2,
			Width: 0,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  true,
	},

	//Split Fails
	FisplarTest{
		name: "Split Fail: incorrect result",
		Fisplar: Fisplar{Depth: 3,
			Width: 3,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: true,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Support utf8, single rune, too short string true",
		Fisplar: Fisplar{
			Depth:                  3,
			Width:                  2,
			ErrorOnTooShortStrings: true,
		},
		input:            "€",
		output:           "€",
		expectSplitError: true,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Split Fail: string too short",
		Fisplar: Fisplar{
			Depth:                  3,
			Width:                  3,
			ErrorOnTooShortStrings: true,
		},
		input:            "1",
		output:           "1",
		expectSplitError: true,
		expectInitError:  false,
	},
}

func TestMain(t *testing.T) {
	for _, ft := range tests {
		t.Run(ft.name, func(t *testing.T) {
			// init
			err := ft.Fisplar.Init()
			if err != nil && !ft.expectInitError {
				t.Fatal("Init() error: Expecting no error, go error. ", err, ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}
			if err == nil && ft.expectInitError {
				t.Fatal("Init() error: Expecting error, got no error", err, ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}
			// Expected Init error
			if err != nil && ft.expectInitError {
				return
			}

			// Split
			v, err := ft.Fisplar.Split(ft.input)
			//t.Log("Result: ", ft.Width, ft.Depth, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)

			// Expected Split error
			if err != nil && ft.expectSplitError {
				return
			}
			// Not expected Split error
			if err != nil {
				// Due to unknown internal error
				if !ft.expectSplitError {
					t.Fatal("Split: Not expected error, got error.", err, ft.Width, ft.Depth, v, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
				} else {
					return
				}
			}

			// Due to unknown error input != output
			if v != ft.output && !ft.expectSplitError {
				t.Fatalf("**Not expected error, got error [input != output].  width=%d depth=%d result=%s input=%s expectedResult=%s  expectSplitE=%t expecteInitE=%t", ft.Width, ft.Depth, v, ft.input, ft.output, ft.expectSplitError, ft.expectInitError)
			}

		})
	}
}
