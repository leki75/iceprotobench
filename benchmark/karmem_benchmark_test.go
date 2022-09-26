//go:build karmem
// +build karmem

package benchmark

import (
	"testing"

	karmem "karmem.org/golang"

	"github.com/leki75/iceprotobench/schema"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := createKarmemTrade()
	size := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trade.Timestamp = uint64(i)
		writer := karmem.NewWriter(64)
		if _, err := trade.WriteAsRoot(writer); err != nil {
			b.Fatal("trade marshal failed")
		}
		data := writer.Bytes()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	trade := createKarmemTrade()
	writer := karmem.NewWriter(64)
	if _, err := trade.WriteAsRoot(writer); err != nil {
		b.Fatal("trade marshal failed")
	}
	data := writer.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := karmem.NewReader(data)
		trade := &schema.KarmemTrade{}
		trade.ReadAsRoot(reader)
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := createKarmemQuote()
	size := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quote.Timestamp = uint64(i)
		writer := karmem.NewWriter(64)
		if _, err := quote.WriteAsRoot(writer); err != nil {
			b.Fatal("quote marshal failed")
		}
		data := writer.Bytes()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	quote := createKarmemQuote()
	writer := karmem.NewWriter(64)
	if _, err := quote.WriteAsRoot(writer); err != nil {
		b.Fatal("quote marshal failed")
	}
	data := writer.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := karmem.NewReader(data)
		quote := &schema.KarmemQuote{}
		quote.ReadAsRoot(reader)
	}
}
