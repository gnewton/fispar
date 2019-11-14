package fisplar

import (
	"testing"
)

type FisplarTest struct {
	Fisplar
	name        string
	input       string
	output      string
	isError     bool
	isInitError bool
}

var tests []FisplarTest = []FisplarTest{
	//// Correct values
	FisplarTest{
		name: "simple test",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:       "1234567",
		output:      "12/34/567",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "simple test, support utf8",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:       "€234567",
		output:      "€2/34/567",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "Test single letter at end",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:       "1234567",
		output:      "12/34/56/7",
		isError:     false,
		isInitError: false,
	},
	//// Correct values
	FisplarTest{
		name: "Handle UTF8 runes",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:       "M1234567",
		output:      "M1/23/4567",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "Test full string length",
		Fisplar: Fisplar{Depth: 4,
			Width: 2,
		},
		input:       "12345678",
		output:      "12/34/56/78",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "Test every letter, full length",
		Fisplar: Fisplar{Depth: 8,
			Width: 1,
		},
		input:       "12345678",
		output:      "1/2/3/4/5/6/7/8",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "Test every letter, full length (no trailing /)",
		Fisplar: Fisplar{Depth: 1,
			Width: 8,
		},
		input:       "12345678",
		output:      "12345678",
		isError:     false,
		isInitError: false,
	},

	FisplarTest{
		name: "Test every letter, full length minus 1",
		Fisplar: Fisplar{Depth: 1,
			Width: 7,
		},
		input:       "12345678",
		output:      "1234567/8",
		isError:     false,
		isInitError: false,
	},

	//// Init Fails
	FisplarTest{
		name: "Init Fail: depth, width are zero",
		Fisplar: Fisplar{Depth: 0,
			Width: 0,
		},
		input:       "1234567",
		output:      "12/34/567",
		isError:     false,
		isInitError: true,
	},

	//// Split Fails
	FisplarTest{
		name: "Split Fail: incorrect result",
		Fisplar: Fisplar{Depth: 3,
			Width: 3,
		},
		input:       "1234567",
		output:      "12/34/567",
		isError:     true,
		isInitError: false,
	},
}

func TestMain(t *testing.T) {
	for _, ft := range tests {

		t.Run(ft.name, func(t *testing.T) {
			// init
			err := ft.Fisplar.Init()
			if (err != nil && !ft.isInitError) || (err == nil && ft.isInitError) {
				t.Fatal(err, ft.Width, ft.Depth, ft.input, ft.output, ft.isError, ft.isInitError)
			}
			// Expected Init error
			if err != nil && ft.isInitError {
				return
			}

			// Split
			v, err := ft.Fisplar.Split(ft.input)

			// Expected Split error
			if err != nil && ft.isError {
				return
			}
			// Not expected Split error
			if err != nil && !ft.isError {

				t.Fatal(err, ft.Width, ft.Depth, v, ft.input, ft.output, ft.isError, ft.isInitError)
			}

			// Not expected value error
			if v != ft.output {
				t.Fatal("input != output", ft.Width, ft.Depth, v, ft.input, ft.output, ft.isError, ft.isInitError)
			}

		})

	}
}
