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
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor_Nos(t *testing.T) {

	for _, zero := range []bool{
		false, true,
	} {
		for i, val := range []struct {
			Nos   string
			Color Color
		}{
			{"1", BoldFm},
			{"2", FaintFm},
			{"3", ItalicFm},
			{"4", UnderlineFm},
			{"5", SlowBlinkFm},
			{"6", RapidBlinkFm},
			{"7", ReverseFm},
			{"8", ConcealFm},
			{"9", CrossedOutFm},

			{"20", FrakturFm},
			{"21", DoublyUnderlineFm},

			{"51", FramedFm},
			{"52", EncircledFm},
			{"53", OverlinedFm},

			{"30", BlackFg},
			{"31", RedFg},
			{"32", GreenFg},
			{"33", YellowFg},
			{"34", BlueFg},
			{"35", MagentaFg},
			{"36", CyanFg},
			{"37", WhiteFg},

			{"90", BlackFg | BrightFg},
			{"91", RedFg | BrightFg},
			{"92", GreenFg | BrightFg},
			{"93", YellowFg | BrightFg},
			{"94", BlueFg | BrightFg},
			{"95", MagentaFg | BrightFg},
			{"96", CyanFg | BrightFg},
			{"97", WhiteFg | BrightFg},

			{"90", BrightFg},

			{"40", BlackBg},
			{"41", RedBg},
			{"42", GreenBg},
			{"43", YellowBg},
			{"44", BlueBg},
			{"45", MagentaBg},
			{"46", CyanBg},
			{"47", WhiteBg},

			{"100", BlackBg | BrightBg},
			{"101", RedBg | BrightBg},
			{"102", GreenBg | BrightBg},
			{"103", YellowBg | BrightBg},
			{"104", BlueBg | BrightBg},
			{"105", MagentaBg | BrightBg},
			{"106", CyanBg | BrightBg},
			{"107", WhiteBg | BrightBg},

			{"100", BrightBg},

			// bold and faint

			{"1", BoldFm | FaintFm},

			// slow blink and rapid blink

			{"5", SlowBlinkFm | RapidBlinkFm},

			// index

			{"38;5;100", (100 << shiftFg) | flagFg},
			{"48;5;100", (100 << shiftBg) | flagBg},

			// longest combination

			{"1;3;4;5;7;8;9;20;21;51;52;53;38;5;123;48;5;231",
				BoldFm | FaintFm |
					ItalicFm | UnderlineFm |
					SlowBlinkFm | RapidBlinkFm |
					ReverseFm | ConcealFm |
					CrossedOutFm | FrakturFm | DoublyUnderlineFm |
					FramedFm | EncircledFm | OverlinedFm |
					Color(123)<<shiftFg | flagFg |
					Color(231)<<shiftBg | flagBg},
		} {
			var (
				nos  = val.Color.Nos(zero)
				want = val.Nos
			)
			if zero {
				if want != "" {
					want = "0;" + want
				} else {
					want = "0"
				}
			}
			assert.Equalf(t, want, nos, "%t %d: wrong nos string %q, want %q",
				zero, i, nos, want)
		}
	}
}

func Test_itoa(t *testing.T) {
	for i := 0; i < 256; i++ {
		var a = itoa(byte(i))
		assert.Equalf(t, a, strconv.Itoa(i), "wrong %q, want %d", a, i)
	}
}

func TestColor_Reset(t *testing.T) {
	var c Color = math.MaxUint // all bits set to 1
	assert.Zero(t, c.Reset())
}

func TestColor_Bold(t *testing.T) {
	assert.True(t, Color(0).Bold()&BoldFm != 0, "not a bold")
	assert.True(t, Color(FaintFm).Bold()&FaintFm == 0, "contains faint")
}

func TestColor_Faint(t *testing.T) {
	assert.True(t, Color(0).Faint()&FaintFm != 0, "not a faint")
	assert.True(t, Color(BoldFm).Faint()&BoldFm == 0, "contains bold")
}

