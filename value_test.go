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
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func q(s string) string {
	return strconv.Quote(s)
}

func TestValue_String(t *testing.T) {
	var au *Aurora
	// colorized
	au = New()
	assert.Equal(t, "x", au.Clear("x").String())
	assert.Equal(t, "\033[30;41mx\033[0m", au.Black("x").BgRed().String())
	// clear
	au = New(WithColors(false))
	assert.Equal(t, "x", au.Black("x").BgRed().String())
	// colorized hyperlink
	au = New()
	assert.Equal(t, `]8;;http://example.com`+
		`\[31mx[0m]8;;\`,
		au.Red("x").Hyperlink("http://example.com").String())
	// clear hyperlink
	assert.Equal(t, `]8;;http://example.com\x]8;;\`,
		au.Clear("x").Hyperlink("http://example.com").String())
}

func TestValue_Color(t *testing.T) {
	// colorized
	assert.Equal(t, RedFg, New().Red("x").Color())
	// clear, also red
	assert.Equal(t, RedFg, New().Red("x").Color())
}

func TestValue_Value(t *testing.T) {
	// colorized
	assert.Equal(t, "x", New().Clear("x").Value())
	// clear
	assert.Equal(t, "x", New(WithColors(false)).Black("x").BgRed().Value())
}

func TestValue_Reset(t *testing.T) {
	// colorized
	var au = New()
	assert.Equal(t, Value{
		cc:    au.cc,
		value: "x",
	}, au.Red("x").BgBlack().Reset())
	// clear
	au = New(WithColors(false))
	assert.Equal(t, Value{
		cc:    au.cc,
		value: "x",
	}, au.Red("x").BgBlack().Reset())
	// a hyperlink
	au = New()
	assert.Equal(t,
		Value{
			cc:    au.cc,
			value: "x",
		}, au.Red("x").
			BgBlack().
			Hyperlink("http://example.com/path", HyperlinkID("10")).
			Reset())
}

func TestValue_Clear(t *testing.T) {
	// colorized
	var au = New()
	assert.Equal(t, Value{
		cc:    au.cc,
		value: "x",
	}, au.Red("x").BgBlack().Clear())
	// clear
	au = New(WithColors(false))
	assert.Equal(t, Value{
		cc:    au.cc,
		value: "x",
	}, au.Red("x").BgBlack().Clear())
	// a hyperlink
	au = New()
	assert.EqualValues(t,
		Value{
			cc:    au.cc,
			value: "x",
			hyperlink: &hyperlink{
				target: "http://example.com/path",
				params: []HyperlinkParam{{
					Key:   HyperlinkIDKey,
					Value: "10",
				}},
			},
		}, au.Red("x").
			BgBlack().
			Hyperlink("http://example.com/path", HyperlinkID("10")).
			Clear())
}

func TestValue_Format(t *testing.T) {
	// colorized
	var au = New()
	assert.Equal(t, "\033[31;44m"+fmt.Sprintf("%+1.3g", 3.14)+"\033[0m",
		fmt.Sprintf("%+1.3g", au.Red(3.14).BgBlue()))
	const utf8Verb = "%+1.3世" // verb that fit more then 1 byte
	assert.Equal(t, "\033[31;44m"+"%!世(float64=+3.14)"+"\033[0m",
		fmt.Sprintf(utf8Verb, au.Red(3.14).BgBlue()))
	// clear
	au = New(WithColors(false))
	assert.Equal(t, fmt.Sprintf("%+1.3g", 3.14),
		fmt.Sprintf("%+1.3g", au.Red(3.14).BgBlue()))
	assert.Equal(t, "%!世(float64=+3.14)",
		fmt.Sprintf(utf8Verb, au.Red(3.14).BgBlue()))
}

