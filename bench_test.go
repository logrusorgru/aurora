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
		a.Black(x),
		a.Red(x),
		a.Green(x),
		a.Brown(x),
		a.Blue(x),
		a.Magenta(x),
		a.Cyan(x),
		a.Gray(x),
		a.BgBlack(x),
		a.BgRed(x),
		a.BgGreen(x),
		a.BgBrown(x),
		a.BgBlue(x),
		a.BgMagenta(x),
		a.BgCyan(x),
		a.BgGray(x),
		a.Bold(x),
		a.Inverse(x),
	}
}

// 18 values
func complexValues(a Aurora, x string) []Value {
	return []Value{
		a.Black(x).BgBlack().Bold().Inverse(),
		a.Red(x).BgRed().Bold().Inverse(),
		a.Green(x).BgGreen().Bold().Inverse(),
		a.Brown(x).BgBrown().Bold().Inverse(),
		a.Blue(x).BgBlue().Bold().Inverse(),
		a.Magenta(x).BgMagenta().Bold().Inverse(),
		a.Cyan(x).BgCyan().Bold().Inverse(),
		a.Gray(x).BgGray().Bold().Inverse(),
		// 8 above
		// we need 18 values to compare with simple values
		// 8 below
		a.Black(x).BgBlack().Bold().Inverse(),
		a.Red(x).BgRed().Bold().Inverse(),
		a.Green(x).BgGreen().Bold().Inverse(),
		a.Brown(x).BgBrown().Bold().Inverse(),
		a.Blue(x).BgBlue().Bold().Inverse(),
		a.Magenta(x).BgMagenta().Bold().Inverse(),
		a.Cyan(x).BgCyan().Bold().Inverse(),
		a.Gray(x).BgGray().Bold().Inverse(),
		// and two
		a.Black(x).BgBlack().Bold().Inverse(),
		a.Red(x).BgRed().Bold().Inverse(),
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
				benchSprintf(b, a, "%s", a.Red("x").BgGray().Bold().Inverse())
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
					a.Red("x").BgGray().Bold().Inverse())
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
				benchSprintf(b, a, a.Magenta("%s").BgGray().Bold().Inverse(),
					a.Red("x"))
			})
			// string %s %s ..., a.Red("x"), ... (one-color many arguments)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18).BgGray().Bold().Inverse(),
					toInterfaces(simpleValues(a, "x"))...)
			})
		})

		b.Run("complex arg (complex format)", func(b *testing.B) {
			// string %s, a.Red("x").BgRed()... (many colors single argument)
			b.Run("1 arg", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta("%s").BgGray().Bold().Inverse(),
					a.Red("x").BgGray().Bold().Inverse())
			})
			// string %s %s..., a.Red("x").BgRed()..., ... (many colors 18 args)
			b.Run("18 args", func(b *testing.B) {
				benchSprintf(b, a, a.Magenta(s18).BgGray().Bold().Inverse(),
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
		{"Black", Black},
		{"Red", Red},
		{"Green", Green},
		{"Brown", Brown},
		{"Blue", Blue},
		{"Magenta", Magenta},
		{"Cyan", Cyan},
		{"Gray", Gray},
		{"BgBlack", BgBlack},
		{"BgRed", BgRed},
		{"BgGreen", BgGreen},
		{"BgBrown", BgBrown},
		{"BgBlue", BgBlue},
		{"BgMagenta", BgMagenta},
		{"BgCyan", BgCyan},
		{"BgGray", BgGray},
		{"Bold", Bold},
		{"Inverse", Inverse},
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
}
