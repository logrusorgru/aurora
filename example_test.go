//
// Copyright (c) 2016-2019 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Unlicense. See LICENSE
// file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package aurora

import (
	"fmt"
)

func ExampleRed() {
	fmt.Println("value exceeds min-threshold:", Red(3.14))

	// Output: value exceeds min-threshold: [31m3.14[0m
}

func ExampleBold() {
	fmt.Println("value:", Bold(Green(99)))

	// Output: value: [1;32m99[0m
}

func ExampleNewAurora_no_colors() {
	a := NewAurora(false)
	fmt.Println(a.Red("Not red"))

	// Output: Not red
}

func ExampleNewAurora_colors() {
	a := NewAurora(true)
	fmt.Println(a.Red("Red"))

	// Output: [31mRed[0m
}

func Example_printf() {
	fmt.Printf("%d %s", Blue(100), BgBlue("cats"))

	// Output: [34m100[0m [44mcats[0m
}

func ExampleSprintf() {
	fmt.Print(
		Sprintf(
			Blue("we've got %d cats, but want %d"), // <- blue format
			Cyan(5),
			Bold(Magenta(25)),
		),
	)

	// Output: [34mwe've got [0;36m5[0;34m cats, but want [0;1;35m25[0;34m[0m
}
