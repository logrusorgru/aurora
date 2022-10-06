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
	"unicode/utf8"
)

// compile-time check
var (
	_ fmt.Stringer  = Value{}
	_ fmt.Formatter = Value{}
	_ Colored       = Value{}
)

func coloredFormat(color Color, s fmt.State, verb rune) string {
	// it's enough for many cases (%-+020.10f)
	// %          - 1
	// availFlags - 3 (5)
	// width      - 2
	// prec       - 3 (.23)
	// verb       - 1
	// --------------
	//             10
	// +
	// \033[                            5
	// 0;1;3;4;5;7;8;9;20;21;51;52;53  30
	// 38;5;216                         8
	// 48;5;216                         8
	// m                                1
	// +
	// \033[0m                          7
	//
	// x2 (possible tail color)
	//
	// 10 + 59 * 2 = 128

	var format = make([]byte, 0, 128)

	if color != 0 {
		format = append(format, esc...)
		format = color.appendNos(format, false)
		format = append(format, 'm')
	}

	format = append(format, '%')

	var f byte
	for i := 0; i < len(availFlags); i++ {
		if f = availFlags[i]; s.Flag(int(f)) {
			format = append(format, f)
		}
	}

	var (
		width, prec int
		ok          bool
	)
	if width, ok = s.Width(); ok {
		format = strconv.AppendInt(format, int64(width), 10)
	}

	if prec, ok = s.Precision(); ok {
		format = append(format, '.')
		format = strconv.AppendInt(format, int64(prec), 10)
	}

	if verb > utf8.RuneSelf {
		format = append(format, string(verb)...)
	} else {
		format = append(format, byte(verb))
	}

	if color != 0 {
		format = append(format, clear...) // just clear
	}

	return string(format)
}

// A Value represents any printable value
// with or without colors, formats and a link.
type Value struct {
	conf      *Config     // back reference to configurations
	value     interface{} // value as is
	color     Color       // color of the value
	hyperlink             // hyperlink target and parameters
}

// String implements standard fmt.Stringer interface.
func (v Value) String() string {

	var (
		t   []byte
		val = fmt.Sprint(v.value)
	)

	if v.target != "" {
		var (
			ln  = len(val)
			nos string
		)
		// calculate length
		ln += v.hyperlink.headLen()
		if v.conf.Colors && v.color != 0 {
			ln += len(esc)
			nos = v.color.Nos(false)
			ln += len(nos) + len("m") + len(clear)
		}
		ln += v.hyperlink.tailLen()
		// fill
		t = make([]byte, 0, ln)
		t = append(t, v.hyperlink.headBytes()...)
		if v.conf.Colors && v.color != 0 {
			t = append(t, esc...)
			t = append(t, nos...)
			t = append(t, 'm')
			t = append(t, val...)
			t = append(t, clear...)
		} else {
			t = append(t, val...)
		}
		t = append(t, v.hyperlink.tailBytes()...)
		return string(t)
	}

	// no links, only colors & formats
	if v.conf.Colors && v.color != 0 {
		return esc + v.color.Nos(false) + "m" + val + clear
	}

	// no links, no colors, no formats, just the value
	return val
}

// Color returns colors and formats of the Value.
func (v Value) Color() Color {
	return v.color
}

// Bleach returns copy of original value without colors
//
// Deprecated: use Reset instead.
func (v Value) Bleach() Value {
	return v.Reset()
}

// Reset colors, formats and links.
func (v Value) Reset() Value {
	v.color, v.hyperlink = 0, hyperlink{}
	return v
}

// Clear colors and formats, preserving links.
func (v Value) Clear() Value {
	v.color = 0
	return v
}

// Value returns value's value (welcome to the tautology club)
func (v Value) Value() interface{} {
	return v.value
}

// Format implements standard fmt.Formatter interface.
func (v Value) Format(s fmt.State, verb rune) {
	v.hyperlink.writeHead(s)
	fmt.Fprintf(s, coloredFormat(v.color, s, verb), v.value)
	v.hyperlink.writeTail(s)
}

// Formats
//
// Bold or increased intensity (1).
func (v Value) Bold() Value {
	v.color = v.color.Bold()
	return v
}

// Faint, decreased intensity, reset the Bold (2).
func (v Value) Faint() Value {
	v.color = v.color.Faint()
	return v
}

