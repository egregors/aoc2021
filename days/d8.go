/*
	https://adventofcode.com/2021/day/8
*/

package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	_1 = "cf"
	_7 = "acf"
	_4 = "bcdf"
	_8 = "abcdefg"

	_2 = "acdeg"
	_3 = "acdfg"
	_5 = "abdfg"
	_0 = "abcefg"
	_6 = "abdefg"
	_9 = "abcdfg"
)

func contain(a, b map[rune]bool) bool {
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

func decodeMapping(s string) map[rune]rune {
	xs := strings.Split(s, " ")
	var sortXs []string
	for _, x := range xs {
		sortXs = append(sortXs, SortString(x))
	}
	xs = sortXs
	sort.Slice(xs, func(i, j int) bool {
		return len(xs[i]) < len(xs[j])
	})

	ds := map[rune]rune{
		'a': '_',
		'b': '_',
		'c': '_',
		'd': '_',
		'e': '_',
		'f': '_',
		'g': '_',
	}

	//var vierW, siebenW, achtW string
	eins, vier, sieben, acht := make(map[rune]bool, 9), make(map[rune]bool, 9), make(map[rune]bool, 9), make(map[rune]bool, 9)

	for _, w := range xs {
		switch len(w) {
		case 2:
			for _, ch := range w {
				eins[ch] = true
				//einsW = w
			}
		case 3:
			for _, ch := range w {
				sieben[ch] = true
				//siebenW = w
			}
		case 4:
			for _, ch := range w {
				vier[ch] = true
				//vierW = w
			}
		case 7:
			for _, ch := range w {
				acht[ch] = true
				//achtW = w
			}
		}
	}

	//fmt.Println("1", einsW)
	//fmt.Println("4", vierW)
	//fmt.Println("7", siebenW)
	//fmt.Println("8", achtW)

	//fmt.Println("\\\\\\\\\\")

	// A is 7 without 1
	for k := range sieben {
		if _, ok := eins[k]; !ok {
			ds['a'] = k
		}
	}

	// looking for 0
	// should have len == 6 and contain full 7 and can't contain full 4
	null := make(map[rune]bool, 9)
	var nullW string
	for _, w := range xs {
		if len(w) == 6 {
			tmpNull := make(map[rune]bool, 9)
			for _, ch := range w {
				tmpNull[ch] = true
			}
			if contain(tmpNull, sieben) && !contain(tmpNull, vier) {
				null = tmpNull
				nullW = SortString(w)
			}
		}
	}

	// D is 4 without 0
	for k := range vier {
		if _, ok := null[k]; !ok {
			ds['d'] = k
		}
	}

	// B is 4 without 1 and D
	for k := range vier {
		if _, ok := eins[k]; !ok && k != ds['d'] {
			ds['b'] = k
		}
	}

	// looking for 6
	// should have len == 6 and contain all letters from 8 exclude one of two 1 letter
	// C is 8 - (6 + 1)
	//sechs := make(map[rune]bool, 9)
	var sechsW string
	for _, w := range xs {
		if len(w) == 6 && w != nullW {
			tmpSechs := make(map[rune]bool, 9)
			for _, ch := range w {
				tmpSechs[ch] = true
			}

			diff := make(map[rune]bool, 9)

			for k := range acht {
				if _, ok := tmpSechs[k]; !ok {
					diff[k] = true
				}
			}

			if len(diff) == 1 {
				for k := range eins {
					if _, ok := diff[k]; ok {
						//sechs = tmpSechs
						sechsW = SortString(w)
						ds['c'] = k
					} else {
						ds['f'] = k
					}
				}
			}

		}
	}

	// looking for 9
	// word with len==6, exclude 0 and 6
	// E is 8 - 9
	neun := make(map[rune]bool)
	for _, w := range xs {
		if len(w) == 6 && w != nullW && w != sechsW {
			for _, ch := range w {
				neun[ch] = true
			}

			diff := make(map[rune]bool, 9)

			for k := range acht {
				if _, ok := neun[k]; !ok {
					diff[k] = true
				}
			}

			if len(diff) == 1 {
				for k := range diff {
					ds['e'] = k
				}
			}

		}
	}

	// G is just last unused letter
	var vals string
	for _, v := range ds {
		vals += string(v)
	}
	vals = SortString(vals)
	for _, ch := range "abcdefg" {
		if !strings.Contains(vals, string(ch)) {
			ds['g'] = ch
		}
	}

	// [ab dab eafb gcdfa cdfbe fbcad cdfgeb cefabd cagedb acedgfb]
	//  1! 7!  4!				      0?     0?		0!	   8!
	//								  6!     6?
	//										 9!
	//  2..3...4....5.................6....................7...... – len
	//
	// ab
	//
	//					    ....       	....  	-- []
	//					   .	a	   .	b
	//					   .	a	   .	b
	//					    ....   OR   ....
	//					   .	b	   .	a
	//					   .	b	   .	a
	//					    ....		....

	// dab
	//
	// 					    dddd        dddd     -- sure for [d: a]
	// 					   .	a      .	b
	// 					   .	a      .	b
	// 					    ....   OR   ....
	// 					   .	b      .	a
	// 					   .	b      .	a
	// 					    ....        ....

	// eafb
	//
	// 				        .... 	    ....  	  |     ....        ....	- []
	// 				       e	a	   f	a 	  |    e	b      f	b
	// 				       e	a	   f	a 	  |    e	b      f	b
	// 				        ffff   OR   eeee  	  |     ffff   OR   eeee
	// 				       .	b	   .	b 	  |    .	a      .	a
	// 				       .	b	   .	b 	  |    .	a      .	a
	// 				        .... 	    ....  	  |     ....        ....
	//
	// looking for 0, should have len == 6 and contain letters form 7 and NOT contain letters from 7
	// cdfgeb VS cefabd VS cagedb -> cagedb
	//
	// cagedb
	//
	//                      dddd
	//                     e!	a~
	//                     e!	a~
	//                      ....
	//                     ?	b~
	//                     ?	b~
	//                      ????
	// c[a]ge[db]
	// 4 has e, so now we sure for [d: a, e: b, f: d]
	// cg is left for now
	//
	// 8 without one letter but with 1ns letters is 6
	// looking for 6: len == 6, without a|b
	//
	// cdfgeb
	//
	//                      dddd
	//                     e	.
	//                     e	.
	//                      ffff
	//                     ?	b
	//                     ?	b
	//                     	????
	//
	// so now we know [b: f], => [d: a, e: b, a: c, f: d, b: f]
	//
	// now, if we know 6, we could find 9
	// looking for 9: len == 6, with abdef OR just left len==6 num
	//
	// c[efa]b[d]
	//
	//  dddd
	// e	a
	// e	a
	//  ffff
	// .	b
	// .	b
	//  cccc
	//
	// now we know [c: g] =>  [d: a, e: b, a: c, f: d, b: f, c: g]
	//
	// and last letter is just 9-8: acedgfb - cefabd -> g, [g: e]
	//
	// [d: a, e: b, a: c, f: d, g: e, b: f, c: g]
	//
	// 				[
	// 				a: c, -
	// 				f: d, -
	// 				g: e, -
	// 				b: f, -
	// 				c: g, -
	// 				d: a, -
	// 				e: b, -
	// 				]
	//
	//
	// way: 1 -> 7 -> 4 -> 0 -> 6 -> 9
	//
	// just for check:
	//
	// cdfbe
	//
	// 								    dddd
	// 								   e	.
	// 								   e	.
	// 								    ffff
	// 								   .	b
	// 								   .	b
	// 								    gggg

	return ds
}

func printDs(ds map[rune]rune) {
	for k, v := range ds {
		fmt.Println(string(k), ":", string(v))
	}
}

func segmentToInt(ds map[rune]rune, s string) int {
	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	// convert and sort
	m := make(map[rune]rune, 9)
	for k, v := range ds {
		m[v] = k
	}

	seg := ""
	for _, ch := range s {
		seg += string(m[ch])
	}
	seg = SortString(seg)

	switch seg {
	case _2:
		return 2
	case _3:
		return 3
	case _5:
		return 5
	case _0:
		return 0
	case _6:
		return 6
	case _9:
		return 9
	}

	return -1
}

func sevenSegmentSearchP1(xs []string) int {
	res := 0
	for _, x := range xs {
		r := strings.Split(x, " | ")[1]
		for _, w := range strings.Split(r, " ") {
			switch len(w) {
			case 2, 3, 4, 7:
				res++
			}
		}
	}
	return res
}

func sevenSegmentSearchP2(xs []string) int {
	sum := 0
	for _, x := range xs {
		parts := strings.Split(x, " | ")
		l, r := parts[0], parts[1]
		ds := decodeMapping(l)

		fmt.Println("===")
		printDs(ds)

		localRes := ""
		for _, w := range strings.Split(r, " ") {
			localRes += strconv.Itoa(segmentToInt(ds, w))
		}
		localResInt, err := strconv.Atoi(localRes)
		if err != nil {
			panic(err)
		}
		sum += localResInt
	}

	return sum
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
