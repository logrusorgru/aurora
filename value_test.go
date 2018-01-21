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
	want := "\033[30mx\033[0m\033[41m"
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
	want = "\033[31m" + fmt.Sprintf("%+1.3g", 3.14) +
		"\033[0m\033[44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	var utf8Verb = "%+1.3世" // verb that fit more then 1 byte
	got = fmt.Sprintf(utf8Verb, v)
	want = "\033[31m" + fmt.Sprintf(utf8Verb, 3.14) +
		"\033[0m\033[44m"
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
		if c := v.Color(); c != clr {
			t.Errorf("wrong color for %s: want %d, got %d", name, clr, c)
		}
	}
	// colorized
	test("Black", Red("x").Black(), BlackFg)
	test("Red", Bold("x").Red(), RedFg|BoldFm)
	test("Green", Inverse("x").Green(), GreenFg|InverseFm)
	test("Inverse&Brown", Bold("x").Inverse().Brown(), BoldFm|InverseFm|BrownFg)
	test("Blue", BgBlue("x").Blue(), BlueFg|BlueBg)
	test("Magenta", Magenta("x").Magenta(), MagentaFg)
	test("Cyan", Red("x").Cyan(), CyanFg)
	test("Gray", Green("x").Gray(), GrayFg)
	test("BgBlack", Black("x").BgBlack(), BlackFg|BlackBg)
	test("BgRed", Red("x").BgRed(), RedFg|RedBg)
	test("BgGreen", Green("x").BgGreen(), GreenFg|GreenBg)
	test("BgBrown", Brown("x").BgBrown(), BrownFg|BrownBg)
	test("BgBlue", Blue("x").BgBlue(), BlueFg|BlueBg)
	test("BgMagenta", BgCyan("x").BgMagenta(), MagentaBg)
	test("BgCyan", Cyan("x").BgCyan(), CyanFg|CyanBg)
	test("BgGray", Gray("x").BgGray(), GrayFg|GrayBg)
	test("Bold & BlueBg", Red("x").BgBlue().Bold(), RedFg|BoldFm|BlueBg)
	test("Inverse", Black("x").Inverse(), BlackFg|InverseFm)
	// clear
	test("Black", valueClear{"x"}.Black(), 0)
	test("Red", valueClear{"x"}.Red(), 0)
	test("Green", valueClear{"x"}.Green(), 0)
	test("Inverse&Brown", valueClear{"x"}.Inverse().Brown(), 0)
	test("Blue", valueClear{"x"}.Blue(), 0)
	test("Magenta", valueClear{"x"}.Magenta(), 0)
	test("Cyan", valueClear{"x"}.Cyan(), 0)
	test("Gray", valueClear{"x"}.Gray(), 0)
	test("BgBlack", valueClear{"x"}.BgBlack(), 0)
	test("BgRed", valueClear{"x"}.BgRed(), 0)
	test("BgGreen", valueClear{"x"}.BgGreen(), 0)
	test("BgBrown", valueClear{"x"}.BgBrown(), 0)
	test("BgBlue", valueClear{"x"}.BgBlue(), 0)
	test("BgMagenta", valueClear{"x"}.BgMagenta(), 0)
	test("BgCyan", valueClear{"x"}.BgCyan(), 0)
	test("BgGray", valueClear{"x"}.BgGray(), 0)
	test("Bold & BlueBg", valueClear{"x"}.BgBlue().Bold(), 0)
	test("Inverse", valueClear{"x"}.Inverse(), 0)
}
