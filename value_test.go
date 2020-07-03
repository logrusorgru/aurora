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
	"fmt"
	"testing"
)

func TestValue_String(t *testing.T) {
	var v Value
	// colorized
	v = value{"x", 0, 0}
	if x := v.String(); x != "x" {
		t.Errorf("(value).String: want %q, got %q", "x", x)
	}
	v = value{"x", BlackFg, RedBg}
	want := "\033[0;30mx\033[0;41m"
	if got := v.String(); want != got {
		t.Errorf("(value).String: want %q, got %q", want, got)
	}
	// clear
	v = valueClear{"x"}
	if x := v.String(); x != "x" {
		t.Errorf("(value).String: want %q, got %q", "x", x)
	}

}

func TestValue_Color(t *testing.T) {
	// colorized
	if (value{"", RedFg, 0}).Color() != RedFg {
		t.Error("wrong color")
	}
	// clear
	if (valueClear{0}).Color() != 0 {
		t.Error("wrong color")
	}
}

func TestValue_Value(t *testing.T) {
	// colorized
	if (value{"x", RedFg, BlueBg}).Value() != "x" {
		t.Error("wrong value")
	}
	// clear
	if (valueClear{"x"}).Value() != "x" {
		t.Error("wrong value")
	}
}

func TestValue_Bleach(t *testing.T) {
	// colorized
	if (value{"x", RedFg, BlueBg}).Bleach() != (value{value: "x"}) {
		t.Error("wrong bleached")
	}
	// clear
	if (valueClear{"x"}).Bleach() != (valueClear{"x"}) {
		t.Error("wrong bleached")
	}
}

func TestValue_Format(t *testing.T) {
	var v Value
	var want, got string
	//
	// colorized
	//
	v = value{3.14, RedFg, BlueBg}
	got = fmt.Sprintf("%+1.3g", v)
	want = "\033[0;31m" + fmt.Sprintf("%+1.3g", 3.14) + "\033[0;44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	var utf8Verb = "%+1.3世" // verb that fit more then 1 byte
	got = fmt.Sprintf(utf8Verb, v)
	want = "\033[0;31m" + fmt.Sprintf(utf8Verb, 3.14) + "\033[0;44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	// clear
	//
	v = valueClear{3.14}
	got = fmt.Sprintf("%+1.3g", v)
	want = fmt.Sprintf("%+1.3g", 3.14)
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	got = fmt.Sprintf(utf8Verb, v)
	want = fmt.Sprintf(utf8Verb, 3.14)
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
}

func Test_tail(t *testing.T) {
	// colorized
	if (value{"x", 0, BlueBg}).tail() != BlueBg {
		t.Error("wrong tail color")
	}
	// clear
	if (valueClear{"x"}).tail() != 0 {
		t.Error("wrong tail color")
	}
}

func Test_setTail(t *testing.T) {
	// colorized
	if (value{"x", 0, 0}).setTail(RedFg) != (value{"x", 0, RedFg}) {
		t.Error("wrong setTail behavior")
	}
	// clear
	if (valueClear{"x"}).setTail(RedFg) != (valueClear{"x"}) {
		t.Error("wrong setTail behavior")
	}
}

