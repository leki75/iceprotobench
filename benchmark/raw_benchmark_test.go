//go:build raw
// +build raw

package benchmark

import (
	"testing"

	"github.com/leki75/iceprotobench/schema/raw"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := createRawTrade()
	size := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trade.Timestamp = uint64(i)
		data := trade.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	trade := createRawTrade()
	data := trade.Marshal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &raw.Trade{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := createRawQuote()
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
	quote := createRawQuote()
	data := quote.Marshal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &raw.Quote{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("quote unmarshal failed")
		}
	}
}
