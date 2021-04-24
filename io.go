package util

import (
	"context"
	"io"
)

type wrapReader struct {
	ctx context.Context
	io.Reader
}

func (r wrapReader) Read(p []byte) (int, error) {
	select {
	case <-r.ctx.Done():
		return 0, r.ctx.Err()
	default:
		return r.Read(p)
	}
}

// Copy is similar to io.Copy but context aware.
func Copy(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	wrapSrc := wrapReader{ctx, src}
	return io.Copy(dst, wrapSrc)
}

// CopyBuffer is similar to io.CopyBuffer but context aware.
func CopyBuffer(ctx context.Context, dst io.Writer, src io.Reader, buf []byte) (int64, error) {
	wrapSrc := wrapReader{ctx, src}
	return io.CopyBuffer(dst, wrapSrc, buf)
}

// CopyN is similar to io.CopyN but context aware.
func CopyN(ctx context.Context, dst io.Writer, src io.Reader, n int64) (int64, error) {
	wrapSrc := wrapReader{ctx, src}
	return io.CopyN(dst, wrapSrc, n)
}

// ReadAll is similar to io.ReadAll but context aware.
func ReadAll(ctx context.Context, r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		select {
		case <-ctx.Done():
			return b, ctx.Err()
		default:
		}
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}
