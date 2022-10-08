//
// Copyright (c) 2016-2022 The Aurora Authors. All rights reserved.
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

	"github.com/stretchr/testify/assert"
)

func isColor(v Value, clr Color) bool {
	return v.Color() == clr
}

func Test_New(t *testing.T) {
	var (
		dconf = NewConfig() // default configurations
		a     = New()
	)
	assert.Equal(t, dconf, a.conf, "non-default configurations")
	// options
	a = New(WithColors(true), WithHyperlinks(true))
	assert.True(t, a.conf.Colors)
	assert.True(t, a.conf.Hyperlinks)
	// colors
	a = New(WithColors(false), WithHyperlinks(false))
	assert.False(t, a.conf.Colors)
	assert.False(t, a.conf.Hyperlinks)

}

func TestAurora_Config(t *testing.T) {
	assert.Equal(t, NewConfig(), New().Config())
}

func TestAurora_no_colors(t *testing.T) {

	var a = New(WithColors(false), WithHyperlinks(false))

	var test = func(mn string, v Value) {
		t.Helper()
		t.Log(mn)
		assert.Zero(t, v.Color(), "colored")
		assert.Equal(t, "x", v.Value(), "wrong value")
	}

	test("Reset", a.Reset("x"))
	test("Clear", a.Clear("x"))

	test("Bold", a.Bold("x"))
	test("Faint", a.Faint("x"))
	test("DoublyUnderline", a.DoublyUnderline("x"))
	test("Fraktur", a.Fraktur("x"))
	test("Italic", a.Italic("x"))
	test("Underline", a.Underline("x"))
	test("SlowBlink", a.SlowBlink("x"))
	test("RapidBlink", a.RapidBlink("x"))
	test("Blink", a.Blink("x"))
	test("Reverse", a.Reverse("x"))
	test("Inverse", a.Inverse("x"))
	test("Conceal", a.Conceal("x"))
	test("Hidden", a.Hidden("x"))
	test("CrossedOut", a.CrossedOut("x"))
	test("StrikeThrough", a.StrikeThrough("x"))
	test("Framed", a.Framed("x"))
	test("Encircled", a.Encircled("x"))
	test("Overlined", a.Overlined("x"))

	test("Black", a.Black("x"))
	test("Red", a.Red("x"))
	test("Green", a.Green("x"))
	test("Yellow", a.Yellow("x"))
	test("Blue", a.Blue("x"))
	test("Magenta", a.Magenta("x"))
	test("Cyan", a.Cyan("x"))
	test("White", a.White("x"))
	test("BrightBlack", a.BrightBlack("x"))
	test("BrightRed", a.BrightRed("x"))
	test("BrightGreen", a.BrightGreen("x"))
	test("BrightYellow", a.BrightYellow("x"))
	test("BrightBlue", a.BrightBlue("x"))
	test("BrightMagenta", a.BrightMagenta("x"))
	test("BrightCyan", a.BrightCyan("x"))
	test("BrightWhite", a.BrightWhite("x"))
	test("Index", a.Index(178, "x"))
	test("Gray", a.Gray(14, "x"))

	test("BgBlack", a.BgBlack("x"))
	test("BgRed", a.BgRed("x"))
	test("BgGreen", a.BgGreen("x"))
	test("BgYellow", a.BgYellow("x"))
	test("BgBlue", a.BgBlue("x"))
	test("BgMagenta", a.BgMagenta("x"))
	test("BgCyan", a.BgCyan("x"))
	test("BgWhite", a.BgWhite("x"))
	test("BgBrightBlack", a.BgBrightBlack("x"))
	test("BgBrightRed", a.BgBrightRed("x"))
	test("BgBrightGreen", a.BgBrightGreen("x"))
	test("BgBrightYellow", a.BgBrightYellow("x"))
	test("BgBrightBlue", a.BgBrightBlue("x"))
	test("BgBrightMagenta", a.BgBrightMagenta("x"))
	test("BgBrightCyan", a.BgBrightCyan("x"))
	test("BgBrightWhite", a.BgBrightWhite("x"))
	test("BgIndex", a.BgIndex(187, "x"))
	test("BgGray", a.BgGray(15, "x"))

	test("Colorize", a.Colorize("x", RedFg|BlueBg|BrightBg|BoldFm))

}

func Test_noColors_sprintf(t *testing.T) {
	var au = New(WithColors(false))
	assert.Equal(t, "x: 2", au.Sprintf(au.Black("x: %d"), au.Blue(2)))
	assert.Equal(t, "x: 2", au.Sprintf("x: %d", au.Blue(2)))
}