func TestColor_DoublyUnderline(t *testing.T) {
	assert.True(t, Color(0).DoublyUnderline()&DoublyUnderlineFm != 0,
		"not a doubly-underlined")
	assert.True(t, Color(UnderlineFm).DoublyUnderline()&UnderlineFm == 0,
		"contains underline")
}

func TestColor_Fraktur(t *testing.T) {
	assert.True(t, Color(0).Fraktur()&FrakturFm != 0, "not a fraktur")
}

func TestColor_Italic(t *testing.T) {
	assert.True(t, Color(0).Italic()&ItalicFm != 0, "not a italic")
}

func TestColor_Underline(t *testing.T) {
	assert.True(t, Color(0).Underline()&UnderlineFm != 0, "not a underlined")
}

func TestColor_SlowBlink(t *testing.T) {
	assert.True(t, Color(0).SlowBlink()&SlowBlinkFm != 0, "not a slow blinking")
}

func TestColor_RapidBlink(t *testing.T) {
	assert.True(t, Color(0).RapidBlink()&RapidBlinkFm != 0,
		"not a rapid blinking")
}

func TestColor_Blink(t *testing.T) {
	assert.True(t, Color(0).Blink()&BlinkFm != 0, "not a blinking")
}

func TestColor_Reverse(t *testing.T) {
	assert.True(t, Color(0).Reverse()&ReverseFm != 0, "not a reversed")
}

func TestColor_Inverse(t *testing.T) {
	assert.True(t, Color(0).Inverse()&InverseFm != 0, "not a inversed")
}

func TestColor_Conceal(t *testing.T) {
	assert.True(t, Color(0).Conceal()&ConcealFm != 0, "not a concealed")
}

func TestColor_Hidden(t *testing.T) {
	assert.True(t, Color(0).Hidden()&HiddenFm != 0, "not a hidden")
}

func TestColor_CrossedOut(t *testing.T) {
	assert.True(t, Color(0).CrossedOut()&CrossedOutFm != 0, "not a crossed out")
}

func TestColor_StrikeThrough(t *testing.T) {
	assert.True(t, Color(0).StrikeThrough()&StrikeThroughFm != 0,
		"not a striked through")
}

func TestColor_Framed(t *testing.T) {
	assert.True(t, Color(0).Framed()&FramedFm != 0, "not a framed")
}

func TestColor_Encircled(t *testing.T) {
	assert.True(t, Color(0).Encircled()&EncircledFm != 0, "not a encircled")
}

func TestColor_Overlined(t *testing.T) {
	assert.True(t, Color(0).Overlined()&OverlinedFm != 0, "not a overlined")
}

func TestColor_Black(t *testing.T) {
	assert.True(t, Color(0).Black()&BlackFg != 0, "not a black")
	assert.True(t, Color(RedFg).Black()&RedFg == flagFg, "contains red")
}

func TestColor_Red(t *testing.T) {
	assert.True(t, Color(0).Red()&RedFg != 0, "not a red")
	assert.True(t, Color(BlackFg).Red()&BlackFg == flagFg, "contains black")
}

func TestColor_Green(t *testing.T) {
	assert.True(t, Color(0).Green()&GreenFg != 0, "not a green")
	assert.True(t, Color(BlackFg).Green()&BlackFg == flagFg, "contains black")
}

func TestColor_Yellow(t *testing.T) {
	assert.True(t, Color(0).Yellow()&YellowFg != 0, "not a yellow")
	assert.True(t, Color(BlackFg).Yellow()&BlackFg == flagFg, "contains black")
}

func TestColor_Brown(t *testing.T) {
	assert.True(t, Color(0).Brown()&BrownFg != 0, "not a brown")
	assert.True(t, Color(BlackFg).Brown()&BlackFg == flagFg, "contains black")
}

func TestColor_Blue(t *testing.T) {
	assert.True(t, Color(0).Blue()&BlueFg != 0, "not a blue")
	assert.True(t, Color(BlackFg).Blue()&BlackFg == flagFg, "contains black")
}

