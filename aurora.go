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

// Package aurora implements ANSI-colors
package aurora

type Aurora struct {
	conf Config
}

// New returns new colorizer by given Options.
func New(opts ...Option) (a Aurora) {
	a.conf.Apply(opts...)
	return
}

// Reset wraps given argument returning Value without formats, colors and links.
func (a Aurora) Reset(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Reset()
	}
	return Value{conf: &a.conf, value: arg}
}

// Clear wraps given argument returning Value without formats and colors. But
// preserving links.
func (a Aurora) Clear(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Reset()
	}
	return Value{conf: &a.conf, value: arg}
}

// Formats
//
// Bold or increased intensity (1).
func (a Aurora) Bold(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Bold()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Bold()}
}

// Faint, decreased intensity (2).
func (a Aurora) Faint(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Faint()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Faint()}
}

// DoublyUnderline or Bold off, double-underline per ECMA-48 (21).
func (a Aurora) DoublyUnderline(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.DoublyUnderline()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).DoublyUnderline()}
}

// Fraktur, rarely supported (20).
func (a Aurora) Fraktur(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Fraktur()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Fraktur()}
}

// Italic, not widely supported, sometimes treated as inverse (3).
func (a Aurora) Italic(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Italic()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Italic()}
}

// Underline (4).
func (a Aurora) Underline(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Underline()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Underline()}
}

// SlowBlink, blinking less than 150 per minute (5).
func (a Aurora) SlowBlink(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.SlowBlink()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).SlowBlink()}
}

// RapidBlink, blinking 150+ per minute, not widely supported (6).
func (a Aurora) RapidBlink(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.RapidBlink()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).RapidBlink()}
}

// Blink is alias for the SlowBlink.
func (a Aurora) Blink(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Blink()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Blink()}
}

// Reverse video, swap foreground and background colors (7).
func (a Aurora) Reverse(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Reverse()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Reverse()}
}

// Inverse is alias for the Reverse
func (a Aurora) Inverse(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Inverse()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Inverse()}
}

// Conceal, hidden, not widely supported (8).
func (a Aurora) Conceal(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Conceal()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Conceal()}
}

// Hidden is alias for the Conceal.
func (a Aurora) Hidden(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Hidden()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Hidden()}
}

// CrossedOut, characters legible, but marked for deletion (9).
func (a Aurora) CrossedOut(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.CrossedOut()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).CrossedOut()}
}

// StrikeThrough is alias for the CrossedOut.
func (a Aurora) StrikeThrough(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.StrikeThrough()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).StrikeThrough()}
}

// Framed (51).
func (a Aurora) Framed(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Framed()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Framed()}
}

// Encircled (52).
func (a Aurora) Encircled(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Encircled()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Encircled()}
}

// Overlined (53).
func (a Aurora) Overlined(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Overlined()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Overlined()}
}

// Foreground colors
//
// Black foreground color (30).
func (a Aurora) Black(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Black()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Black()}
}

// Red foreground color (31).
func (a Aurora) Red(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Red()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Red()}
}

// Green foreground color (32).
func (a Aurora) Green(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Green()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Green()}
}

// Yellow foreground color (33).
func (a Aurora) Yellow(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Yellow()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Yellow()}
}

// Brown foreground color (33).
//
// Deprecated: use Yellow instead, following specification.
func (a Aurora) Brown(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Brown()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Brown()}
}

// Blue foreground color (34).
func (a Aurora) Blue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Blue()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Blue()}
}

// Magenta foreground color (35).
func (a Aurora) Magenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Magenta()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Magenta()}
}

// Cyan foreground color (36).
func (a Aurora) Cyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Cyan()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Cyan()}
}

// White foreground color (37).
func (a Aurora) White(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.White()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).White()}
}

// Bright foreground colors.
//
// BrightBlack foreground color (90).
func (a Aurora) BrightBlack(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightBlack()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightBlack()}
}

// BrightRed foreground color (91).
func (a Aurora) BrightRed(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightRed()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightRed()}
}

// BrightGreen foreground color (92).
func (a Aurora) BrightGreen(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightGreen()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightGreen()}
}

// BrightYellow foreground color (93).
func (a Aurora) BrightYellow(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightYellow()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightYellow()}
}

// BrightBlue foreground color (94).
func (a Aurora) BrightBlue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightBlue()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightBlue()}
}

// BrightMagenta foreground color (95).
func (a Aurora) BrightMagenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightMagenta()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightMagenta()}
}

// BrightCyan foreground color (96).
func (a Aurora) BrightCyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightCyan()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightCyan()}
}

// BrightWhite foreground color (97).
func (a Aurora) BrightWhite(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BrightWhite()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BrightWhite()}
}

// Other colors.
//
// Index of pre-defined 8-bit foreground color from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (a Aurora) Index(n ColorIndex, arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Index(n)
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Index(n)}
}

// Gray from 0 to 23.
func (a Aurora) Gray(n GrayIndex, arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.Gray(n)
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).Gray(n)}
}