func TestAurora_colored(t *testing.T) {

	var a = New() // with colors, with hyperlinks

	var test = func(mn string, v Value, clr Color) {
		t.Helper()
		t.Log(mn)
		assert.Equal(t, clr, v.Color())
		assert.Equal(t, "x", v.Value(), "wrong value")
	}
	test("Reset", a.Reset("x"), 0)

	test("Bold", a.Bold("x"), BoldFm)
	test("Faint", a.Faint("x"), FaintFm)
	test("DoublyUnderline", a.DoublyUnderline("x"), DoublyUnderlineFm)
	test("Fraktur", a.Fraktur("x"), FrakturFm)
	test("Italic", a.Italic("x"), ItalicFm)
	test("Underline", a.Underline("x"), UnderlineFm)
	test("SlowBlink", a.SlowBlink("x"), SlowBlinkFm)
	test("RapidBlink", a.RapidBlink("x"), RapidBlinkFm)
	test("Blink", a.Blink("x"), BlinkFm)
	test("Reverse", a.Reverse("x"), ReverseFm)
	test("Inverse", a.Inverse("x"), InverseFm)
	test("Conceal", a.Conceal("x"), ConcealFm)
	test("Hidden", a.Hidden("x"), HiddenFm)
	test("CrossedOut", a.CrossedOut("x"), CrossedOutFm)
	test("StrikeThrough", a.StrikeThrough("x"), StrikeThroughFm)
	test("Framed", a.Framed("x"), FramedFm)
	test("Encircled", a.Encircled("x"), EncircledFm)
	test("Overlined", a.Overlined("x"), OverlinedFm)

	test("Black", a.Black("x"), BlackFg)
	test("Red", a.Red("x"), RedFg)
	test("Green", a.Green("x"), GreenFg)
	test("Yellow", a.Yellow("x"), YellowFg)
	test("Blue", a.Blue("x"), BlueFg)
	test("Magenta", a.Magenta("x"), MagentaFg)
	test("Cyan", a.Cyan("x"), CyanFg)
	test("White", a.White("x"), WhiteFg)
	test("BrightBlack", a.BrightBlack("x"), BrightFg|BlackFg)
	test("BrightRed", a.BrightRed("x"), BrightFg|RedFg)
	test("BrightGreen", a.BrightGreen("x"), BrightFg|GreenFg)
	test("BrightYellow", a.BrightYellow("x"), BrightFg|YellowFg)
	test("BrightBlue", a.BrightBlue("x"), BrightFg|BlueFg)
	test("BrightMagenta", a.BrightMagenta("x"), BrightFg|MagentaFg)
	test("BrightCyan", a.BrightCyan("x"), BrightFg|CyanFg)
	test("BrightWhite", a.BrightWhite("x"), BrightFg|WhiteFg)
	test("Index", a.Index(178, "x"), (Color(178)<<shiftFg)|flagFg)
	test("Gray", a.Gray(14, "x"), (Color(232+14)<<shiftFg)|flagFg)

	test("BgBlack", a.BgBlack("x"), BlackBg)
	test("BgRed", a.BgRed("x"), RedBg)
	test("BgGreen", a.BgGreen("x"), GreenBg)
	test("BgYellow", a.BgYellow("x"), YellowBg)
	test("BgBlue", a.BgBlue("x"), BlueBg)
	test("BgMagenta", a.BgMagenta("x"), MagentaBg)
	test("BgCyan", a.BgCyan("x"), CyanBg)
	test("BgWhite", a.BgWhite("x"), WhiteBg)
	test("BgBrightBlack", a.BgBrightBlack("x"), BrightBg|BlackBg)
	test("BgBrightRed", a.BgBrightRed("x"), BrightBg|RedBg)
	test("BgBrightGreen", a.BgBrightGreen("x"), BrightBg|GreenBg)
	test("BgBrightYellow", a.BgBrightYellow("x"), BrightBg|YellowBg)
	test("BgBrightBlue", a.BgBrightBlue("x"), BrightBg|BlueBg)
	test("BgBrightMagenta", a.BgBrightMagenta("x"), BrightBg|MagentaBg)
	test("BgBrightCyan", a.BgBrightCyan("x"), BrightBg|CyanBg)
	test("BgBrightWhite", a.BgBrightWhite("x"), BrightBg|WhiteBg)
	test("BgIndex", a.BgIndex(187, "x"), (Color(187)<<shiftBg)|flagBg)
	test("BgGray", a.BgGray(15, "x"), (Color(15+232)<<shiftBg)|flagBg)

	test("Colorize", a.Colorize("x", RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm)
}

func TestAurora_Sprintf(t *testing.T) {
	var a = New()
	assert.Equal(t, "\033[30mx: \033[0;34m2\033[0;30mB\033[0m",
		a.Sprintf(a.Black("x: %dB"), a.Blue(2)))
	assert.Equal(t, "x: \033[34m2\033[0mB",
		a.Sprintf("x: %dB", a.Blue(2)))
}
