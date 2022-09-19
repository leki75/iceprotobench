//go:build raw
// +build raw

package benchmark

import (
	"math"
	"testing"
	"time"

	"github.com/leki75/iceprotobench/schema"
)

var (
	conditions = [4]byte{'A', 'B', 'C', 'D'}
	symbol     = *(*[11]byte)([]byte("12345678901"))
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := &schema.RawTrade{
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
		trade.Exchange = 'A' + byte(i%26)
		trade.ReceivedAt = now

		data := trade.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	trade := &schema.RawTrade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
	data := trade.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &schema.RawTrade{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := &schema.RawQuote{
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
		exchange := 'A' + byte(i%26)
		quote.Timestamp = now
		quote.BidExchange = exchange
		quote.AskExchange = exchange
		quote.Nbbo = (i % 2) == 0
		quote.ReceivedAt = now

		data := quote.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	quote := &schema.RawQuote{
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
		ReceivedAt:  now,
	}
	data := quote.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &schema.RawQuote{}
		if err := result.Unmarshal(data); err != nil {
			b.Fatal("quote unmarshal failed")
		}
	}
}
