package mention

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MentionSuite struct{}

var _ = Suite(&MentionSuite{})

func (s *MentionSuite) TestGetTags(c *C) {

	sample := []struct {
		src  string
		tags []Tag
	}{
		{
			"@gernest",
			[]Tag{
				{'@', "gernest", []uint16{}, 0},
			},
		},
		{
			"@gernest ",
			[]Tag{
				{'@', "gernest", []uint16{}, 0},
			},
		},
		{
			"@gernest@mwanza hello",
			[]Tag{
				{'@', "gernest@mwanza", []uint16{}, 0},
			},
		},
		{
			"please email support@example.com to contact @martin",
			[]Tag{
				{'@', "martin", []uint16{}, 44},
			},
		},
		{
			"please email العَرَبِيَّة@example.com to contact @martin",
			[]Tag{
				{'@', "martin", []uint16{}, 61},
			},
		},
		{
			"@gernest @mwanza @mwanza",
			[]Tag{
				{'@', "gernest", []uint16{}, 0},
				{'@', "mwanza", []uint16{}, 9},
				{'@', "mwanza", []uint16{}, 17},
			},
		},
		{
			"Hello to @gernest. Maybe we can do it together @mwanza",
			[]Tag{
				{'@', "gernest", []uint16{}, 9},
				{'@', "mwanza", []uint16{}, 47},
			},
		},
		{
			" @gernest @mwanza",
			[]Tag{
				{'@', "gernest", []uint16{}, 1},
				{'@', "mwanza", []uint16{}, 10},
			},
		},
		{
			" @gernest @mwanza ",
			[]Tag{
				{'@', "gernest", []uint16{}, 1},
				{'@', "mwanza", []uint16{}, 10},
			},
		},
		{
			" @gernest @mwanza @tanzania",
			[]Tag{
				{'@', "gernest", []uint16{}, 1},
				{'@', "mwanza", []uint16{}, 10},
				{'@', "tanzania", []uint16{}, 18},
			},
		},
		{
			" @gernest,@mwanza/Tanzania ",
			[]Tag{
				{'@', "gernest", []uint16{}, 1},
				{'@', "mwanza", []uint16{}, 10},
			},
		},
		{
			"how does it feel to be rejected? @ it is @loner tt ggg sjdsj dj @linker ",
			[]Tag{
				{'@', "loner", []uint16{}, 41},
				{'@', "linker", []uint16{}, 64},
			},
		},
		{
			"This @gernest is @@@@ @@@ @@ @ @,, @, @mwanza,",
			[]Tag{
				{'@', "gernest", []uint16{}, 5},
				{'@', "mwanza", []uint16{}, 38},
			},
		},
		{
			"hello@world",
			nil,
		},
		{
			"@hello\u2000world", // en space
			[]Tag{{'@', "hello", []uint16{}, 0}},
		},
		{
			"@hello\u200dworld", // zero width joiner (unprintable)
			[]Tag{{'@', "hello", []uint16{}, 0}},
		},
		{
			"@hello\x1eworld", // control character
			[]Tag{{'@', "hello", []uint16{}, 0}},
		},
		{
			"Hello @العَرَبِيَّة there",
			[]Tag{{'@', "العَرَبِيَّة", []uint16{}, 6}},
		},
		{
			"ﺎﻠﻋﺮﺒﻳﺓ @العَرَبِيَّة there",
			[]Tag{{'@', "العَرَبِيَّة", []uint16{}, 22}},
		},
	}
	terms := []rune(",/.")

	for _, v := range sample {
		c.Assert(GetTags('@', v.src, terms...), DeepEquals, v.tags, Commentf("Failed: %+v", v))
	}

	// use default terminators
	c.Assert(GetTags('@', "hello @test"), DeepEquals, []Tag{{'@', "test", []uint16{}, 6}})
}

func BenchmarkGetTags(b *testing.B) {
	terms := []rune(",/. ")
	src := "This @gernest is @hello\u2000world @hello\u200d"
	for i := 0; i < b.N; i++ {
		GetTags('@', src, terms...)
	}
}
