package mt

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"

	"github.com/go-faster/errors"
	"github.com/klauspost/compress/gzip"
	"go.uber.org/multierr"
)

type gzipPool struct {
	writers sync.Pool
	readers sync.Pool
}

type pooledGzipReader struct {
	*gzip.Reader
	pool *gzipPool
}

func (p *pooledGzipReader) Close() error {
	err := p.Reader.Close()
	p.pool.PutReader(p.Reader)
	return err
}

func newGzipPool() *gzipPool {
	return &gzipPool{
		writers: sync.Pool{
			New: func() interface{} {
				return gzip.NewWriter(nil)
			},
		},
		readers: sync.Pool{},
	}
}

func (g *gzipPool) GetWriter(w io.Writer) *gzip.Writer {
	writer := g.writers.Get().(*gzip.Writer)
	writer.Reset(w)
	return writer
}

func (g *gzipPool) PutWriter(w *gzip.Writer) {
	g.writers.Put(w)
}

func (g *gzipPool) GetReader(r io.Reader) (*gzip.Reader, error) {
	reader, ok := g.readers.Get().(*gzip.Reader)
	if !ok {
		r, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	if err := reader.Reset(r); err != nil {
		g.readers.Put(reader)
		return nil, err
	}
	return reader, nil
}

func (g *gzipPool) PutReader(w *gzip.Reader) {
	g.readers.Put(w)
}

func newCompressedReader(buf []byte) (io.ReadCloser, error) {
	r, err := gzipRWPool.GetReader(bytes.NewReader(buf))
	if err == nil {
		return &pooledGzipReader{Reader: r, pool: gzipRWPool}, nil
	}

	// zlib is a compatibility fallback, so keep it simple unless metrics show it is hot.
	// Track decode counts by codec and only consider pooling if zlib becomes a meaningful share.
	zr, zErr := zlib.NewReader(bytes.NewReader(buf))
	if zErr != nil {
		return nil, fmt.Errorf("create decompressor: gzip: %v; zlib: %w", err, zErr)
	}
	return zr, nil
}

// TLGzipPacked represents a Packed Object.
//
// Used to replace any other object (or rather, a serialization thereof)
// with its archived (gzipped) representation.
type TLGzipPacked struct {
	PackedData []byte
	Obj        iface.TLObject
}

var (
	gzipRWPool  = newGzipPool()
	gzipBufPool = sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(nil)
	}}
)

func (g *TLGzipPacked) ClazzName() string {
	return "gzip_packed"
}

func (g *TLGzipPacked) Encode(b *bin.Encoder, layer int32) (rErr error) {
	_ = layer

	b.PutClazzID(ClazzID_gzip_packed)
	if len(g.PackedData) == 0 {
		return fmt.Errorf("unable to encode gzip_packed: field packed_data is empty")
	}

	// Writing compressed data to buf.
	buf := gzipBufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer gzipBufPool.Put(buf)

	w := gzipRWPool.GetWriter(buf)
	defer func() {
		if closeErr := w.Close(); closeErr != nil {
			closeErr = errors.Wrap(closeErr, "close")
			multierr.AppendInto(&rErr, closeErr)
		}
		gzipRWPool.PutWriter(w)
	}()
	if _, err := w.Write(g.PackedData); err != nil {
		return fmt.Errorf("unable to encode gzip_packed: compress packed_data: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("unable to encode gzip_packed: close gzip writer: %w", err)
	}

	// Writing compressed data as bytes.
	b.PutBytes(buf.Bytes())

	return nil
}

type countReader struct {
	reader io.Reader
	read   int64
}

func (c *countReader) Total() int64 {
	return atomic.LoadInt64(&c.read)
}

func (c *countReader) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	atomic.AddInt64(&c.read, int64(n))
	return n, err
}

// DecompressionBombErr means that gzip_packed decode detected decompression bomb
// which decompressed payload is significantly higher than initial compressed
// size and stopped decompression to prevent OOM.
type DecompressionBombErr struct {
	Compressed   int
	Decompressed int
}

func (d *DecompressionBombErr) Error() string {
	return fmt.Sprintf("payload too big (expanded %d bytes to greater than %d)",
		d.Compressed, d.Decompressed,
	)
}

func (g *TLGzipPacked) Decode(b *bin.Decoder) (rErr error) {
	buf, err := b.Bytes()
	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: field packed_data: %w", err)
	}

	r, err := newCompressedReader(buf)
	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: create decompressor: %w", err)
	}
	defer func() {
		if closeErr := r.Close(); closeErr != nil {
			closeErr = errors.Wrap(closeErr, "close")
			multierr.AppendInto(&rErr, closeErr)
		}
	}()

	// Apply mitigation for reading too much data which can result in OOM.
	const maxUncompressedSize = 1024 * 1024 * 10 // 10 mb
	reader := &countReader{
		reader: io.LimitReader(r, maxUncompressedSize),
	}
	if g.PackedData, err = io.ReadAll(reader); err != nil {
		return fmt.Errorf("unable to decode gzip_packed: decompress packed_data: %w", err)
	}
	if reader.Total() >= maxUncompressedSize {
		// Read limit reached, possible decompression bomb detected.
		return fmt.Errorf("unable to decode gzip_packed: decompress packed_data: %w", &DecompressionBombErr{
			Compressed:   maxUncompressedSize,
			Decompressed: int(reader.Total()),
		})
	}

	g.Obj, err = iface.DecodeObject(bin.NewDecoder(g.PackedData))
	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: field object: %w", err)
	}

	return nil
}