func TestColor_Magenta(t *testing.T) {
	assert.True(t, Color(0).Magenta()&MagentaFg != 0, "not a magenta")
	assert.True(t, Color(BlackFg).Magenta()&BlackFg == flagFg, "contains black")
}

func TestColor_Cyan(t *testing.T) {
	assert.True(t, Color(0).Cyan()&CyanFg != 0, "not a cyan")
	assert.True(t, Color(BlackFg).Cyan()&BlackFg == flagFg, "contains black")
}

func TestColor_White(t *testing.T) {
	assert.True(t, Color(0).White()&WhiteFg != 0, "not a white")
	assert.True(t, Color(BlackFg).White()&BlackFg == flagFg, "contains black")
}

func TestColor_BrightBlack(t *testing.T) {
	assert.True(t, Color(0).BrightBlack()&(BrightFg|BlackFg) != 0,
		"not a bright black")
	assert.True(t, Color(RedFg).BrightBlack()&RedFg == flagFg, "contains red")
}

func TestColor_BrightRed(t *testing.T) {
	assert.True(t, Color(0).BrightRed()&(BrightFg|RedFg) != 0,
		"not a bright red")
	assert.True(t, Color(BlackFg).BrightRed()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightGreen(t *testing.T) {
	assert.True(t, Color(0).BrightGreen()&(BrightFg|GreenFg) != 0,
		"not a bright green")
	assert.True(t, Color(BlackFg).BrightGreen()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightYellow(t *testing.T) {
	assert.True(t, Color(0).BrightYellow()&(BrightFg|YellowFg) != 0,
		"not a bright yellow")
	assert.True(t, Color(BlackFg).BrightYellow()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightBlue(t *testing.T) {
	assert.True(t, Color(0).BrightBlue()&(BrightFg|BlueFg) != 0,
		"not a bright blue")
	assert.True(t, Color(BlackFg).BrightBlue()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightMagenta(t *testing.T) {
	assert.True(t, Color(0).BrightMagenta()&(BrightFg|MagentaFg) != 0,
		"not a bright blue")
	assert.True(t, Color(BlackFg).BrightMagenta()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightCyan(t *testing.T) {
	assert.True(t, Color(0).BrightCyan()&(BrightFg|CyanFg) != 0,
		"not a bright cyan")
	assert.True(t, Color(BlackFg).BrightCyan()&BlackFg == flagFg,
		"contains black")
}

func TestColor_BrightWhite(t *testing.T) {
	assert.True(t, Color(0).BrightWhite()&(BrightFg|WhiteFg) != 0,
		"not a bright white")
	assert.True(t, Color(BlackFg).BrightWhite()&BlackFg == flagFg,
		"contains black")
}

func TestColor_Index(t *testing.T) {
	for i := 0; i < 256; i++ {
		var ci = ColorIndex(i)
		assert.Truef(t, Color(0).Index(ci)&flagFg != 0,
			"missing indexed color, color index %d", i)
		assert.Truef(t, Color(BlackFg).Index(ci)&BlackFg == flagFg,
			"contains black, color index %d", i)
	}
}

func TestColor_Gray(t *testing.T) {
	for i := GrayIndex(0); i < 25; i++ {
		assert.Truef(t, Color(0).Gray(i)&flagFg != 0,
			"missing indexed gray color, gray index %d", i)
		assert.Truef(t, Color(BlackFg).Gray(i)&BlackFg == flagFg,
			"contains black, gray index %d", i)
	}
}

func TestColor_BgBlack(t *testing.T) {
	assert.True(t, Color(0).BgBlack()&BlackBg != 0, "not a black background")
	assert.True(t, Color(RedBg).BgBlack()&RedBg == flagBg,
		"contains red background")
}

func TestColor_BgRed(t *testing.T) {
	assert.True(t, Color(0).BgRed()&RedBg != 0, "not a red background")
	assert.True(t, Color(BlackBg).BgRed()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgGreen(t *testing.T) {
	assert.True(t, Color(0).BgGreen()&GreenBg != 0, "not a green background")
	assert.True(t, Color(BlackBg).BgGreen()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgYellow(t *testing.T) {
	assert.True(t, Color(0).BgYellow()&YellowBg != 0, "not a yellow background")
	assert.True(t, Color(BlackBg).BgYellow()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrown(t *testing.T) {
	assert.True(t, Color(0).BgBrown()&BrownBg != 0, "not a brown background")
	assert.True(t, Color(BlackBg).BgBrown()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBlue(t *testing.T) {
	assert.True(t, Color(0).BgBlue()&BlueBg != 0, "not a blue background")
	assert.True(t, Color(BlackBg).BgBlue()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgMagenta(t *testing.T) {
	assert.True(t, Color(0).BgMagenta()&MagentaBg != 0,
		"not a magenta background")
	assert.True(t, Color(BlackBg).BgMagenta()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgCyan(t *testing.T) {
	assert.True(t, Color(0).BgCyan()&CyanBg != 0, "not a cyan background")
	assert.True(t, Color(BlackBg).BgCyan()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgWhite(t *testing.T) {
	assert.True(t, Color(0).BgWhite()&WhiteBg != 0, "not a white background")
	assert.True(t, Color(BlackBg).BgWhite()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightBlack(t *testing.T) {
	assert.True(t, Color(0).BgBrightBlack()&(BrightBg|BlackBg) != 0,
		"not a bright black background")
	assert.True(t, Color(RedBg).BgBrightBlack()&RedBg == flagBg,
		"contains red background")
}

func TestColor_BgBrightRed(t *testing.T) {
	assert.True(t, Color(0).BgBrightRed()&(BrightBg|RedBg) != 0,
		"not a bright red background")
	assert.True(t, Color(BlackBg).BgBrightRed()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightGreen(t *testing.T) {
	assert.True(t, Color(0).BgBrightGreen()&(BrightBg|GreenBg) != 0,
		"not a bright green background")
	assert.True(t, Color(BlackBg).BgBrightGreen()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightYellow(t *testing.T) {
	assert.True(t, Color(0).BgBrightYellow()&(BrightBg|YellowBg) != 0,
		"not a bright yellow background")
	assert.True(t, Color(BlackBg).BgBrightYellow()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightBlue(t *testing.T) {
	assert.True(t, Color(0).BgBrightBlue()&(BrightBg|BlueBg) != 0,
		"not a bright blue background")
	assert.True(t, Color(BlackBg).BgBrightBlue()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightMagenta(t *testing.T) {
	assert.True(t, Color(0).BgBrightMagenta()&(BrightBg|MagentaBg) != 0,
		"not a bright magenta background")
	assert.True(t, Color(BlackBg).BgBrightMagenta()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightCyan(t *testing.T) {
	assert.True(t, Color(0).BgBrightCyan()&(BrightBg|CyanBg) != 0,
		"not a bright cyan background")
	assert.True(t, Color(BlackBg).BgBrightCyan()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgBrightWhite(t *testing.T) {
	assert.True(t, Color(0).BgBrightWhite()&(BrightBg|WhiteBg) != 0,
		"not a bright white background")
	assert.True(t, Color(BlackBg).BgBrightWhite()&BlackBg == flagBg,
		"contains black background")
}

func TestColor_BgIndex(t *testing.T) {
	for i := 0; i < 256; i++ {
		var ci = ColorIndex(i)
		assert.True(t, Color(0).BgIndex(ci)&(flagBg) != 0,
			"missing indexed background color")
		assert.True(t, Color(BlackBg).BgIndex(ci)&BlackBg == flagBg,
			"contains black background")
	}
}

func TestColor_BgGray(t *testing.T) {
	for i := GrayIndex(0); i <= 25; i++ {
		assert.Truef(t, Color(0).BgGray(i)&(flagBg) != 0,
			"missing indexed gray background color, gray index %d", i)
		assert.Truef(t, Color(BlackBg).BgGray(i)&BlackBg == flagBg,
			"contains black background, gray index %d", i)
	}
}
