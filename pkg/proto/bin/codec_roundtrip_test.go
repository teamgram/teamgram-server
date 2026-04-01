package bin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoderDecoderRoundTripScalars(t *testing.T) {
	a := require.New(t)

	x := AcquireEncoderCap(64)
	defer x.Release()

	x.PutClazzID(0x1cb5c415)
	x.PutInt(-123)
	x.PutUint(456)
	x.PutLong(-789)
	x.PutUlong(987)
	x.PutDouble(3.1415926)

	d := NewDecoder(x.Bytes())

	clazzID, err := d.ClazzID()
	a.NoError(err)
	a.Equal(uint32(0x1cb5c415), clazzID)

	intValue, err := d.Int32()
	a.NoError(err)
	a.Equal(int32(-123), intValue)

	uintValue, err := d.Uint32()
	a.NoError(err)
	a.Equal(uint32(456), uintValue)

	longValue, err := d.Long()
	a.NoError(err)
	a.Equal(int64(-789), longValue)

	ulongValue, err := d.Ulong()
	a.NoError(err)
	a.Equal(uint64(987), ulongValue)

	doubleValue, err := d.Double()
	a.NoError(err)
	a.Equal(3.1415926, doubleValue)
	a.Equal(0, d.Remaining())
}

func TestEncoderDecoderRoundTripStringsAndBytesAtTLBoundaries(t *testing.T) {
	a := require.New(t)

	cases := []int{0, 1, 2, 3, 4, 253, 254, 255}
	for _, size := range cases {
		t.Run(string(rune(size)), func(t *testing.T) {
			x := AcquireEncoderCap(size*2 + 16)
			defer x.Release()

			payload := make([]byte, size)
			for i := range payload {
				payload[i] = byte(i)
			}

			x.PutString(string(payload))
			x.PutBytes(payload)

			d := NewDecoder(x.Bytes())

			stringValue, err := d.String()
			a.NoError(err)
			a.Equal(string(payload), stringValue)

			bytesView, err := d.BytesView()
			a.NoError(err)
			a.Equal(payload, bytesView)

			a.Equal(0, d.Remaining())
		})
	}
}

func TestDecoderBytesReturnsCopyAndBytesViewReturnsBorrowedSlice(t *testing.T) {
	a := require.New(t)

	x := AcquireEncoderCap(32)
	defer x.Release()

	x.PutBytes([]byte{1, 2, 3, 4})

	d := NewDecoder(x.Bytes())

	view, err := d.BytesView()
	a.NoError(err)
	a.Equal([]byte{1, 2, 3, 4}, view)

	d.Reset(x.Bytes())
	copied, err := d.Bytes()
	a.NoError(err)
	a.Equal([]byte{1, 2, 3, 4}, copied)

	view[0] = 9
	a.Equal(byte(1), copied[0])
}

func TestDecoderVectorHeaderAndConsumeClazzID(t *testing.T) {
	a := require.New(t)

	x := AcquireEncoder()
	defer x.Release()

	x.PutVectorHeader(3)
	x.PutInt(1)
	x.PutInt(2)
	x.PutInt(3)

	d := NewDecoder(x.Bytes())
	n, err := d.VectorHeader()
	a.NoError(err)
	a.Equal(int32(3), n)

	v1, err := d.Int()
	a.NoError(err)
	a.Equal(1, v1)
	v2, err := d.Int()
	a.NoError(err)
	a.Equal(2, v2)
	v3, err := d.Int()
	a.NoError(err)
	a.Equal(3, v3)
}

func TestDecoderReportsOffsetOnUnexpectedClazzID(t *testing.T) {
	a := require.New(t)

	x := AcquireEncoder()
	defer x.Release()
	x.PutClazzID(0x11111111)

	d := NewDecoder(x.Bytes())
	err := d.ConsumeClazzID(0x22222222)
	a.Error(err)

	typed, ok := err.(*UnexpectedClazzIDError)
	a.True(ok)
	a.Equal(uint32(0x22222222), typed.Want)
	a.Equal(uint32(0x11111111), typed.Got)
	a.Equal(0, typed.Offset)
}

func TestDecoderSkipChecksBounds(t *testing.T) {
	a := require.New(t)

	d := NewDecoder([]byte{1, 2, 3})
	a.NoError(d.Skip(2))
	a.Equal(1, d.Remaining())
	a.ErrorIs(d.Skip(2), ErrUnexpectedEOF)
}

func TestDecoderConsumeNRejectsTooSmallTarget(t *testing.T) {
	a := require.New(t)

	d := NewDecoder([]byte{1, 2, 3, 4})
	target := make([]byte, 2)

	err := d.ConsumeN(target, 4)
	a.ErrorIs(err, ErrUnexpectedEOF)
	a.Equal(0, d.Offset())
	a.Equal([]byte{0, 0}, target)
}

func TestUnexpectedClazzIDErrorIncludesOffsetZero(t *testing.T) {
	a := require.New(t)

	err := &UnexpectedClazzIDError{Want: 0x2, Got: 0x1, Offset: 0}
	a.Contains(err.Error(), "offset 0")
}

func TestInvalidLengthErrorIncludesOffsetZero(t *testing.T) {
	a := require.New(t)

	err := &InvalidLengthError{Type: "vector", Length: -1, Offset: 0}
	a.Contains(err.Error(), "offset 0")
}

func TestEncoderDoubleReleaseIsSafe(t *testing.T) {
	x := AcquireEncoder()
	x.PutInt(1)
	x.Release()
	x.Release()
}

func TestEncoderUseAfterReleasePanics(t *testing.T) {
	x := AcquireEncoder()
	x.Release()

	require.Panics(t, func() {
		x.PutInt(1)
	})
	require.Panics(t, func() {
		_ = x.Bytes()
	})
}

func TestDecoderRejectsNonZeroStringPadding(t *testing.T) {
	d := NewDecoder([]byte{1, 'a', 1, 0})

	_, err := d.String()
	require.Error(t, err)
}

func TestDecoderRejectsNonZeroBytesPadding(t *testing.T) {
	d := NewDecoder([]byte{1, 'a', 1, 0})

	_, err := d.BytesView()
	require.Error(t, err)
}

func TestFieldsPanicsOnInvalidBitIndex(t *testing.T) {
	var f Fields

	require.Panics(t, func() { f.Set(-1) })
	require.Panics(t, func() { f.Set(32) })
	require.Panics(t, func() { f.Unset(-1) })
	require.Panics(t, func() { f.Unset(32) })
	require.Panics(t, func() { _ = f.Has(-1) })
	require.Panics(t, func() { _ = f.Has(32) })
}

func TestVectorClazzIDUsesPredicateNaming(t *testing.T) {
	require.Equal(t, uint32(0x1cb5c415), ClazzID_vector)
}
