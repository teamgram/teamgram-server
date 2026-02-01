// Package mention provides function for parsing twitter like mentions and hashtags
package mention

import (
	"strings"
	"unicode"
)

// Tag is string that is prefixed with a marker. Often used to mark users like
// @genrest.
type Tag struct {

	// The character used to mark the beginning of the tag.
	Char rune

	// Tag non space string that follows after the tag character mark.
	Tag string

	// Tag non space string that follows after the tag character mark.
	TagUTF16 []uint16

	// Index is the byte position in the source string where the tag was found.
	Index int
}

// GetTags returns a slice of Tags, that is all characters after rune char up
// to occurrence of space or another occurrence of rune char. Additionally you
// can provide a coma separated unicode characters to be used as terminating
// sequence.
func GetTags(prefix rune, str string, terminator ...rune) (tags []Tag) {
	// If we have no terminators given, default to only whitespace
	if len(terminator) == 0 {
		terminator = []rune(" ")
	}
	// get list of indexes in our str that is a terminator
	// Always include the beginning of our str a terminator. This is so we can
	// detect the first character as a prefix
	termIndexes := []int{-1}
	for i, char := range str {
		if isTerminator(char, terminator...) {
			termIndexes = append(termIndexes, i)
		}
	}
	// Always include last character as a terminator
	termIndexes = append(termIndexes, len(str))

	// check if the character AFTER our term index is our prefix
	for i, t := range termIndexes {
		// ensure term index is not the last character in str
		if t >= (len(str) - 1) {
			break
		}
		if str[t+1] == byte(prefix) {
			tagText := strings.TrimLeft(str[t+2:termIndexes[i+1]], string(prefix))
			if tagText == "" {
				continue
			}
			index := t + 1
			tags = append(tags, Tag{prefix, tagText, []uint16{}, index})
		}
	}

	return
}

// GetTagsAsUniqueStrings gets all tags as a slice of unique strings. This is
// here to have a means of being somewhat backwards compatible with previous
// versions of mention
func GetTagsAsUniqueStrings(prefix rune, str string, terminator ...rune) (strs []string) {
	tags := GetTags(prefix, str, terminator...)
	for _, tag := range tags {
		strs = append(strs, tag.Tag)
	}
	return uniquify(strs)
}

// Is given rune listed as a terminator
func isTerminator(r rune, terminator ...rune) bool {
	for _, t := range terminator {
		if r == t {
			return true
		}
	}
	return unicode.IsSpace(r) || !unicode.IsPrint(r)
}

// Ensures the given slice of strings are unique and that none are empty
// strings
func uniquify(in []string) (out []string) {
	for _, i := range in {
		if i == "" {
			continue
		}
		for _, o := range out {
			if i == o {
				continue
			}
		}
		out = append(out, i)
	}
	return
}

func GetUTF16Tags(prefix rune, str []uint16, terminator ...rune) (tags []Tag) {
	// If we have no terminators given, default to only whitespace
	if len(terminator) == 0 {
		terminator = []rune(" ")
	}
	// get list of indexes in our str that is a terminator
	// Always include the beginning of our str a terminator. This is so we can
	// detect the first character as a prefix
	termIndexes := []int{-1}
	for i, char := range str {
		if isTerminator(rune(char), terminator...) {
			termIndexes = append(termIndexes, i)
		}
	}
	// Always include last character as a terminator
	termIndexes = append(termIndexes, len(str))

	// check if the character AFTER our term index is our prefix
	for i, t := range termIndexes {
		// ensure term index is not the last character in str
		if t >= (len(str) - 1) {
			break
		}
		if str[t+1] == uint16(prefix) {
			//utf16.EncodeRune()
			// tagText := strings.TrimLeft(str[t+2:termIndexes[i+1]], string(prefix))
			tagText := str[t+2 : termIndexes[i+1]]
			if len(tagText) == 0 {
				continue
			}
			index := t + 1
			tags = append(tags, Tag{prefix, "", tagText, index})
		}
	}

	return
}

// GetTagsAsUniqueStrings gets all tags as a slice of unique strings. This is
// here to have a means of being somewhat backwards compatible with previous
// versions of mention
func GetTagsAsUniqueUTF16Strings(prefix rune, str []uint16, terminator ...rune) (strs [][]uint16) {
	tags := GetUTF16Tags(prefix, str, terminator...)
	for _, tag := range tags {
		strs = append(strs, tag.TagUTF16)
	}
	return uniquifyUTF16(strs)
}

// Ensures the given slice of strings are unique and that none are empty
// strings
func uniquifyUTF16(in [][]uint16) (out [][]uint16) {
	for _, i := range in {
		if len(i) == 0 {
			continue
		}
		for _, o := range out {
			if equalUTf16(i, o) {
				continue
			}
		}
		out = append(out, i)
	}
	return
}

func equalUTf16(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