func TestValue_colors(t *testing.T) {
	var test = func(name string, v Value, clr Color) {
		t.Helper()
		assert.Equal(t, clr, v.Color())
	}
	// colorized
	var au = New()
	test("Reset", au.Reset("x"), 0)
	test("Clear", au.Clear("x"), 0)
	test("Bold", au.Bold("x"), BoldFm)
	test("Faint", au.Faint("x"), FaintFm)
	test("DoublyUnderline", au.DoublyUnderline("x"), DoublyUnderlineFm)
	test("Fraktur & au.Yellow", Yellow(Fraktur("x")), FrakturFm|YellowFg)
	test("Italic", au.Italic("x"), ItalicFm)
	test("Underline", au.Underline("x"), UnderlineFm)
	test("SlowBlink", au.SlowBlink("x"), SlowBlinkFm)
	test("RapidBlink", au.RapidBlink("x"), RapidBlinkFm)
	test("Blink", au.Blink("x"), BlinkFm)
	test("Reverse", au.Reverse("x"), ReverseFm)
	test("Inverse", au.Inverse("x"), InverseFm)
	test("Conceal", au.Conceal("x"), ConcealFm)
	test("Hidden", au.Hidden("x"), HiddenFm)
	test("CrossedOut", au.CrossedOut("x"), CrossedOutFm)
	test("StrikeThrough", au.StrikeThrough("x"), StrikeThroughFm)
	test("Framed", au.Framed("x"), FramedFm)
	test("Encircled", au.Encircled("x"), EncircledFm)
	test("Overlined", au.Overlined("x"), OverlinedFm)
	test("Black", au.Black("x"), BlackFg)
	test("Red", au.Red("x"), RedFg)
	test("Green", au.Green("x"), GreenFg)
	test("Yellow", au.Yellow("x"), YellowFg)
	test("Blue", au.Blue("x"), BlueFg)
	test("Magenta", au.Magenta("x"), MagentaFg)
	test("Cyan", au.Cyan("x"), CyanFg)
	test("White", au.White("x"), WhiteFg)
	test("BrightBlack", au.BrightBlack("x"), BrightFg|BlackFg)
	test("BrightRed", au.BrightRed("x"), BrightFg|RedFg)
	test("BrightGreen", au.BrightGreen("x"), BrightFg|GreenFg)
	test("BrightYellow", au.BrightYellow("x"), BrightFg|YellowFg)
	test("BrightBlue", au.BrightBlue("x"), BrightFg|BlueFg)
	test("BrightMagenta", au.BrightMagenta("x"), BrightFg|MagentaFg)
	test("BrightCyan", au.BrightCyan("x"), BrightFg|CyanFg)
	test("BrightWhite", au.BrightWhite("x"), BrightFg|WhiteFg)
	test("Index", au.Index(178, "x"), (Color(178)<<shiftFg)|flagFg)
	test("Gray", au.Gray(14, "x"), (Color(14+232)<<shiftFg)|flagFg)
	test("BgBlack", au.BgBlack("x"), BlackBg)
	test("BgRed", au.BgRed("x"), RedBg)
	test("BgGreen", au.BgGreen("x"), GreenBg)
	test("BgYellow", au.BgYellow("x"), YellowBg)
	test("BgBlue", au.BgBlue("x"), BlueBg)
	test("BgMagenta", au.BgMagenta("x"), MagentaBg)
	test("BgCyan", au.BgCyan("x"), CyanBg)
	test("BgWhite", au.BgWhite("x"), WhiteBg)
	test("BgBrightBlack", au.BgBrightBlack("x"), BrightBg|BlackBg)
	test("BgBrightRed", au.BgBrightRed("x"), BrightBg|RedBg)
	test("BgBrightGreen", au.BgBrightGreen("x"), BrightBg|GreenBg)
	test("BgBrightYellow", au.BgBrightYellow("x"), BrightBg|YellowBg)
	test("BgBrightBlue", au.BgBrightBlue("x"), BrightBg|BlueBg)
	test("BgBrightMagenta", au.BgBrightMagenta("x"), BrightBg|MagentaBg)
	test("BgBrightCyan", au.BgBrightCyan("x"), BrightBg|CyanBg)
	test("BgBrightWhite", au.BgBrightWhite("x"), BrightBg|WhiteBg)
	test("BgIndex", au.BgIndex(187, "x"), Color(187)<<shiftBg|flagBg)
	test("BgGray", au.BgGray(15, "x"), Color(232+15)<<shiftBg|flagBg)
	test("Colorize", au.Colorize("x", RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm)
	// clear
	au = New(WithColors(false))
	test("Reset", au.Clear("x").Reset(), 0)
	test("Clear", au.Bold("x").Clear(), 0)
	test("Bold", au.Clear("x").Bold(), 0)
	test("Faint", au.Clear("x").Faint(), 0)
	test("DoublyUnderline", au.Clear("x").DoublyUnderline(), 0)
	test("Fraktur & Yellow", au.Clear("x").Fraktur(), 0)
	test("Italic", au.Clear("x").Italic(), 0)
	test("Underline", au.Clear("x").Underline(), 0)
	test("SlowBlink", au.Clear("x").SlowBlink(), 0)
	test("RapidBlink", au.Clear("x").RapidBlink(), 0)
	test("Blink", au.Clear("x").Blink(), 0)
	test("Reverse", au.Clear("x").Reverse(), 0)
	test("Inverse", au.Clear("x").Inverse(), 0)
	test("Conceal", au.Clear("x").Conceal(), 0)
	test("Hidden", au.Clear("x").Hidden(), 0)
	test("CrossedOut", au.Clear("x").CrossedOut(), 0)
	test("StrikeThrough", au.Clear("x").StrikeThrough(), 0)
	test("Framed", au.Clear("x").Framed(), 0)
	test("Encircled", au.Clear("x").Encircled(), 0)
	test("Overlined", au.Clear("x").Overlined(), 0)
	test("Black", au.Clear("x").Black(), 0)
	test("Red", au.Clear("x").Red(), 0)
	test("Green", au.Clear("x").Green(), 0)
	test("Yellow", au.Clear("x").Yellow(), 0)
	test("Blue", au.Clear("x").Blue(), 0)
	test("Magenta", au.Clear("x").Magenta(), 0)
	test("Cyan", au.Clear("x").Cyan(), 0)
	test("White", au.Clear("x").White(), 0)
	test("BrightBlack", au.Clear("x").BrightBlack(), 0)
	test("BrightRed", au.Clear("x").BrightRed(), 0)
	test("BrightGreen", au.Clear("x").BrightGreen(), 0)
	test("BrightYellow", au.Clear("x").BrightYellow(), 0)
	test("BrightBlue", au.Clear("x").BrightBlue(), 0)
	test("BrightMagenta", au.Clear("x").BrightMagenta(), 0)
	test("BrightCyan", au.Clear("x").BrightCyan(), 0)
	test("BrightWhite", au.Clear("x").BrightWhite(), 0)
	test("Index", au.Clear("x").Index(178), 0)
	test("Gray", au.Clear("x").Gray(14), 0)
	test("BgBlack", au.Clear("x").BgBlack(), 0)
	test("BgRed", au.Clear("x").BgRed(), 0)
	test("BgGreen", au.Clear("x").BgGreen(), 0)
	test("BgYellow", au.Clear("x").BgYellow(), 0)
	test("BgBlue", au.Clear("x").BgBlue(), 0)
	test("BgMagenta", au.Clear("x").BgMagenta(), 0)
	test("BgCyan", au.Clear("x").BgCyan(), 0)
	test("BgWhite", au.Clear("x").BgWhite(), 0)
	test("BgBrightBlack", au.Clear("x").BgBrightBlack(), 0)
	test("BgBrightRed", au.Clear("x").BgBrightRed(), 0)
	test("BgBrightGreen", au.Clear("x").BgBrightGreen(), 0)
	test("BgBrightYellow", au.Clear("x").BgBrightYellow(), 0)
	test("BgBrightBlue", au.Clear("x").BgBrightBlue(), 0)
	test("BgBrightMagenta", au.Clear("x").BgBrightMagenta(), 0)
	test("BgBrightCyan", au.Clear("x").BgBrightCyan(), 0)
	test("BgBrightWhite", au.Clear("x").BgBrightWhite(), 0)
	test("BgIndex", au.Clear("x").BgIndex(187), 0)
	test("BgGray", au.Clear("x").BgGray(15), 0)
	test("Colorize", au.Clear("x").Colorize(RedFg|BlueBg|BrightBg|BoldFm), 0)
	// change
	au = New()
	test("Reset", au.Bold("x").Reset(), 0)
	test("Clear", au.Bold("x").Clear(), 0)
	test("Bold", au.Faint("x").Bold(), BoldFm)
	test("Faint", au.Bold("x").Faint(), FaintFm)
	test("DoublyUnderline", au.Reset("x").DoublyUnderline(), DoublyUnderlineFm)
	test("Fraktur & Yellow", au.Reset("x").Yellow().Fraktur(), FrakturFm|YellowFg)
	test("Italic", au.Reset("x").Italic(), ItalicFm)
	test("Underline", au.Reset("x").Underline(), UnderlineFm)
	test("SlowBlink", au.RapidBlink("x").SlowBlink(), SlowBlinkFm)
	test("RapidBlink", au.SlowBlink("x").RapidBlink(), RapidBlinkFm)
	test("Blink", au.Reset("x").Blink(), BlinkFm)
	test("Reverse", au.Reset("x").Reverse(), ReverseFm)
	test("Inverse", au.Reset("x").Inverse(), InverseFm)
	test("Conceal", au.Reset("x").Conceal(), ConcealFm)
	test("Hidden", au.Reset("x").Hidden(), HiddenFm)
	test("CrossedOut", au.Reset("x").CrossedOut(), CrossedOutFm)
	test("StrikeThrough", au.Reset("x").StrikeThrough(), StrikeThroughFm)
	test("Framed", au.Reset("x").Framed(), FramedFm)
	test("Encircled", au.Reset("x").Encircled(), EncircledFm)
	test("Overlined", au.Reset("x").Overlined(), OverlinedFm)
	test("Black", au.Reset("x").Black(), BlackFg)
	test("Red", au.Reset("x").Red(), RedFg)
	test("Green", au.Reset("x").Green(), GreenFg)
	test("Yellow", au.Reset("x").Yellow(), YellowFg)
	test("Blue", au.Reset("x").Blue(), BlueFg)
	test("Magenta", au.Reset("x").Magenta(), MagentaFg)
	test("Cyan", au.Reset("x").Cyan(), CyanFg)
	test("White", au.Reset("x").White(), WhiteFg)
	test("BrightBlack", au.Reset("x").BrightBlack(), BrightFg|BlackFg)
	test("BrightRed", au.Reset("x").BrightRed(), BrightFg|RedFg)
	test("BrightGreen", au.Reset("x").BrightGreen(), BrightFg|GreenFg)
	test("BrightYellow", au.Reset("x").BrightYellow(), BrightFg|YellowFg)
	test("BrightBlue", au.Reset("x").BrightBlue(), BrightFg|BlueFg)
	test("BrightMagenta", au.Reset("x").BrightMagenta(), BrightFg|MagentaFg)
	test("BrightCyan", au.Reset("x").BrightCyan(), BrightFg|CyanFg)
	test("BrightWhite", au.Reset("x").BrightWhite(), BrightFg|WhiteFg)
	test("Index", au.Reset("x").Index(178), (Color(178)<<shiftFg)|flagFg)
	test("Gray", au.Reset("x").Gray(14), (Color(14+232)<<shiftFg)|flagFg)
	test("BgBlack", au.Reset("x").BgBlack(), BlackBg)
	test("BgRed", au.Reset("x").BgRed(), RedBg)
	test("BgGreen", au.Reset("x").BgGreen(), GreenBg)
	test("BgYellow", au.Reset("x").BgYellow(), YellowBg)
	test("BgBlue", au.Reset("x").BgBlue(), BlueBg)
	test("BgMagenta", au.Reset("x").BgMagenta(), MagentaBg)
	test("BgCyan", au.Reset("x").BgCyan(), CyanBg)
	test("BgWhite", au.Reset("x").BgWhite(), WhiteBg)
	test("BgBrightBlack", au.Reset("x").BgBrightBlack(), BrightBg|BlackBg)
	test("BgBrightRed", au.Reset("x").BgBrightRed(), BrightBg|RedBg)
	test("BgBrightGreen", au.Reset("x").BgBrightGreen(), BrightBg|GreenBg)
	test("BgBrightYellow", au.Reset("x").BgBrightYellow(), BrightBg|YellowBg)
	test("BgBrightBlue", au.Reset("x").BgBrightBlue(), BrightBg|BlueBg)
	test("BgBrightMagenta", au.Reset("x").BgBrightMagenta(), BrightBg|MagentaBg)
	test("BgBrightCyan", au.Reset("x").BgBrightCyan(), BrightBg|CyanBg)
	test("BgBrightWhite", au.Reset("x").BgBrightWhite(), BrightBg|WhiteBg)
	test("BgIndex", au.Reset("x").BgIndex(187), Color(187)<<shiftBg|flagBg)
	test("BgGray", au.Reset("x").BgGray(15), Color(232+15)<<shiftBg|flagBg)
	test("Colorize", au.Reset("x").Colorize(RedFg|BlueBg|BrightBg|BoldFm),
		RedFg|BlueBg|BrightBg|BoldFm)
	// overflow
	test("Gray", au.Reset("x").Gray(151), Color(232+23)<<shiftFg|flagFg)
	test("BgGray", au.Reset("x").BgGray(115), Color(232+23)<<shiftBg|flagBg)
}

func TestValue_hyperlinks(t *testing.T) {
	const target = "http://example.com/path?query=value"
	var (
		au     = New()
		params = []HyperlinkParam{
			{
				Key:   HyperlinkIDKey,
				Value: "10",
			},
			{
				Key:   "another",
				Value: "custom",
			},
		}
		val  Value
		want string
	)
	val = au.Red("x").Hyperlink(target, HyperlinkID("10"),
		HyperlinkParam{"another", "custom"})
	assert.Equal(t, target, val.HyperlinkTarget())
	assert.EqualValues(t, params, val.HyperlinkParams())
	want = `]8;id=10:another=custom;http://example.com/path?` +
		`query=value\[31mx[0m]8;;\`
	assert.Equal(t, want, val.String())

	// wrap a hyperlink
	au = New(WithHyperlinks(false))
	want = `[31mhttp://example.com[0m`
	val = au.Red("x").Hyperlink("http://example.com")
	assert.Equal(t, want, val.String())
	assert.Equal(t, "http://example.com", val.HyperlinkTarget())
	assert.Equal(t, "http://example.com", val.Value())

	// no target, no parameters
	val = au.Clear("x")
	assert.Equal(t, "", val.HyperlinkTarget())
	assert.Nil(t, val.HyperlinkParams())
}
