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

var (
	gVal  Value
	gVals []Value
	gStr  string
)

var (
	short  = "x"
	long   = "some long and complicated string to colorize"
	shortf = "%s"
	longf  = long + " %s"
	s18    = "%s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s"
)

// 18 values
func simpleValues(a Aurora, x string) []Value {
	return []Value{
		a.Reset("x"),

		a.Bold("x"),
		a.Faint("x"),
		a.DoublyUnderline("x"),
		a.Fraktur("x"),
		a.Italic("x"),
		a.Underline("x"),
		a.SlowBlink("x"),
		a.RapidBlink("x"),
		a.Blink("x"),
		a.Reverse("x"),
		a.Inverse("x"),
		a.Conceal("x"),
		a.Hidden("x"),
		a.CrossedOut("x"),
		a.StrikeThrough("x"),
		a.Framed("x"),
		a.Encircled("x"),
		a.Overlined("x"),

		a.Black("x"),
		a.Red("x"),
		a.Green("x"),
		a.Yellow("x"),
		a.Brown("x"),
		a.Blue("x"),
		a.Magenta("x"),
		a.Cyan("x"),
		a.White("x"),
		a.BrightBlack("x"),
		a.BrightRed("x"),
		a.BrightGreen("x"),
		a.BrightYellow("x"),
		a.BrightBlue("x"),
		a.BrightMagenta("x"),
		a.BrightCyan("x"),
		a.BrightWhite("x"),
		a.Index(178, "x"),
		a.Gray(14, "x"),

		a.BgBlack("x"),
		a.BgRed("x"),
		a.BgGreen("x"),
		a.BgYellow("x"),
		a.BgBrown("x"),
		a.BgBlue("x"),
		a.BgMagenta("x"),
		a.BgCyan("x"),
		a.BgWhite("x"),
		a.BgBrightBlack("x"),
		a.BgBrightRed("x"),
		a.BgBrightGreen("x"),
		a.BgBrightYellow("x"),
		a.BgBrightBlue("x"),
		a.BgBrightMagenta("x"),
		a.BgBrightCyan("x"),
		a.BgBrightWhite("x"),
		a.BgIndex(187, "x"),
		a.BgGray(15, "x"),
	}
}

// 18 values
func complexValues(a Aurora, x string) []Value {

	var allFormats = func(val Value) Value {
		return val.Bold().DoublyUnderline().Fraktur().Blink().Italic().
			Underline().Blink().Reverse().Conceal().CrossedOut().Framed().
			Encircled().Overlined()

	}

	return []Value{
		allFormats(a.Reset(x).BrightBlack().BgBrightWhite()),

		allFormats(a.Bold(x).BrightRed().BgBrightMagenta()),
		allFormats(a.Faint(x).Index(128).BgIndex(135)),
		allFormats(a.DoublyUnderline(x).Gray(13).BgRed()),
		allFormats(a.Fraktur(x).BgYellow().Black()),
		allFormats(a.Italic(x).Red().BgBrightMagenta()),
		allFormats(a.Underline(x).BrightBlue().BgYellow()),
		allFormats(a.SlowBlink(x).BrightCyan().BgBrightRed()),
		allFormats(a.RapidBlink(x).Green().BgYellow()),
		allFormats(a.Blink(x).Red().BgBlack()),
		allFormats(a.Reverse(x).BgBlack().White()),
		allFormats(a.Inverse(x).White().BgBlack()),
		allFormats(a.Conceal(x).BrightBlue().BgYellow()),
		allFormats(a.Hidden(x).BrightYellow().BgBrightGreen()),
		allFormats(a.CrossedOut(x).BrightGreen()),
		allFormats(a.StrikeThrough(x).Index(128)),
		allFormats(a.Framed(x).Index(128).BgIndex(55)),
		allFormats(a.Encircled(x).Index(216).BgIndex(20)),
		allFormats(a.Overlined(x).Index(224).BgIndex(15)),

		allFormats(a.Black(x).BgBrightRed()),
		allFormats(a.Red(x).BgBrightRed()),
		allFormats(a.Green(x).BgBrightRed()),
		allFormats(a.Yellow(x).BgBrightRed()),
		allFormats(a.Brown(x).BgBrightRed()),
		allFormats(a.Blue(x).BgBrightRed()),
		allFormats(a.Magenta(x).BgBrightRed()),
		allFormats(a.Cyan(x).BgBrightRed()),
		allFormats(a.White(x).BgBrightRed()),
		allFormats(a.BrightBlack(x).BgBrightRed()),
		allFormats(a.BrightRed(x).BgBrightRed()),
		allFormats(a.BrightGreen(x).BgBrightRed()),
		allFormats(a.BrightYellow(x).BgWhite()),
		allFormats(a.BrightBlue(x).BgMagenta()),
		allFormats(a.BrightMagenta(x).BgGreen()),
		allFormats(a.BrightCyan(x).BgRed()),
		allFormats(a.BrightWhite(x).BgBlack()),
		allFormats(a.Index(178, x).BgGray(1)),
		allFormats(a.Gray(14, x).BgIndex(89)),

		allFormats(a.BgBlack(x).BrightBlack()),
		allFormats(a.BgRed(x).BrightRed()),
		allFormats(a.BgGreen(x).BrightGreen()),
		allFormats(a.BgYellow(x).BrightYellow()),
		allFormats(a.BgBrown(x).BrightYellow()),
		allFormats(a.BgBlue(x).BrightBlue()),
		allFormats(a.BgMagenta(x).BrightMagenta()),
		allFormats(a.BgCyan(x).BrightCyan()),
		allFormats(a.BgWhite(x).BrightWhite()),
		allFormats(a.BgBrightBlack(x).Black()),
		allFormats(a.BgBrightRed(x).Red()),
		allFormats(a.BgBrightGreen(x).Green()),
		allFormats(a.BgBrightYellow(x).Yellow()),
		allFormats(a.BgBrightBlue(x).Blue()),
		allFormats(a.BgBrightMagenta(x).Magenta()),
		allFormats(a.BgBrightCyan(x).Cyan()),
		allFormats(a.BgBrightWhite(x).White()),
		allFormats(a.BgIndex(187, x).Index(16)),
		allFormats(a.BgGray(15, x).Gray(0)),
	}
}

