//go:build protovt
// +build protovt

package benchmark

import (
	"testing"

	"github.com/leki75/iceprotobench/schema"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := createProtoTrade()
	size := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trade.Timestamp = uint64(i)
		data, err := trade.MarshalVT()
		if err != nil {
			b.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	trade := createProtoTrade()
	data, err := trade.MarshalVT()
	if err != nil {
		b.Fatal("trade marshal failed")
	}
	result := &schema.ProtoTrade{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			b.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := createProtoQuote()
	size := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quote.Timestamp = uint64(i)
		data, err := quote.MarshalVT()
		if err != nil {
			b.Fatal("quote marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	quote := createProtoQuote()
	data, err := quote.MarshalVT()
	if err != nil {
		b.Fatal("quote marshal failed")
	}
	result := &schema.ProtoQuote{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			b.Fatal("quote unmarshal failed")
		}
	}
}
