package raw

import (
	"testing"
)

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := newQuote()
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quote.Timestamp = uint64(i)
		data := quote.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	quote := newQuote()
	data := quote.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &Quote{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("quote unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshalUnsafe(b *testing.B) {
	quote := newQuote()
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quote.Timestamp = uint64(i)
		data := quote.MarshalUnsafe()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshalUnsafe(b *testing.B) {
	quote := newQuote()
	data := quote.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &Quote{}
		if err := result.UnmarshalUnsafe(data); err != nil {
			b.Fatal("quote unmarshal failed")
		}
	}
}
