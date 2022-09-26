package raw

import (
	"testing"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := newTrade()
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trade.Id = uint64(i)
		data := trade.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	trade := newTrade()
	data := trade.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &Trade{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkTradeMarshalUnsafe(b *testing.B) {
	trade := newTrade()
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trade.Id = uint64(i)
		data := trade.MarshalUnsafe()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshalUnsafe(b *testing.B) {
	trade := newTrade()
	data := trade.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &Trade{}
		if err := result.UnmarshalUnsafe(data); err != nil {
			b.Fatal("trade unmarshal failed")
		}
	}
}
