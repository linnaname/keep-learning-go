package main

import "testing"

const BLOG = "www.linnana.me 好大一只电风扇"

func initStrings(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func initStringi(N int) []interface{} {
	s := make([]interface{}, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func BenchmarkStringPlus10(b *testing.B) {
	p := initStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt10(b *testing.B) {
	p := initStringi(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin10(b *testing.B) {
	p := initStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer10(b *testing.B) {
	p := initStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder10(b *testing.B) {
	p := initStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p)
	}
}

func BenchmarkStringPlus100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt100(b *testing.B) {
	p := initStringi(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p)
	}
}

func BenchmarkStringBuilderGrowFirst100(b *testing.B) {
	p := initStrings(100)
	cap := 100 * len(BLOG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilderGrowFirst(p, cap)
	}
}

func BenchmarkStringPlus1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt1000(b *testing.B) {
	p := initStringi(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p)
	}
}

func BenchmarkStringBuilderGrowFirst1000(b *testing.B) {
	p := initStrings(1000)
	cap := 1000 * len(BLOG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilderGrowFirst(p, cap)
	}
}

func BenchmarkStringPlus10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt10000(b *testing.B) {
	p := initStringi(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p)
	}
}

func BenchmarkStringBuilderGrowFirst10000(b *testing.B) {
	p := initStrings(10000)
	cap := 10000 * len(BLOG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilderGrowFirst(p, cap)
	}
}
