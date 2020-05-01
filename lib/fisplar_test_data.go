package lib

type FisplarTest struct {
	Fisplar
	name             string
	input            string
	output           string
	expectSplitError bool
	expectInitError  bool
}

var Tests []FisplarTest = []FisplarTest{
	//// Correct values
	FisplarTest{
		name: "Odd #chars; ascii",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Even #chars; ascii",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:            "12345678",
		output:           "12/34/5678",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Odd #chars; odd depth; ascii",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/56/7",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Even #chars; odd depth; ascii",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "12345678",
		output:           "12/34/56/78",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Odd #chars; utf8",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:            "1234€67",
		output:           "12/34/€67",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Even #chars; utf8",
		Fisplar: Fisplar{Depth: 2,
			Width: 2,
		},
		input:            "1234€678",
		output:           "12/34/€678",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Odd #chars; odd depth; utf8",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "1234€67",
		output:           "12/34/€6/7",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Even #chars; odd depth; utf8",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "1234€678",
		output:           "12/34/€6/78",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Support utf8, first rune",
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
		name: "Support utf8, last rune",
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
		name: "Support utf8, last rune, after separator",
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
		name: "Support utf8, last rune, after separator",
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
		name: "Support utf8, all runes",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "€€€€€€",
		output:           "€€/€€/€€/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "3x2; ascii",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "123456",
		output:           "12/34/56/",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Support utf8, almost runes",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "€1€2€3",
		output:           "€1/€2/€3/",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test single letter at end; ascii",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/56/7",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test single letter at end; utf8",
		Fisplar: Fisplar{Depth: 3,
			Width: 2,
		},
		input:            "123456€",
		output:           "12/34/56/€",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 2*4; ascii",
		Fisplar: Fisplar{Depth: 4,
			Width: 2,
		},
		input:            "12345678",
		output:           "12/34/56/78/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 2*4; utf8",
		Fisplar: Fisplar{Depth: 4,
			Width: 2,
		},
		input:            "1234€678",
		output:           "12/34/€6/78/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 4*2; ascii",
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
		name: "Test depth*width = length, 4*2; utf8",
		Fisplar: Fisplar{
			Depth: 2,
			Width: 4,
		},
		input:            "123456€8",
		output:           "1234/56€8/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 8*1; ascii",
		Fisplar: Fisplar{Depth: 8,
			Width: 1,
		},
		input:            "12345678",
		output:           "1/2/3/4/5/6/7/8/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 8*1; utf8",
		Fisplar: Fisplar{Depth: 8,
			Width: 1,
		},
		input:            "€2345678",
		output:           "€/2/3/4/5/6/7/8/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 7*1; ascii",
		Fisplar: Fisplar{Depth: 7,
			Width: 1,
		},
		input:            "12345678",
		output:           "1/2/3/4/5/6/7/8",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 7*1; utf8",
		Fisplar: Fisplar{Depth: 7,
			Width: 1,
		},
		input:            "1€345678",
		output:           "1/€/3/4/5/6/7/8",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width = length, 1*8; ascii",
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
		name: "Test depth*width = length, 1*8; utf8",
		Fisplar: Fisplar{
			Depth: 1,
			Width: 8,
		},
		input:            "123456€8",
		output:           "123456€8/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width, 9*8; ErrorOnTooShortStrings:false; ascii",
		Fisplar: Fisplar{
			Depth:                  9,
			Width:                  8,
			ErrorOnTooShortStrings: false,
		},
		input:            "12345678",
		output:           "12345678/",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test depth*width, 9*8; ErrorOnTooShortStrings:false; utf8",
		Fisplar: Fisplar{
			Depth:                  9,
			Width:                  8,
			ErrorOnTooShortStrings: false,
		},
		input:            "123456€8",
		output:           "123456€8/",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test depth*width, 8*9; ErrorOnTooShortStrings:false; ascii",
		Fisplar: Fisplar{
			Depth: 8,
			Width: 9,
		},
		input:            "12345678",
		output:           "12345678",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test depth*width, 8*9; ErrorOnTooShortStrings:false; utf8",
		Fisplar: Fisplar{
			Depth: 8,
			Width: 9,
		},
		input:            "12345€78",
		output:           "12345€78",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test every letter, full length; ascii",
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
		name: "Test every letter, full length; utf8",
		Fisplar: Fisplar{
			Depth: 1,
			Width: 8,
		},
		input:            "123456€8",
		output:           "123456€8/",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Test last letter, full length minus 1; ascii",
		Fisplar: Fisplar{Depth: 1,
			Width: 7,
		},
		input:            "12345678",
		output:           "1234567/8",
		expectSplitError: false,
		expectInitError:  false,
	},
	FisplarTest{
		name: "Test last letter, full length minus 1; uft8",
		Fisplar: Fisplar{Depth: 1,
			Width: 7,
		},
		input:            "1234567€",
		output:           "1234567/€",
		expectSplitError: false,
		expectInitError:  false,
	},

	//// Init Fails
	FisplarTest{
		name: "Error: Init Fail: depth is zero",
		Fisplar: Fisplar{Depth: 0,
			Width: 2,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  true,
	},

	FisplarTest{
		name: "Error: Init: width is zero",
		Fisplar: Fisplar{Depth: 2,
			Width: 0,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: false,
		expectInitError:  true,
	},

	FisplarTest{
		name: "Error: Init: Test depth*width, 9*8; ErrorOnTooShortStrings:false",
		Fisplar: Fisplar{
			Depth:                  9,
			Width:                  8,
			ErrorOnTooShortStrings: true,
		},
		input:  "12345678",
		output: "12345678",

		expectSplitError: true,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Error: Init: Test depth*width, 8*9; ErrorOnTooShortStrings:false",
		Fisplar: Fisplar{
			Depth:                  8,
			Width:                  9,
			ErrorOnTooShortStrings: true,
		},
		input:            "12345678",
		output:           "12345678",
		expectSplitError: true,
		expectInitError:  false,
	},

	//Split Fails
	FisplarTest{
		name: "Error: Split: incorrect result",
		Fisplar: Fisplar{
			Depth: 3,
			Width: 3,
		},
		input:            "1234567",
		output:           "12/34/567",
		expectSplitError: true,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Error: Split: Support utf8, single rune, too short string true",
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
		name: "Support utf8, single rune, too short string false",
		Fisplar: Fisplar{
			Depth:                  3,
			Width:                  2,
			ErrorOnTooShortStrings: false,
		},
		input:            "€",
		output:           "€",
		expectSplitError: false,
		expectInitError:  false,
	},

	FisplarTest{
		name: "Error: Split: string too short",
		Fisplar: Fisplar{
			Depth:                  3,
			Width:                  3,
			ErrorOnTooShortStrings: true,
		},
		input:            "a",
		output:           "a",
		expectSplitError: true,
		expectInitError:  false,
	},

	FisplarTest{
		name: "String too short but ErrorOnTooShortStrings: false",
		Fisplar: Fisplar{
			Depth:                  3,
			Width:                  3,
			ErrorOnTooShortStrings: false,
		},
		input:            "a",
		output:           "a",
		expectSplitError: false,
		expectInitError:  false,
	},
}