func benchSimpleValue(b *testing.B, a Aurora, x string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gVals = simpleValues(a, x)
	}
	b.ReportAllocs()
}

func benchComplexValue(b *testing.B, a Aurora, x string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gVals = complexValues(a, x)
	}
	b.ReportAllocs()
}

func benchValueString(b *testing.B, vals []Value) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(vals); j++ { // TODO: range and allocs ?
			gStr = vals[j].String()
		}
	}
	b.ReportAllocs()
}

func benchSprintf(b *testing.B, a Aurora, format interface{},
	args ...interface{}) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gStr = a.Sprintf(format, args...)
	}
	b.ReportAllocs()
}

func toInterfaces(vals []Value) []interface{} {
	r := make([]interface{}, 0, len(vals))
	for _, val := range vals {
		r = append(r, val)
	}
	return r
}

func auroraBench(a Aurora, b *testing.B) {
	// a.Red("...")
	b.Run("simple value", func(b *testing.B) {
		b.Run("short", func(b *testing.B) { benchSimpleValue(b, a, short) })
		b.Run("long", func(b *testing.B) { benchSimpleValue(b, a, long) })
	})
	// a.Red("...").BgRed().Bold().Inverse()
	b.Run("complex value", func(b *testing.B) {
		b.Run("short", func(b *testing.B) { benchComplexValue(b, a, short) })
		b.Run("long", func(b *testing.B) { benchComplexValue(b, a, long) })
	})
	// a.Red("...").String()
	b.Run("simple value string", func(b *testing.B) {
		b.Run("short", func(b *testing.B) {
			benchValueString(b, simpleValues(a, short))
		})
		b.Run("long", func(b *testing.B) {
			benchValueString(b, simpleValues(a, long))
		})
	})
	// a.Red("...").BgRed().Bold().Inverse().String()
	b.Run("complex value string", func(b *testing.B) {
		b.Run("short", func(b *testing.B) {
			benchValueString(b, complexValues(a, short))
		})
		b.Run("long", func(b *testing.B) {
			benchValueString(b, complexValues(a, long))
		})
	})

	// Sprintf

	b.Run("sprintf", func(b *testing.B) {
		// a.Sprintf(noString("... %s"), "x")
		b.Run("no-string", func(b *testing.B) {
			b.Run("short format", func(b *testing.B) {
				benchSprintf(b, a, noString(shortf), "x")
			})
			b.Run("long format", func(b *testing.B) {
				benchSprintf(b, a, noString(longf), "x")
			})
		})
		// a.Sprintf("... %s", "x")
		b.Run("usual", func(b *testing.B) {
			b.Run("short format", func(b *testing.B) {
				benchSprintf(b, a, shortf, "x")
			})
			b.Run("long format", func(b *testing.B) {
				benchSprintf(b, a, longf, "x")
			})
		})

		//
		// vary arguments
		//

		b.Run("simple arg", func(b *testing.B) {
			// string %s, a.Red("x") (one-color single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, "%s", a.Red("x"))
			})
			// string %s %s ..., a.Red("x"), ... (one-color many arguments)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, s18, toInterfaces(simpleValues(a, "x"))...)
			})
		})

		b.Run("complex arg", func(b *testing.B) {
			// string %s, a.Red("x").BgRed()... (many colors single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, "%s", a.Red("x").BgGray(5).Bold().Inverse())
			})
			// string %s %s..., a.Red("x").BgRed()..., ... (many colors 18 args)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, s18, toInterfaces(complexValues(a, "x"))...)
			})
		})

		//
		// same with colored format
		//

		b.Run("simple arg (simple format)", func(b *testing.B) {
			// string %s, a.Red("x") (one-color single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta("%s"), a.Red("x"))
			})
			// string %s %s ..., a.Red("x"), ... (one-color many arguments)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18),
					toInterfaces(simpleValues(a, "x"))...)
			})
		})

		b.Run("complex arg (simple format)", func(b *testing.B) {
			// string %s, a.Red("x").BgRed()... (many colors single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta("%s"),
					a.Red("x").BgGray(5).Bold().Inverse())
			})
			// string %s %s..., a.Red("x").BgRed()..., ... (many colors 18 args)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18),
					toInterfaces(complexValues(a, "x"))...)
			})
		})

		//
		// same with complex format
		//

		b.Run("simple arg (complex format)", func(b *testing.B) {
			// string %s, a.Red("x") (one-color single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta("%s").BgGray(5).Bold().Inverse(),
					a.Red("x"))
			})
			// string %s %s ..., a.Red("x"), ... (one-color many arguments)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18).BgGray(5).Bold().Inverse(),
					toInterfaces(simpleValues(a, "x"))...)
			})
		})

		b.Run("complex arg (complex format)", func(b *testing.B) {
			// string %s, a.Red("x").BgRed()... (many colors single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta("%s").BgGray(5).Bold().Inverse(),
					a.Red("x").BgGray(5).Bold().Inverse())
			})
			// string %s %s..., a.Red("x").BgRed()..., ... (many colors 18 args)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18).BgGray(5).Bold().Inverse(),
					toInterfaces(complexValues(a, "x"))...)
			})
		})
	})

}

