//
// Copyright (c) 2016 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
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

func Test_LightGray(t *testing.T) {
	testFunc(t, "LightGray", LightGray("x"), GrayFg)
	testFunc(t, "Complex LightGray", Inverse(Bold(BgLightGray(LightGray("x")))),
		GrayFg|GrayBg|BoldFm|InverseFm)
}

func Test_Gray(t *testing.T) {
	testFunc(t, "Gray", Gray("x"), BlackFg|BrightFm)
	testFunc(t, "BrightBlack", Bright(Black("x")), BlackFg|BrightFm)
	testFunc(t, "Complex Gray", Inverse(Bold(BgGray(Gray("x")))),
		BlackFg|BlackBg|BoldFm|InverseFm|BrightFm)
}

func Test_LightRed(t *testing.T) {
	testFunc(t, "LightRed", LightRed("x"), RedFg|BrightFm)
	testFunc(t, "Complex Red", Inverse(Bold(BgLightRed(LightRed("x")))),
		RedFg|RedBg|BoldFm|InverseFm|BrightFm)
}

func Test_LightGreen(t *testing.T) {
	testFunc(t, "LightGreen", LightGreen("x"), GreenFg|BrightFm)
	testFunc(t, "Complex LightGreen", Inverse(Bold(BgLightGreen(LightGreen("x")))),
		GreenFg|GreenBg|BoldFm|InverseFm|BrightFm)
}

func Test_Yellow(t *testing.T) {
	testFunc(t, "Yellow", Yellow("x"), BrownFg|BrightFm)
	testFunc(t, "Complex Yellow", Inverse(Bold(BgYellow(Yellow("x")))),
		BrownFg|BrownBg|BoldFm|InverseFm|BrightFm)
}

func Test_LightBlue(t *testing.T) {
	testFunc(t, "LightBlue", LightBlue("x"), BlueFg|BrightFm)
	testFunc(t, "Complex LightBlue", Inverse(Bold(BgLightBlue(LightBlue("x")))),
		BlueFg|BlueBg|BoldFm|InverseFm|BrightFm)
}

func Test_LightMagenta(t *testing.T) {
	testFunc(t, "LightMagenta", LightMagenta("x"), MagentaFg|BrightFm)
	testFunc(t, "Complex LightMagenta", Inverse(Bold(BgLightMagenta(LightMagenta("x")))),
		MagentaFg|MagentaBg|BoldFm|InverseFm|BrightFm)
}

func Test_LightCyan(t *testing.T) {
	testFunc(t, "LightCyan", LightCyan("x"), CyanFg|BrightFm)
	testFunc(t, "Complex LightCyan", Inverse(Bold(BgLightCyan(LightCyan("x")))),
		CyanFg|CyanBg|BoldFm|InverseFm|BrightFm)
}

func Test_White(t *testing.T) {
	testFunc(t, "White", White("x"), GrayFg|BrightFm)
	testFunc(t, "Complex White", Inverse(Bold(BgWhite(White("x")))),
		GrayFg|GrayBg|BoldFm|InverseFm|BrightFm)
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

func Test_BgLightGray(t *testing.T) {
	testFunc(t, "BgLightGray", BgLightGray("x"), GrayBg)
	testFunc(t, "Complex BgLightGray", LightGray(BgLightGray("x")), GrayBg|GrayFg)
}

func Test_BgGray(t *testing.T) {
	testFunc(t, "BgGray", BgGray("x"), BlackBg|BrightFm)
	testFunc(t, "Complex BgGray", Gray(BgGray("x")), BlackBg|BlackFg|BrightFm)
}

func Test_BgLightRed(t *testing.T) {
	testFunc(t, "BgLightRed", BgLightRed("x"), RedBg|BrightFm)
	testFunc(t, "Complex BgLightRed", LightRed(BgLightRed("x")), RedBg|RedFg|BrightFm)
}

func Test_BgLightGreen(t *testing.T) {
	testFunc(t, "BgLightGreen", BgLightGreen("x"), GreenBg|BrightFm)
	testFunc(t, "Complex BgLightGreen", LightGreen(BgLightGreen("x")), GreenBg|GreenFg|BrightFm)
}

func Test_BgYellow(t *testing.T) {
	testFunc(t, "BgYellow", BgYellow("x"), BrownBg|BrightFm)
	testFunc(t, "Complex BgYellow", Yellow(BgYellow("x")), BrownBg|BrownFg|BrightFm)
}

func Test_BgLightBlue(t *testing.T) {
	testFunc(t, "BgLightBlue", BgLightBlue("x"), BlueBg|BrightFm)
	testFunc(t, "Complex BgLightBlue", LightBlue(BgLightBlue("x")), BlueBg|BlueFg|BrightFm)
}

func Test_BgLightMagenta(t *testing.T) {
	testFunc(t, "BgLightMagenta", BgLightMagenta("x"), MagentaBg|BrightFm)
	testFunc(t, "Complex BgLightMagenta", LightMagenta(BgLightMagenta("x")), MagentaBg|MagentaFg|BrightFm)
}

func Test_BgLightCyan(t *testing.T) {
	testFunc(t, "BgLightCyan", BgLightCyan("x"), CyanBg|BrightFm)
	testFunc(t, "Complex BgLightCyan", LightCyan(BgLightCyan("x")), CyanBg|CyanFg|BrightFm)
}

func Test_BgWhite(t *testing.T) {
	testFunc(t, "BgWhite", BgWhite("x"), GrayBg|BrightFm)
	testFunc(t, "Complex BgWhite", White(BgWhite("x")), GrayBg|GrayFg|BrightFm)
}

func Test_Bold(t *testing.T) {
	testFunc(t, "Bold", Bold("x"), BoldFm)
}

func Test_Inverse(t *testing.T) {
	testFunc(t, "Inverse", Inverse("x"), InverseFm)
}

func Test_Bright(t *testing.T) {
	testFunc(t, "Bright", Bright("x"), BrightFm)
}
