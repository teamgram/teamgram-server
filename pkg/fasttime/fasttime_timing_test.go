package fasttime

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkUnixTimestamp(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var ts int64
		for pb.Next() {
			ts += UnixTimestamp()
		}
		Sink.Store(ts)
	})
}

func BenchmarkTimeNowUnix(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var ts int64
		for pb.Next() {
			ts += time.Now().Unix()
		}
		Sink.Store(ts)
	})
}

// Sink should prevent from code elimination by optimizing compiler
var Sink atomic.Int64