func Benchmark_auroraClear(b *testing.B) {
	a := NewAurora(false)
	auroraBench(a, b)
}

// create a value
func Benchmark_aurora(b *testing.B) {
	a := NewAurora(true)
	auroraBench(a, b)
}

// Wrap functions

func Benchmark_wrap(b *testing.B) {
	ary := []struct {
		name string
		fn   func(interface{}) Value
	}{
		{"Reset", Reset},

		{"Bold", Bold},
		{"Faint", Faint},
		{"DoublyUnderline", DoublyUnderline},
		{"Fraktur", Fraktur},
		{"Italic", Italic},
		{"Underline", Underline},
		{"SlowBlink", SlowBlink},
		{"RapidBlink", RapidBlink},
		{"Blink", Blink},
		{"Reverse", Reverse},
		{"Inverse", Inverse},
		{"Conceal", Conceal},
		{"Hidden", Hidden},
		{"CrossedOut", CrossedOut},
		{"StrikeThrough", StrikeThrough},
		{"Framed", Framed},
		{"Encircled", Encircled},
		{"Overlined", Overlined},

		{"Black", Black},
		{"Red", Red},
		{"Green", Green},
		{"Yellow", Yellow},
		{"Brown", Brown},
		{"Blue", Blue},
		{"Magenta", Magenta},
		{"Cyan", Cyan},
		{"White", White},
		{"BrightBlack", BrightBlack},
		{"BrightRed", BrightRed},
		{"BrightGreen", BrightGreen},
		{"BrightYellow", BrightYellow},
		{"BrightBlue", BrightBlue},
		{"BrightMagenta", BrightMagenta},
		{"BrightCyan", BrightCyan},
		{"BrightWhite", BrightWhite},

		{"BgBlack", BgBlack},
		{"BgRed", BgRed},
		{"BgGreen", BgGreen},
		{"BgYellow", BgYellow},
		{"BgBrown", BgBrown},
		{"BgBlue", BgBlue},
		{"BgMagenta", BgMagenta},
		{"BgCyan", BgCyan},
		{"BgWhite", BgWhite},
		{"BgBrightBlack", BgBrightBlack},
		{"BgBrightRed", BgBrightRed},
		{"BgBrightGreen", BgBrightGreen},
		{"BgBrightYellow", BgBrightYellow},
		{"BgBrightBlue", BgBrightBlue},
		{"BgBrightMagenta", BgBrightMagenta},
		{"BgBrightCyan", BgBrightCyan},
		{"BgBrightWhite", BgBrightWhite},
	}
	for _, wf := range ary {
		b.Run(wf.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gVal = wf.fn(0)
			}
			b.ReportAllocs()
		})
	}

	b.Run("Colorize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gVal = Colorize(0, RedFg)
		}
		b.ReportAllocs()
	})

	b.Run("Index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gVal = Index(216, 0)
		}
		b.ReportAllocs()
	})

	b.Run("BgIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gVal = Index(216, 0)
		}
		b.ReportAllocs()
	})

	b.Run("Gray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gVal = Gray(15, 0)
		}
		b.ReportAllocs()
	})

	b.Run("BgGray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gVal = BgGray(15, 0)
		}
		b.ReportAllocs()
	})
}
