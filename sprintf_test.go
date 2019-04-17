//
// Copyright (c) 2016-2019 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Do What The Fuck You
// Want To Public License, Version 2, as published by Sam Hocevar.
// See LICENSE file for more details or see below.
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

type noString string

func Test_Sprintf(t *testing.T) {
	//b := Black("x")
	v := Sprintf(noString("delta: +%d"), 3)
	if v != "delta: +3" {
		t.Error("Sprintf: wrong result")
	}
	v = Sprintf(Red("deltas: +%d, %d, %d points"), 3, 5, 9)
	want := "\033[31mdeltas: +3, 5, 9 points\033[0m"
	if v != want {
		t.Errorf("Sprintf: want: %q, got %q", want, v)
	}
}
