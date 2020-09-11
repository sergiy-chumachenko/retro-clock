package main

type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	sep = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	digits = [10]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}
)