// DoublyUnderline or Bold off, double-underline per ECMA-48 (21). It depends.
func (v Value) DoublyUnderline() Value {
	v.color = v.color.DoublyUnderline()
	return v
}

// Fraktur, rarely supported (20).
func (v Value) Fraktur() Value {
	v.color = v.color.Fraktur()
	return v
}

// Italic, not widely supported, sometimes treated as inverse (3).
func (v Value) Italic() Value {
	v.color = v.color.Italic()
	return v
}

// Underline (4).
func (v Value) Underline() Value {
	v.color = v.color.Underline()
	return v
}

// SlowBlink, blinking less than 150 per minute (5).
func (v Value) SlowBlink() Value {
	v.color = v.color.SlowBlink()
	return v
}

// RapidBlink, blinking 150+ per minute, not widely supported (6).
func (v Value) RapidBlink() Value {
	v.color = v.color.RapidBlink()
	return v
}

// Blink is alias for the SlowBlink.
func (v Value) Blink() Value {
	return v.SlowBlink()
}

// Reverse video, swap foreground and background colors (7).
func (v Value) Reverse() Value {
	v.color = v.color.Reverse()
	return v
}

// Inverse is alias for the Reverse.
func (v Value) Inverse() Value {
	return v.Reverse()
}

// Conceal, hidden, not widely supported (8).
func (v Value) Conceal() Value {
	v.color = v.color.Conceal()
	return v
}

// Hidden is alias for the Conceal.
func (v Value) Hidden() Value {
	return v.Conceal()
}

// CrossedOut, characters legible, but marked for deletion (9).
func (v Value) CrossedOut() Value {
	v.color = v.color.CrossedOut()
	return v
}

// StrikeThrough is alias for the CrossedOut.
func (v Value) StrikeThrough() Value {
	return v.CrossedOut()
}

// Framed (51).
func (v Value) Framed() Value {
	v.color = v.color.Framed()
	return v
}

// Encircled (52).
func (v Value) Encircled() Value {
	v.color = v.color.Encircled()
	return v
}

// Overlined (53).
func (v Value) Overlined() Value {
	v.color = v.color.Overlined()
	return v
}

// Foreground colors.
//
// Black foreground color (30).
func (v Value) Black() Value {
	v.color = v.color.Black()
	return v
}

// Red foreground color (31).
func (v Value) Red() Value {
	v.color = v.color.Red()
	return v
}

// Green foreground color (32).
func (v Value) Green() Value {
	v.color = v.color.Green()
	return v
}

// Yellow foreground color (33).
func (v Value) Yellow() Value {
	v.color = v.color.Yellow()
	return v
}

// Brown foreground color (33)
//
// Deprecated: use Yellow instead, following specification.
func (v Value) Brown() Value {
	return v.Yellow()
}

// Blue foreground color (34).
func (v Value) Blue() Value {
	v.color = v.color.Blue()
	return v
}

// Magenta foreground color (35).
func (v Value) Magenta() Value {
	v.color = v.color.Magenta()
	return v
}

// Cyan foreground color (36).
func (v Value) Cyan() Value {
	v.color = v.color.Cyan()
	return v
}

// White foreground color (37).
func (v Value) White() Value {
	v.color = v.color.White()
	return v
}

// Bright foreground colors.
//
// BrightBlack foreground color (90).
func (v Value) BrightBlack() Value {
	v.color = v.color.BrightBlack()
	return v
}

// BrightRed foreground color (91).
func (v Value) BrightRed() Value {
	v.color = v.color.BrightRed()
	return v
}

// BrightGreen foreground color (92).
func (v Value) BrightGreen() Value {
	v.color = v.color.BrightGreen()
	return v
}

// BrightYellow foreground color (93).
func (v Value) BrightYellow() Value {
	v.color = v.color.BrightYellow()
	return v
}

// BrightBlue foreground color (94).
func (v Value) BrightBlue() Value {
	v.color = v.color.BrightBlue()
	return v
}

// BrightMagenta foreground color (95).
func (v Value) BrightMagenta() Value {
	v.color = v.color.BrightMagenta()
	return v
}

// BrightCyan foreground color (96).
func (v Value) BrightCyan() Value {
	v.color = v.color.BrightCyan()
	return v
}

// BrightWhite foreground color (97).
func (v Value) BrightWhite() Value {
	v.color = v.color.BrightWhite()
	return v
}

// Other colors.
//
// Index of pre-defined 8-bit foreground color from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (v Value) Index(n ColorIndex) Value {
	v.color = v.color.Index(n)
	return v
}

