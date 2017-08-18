//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
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

func testFunc(t *testing.T, name string, v Value, clr Color) {
	if str, ok := v.Value().(string); !ok {
		t.Errorf("%s wrong value type: %t, expected string", name, v.Value())
	} else if str != "x" {
		t.Errorf("%s wrong value: '%v', expected 'x'", name, v.Value())
	}
	if !isTail(v, 0) {
		t.Errorf("%s wrong tail: %d", name, v.tail())
	}
	if !isColor(v, clr) {
		t.Errorf("%s wrong color: %d, want: %d", name, v.Color(), clr)
	}
}

func Test_Colorize(t *testing.T) {
	testFunc(t, "Colorize", Colorize("x", BlackFg|BoldFm), BlackFg|BoldFm)
	testFunc(t, "Complex Colorize", Colorize(BgRed("x"), BlackFg|BoldFm),
		BlackFg|BoldFm)
}

func Test_Black(t *testing.T) {
	testFunc(t, "Black", Black("x"), BlackFg)
	testFunc(t, "Complex Black", Inverse(Bold(BgBlack(Black("x")))),
		BlackFg|BlackBg|BoldFm|InverseFm)
}

func Test_Red(t *testing.T) {
	testFunc(t, "Red", Red("x"), RedFg)
	testFunc(t, "Complex Red", Inverse(Bold(BgRed(Red("x")))),
		RedFg|RedBg|BoldFm|InverseFm)
}

func Test_Green(t *testing.T) {
	testFunc(t, "Green", Green("x"), GreenFg)
	testFunc(t, "Complex Green", Inverse(Bold(BgGreen(Green("x")))),
		GreenFg|GreenBg|BoldFm|InverseFm)
}

func Test_Brown(t *testing.T) {
	testFunc(t, "Brown", Brown("x"), BrownFg)
	testFunc(t, "Complex Brown", Inverse(Bold(BgBrown(Brown("x")))),
		BrownFg|BrownBg|BoldFm|InverseFm)
}

func Test_Blue(t *testing.T) {
	testFunc(t, "Blue", Blue("x"), BlueFg)
	testFunc(t, "Complex Blue", Inverse(Bold(BgBlue(Blue("x")))),
		BlueFg|BlueBg|BoldFm|InverseFm)
}

func Test_Magenta(t *testing.T) {
	testFunc(t, "Magenta", Magenta("x"), MagentaFg)
	testFunc(t, "Complex Magenta", Inverse(Bold(BgMagenta(Magenta("x")))),
		MagentaFg|MagentaBg|BoldFm|InverseFm)
}

func Test_Cyan(t *testing.T) {
	testFunc(t, "Cyan", Cyan("x"), CyanFg)
	testFunc(t, "Complex Cyan", Inverse(Bold(BgCyan(Cyan("x")))),
		CyanFg|CyanBg|BoldFm|InverseFm)
}

func Test_Gray(t *testing.T) {
	testFunc(t, "Gray", Gray("x"), GrayFg)
	testFunc(t, "Complex Gray", Inverse(Bold(BgGray(Gray("x")))),
		GrayFg|GrayBg|BoldFm|InverseFm)
}

func Test_BgBlack(t *testing.T) {
	testFunc(t, "BgBlack", BgBlack("x"), BlackBg)
	testFunc(t, "Complex BgBlack", Black(BgBlack("x")), BlackBg|BlackFg)
}

func Test_BgRed(t *testing.T) {
	testFunc(t, "BgRed", BgRed("x"), RedBg)
	testFunc(t, "Complex BgRed", Red(BgRed("x")), RedBg|RedFg)
}

func Test_BgGreen(t *testing.T) {
	testFunc(t, "BgGreen", BgGreen("x"), GreenBg)
	testFunc(t, "Complex BgGreen", Green(BgGreen("x")), GreenBg|GreenFg)
}

func Test_BgBrown(t *testing.T) {
	testFunc(t, "BgBrown", BgBrown("x"), BrownBg)
	testFunc(t, "Complex BgBrown", Brown(BgBrown("x")), BrownBg|BrownFg)
}

func Test_BgBlue(t *testing.T) {
	testFunc(t, "BgBlue", BgBlue("x"), BlueBg)
	testFunc(t, "Complex BgBlue", Blue(BgBlue("x")), BlueBg|BlueFg)
}

func Test_BgMagenta(t *testing.T) {
	testFunc(t, "BgMagenta", BgMagenta("x"), MagentaBg)
	testFunc(t, "Complex BgMagenta", Magenta(BgMagenta("x")), MagentaBg|MagentaFg)
}

func Test_BgCyan(t *testing.T) {
	testFunc(t, "BgCyan", BgCyan("x"), CyanBg)
	testFunc(t, "Complex BgCyan", Cyan(BgCyan("x")), CyanBg|CyanFg)
}

func Test_BgGray(t *testing.T) {
	testFunc(t, "BgGray", BgGray("x"), GrayBg)
	testFunc(t, "Complex BgGray", Gray(BgGray("x")), GrayBg|GrayFg)
}

func Test_Bold(t *testing.T) {
	testFunc(t, "Bold", Bold("x"), BoldFm)
}

func Test_Inverse(t *testing.T) {
	testFunc(t, "Inverse", Inverse("x"), InverseFm)
}
