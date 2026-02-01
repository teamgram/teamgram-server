// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

// WordLen represents 4-byte sequence.
// Values in TL are generally aligned to Word.
const WordLen = 4

func nearestPaddedValueLength(l int) int {
	n := WordLen * (l / WordLen)
	if n < l {
		n += WordLen
	}
	return n
}