// Gray from 0 to 24.
func (v Value) Gray(n GrayIndex) Value {
	v.color = v.color.Gray(n)
	return v
}

// Background colors
//
// BgBlack background color (40).
func (v Value) BgBlack() Value {
	v.color = v.color.BgBlack()
	return v
}

// BgRed background color (41).
func (v Value) BgRed() Value {
	v.color = v.color.BgRed()
	return v
}

// BgGreen background color (42).
func (v Value) BgGreen() Value {
	v.color = v.color.BgGreen()
	return v
}

// BgYellow background color (43).
func (v Value) BgYellow() Value {
	v.color = v.color.BgYellow()
	return v
}

// BgBrown background color (43).
//
// Deprecated: use BgYellow instead, following specification.
func (v Value) BgBrown() Value {
	return v.BgYellow()
}

// BgBlue background color (44).
func (v Value) BgBlue() Value {
	v.color = v.color.BgBlue()
	return v
}

// BgMagenta background color (45).
func (v Value) BgMagenta() Value {
	v.color = v.color.BgMagenta()
	return v
}

// BgCyan background color (46).
func (v Value) BgCyan() Value {
	v.color = v.color.BgCyan()
	return v
}

// BgWhite background color (47).
func (v Value) BgWhite() Value {
	v.color = v.color.BgWhite()
	return v
}

// Bright background colors.
//
// BgBrightBlack background color (100).
func (v Value) BgBrightBlack() Value {
	v.color = v.color.BgBrightBlack()
	return v
}

// BgBrightRed background color (101).
func (v Value) BgBrightRed() Value {
	v.color = v.color.BgBrightRed()
	return v
}

// BgBrightGreen background color (102).
func (v Value) BgBrightGreen() Value {
	v.color = v.color.BgBrightGreen()
	return v
}

// BgBrightYellow background color (103).
func (v Value) BgBrightYellow() Value {
	v.color = v.color.BgBrightYellow()
	return v
}

// BgBrightBlue background color (104).
func (v Value) BgBrightBlue() Value {
	v.color = v.color.BgBrightBlue()
	return v
}

// BgBrightMagenta background color (105).
func (v Value) BgBrightMagenta() Value {
	v.color = v.color.BgBrightMagenta()
	return v
}

// BgBrightCyan background color (106).
func (v Value) BgBrightCyan() Value {
	v.color = v.color.BgBrightCyan()
	return v
}

// BgBrightWhite background color (107).
func (v Value) BgBrightWhite() Value {
	v.color = v.color.BgBrightWhite()
	return v
}

// Other background colors.
//
// BgIndex of 8-bit pre-defined background color from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (v Value) BgIndex(n ColorIndex) Value {
	v.color = v.color.BgIndex(n)
	return v
}

// BgGray from 0 to 24.
func (v Value) BgGray(n GrayIndex) Value {
	v.color = v.color.BgGray(n)
	return v
}

// Special colorization method.
//
// Colorize removes existing colors and formats of the argument and applies
// given.
func (v Value) Colorize(color Color) Value {
	v.color = color
	return v
}

// Hyperlinks feature
//
// Hyperlink with given target and parameters. If hyperlinks feature is
// disabled, then the 'arg' argument dropped and the 'target' used instead,
// inheriting all colors and format from the 'arg' (if it's a Colored).
//
// See https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
// for details about the hyperlinks feature.
//
// The Hyperlink doesn't escape the target and the parameters. They should be
// checked and escaped before. See HyperlinkEscape function.
//
// See also HyperlinkID function.
//
// For a simple example
//
//	val.Hyperlink("http://example.com")
//
// and an example with ID
//
//	val.Hyperlink("http://example.com", aurora.HyperlinkID("10"))
//
// Successive calls replace previously set target and parameters.
func (v Value) Hyperlink(target string, params ...HyperlinkParam) Value {
	if !v.conf.Hyperlinks {
		v.value = target            // drop value, use the target
		v.hyperlink.target = target // keep for the HyperlinkTarget method
		return v
	}
	v.hyperlink.target = target
	v.hyperlink.params = params
	return v
}

// HyperlinkTarget if any.
func (v Value) HyperlinkTarget() (target string) {
	return v.hyperlink.target
}

// HyperlinkParams if any.
func (v Value) HyperlinkParams() []HyperlinkParam {
	return v.hyperlink.params
}
