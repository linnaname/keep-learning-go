package gzip

import "testing"

func TestRequestGzip(t *testing.T) {
	length, err := RequestGzip(true) // 开启 gzip
	if err != nil {
		t.Error(err)
	}
	t.Logf("Request with gzip length: %d", length)

	length, err = RequestGzip(false) // 未开启 gzip
	if err != nil {
		t.Error(err)
	}
	t.Logf("Request without gzip length: %d", length)
}

func BenchmarkClientUncompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClientUncompress()
	}
}

func BenchmarkClientNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClientNormal()
	}
}