func TestValue_colors(t *testing.T) {
	test := func(name string, v Value, clr Color) {
		t.Helper()
		if c := v.Color(); c != clr {
			t.Errorf("wrong color for %s: want %d, got %d", name, clr, c)
		}
	}
	// colorized
	test("Reset", Reset("x"), 0)
	test("Bold", Bold("x"), BoldFm)
	test("Faint", Faint("x"), FaintFm)
	test("DoublyUnderline", DoublyUnderline("x"), DoublyUnderlineFm)
	test("Fraktur & Yellow", Yellow(Fraktur("x")), FrakturFm|YellowFg)
	test("Italic", Italic("x"), ItalicFm)
	test("Underline", Underline("x"), UnderlineFm)
	test("SlowBlink", SlowBlink("x"), SlowBlinkFm)
	test("RapidBlink", RapidBlink("x"), RapidBlinkFm)
	test("Blink", Blink("x"), BlinkFm)
	test("Reverse", Reverse("x"), ReverseFm)
	test("Inverse", Inverse("x"), InverseFm)
	test("Conceal", Conceal("x"), ConcealFm)
	test("Hidden", Hidden("x"), HiddenFm)
	test("CrossedOut", CrossedOut("x"), CrossedOutFm)
	test("StrikeThrough", StrikeThrough("x"), StrikeThroughFm)
	test("Framed", Framed("x"), FramedFm)
	test("Encircled", Encircled("x"), EncircledFm)
	test("Overlined", Overlined("x"), OverlinedFm)
	test("Black", Black("x"), BlackFg)
	test("Red", Red("x"), RedFg)
	test("Green", Green("x"), GreenFg)
	test("Yellow", Yellow("x"), YellowFg)
	test("Brown", Brown("x"), BrownFg)
	test("Blue", Blue("x"), BlueFg)
	test("Magenta", Magenta("x"), MagentaFg)
	test("Cyan", Cyan("x"), CyanFg)
	test("White", White("x"), WhiteFg)
	test("BrightBlack", BrightBlack("x"), BrightFg|BlackFg)
	test("BrightRed", BrightRed("x"), BrightFg|RedFg)
	test("BrightGreen", BrightGreen("x"), BrightFg|GreenFg)
	test("BrightYellow", BrightYellow("x"), BrightFg|YellowFg)
	test("BrightBlue", BrightBlue("x"), BrightFg|BlueFg)
	test("BrightMagenta", BrightMagenta("x"), BrightFg|MagentaFg)
	test("BrightCyan", BrightCyan("x"), BrightFg|CyanFg)
	test("BrightWhite", BrightWhite("x"), BrightFg|WhiteFg)
	test("Index", Index(178, "x"), (Color(178)<<shiftFg)|flagFg)
	test("Gray", Gray(14, "x"), (Color(14+232)<<shiftFg)|flagFg)
	test("BgBlack", BgBlack("x"), BlackBg)
	test("BgRed", BgRed("x"), RedBg)
	test("BgGreen", BgGreen("x"), GreenBg)
	test("BgYellow", BgYellow("x"), YellowBg)
	test("BgBrown", BgBrown("x"), BrownBg)
	test("BgBlue", BgBlue("x"), BlueBg)
	test("BgMagenta", BgMagenta("x"), MagentaBg)
	test("BgCyan", BgCyan("x"), CyanBg)
	test("BgWhite", BgWhite("x"), WhiteBg)
	test("BgBrightBlack", BgBrightBlack("x"), BrightBg|BlackBg)
	test("BgBrightRed", BgBrightRed("x"), BrightBg|RedBg)
	test("BgBrightGreen", BgBrightGreen("x"), BrightBg|GreenBg)
	test("BgBrightYellow", BgBrightYellow("x"), BrightBg|YellowBg)
	test("BgBrightBlue", BgBrightBlue("x"), BrightBg|BlueBg)
	test("BgBrightMagenta", BgBrightMagenta("x"), BrightBg|MagentaBg)
	test("BgBrightCyan", BgBrightCyan("x"), BrightBg|CyanBg)
	test("BgBrightWhite", BgBrightWhite("x"), BrightBg|WhiteBg)
	test("BgIndex", BgIndex(187, "x"), Color(187)<<shiftBg|flagBg)
	test("BgGray", BgGray(15, "x"), Color(232+15)<<shiftBg|flagBg)
	test("Colorize", Colorize("x", RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm)
	// clear
	test("Reset", valueClear{"x"}.Reset(), 0)
	test("Bold", valueClear{"x"}.Bold(), 0)
	test("Faint", valueClear{"x"}.Faint(), 0)
	test("DoublyUnderline", valueClear{"x"}.DoublyUnderline(), 0)
	test("Fraktur & Yellow", valueClear{"x"}.Fraktur(), 0)
	test("Italic", valueClear{"x"}.Italic(), 0)
	test("Underline", valueClear{"x"}.Underline(), 0)
	test("SlowBlink", valueClear{"x"}.SlowBlink(), 0)
	test("RapidBlink", valueClear{"x"}.RapidBlink(), 0)
	test("Blink", valueClear{"x"}.Blink(), 0)
	test("Reverse", valueClear{"x"}.Reverse(), 0)
	test("Inverse", valueClear{"x"}.Inverse(), 0)
	test("Conceal", valueClear{"x"}.Conceal(), 0)
	test("Hidden", valueClear{"x"}.Hidden(), 0)
	test("CrossedOut", valueClear{"x"}.CrossedOut(), 0)
	test("StrikeThrough", valueClear{"x"}.StrikeThrough(), 0)
	test("Framed", valueClear{"x"}.Framed(), 0)
	test("Encircled", valueClear{"x"}.Encircled(), 0)
	test("Overlined", valueClear{"x"}.Overlined(), 0)
	test("Black", valueClear{"x"}.Black(), 0)
	test("Red", valueClear{"x"}.Red(), 0)
	test("Green", valueClear{"x"}.Green(), 0)
	test("Yellow", valueClear{"x"}.Yellow(), 0)
	test("Brown", valueClear{"x"}.Brown(), 0)
	test("Blue", valueClear{"x"}.Blue(), 0)
	test("Magenta", valueClear{"x"}.Magenta(), 0)
	test("Cyan", valueClear{"x"}.Cyan(), 0)
	test("White", valueClear{"x"}.White(), 0)
	test("BrightBlack", valueClear{"x"}.BrightBlack(), 0)
	test("BrightRed", valueClear{"x"}.BrightRed(), 0)
	test("BrightGreen", valueClear{"x"}.BrightGreen(), 0)
	test("BrightYellow", valueClear{"x"}.BrightYellow(), 0)
	test("BrightBlue", valueClear{"x"}.BrightBlue(), 0)
	test("BrightMagenta", valueClear{"x"}.BrightMagenta(), 0)
	test("BrightCyan", valueClear{"x"}.BrightCyan(), 0)
	test("BrightWhite", valueClear{"x"}.BrightWhite(), 0)
	test("Index", valueClear{"x"}.Index(178), 0)
	test("Gray", valueClear{"x"}.Gray(14), 0)
	test("BgBlack", valueClear{"x"}.BgBlack(), 0)
	test("BgRed", valueClear{"x"}.BgRed(), 0)
	test("BgGreen", valueClear{"x"}.BgGreen(), 0)
	test("BgYellow", valueClear{"x"}.BgYellow(), 0)
	test("BgBrown", valueClear{"x"}.BgBrown(), 0)
	test("BgBlue", valueClear{"x"}.BgBlue(), 0)
	test("BgMagenta", valueClear{"x"}.BgMagenta(), 0)
	test("BgCyan", valueClear{"x"}.BgCyan(), 0)
	test("BgWhite", valueClear{"x"}.BgWhite(), 0)
	test("BgBrightBlack", valueClear{"x"}.BgBrightBlack(), 0)
	test("BgBrightRed", valueClear{"x"}.BgBrightRed(), 0)
	test("BgBrightGreen", valueClear{"x"}.BgBrightGreen(), 0)
	test("BgBrightYellow", valueClear{"x"}.BgBrightYellow(), 0)
	test("BgBrightBlue", valueClear{"x"}.BgBrightBlue(), 0)
	test("BgBrightMagenta", valueClear{"x"}.BgBrightMagenta(), 0)
	test("BgBrightCyan", valueClear{"x"}.BgBrightCyan(), 0)
	test("BgBrightWhite", valueClear{"x"}.BgBrightWhite(), 0)
	test("BgIndex", valueClear{"x"}.BgIndex(187), 0)
	test("BgGray", valueClear{"x"}.BgGray(15), 0)
	test("Colorize", valueClear{"x"}.Colorize(RedFg|BlueBg|BrightBg|BoldFm), 0)
	// change
	test("Reset", Bold("x").Reset(), 0)
	test("Bold", Faint("x").Bold(), BoldFm)
	test("Faint", Bold("x").Faint(), FaintFm)
	test("DoublyUnderline", Reset("x").DoublyUnderline(), DoublyUnderlineFm)
	test("Fraktur & Yellow", Reset("x").Yellow().Fraktur(), FrakturFm|YellowFg)
	test("Italic", Reset("x").Italic(), ItalicFm)
	test("Underline", Reset("x").Underline(), UnderlineFm)
	test("SlowBlink", RapidBlink("x").SlowBlink(), SlowBlinkFm)
	test("RapidBlink", SlowBlink("x").RapidBlink(), RapidBlinkFm)
	test("Blink", Reset("x").Blink(), BlinkFm)
	test("Reverse", Reset("x").Reverse(), ReverseFm)
	test("Inverse", Reset("x").Inverse(), InverseFm)
	test("Conceal", Reset("x").Conceal(), ConcealFm)
	test("Hidden", Reset("x").Hidden(), HiddenFm)
	test("CrossedOut", Reset("x").CrossedOut(), CrossedOutFm)
	test("StrikeThrough", Reset("x").StrikeThrough(), StrikeThroughFm)
	test("Framed", Reset("x").Framed(), FramedFm)
	test("Encircled", Reset("x").Encircled(), EncircledFm)
	test("Overlined", Reset("x").Overlined(), OverlinedFm)
	test("Black", Reset("x").Black(), BlackFg)
	test("Red", Reset("x").Red(), RedFg)
	test("Green", Reset("x").Green(), GreenFg)
	test("Yellow", Reset("x").Yellow(), YellowFg)
	test("Brown", Reset("x").Brown(), BrownFg)
	test("Blue", Reset("x").Blue(), BlueFg)
	test("Magenta", Reset("x").Magenta(), MagentaFg)
	test("Cyan", Reset("x").Cyan(), CyanFg)
	test("White", Reset("x").White(), WhiteFg)
	test("BrightBlack", Reset("x").BrightBlack(), BrightFg|BlackFg)
	test("BrightRed", Reset("x").BrightRed(), BrightFg|RedFg)
	test("BrightGreen", Reset("x").BrightGreen(), BrightFg|GreenFg)
	test("BrightYellow", Reset("x").BrightYellow(), BrightFg|YellowFg)
	test("BrightBlue", Reset("x").BrightBlue(), BrightFg|BlueFg)
	test("BrightMagenta", Reset("x").BrightMagenta(), BrightFg|MagentaFg)
	test("BrightCyan", Reset("x").BrightCyan(), BrightFg|CyanFg)
	test("BrightWhite", Reset("x").BrightWhite(), BrightFg|WhiteFg)
	test("Index", Reset("x").Index(178), (Color(178)<<shiftFg)|flagFg)
	test("Gray", Reset("x").Gray(14), (Color(14+232)<<shiftFg)|flagFg)
	test("BgBlack", Reset("x").BgBlack(), BlackBg)
	test("BgRed", Reset("x").BgRed(), RedBg)
	test("BgGreen", Reset("x").BgGreen(), GreenBg)
	test("BgYellow", Reset("x").BgYellow(), YellowBg)
	test("BgBrown", Reset("x").BgBrown(), BrownBg)
	test("BgBlue", Reset("x").BgBlue(), BlueBg)
	test("BgMagenta", Reset("x").BgMagenta(), MagentaBg)
	test("BgCyan", Reset("x").BgCyan(), CyanBg)
	test("BgWhite", Reset("x").BgWhite(), WhiteBg)
	test("BgBrightBlack", Reset("x").BgBrightBlack(), BrightBg|BlackBg)
	test("BgBrightRed", Reset("x").BgBrightRed(), BrightBg|RedBg)
	test("BgBrightGreen", Reset("x").BgBrightGreen(), BrightBg|GreenBg)
	test("BgBrightYellow", Reset("x").BgBrightYellow(), BrightBg|YellowBg)
	test("BgBrightBlue", Reset("x").BgBrightBlue(), BrightBg|BlueBg)
	test("BgBrightMagenta", Reset("x").BgBrightMagenta(), BrightBg|MagentaBg)
	test("BgBrightCyan", Reset("x").BgBrightCyan(), BrightBg|CyanBg)
	test("BgBrightWhite", Reset("x").BgBrightWhite(), BrightBg|WhiteBg)
	test("BgIndex", Reset("x").BgIndex(187), Color(187)<<shiftBg|flagBg)
	test("BgGray", Reset("x").BgGray(15), Color(232+15)<<shiftBg|flagBg)
	test("Colorize", Reset("x").Colorize(RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm)
	// overflow
	test("Gray", Reset("x").Gray(151), Color(232+23)<<shiftFg|flagFg)
	test("BgGray", Reset("x").BgGray(115), Color(232+23)<<shiftBg|flagBg)
}
