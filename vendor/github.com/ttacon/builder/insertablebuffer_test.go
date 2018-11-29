package builder

import "testing"

func TestInsertString(t *testing.T) {
	buf := NewBuilderString("hello")
	buf.InsertString(0, "world")
	if buf.String() != "worldhello" {
		t.Error("didn't get back expected string")
	}
}

func TestInsertString_InMiddle(t *testing.T) {
	buf := NewBuilderString("hello")
	buf.InsertString(1, "world")
	if buf.String() != "hworldello" {
		t.Error("didn't get back expected string, got: ", buf.String())
	}
}

func TestInsert(t *testing.T) {
	buf := NewBuilderString("hello")
	buf.Insert(0, []byte("world"))
	if buf.String() != "worldhello" {
		t.Error("didn't get back expected string")
	}
}

func TestInsert_InMiddle(t *testing.T) {
	buf := NewBuilderString("hello")
	buf.Insert(1, []byte("world"))
	if buf.String() != "hworldello" {
		t.Error("didn't get back expected string, got: ", buf.String())
	}
}

func BenchmarkInsertString_EmptyBuf(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewBuilder(nil)
		buf.InsertString(0, "hello")
	}
}

func BenchmarkInsertString_NonEmptyBuf(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewBuilderString("world")
		buf.InsertString(0, "hello")
	}
}

func BenchmarkInsertString_NonEmptyBufTwoInsertions(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewBuilderString("world")
		buf.InsertString(0, "hello")
		buf.InsertString(0, "hello")
	}
}

func BenchmarkInsertString_NonEmptyBufThreeInsertions(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewBuilderString("world")
		buf.InsertString(0, "hello")
		buf.InsertString(0, "hello")
		buf.InsertString(0, "hello")
	}
}
