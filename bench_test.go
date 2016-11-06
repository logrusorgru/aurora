//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE.md file for more details or see below.
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
	"testing"
)

var (
	gVal Value
	gStr string
)

func Benchmark_auroraClear(b *testing.B) {
	a := NewAurora(false)
	x := "x"
	for i := 0; i < b.N; i++ {
		gVal = a.Black(x)
		gVal = a.Red(x)
		gVal = a.Green(x)
		gVal = a.Brown(x)
		gVal = a.Blue(x)
		gVal = a.Magenta(x)
		gVal = a.Cyan(x)
		gVal = a.Gray(x)
		gVal = a.BgBlack(x)
		gVal = a.BgRed(x)
		gVal = a.BgGreen(x)
		gVal = a.BgBrown(x)
		gVal = a.BgBlue(x)
		gVal = a.BgMagenta(x)
		gVal = a.BgCyan(x)
		gVal = a.BgGray(x)
		gVal = a.Bold(x)
		gVal = a.Inverse(x)
	}
}

func Benchmark_auroraClear_sprintf(b *testing.B) {
	a := NewAurora(false)
	for i := 0; i < b.N; i++ {
		gStr = a.Sprintf(noString("%f"), 3.14)
		gStr = a.Sprintf("%f", 3.14)
		gStr = a.Sprintf("%f", a.Red(3.14))
		gStr = a.Sprintf(Blue("%f"), a.Red(3.14))
	}
}

func Benchmark_aurora(b *testing.B) {
	a := NewAurora(true)
	x := "x"
	for i := 0; i < b.N; i++ {
		gVal = a.Black(x)
		gVal = a.Red(x)
		gVal = a.Green(x)
		gVal = a.Brown(x)
		gVal = a.Blue(x)
		gVal = a.Magenta(x)
		gVal = a.Cyan(x)
		gVal = a.Gray(x)
		gVal = a.BgBlack(x)
		gVal = a.BgRed(x)
		gVal = a.BgGreen(x)
		gVal = a.BgBrown(x)
		gVal = a.BgBlue(x)
		gVal = a.BgMagenta(x)
		gVal = a.BgCyan(x)
		gVal = a.BgGray(x)
		gVal = a.Bold(x)
		gVal = a.Inverse(x)
	}
}

func Benchmark_aurora_sprintf(b *testing.B) {
	a := NewAurora(true)
	for i := 0; i < b.N; i++ {
		gStr = a.Sprintf(noString("%f"), 3.14)
		gStr = a.Sprintf("%f", 3.14)
		gStr = a.Sprintf("%f", a.Red(3.14))
		gStr = a.Sprintf(Blue("%f"), a.Red(3.14))
	}
}