// Background colors.
//
// BgBlack background color (40).
func (a Aurora) BgBlack(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBlack()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBlack()}
}

// BgRed background color (41).
func (a Aurora) BgRed(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgRed()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgRed()}
}

// BgGreen background color (42).
func (a Aurora) BgGreen(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgGreen()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgGreen()}
}

// BgYellow background color (43).
func (a Aurora) BgYellow(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgYellow()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgYellow()}
}

// BgBrown background color (43).
//
// Deprecated: use BgYellow instead, following specification.
func (a Aurora) BgBrown(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrown()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrown()}
}

// BgBlue background color (44).
func (a Aurora) BgBlue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBlue()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBlue()}
}

// BgMagenta background color (45).
func (a Aurora) BgMagenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgMagenta()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgMagenta()}
}

// BgCyan background color (46).
func (a Aurora) BgCyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgCyan()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgCyan()}
}

// BgWhite background color (47).
func (a Aurora) BgWhite(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgWhite()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgWhite()}
}

// Bright background colors.
//
// BgBrightBlack background color (100).
func (a Aurora) BgBrightBlack(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightBlack()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightBlack()}
}

// BgBrightRed background color (101).
func (a Aurora) BgBrightRed(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightRed()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightRed()}
}

// BgBrightGreen background color (102).
func (a Aurora) BgBrightGreen(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightGreen()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightGreen()}
}

// BgBrightYellow background color (103).
func (a Aurora) BgBrightYellow(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightYellow()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightYellow()}
}

// BgBrightBlue background color (104).
func (a Aurora) BgBrightBlue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightBlue()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightBlue()}
}

// BgBrightMagenta background color (105).
func (a Aurora) BgBrightMagenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightMagenta()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightMagenta()}
}

// BgBrightCyan background color (106).
func (a Aurora) BgBrightCyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightCyan()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightCyan()}
}

// BgBrightWhite background color (107).
func (a Aurora) BgBrightWhite(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgBrightWhite()
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgBrightWhite()}
}

// Other background colors.
//
// BgIndex of 8-bit pre-defined background color from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (a Aurora) BgIndex(n ColorIndex, arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgIndex(n)
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgIndex(n)}
}

// BgGray from 0 to 23.
func (a Aurora) BgGray(n GrayIndex, arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		return val.BgGray(n)
	}
	return Value{conf: &a.conf, value: arg, color: Color(0).BgGray(n)}
}

// Special color functions.
//
// Colorize removes existing colors and
// formats of the argument and applies given.
func (a Aurora) Colorize(arg interface{}, color Color) Value {
	if val, ok := arg.(Value); ok {
		return val.Colorize(color)
	}
	return Value{conf: &a.conf, value: arg, color: color}
}

// Hyperlinks feature
//
// Hyperlink with given target and parameters. If hyperlinks feature is
// disabled, then the 'arg' argument dropped and the 'target' used instead
// inheriting all colors and format from the 'arg' (if it's a Colored).
//
// See https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
// for details about the hyperlinks feature.
//
// The Hyperlink doesn't escape the target and the params. They should be
// checked and escaped before.
//
// See also HyperlinkID function.
//
// For a simple example
//
//	au.Hyperlink("Example", "http://example.com")
//
// and an example with ID
//
//	au.Hyperlink("Example", "http://example.com", aurora.HyperlinkID("10"))
func (a Aurora) Hyperlink(arg interface{}, target string,
	params ...HyperlinkParam) Value {

	if val, ok := arg.(Value); ok {
		return val.Hyperlink(target, params...)
	}
	return Value{conf: &a.conf, value: arg}.Hyperlink(target, params...)
}

// HyperlinkTarget of the argument if it's a Value.
func (a Aurora) HyperlinkTarget(arg interface{}) (target string) {
	if val, ok := arg.(Value); ok {
		return val.HyperlinkTarget()
	}
	return // no target
}

// HyperlinkParams of the argument if it's a Value.
func (a Aurora) HyperlinkParams(arg interface{}) (params []HyperlinkParam) {
	if val, ok := arg.(Value); ok {
		return val.HyperlinkParams()
	}
	return // no target
}

func (a Aurora) transform(arg interface{}) (val Value, ok bool) {
	var ai Value
	ai, ok = arg.(Value)
	if ok {
		return // Value{}, false
	}
	val = Value{conf: &a.conf, value: ai.value, color: ai.color}
	if a.conf.Hyperlinks {
		val.hyperlink = ai.hyperlink
	}
	return // transformed value, true
}

// Support methods.
//
// Sprintf allows to use colored format. It allies own configurations to
// all given values (if there is a Value).
func (a Aurora) Sprintf(format interface{}, args ...interface{}) string {
	// clear colors & links as configured by the a
	if f, ok := a.transform(format); ok {
		format = f
	}
	for i := range args {
		if ax, ok := a.transform(args[i]); ok {
			args[i] = ax
		}
	}
	return Sprintf(format, args...)
}
