//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE.md file for more details or see below.
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
	v := Value{"x", 0, 0}
	if x := v.String(); x != "x" {
		t.Errorf("(Value).String: want %q, got %q", "x", x)
	}
	v = Value{"x", BlackFg, RedBg}
	want := "\033[30mx\033[0m\033[41m"
	if got := v.String(); want != got {
		t.Errorf("(Value).String: want %q, got %q", want, got)
	}
}

func TestValue_Color(t *testing.T) {
	if (Value{"", RedFg, 0}).Color() != RedFg {
		t.Error("wrong color")
	}
}

func TestValue_Value(t *testing.T) {
	if (Value{"x", RedFg, BlueBg}).Value() != "x" {
		t.Error("wrong value")
	}
}

func TestValue_Bleach(t *testing.T) {
	if (Value{"x", RedFg, BlueBg}).Bleach() != (Value{value: "x"}) {
		t.Error("wrong bleached")
	}
}

func TestValue_Format(t *testing.T) {
	v := Value{3.14, RedFg, BlueBg}
	got := fmt.Sprintf("%+1.3g", v)
	want := "\033[31m" + fmt.Sprintf("%+1.3g", 3.14) +
		"\033[0m\033[44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	got = fmt.Sprintf("%+1.3世", v)
	want = "\033[31m" + fmt.Sprintf("%+1.3世", 3.14) +
		"\033[0m\033[44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
}
