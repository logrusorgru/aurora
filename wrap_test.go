//
// Copyright (c) 2016-2020 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Unlicense. See LICENSE
// file for more details or see below.
//

//
// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>
//

package aurora

import (
	"testing"
)

func testFunc(t *testing.T, name string, v Value, clr Color) {
	t.Helper()
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

func Test_Reset(t *testing.T) {
	testFunc(t, "Reset", Reset("x"), 0)
	testFunc(t, "Complex Reset", Reset(DoublyUnderline(Underline("x"))),
		0)
}

func Test_Bold(t *testing.T) {
	testFunc(t, "Bold", Bold("x"), BoldFm)
	testFunc(t, "Complex Bold", Bold(Italic(Red("x"))),
		ItalicFm|RedFg|BoldFm)
}

func Test_Faint(t *testing.T) {
	testFunc(t, "Faint", Faint("x"), FaintFm)
	testFunc(t, "Complex Faint", Faint(BrightBlack("x").BgBrightGreen()),
		BrightFg|BlackFg|FaintFm|BrightBg|GreenBg)
}

func Test_DoublyUnderline(t *testing.T) {
	testFunc(t, "DoublyUnderline", DoublyUnderline("x"), DoublyUnderlineFm)
	testFunc(t, "Complex DoublyUnderline", DoublyUnderline(BrightBlue("x")),
		BrightFg|BlueFg|DoublyUnderlineFm)
}

func Test_Fraktur(t *testing.T) {
	testFunc(t, "Fraktur", Fraktur("x"), FrakturFm)
	testFunc(t, "Complex Fraktur", Fraktur(Faint("x").BgBrightCyan()),
		FaintFm|FrakturFm|BrightBg|CyanBg)
}

func Test_Italic(t *testing.T) {
	testFunc(t, "Italic", Italic("x"), ItalicFm)
	testFunc(t, "Complex Italic", Italic(BgBrightBlack("x")),
		ItalicFm|BrightBg|BlackBg)
}

func Test_Underline(t *testing.T) {
	testFunc(t, "Underline", Underline("x"), UnderlineFm)
	testFunc(t, "Complex Underline", Underline(Faint("x").Bold()),
		UnderlineFm|BoldFm)
}

func Test_SlowBlink(t *testing.T) {
	testFunc(t, "SlowBlink", SlowBlink("x"), SlowBlinkFm)
	testFunc(t, "Complex SlowBlink", SlowBlink(Reverse("x").Bold().Faint()),
		SlowBlinkFm|ReverseFm|FaintFm)
}

func Test_RapidBlink(t *testing.T) {
	testFunc(t, "RapidBlink", RapidBlink("x"), RapidBlinkFm)
	testFunc(t, "Complex RapidBlink", RapidBlink(Italic(Bold("x"))),
		RapidBlinkFm|ItalicFm|BoldFm)
}

func Test_Blink(t *testing.T) {
	testFunc(t, "Blink", Blink("x"), BlinkFm)
	testFunc(t, "Complex Blink", Blink(Reverse("x").Italic()),
		BlinkFm|ReverseFm|ItalicFm)
}

func Test_Reverse(t *testing.T) {
	testFunc(t, "Reverse", Reverse("x"), ReverseFm)
	testFunc(t, "Complex Reverse", Reverse(Italic("x").Reverse()),
		ReverseFm|ItalicFm|ReverseFm)
}

func Test_Inverse(t *testing.T) {
	testFunc(t, "Inverse", Inverse("x"), InverseFm)
	testFunc(t, "Complex Inverse", Inverse(Italic("x").Overlined()),
		InverseFm|ItalicFm|OverlinedFm)
}

func Test_Conceal(t *testing.T) {
	testFunc(t, "Conceal", Conceal("x"), ConcealFm)
	testFunc(t, "Complex Conceal", Conceal(Underline("x").Italic()),
		ConcealFm|UnderlineFm|ItalicFm)
}

func Test_Hidden(t *testing.T) {
	testFunc(t, "Hidden", Hidden("x"), HiddenFm)
	testFunc(t, "Complex Hidden", Hidden(Overlined("x")),
		HiddenFm|OverlinedFm)
}

func Test_CrossedOut(t *testing.T) {
	testFunc(t, "CrossedOut", CrossedOut("x"), CrossedOutFm)
	testFunc(t, "Complex CrossedOut", CrossedOut(BgRed("x")),
		CrossedOutFm|RedBg)
}

func Test_StrikeThrough(t *testing.T) {
	testFunc(t, "StrikeThrough", StrikeThrough("x"), StrikeThroughFm)
	testFunc(t, "Complex StrikeThrough", StrikeThrough(BgBrightCyan("x")),
		StrikeThroughFm|BrightBg|CyanBg)
}

func Test_Framed(t *testing.T) {
	testFunc(t, "Framed", Framed("x"), FramedFm)
	testFunc(t, "Complex Framed", Framed(Overlined("x").Underline()),
		FramedFm|OverlinedFm|UnderlineFm)
}

func Test_Encircled(t *testing.T) {
	testFunc(t, "Encircled", Encircled("x"), EncircledFm)
	testFunc(t, "Complex Encircled", Encircled(Italic("x").Overlined()),
		EncircledFm|ItalicFm|OverlinedFm)
}

func Test_Overlined(t *testing.T) {
	testFunc(t, "Overlined", Overlined("x"), OverlinedFm)
	testFunc(t, "Complex Overlined", Overlined(BgBrightWhite("x").Reverse()),
		OverlinedFm|BrightBg|WhiteBg|ReverseFm)
}

func Test_Black(t *testing.T) {
	testFunc(t, "Black", Black("x"), BlackFg)
	testFunc(t, "Complex Black", Black(Overlined("x").BgRed()),
		BlackFg|OverlinedFm|RedBg)
}

func Test_Red(t *testing.T) {
	testFunc(t, "Red", Red("x"), RedFg)
	testFunc(t, "Complex Red", Red(Inverse("x").DoublyUnderline()),
		RedFg|InverseFm|DoublyUnderlineFm)
}

func Test_Green(t *testing.T) {
	testFunc(t, "Green", Green("x"), GreenFg)
	testFunc(t, "Complex Green", Green(BgMagenta("x").Inverse().Reverse()),
		GreenFg|MagentaBg|InverseFm|ReverseFm)
}

func Test_Yellow(t *testing.T) {
	testFunc(t, "Yellow", Yellow("x"), YellowFg)
	testFunc(t, "Complex Yellow", Yellow(Bold("x").BgBrightMagenta()),
		YellowFg|BoldFm|BrightBg|MagentaBg)
}

func Test_Brown(t *testing.T) {
	testFunc(t, "Brown", Brown("x"), BrownFg)
	testFunc(t, "Complex Brown", Brown(BgBrightBlue("x").Bold()),
		BrownFg|BrightBg|BlueBg|BoldFm)
}

func Test_Blue(t *testing.T) {
	testFunc(t, "Blue", Blue("x"), BlueFg)
	testFunc(t, "Complex Blue", Blue(Fraktur("x").Underline()),
		BlueFg|FrakturFm|UnderlineFm)
}

func Test_Magenta(t *testing.T) {
	testFunc(t, "Magenta", Magenta("x"), MagentaFg)
	testFunc(t, "Complex Magenta", Magenta(DoublyUnderline("x").Fraktur()),
		MagentaFg|DoublyUnderlineFm|FrakturFm)
}

func Test_Cyan(t *testing.T) {
	testFunc(t, "Cyan", Cyan("x"), CyanFg)
	testFunc(t, "Complex Cyan", Cyan(StrikeThrough("x")),
		CyanFg|StrikeThroughFm)
}

func Test_White(t *testing.T) {
	testFunc(t, "White", White("x"), WhiteFg)
	testFunc(t, "Complex White", White(Inverse("x").Blink()),
		WhiteFg|InverseFm|BlinkFm)
}

func Test_BrightBlack(t *testing.T) {
	testFunc(t, "BrightBlack", BrightBlack("x"), BrightFg|BlackFg)
	testFunc(t, "Complex BrightBlack", BrightBlack(BgBrightGreen("x")),
		BrightFg|BlackFg|BrightBg|GreenBg)
}

func Test_BrightRed(t *testing.T) {
	testFunc(t, "BrightRed", BrightRed("x"), BrightFg|RedFg)
	testFunc(t, "Complex BrightRed", BrightRed(Bold("x").BgYellow()),
		BrightFg|RedFg|BoldFm|YellowBg)
}

func Test_BrightGreen(t *testing.T) {
	testFunc(t, "BrightGreen", BrightGreen("x"), BrightFg|GreenFg)
	testFunc(t, "Complex BrightGreen", BrightGreen(Faint("x")),
		BrightFg|GreenFg|FaintFm)
}

func Test_BrightYellow(t *testing.T) {
	testFunc(t, "BrightYellow", BrightYellow("x"), BrightFg|YellowFg)
	testFunc(t, "Complex BrightYellow", BrightYellow(BgRed("x")),
		BrightFg|YellowFg|RedBg)
}

func Test_BrightBlue(t *testing.T) {
	testFunc(t, "BrightBlue", BrightBlue("x"), BrightFg|BlueFg)
	testFunc(t, "Complex BrightBlue", BrightBlue(DoublyUnderline("x")),
		BrightFg|BlueFg|DoublyUnderlineFm)
}

func Test_BrightMagenta(t *testing.T) {
	testFunc(t, "BrightMagenta", BrightMagenta("x"), BrightFg|MagentaFg)
	testFunc(t, "Complex BrightMagenta", BrightMagenta(Underline("x")),
		BrightFg|MagentaFg|UnderlineFm)
}

func Test_BrightCyan(t *testing.T) {
	testFunc(t, "BrightCyan", BrightCyan("x"), BrightFg|CyanFg)
	testFunc(t, "Complex BrightCyan", BrightCyan(BgGreen("x").Italic()),
		BrightFg|CyanFg|GreenBg|ItalicFm)
}

func Test_BrightWhite(t *testing.T) {
	testFunc(t, "BrightWhite", BrightWhite("x"), BrightFg|WhiteFg)
	testFunc(t, "Complex BrightWhite", BrightWhite(Inverse("x").BgRed()),
		BrightFg|WhiteFg|InverseFm|RedBg)
}

func Test_Index(t *testing.T) {
	testFunc(t, "Index", Index(178, "x"), (Color(178)<<shiftFg)|flagFg)
	testFunc(t, "Complex Index", Index(178, Inverse("x")),
		(Color(178)<<shiftFg|flagFg)|InverseFm)
}

func Test_Gray(t *testing.T) {
	testFunc(t, "Gray", Gray(14, "x"), Color(232+14)<<shiftFg|flagFg)
	testFunc(t, "Complex Gray", Gray(14, Index(19, "x").Bold()),
		flagFg|Color(14+232)<<shiftFg|BoldFm)
}

func Test_BgBlack(t *testing.T) {
	testFunc(t, "BgBlack", BgBlack("x"), BlackBg)
	testFunc(t, "Complex BgBlack", BgBlack(BgGray(15, "x")),
		BlackBg)
}

func Test_BgRed(t *testing.T) {
	testFunc(t, "BgRed", BgRed("x"), RedBg)
	testFunc(t, "Complex BgRed", BgRed(BrightBlack("x")),
		RedBg|BrightFg|BlackFg)
}

func Test_BgGreen(t *testing.T) {
	testFunc(t, "BgGreen", BgGreen("x"), GreenBg)
	testFunc(t, "Complex BgGreen", BgGreen(BrightGreen("x")),
		GreenBg|BrightFg|GreenFg)
}

func Test_BgYellow(t *testing.T) {
	testFunc(t, "BgYellow", BgYellow("x"), YellowBg)
	testFunc(t, "Complex BgYellow", BgYellow(Blink("x")),
		YellowBg|BlinkFm)
}

func Test_BgBrown(t *testing.T) {
	testFunc(t, "BgBrown", BgBrown("x"), BrownBg)
	testFunc(t, "Complex BgBrown", BgBrown(Hidden("x")),
		BrownBg|HiddenFm)
}

func Test_BgBlue(t *testing.T) {
	testFunc(t, "BgBlue", BgBlue("x"), BlueBg)
	testFunc(t, "Complex BgBlue", BgBlue(Framed("x")),
		BlueBg|FramedFm)
}

func Test_BgMagenta(t *testing.T) {
	testFunc(t, "BgMagenta", BgMagenta("x"), MagentaBg)
	testFunc(t, "Complex BgMagenta", BgMagenta(Encircled("x")),
		MagentaBg|EncircledFm)
}

func Test_BgCyan(t *testing.T) {
	testFunc(t, "BgCyan", BgCyan("x"), CyanBg)
	testFunc(t, "Complex BgCyan", BgCyan(StrikeThrough("x")),
		CyanBg|StrikeThroughFm)
}

func Test_BgWhite(t *testing.T) {
	testFunc(t, "BgWhite", BgWhite("x"), WhiteBg)
	testFunc(t, "Complex BgWhite", BgWhite(Inverse("x")),
		WhiteBg|InverseFm)
}

func Test_BgBrightBlack(t *testing.T) {
	testFunc(t, "BgBrightBlack", BgBrightBlack("x"), BrightBg|BlackBg)
	testFunc(t, "Complex BgBrightBlack", BgBrightBlack(BrightWhite("x")),
		BrightBg|BlackBg|BrightFg|WhiteFg)
}

func Test_BgBrightRed(t *testing.T) {
	testFunc(t, "BgBrightRed", BgBrightRed("x"), BrightBg|RedBg)
	testFunc(t, "Complex BgBrightRed", BgBrightRed(White("x")),
		BrightBg|RedBg|WhiteFg)
}

func Test_BgBrightGreen(t *testing.T) {
	testFunc(t, "BgBrightGreen", BgBrightGreen("x"), BrightBg|GreenBg)
	testFunc(t, "Complex BgBrightGreen", BgBrightGreen(BrightCyan("x")),
		BrightBg|GreenBg|BrightFg|CyanFg)
}

func Test_BgBrightYellow(t *testing.T) {
	testFunc(t, "BgBrightYellow", BgBrightYellow("x"), BrightBg|YellowBg)
	testFunc(t, "Complex BgBrightYellow", BgBrightYellow(Cyan("x")),
		BrightBg|YellowBg|CyanFg)
}

func Test_BgBrightBlue(t *testing.T) {
	testFunc(t, "BgBrightBlue", BgBrightBlue("x"), BrightBg|BlueBg)
	testFunc(t, "Complex BgBrightBlue", BgBrightBlue(Green("x")),
		BrightBg|BlueBg|GreenFg)
}

func Test_BgBrightMagenta(t *testing.T) {
	testFunc(t, "BgBrightMagenta", BgBrightMagenta("x"), BrightBg|MagentaBg)
	testFunc(t, "Complex BgBrightMagenta", BgBrightMagenta(Black("x")),
		BrightBg|MagentaBg|BlackFg)
}

func Test_BgBrightCyan(t *testing.T) {
	testFunc(t, "BgBrightCyan", BgBrightCyan("x"), BrightBg|CyanBg)
	testFunc(t, "Complex BgBrightCyan", BgBrightCyan(Yellow("x")),
		BrightBg|CyanBg|YellowFg)
}

func Test_BgBrightWhite(t *testing.T) {
	testFunc(t, "BgBrightWhite", BgBrightWhite("x"), BrightBg|WhiteBg)
	testFunc(t, "Complex BgBrightWhite", BgBrightWhite(Overlined("x")),
		BrightBg|WhiteBg|OverlinedFm)
}

func Test_BgIndex(t *testing.T) {
	testFunc(t, "BgIndex", BgIndex(187, "x"), Color(187)<<shiftBg|flagBg)
	testFunc(t, "Complex BgIndex", BgIndex(187, Italic("x")),
		Color(187)<<shiftBg|flagBg|ItalicFm)
}

func Test_BgGray(t *testing.T) {
	testFunc(t, "BgGray", BgGray(15, "x"), Color(232+15)<<shiftBg|flagBg)
	testFunc(t, "Complex BgGray", BgGray(15, Index(216, "x")),
		(Color(15+232)<<shiftBg)|flagBg|(Color(216)<<shiftFg)|flagFg)
}

func Test_Colorize(t *testing.T) {
	testFunc(t, "Colorize", Colorize("x", RedFg|BoldFm), RedFg|BoldFm)
	testFunc(t, "Complex Colorize",
		Colorize(Italic("x"), RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm,
	)
}

func Test_bigGray(t *testing.T) {
	testFunc(t, "Gray", Gray(115, "x"), Color(232+23)<<shiftFg|flagFg)
	testFunc(t, "BgGray", BgGray(215, "x"), Color(232+23)<<shiftBg|flagBg)
}
