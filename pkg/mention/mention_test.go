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
				{'@', "gernest", 0},
			},
		},
		{
			"@gernest ",
			[]Tag{
				{'@', "gernest", 0},
			},
		},
		{
			"@gernest@mwanza hello",
			[]Tag{
				{'@', "gernest@mwanza", 0},
			},
		},
		{
			"please email support@example.com to contact @martin",
			[]Tag{
				{'@', "martin", 44},
			},
		},
		{
			"please email العَرَبِيَّة@example.com to contact @martin",
			[]Tag{
				{'@', "martin", 61},
			},
		},
		{
			"@gernest @mwanza @mwanza",
			[]Tag{
				{'@', "gernest", 0},
				{'@', "mwanza", 9},
				{'@', "mwanza", 17},
			},
		},
		{
			"Hello to @gernest. Maybe we can do it together @mwanza",
			[]Tag{
				{'@', "gernest", 9},
				{'@', "mwanza", 47},
			},
		},
		{
			" @gernest @mwanza",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			" @gernest @mwanza ",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			" @gernest @mwanza @tanzania",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
				{'@', "tanzania", 18},
			},
		},
		{
			" @gernest,@mwanza/Tanzania ",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			"how does it feel to be rejected? @ it is @loner tt ggg sjdsj dj @linker ",
			[]Tag{
				{'@', "loner", 41},
				{'@', "linker", 64},
			},
		},
		{
			"This @gernest is @@@@ @@@ @@ @ @,, @, @mwanza,",
			[]Tag{
				{'@', "gernest", 5},
				{'@', "mwanza", 38},
			},
		},
		{
			"hello@world",
			nil,
		},
		{
			"@hello\u2000world", // en space
			[]Tag{{'@', "hello", 0}},
		},
		{
			"@hello\u200dworld", // zero width joiner (unprintable)
			[]Tag{{'@', "hello", 0}},
		},
		{
			"@hello\x1eworld", // control character
			[]Tag{{'@', "hello", 0}},
		},
		{
			"Hello @العَرَبِيَّة there",
			[]Tag{{'@', "العَرَبِيَّة", 6}},
		},
		{
			"ﺎﻠﻋﺮﺒﻳﺓ @العَرَبِيَّة there",
			[]Tag{{'@', "العَرَبِيَّة", 22}},
		},
	}
	terms := []rune(",/.")

	for _, v := range sample {
		c.Assert(GetTags('@', v.src, terms...), DeepEquals, v.tags, Commentf("Failed: %+v", v))
	}

	// use default terminators
	c.Assert(GetTags('@', "hello @test"), DeepEquals, []Tag{{'@', "test", 6}})
}

func BenchmarkGetTags(b *testing.B) {
	terms := []rune(",/. ")
	src := "This @gernest is @hello\u2000world @hello\u200d"
	for i := 0; i < b.N; i++ {
		GetTags('@', src, terms...)
	}
}
