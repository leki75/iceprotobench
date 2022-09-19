//go:build protovt
// +build protovt

package benchmark

import (
	"math"
	"testing"
	"time"

	"github.com/leki75/iceprotobench/schema"
)

var (
	conditions = []byte{'A', 'B', 'C', 'D'}
	symbol     = "12345678901"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := &schema.ProtoTrade{
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Tape:       'A',
	}
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := uint64(time.Now().UnixNano())
		trade.Id = uint64(i)
		trade.Timestamp = now
		trade.Exchange = 'A' + int32(i%26)
		trade.ReceivedAt = int64(now)

		data, err := trade.MarshalVT()
		if err != nil {
			b.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	trade := &schema.ProtoTrade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: int64(now),
	}
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
	quote := &schema.ProtoQuote{
		BidPrice:  math.MaxFloat64,
		AskPrice:  math.MaxFloat64,
		BidSize:   math.MaxUint32,
		AskSize:   math.MaxUint32,
		Condition: '@',
		Symbol:    symbol,
		Tape:      'C',
	}
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := uint64(time.Now().UnixNano())
		exchange := 'A' + int32(i%26)
		quote.Timestamp = now
		quote.BidExchange = exchange
		quote.AskExchange = exchange
		quote.Nbbo = (i % 2) == 0
		quote.ReceivedAt = int64(now)

		data, err := quote.MarshalVT()
		if err != nil {
			b.Fatal("quote marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	quote := &schema.ProtoQuote{
		Timestamp:   now,
		BidPrice:    math.MaxFloat64,
		AskPrice:    math.MaxFloat64,
		BidSize:     math.MaxUint32,
		AskSize:     math.MaxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Condition:   '@',
		Nbbo:        false,
		Symbol:      symbol,
		Tape:        'C',
		ReceivedAt:  int64(now),
	}
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
