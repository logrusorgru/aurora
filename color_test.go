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
	"strconv"
	"testing"
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
			if nos != want {
				t.Errorf("%t %d: wrong nos string %q, want %q",
					zero, i, nos, want)
			}
		}
	}
}

func TestColor_IsValid(t *testing.T) {
	if Color(0).IsValid() == false {
		t.Error("invalid")
	}
}

func Test_itoa(t *testing.T) {
	for i := 0; i < 256; i++ {
		if a := itoa(byte(i)); a != strconv.Itoa(i) {
			t.Errorf("wrong %q, want %d", a, i)
		}
	}
}
