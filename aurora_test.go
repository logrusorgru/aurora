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

func isClear(v Value) bool {
	return v.Color() == 0 && v.tail() == 0
}

func isColor(v Value, clr Color) bool {
	return v.Color() == clr
}

func isTail(v Value, tl Color) bool {
	return v.tail() == tl
}

func Test_NewAurora(t *testing.T) {
	if a := NewAurora(false); a == nil {
		t.Error("NewAurora(false) returns nil")
	}
	if a := NewAurora(true); a == nil {
		t.Error("NewAurora(true) returns nil")
	}
	if t.Failed() {
		t.FailNow()
	}
}

func Test_auroraClear_methods(t *testing.T) {
	a := NewAurora(false)
	test := func(mn string, v Value) {
		if !isClear(v) {
			t.Errorf("NewAurora(false).%s is not clear", mn)
		} else if str, ok := v.Value().(string); !ok {
			t.Errorf("NewAurora(false).%s wrong value type", mn)
		} else if str != "x" {
			t.Errorf("NewAurora(false).%s wrong value", mn)
		}
	}
	test("Black", a.Black("x"))
	test("Red", a.Red("x"))
	test("Green", a.Green("x"))
	test("Brown", a.Brown("x"))
	test("Blue", a.Blue("x"))
	test("Magenta", a.Magenta("x"))
	test("Cyan", a.Cyan("x"))
	test("Gray", a.Gray("x"))
	test("BgBlack", a.BgBlack("x"))
	test("BgRed", a.BgRed("x"))
	test("BgGreen", a.BgGreen("x"))
	test("BgBrown", a.BgBrown("x"))
	test("BgBlue", a.BgBlue("x"))
	test("BgMagenta", a.BgMagenta("x"))
	test("BgCyan", a.BgCyan("x"))
	test("BgGray", a.BgGray("x"))
	test("Bold", a.Bold("x"))
	test("Inverse", a.Inverse("x"))
	test("Colorize", a.Colorize("x", RedFg|RedBg))
}

func Test_auroraClear_sprintf(t *testing.T) {
	a := NewAurora(false)
	if s := a.Sprintf(a.Black("x: %d"), a.Blue(2)); s != "x: 2" {
		t.Error("NewAurora(false).Sprintf wrong value")
	}
	if s := a.Sprintf("x: %d", a.Blue(2)); s != "x: 2" {
		t.Error("NewAurora(false).Sprintf wrong value")
	}
}

func Test_aurora_methods(t *testing.T) {
	a := NewAurora(true)
	test := func(mn string, v Value, clr Color) {
		if !isColor(v, clr) {
			t.Errorf("NewAurora(true).%s wrong color: %d", mn, v.Color())
		} else if !isTail(v, 0) {
			t.Errorf("NewAurora(true).%s unexpected tail value", mn)
		} else if str, ok := v.Value().(string); !ok {
			t.Errorf("NewAurora(true).%s wrong value type", mn)
		} else if str != "x" {
			t.Errorf("NewAurora(true).%s wrong value", mn)
		}
	}
	test("Black", a.Black("x"), BlackFg)
	test("Red", a.Red("x"), RedFg)
	test("Green", a.Green("x"), GreenFg)
	test("Brown", a.Brown("x"), BrownFg)
	test("Blue", a.Blue("x"), BlueFg)
	test("Magenta", a.Magenta("x"), MagentaFg)
	test("Cyan", a.Cyan("x"), CyanFg)
	test("Gray", a.Gray("x"), GrayFg)
	test("BgBlack", a.BgBlack("x"), BlackBg)
	test("BgRed", a.BgRed("x"), RedBg)
	test("BgGreen", a.BgGreen("x"), GreenBg)
	test("BgBrown", a.BgBrown("x"), BrownBg)
	test("BgBlue", a.BgBlue("x"), BlueBg)
	test("BgMagenta", a.BgMagenta("x"), MagentaBg)
	test("BgCyan", a.BgCyan("x"), CyanBg)
	test("BgGray", a.BgGray("x"), GrayBg)
	test("Bold", a.Bold("x"), BoldFm)
	test("Inverse", a.Inverse("x"), InverseFm)
	test("Colorize", a.Colorize("x", RedFg|RedBg), RedFg|RedBg)
}

func Test_aurora_Sprintf(t *testing.T) {
	a := NewAurora(true)
	s := a.Sprintf(a.Black("x: %d"), a.Blue(2))
	if s != "\033[30mx: \033[34m2\033[0m\033[30m\033[0m" {
		t.Errorf("NewAurora(true).Sprintf wrong value: %q", s)
	}
	s = a.Sprintf("x: %d", a.Blue(2))
	if s != "x: \033[34m2\033[0m" {
		t.Errorf("NewAurora(true).Sprintf wrong value: %q", s)
	}
}
