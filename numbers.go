// Package funks implements basic functions that I use frequently to save time on replicating code across different projects.
package funks

// Digit defines arrays for numbers
type Digit [7]string

var (
	Zero = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
	}
	One = Digit{
		" 🮔🮔🮔  ",
		"🮔🮔🮔🮔  ",
		"  🮔🮔  ",
		"  🮔🮔  ",
		"  🮔🮔  ",
		"  🮔🮔  ",
		"🮔🮔🮔🮔🮔🮔",
	}
	Two = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"    🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔    ",
		"🮔🮔    ",
		"🮔🮔🮔🮔🮔🮔",
	}
	Three = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		" 🮔🮔🮔🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
	}
	Four = Digit{
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
	}
	Five = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔    ",
		"🮔🮔    ",
		"🮔🮔🮔🮔🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
	}
	Six = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔    ",
		"🮔🮔    ",
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
	}
	Seven = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
	}
	Eight = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
	}
	Nine = Digit{
		"🮔🮔🮔🮔🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔  🮔🮔",
		"🮔🮔🮔🮔🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
		"    🮔🮔",
	}
	Colon = Digit{
		"      ",
		"      ",
		"  🮔   ",
		"      ",
		"  🮔   ",
		"      ",
		"      ",
	}
	Colonempty = Digit{
		"      ",
		"      ",
		"  ▣   ",
		"      ",
		"  ▣   ",
		"      ",
		"      ",
	}
	S = Digit{
		"      ",
		"      ",
		"  🮔🮔🮔🮔",
		" 🮔    ",
		"  🮔🮔🮔 ",
		"     🮔",
		" 🮔🮔🮔🮔 ",
	}
	Blank = Digit{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
)

var Digits = [...]Digit{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine}
