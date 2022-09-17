//go:build raw
// +build raw

package schema

import (
	"log"
	"math"
	"testing"
	"time"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := &RawTrade{
		Price:      math.MaxUint64,
		Volume:     math.MaxUint32,
		Conditions: [4]byte{'A', 'B', 'C', 'D'},
		Symbol:     *(*[11]byte)([]byte("12345678901")),
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
	trade := &RawTrade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxUint64,
		Volume:     math.MaxUint32,
		Conditions: [4]byte{'A', 'B', 'C', 'D'},
		Symbol:     *(*[11]byte)([]byte("12345678901")),
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
	data := trade.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data[43] = 'A' + byte(i)
		result := &RawTrade{}
		if err := result.Unmarshal(data); err != nil || result.Exchange != 'A'+byte(i) {
			log.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := &RawQuote{
		BidPrice:  math.MaxUint64,
		AskPrice:  math.MaxUint64,
		BidSize:   math.MaxUint32,
		AskSize:   math.MaxUint32,
		Condition: '@',
		Symbol:    *(*[11]byte)([]byte("12345678901")),
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
		quote.CreatedAt = now
		data := quote.Marshal()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	quote := &RawQuote{
		Timestamp:   now,
		BidPrice:    math.MaxUint64,
		AskPrice:    math.MaxUint64,
		BidSize:     math.MaxUint32,
		AskSize:     math.MaxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Condition:   '@',
		Nbbo:        false,
		Symbol:      *(*[11]byte)([]byte("12345678901")),
		Tape:        'C',
		CreatedAt:   now,
	}
	data := quote.Marshal()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data[32] = 'A' + byte(i)
		result := &RawQuote{}
		if err := result.Unmarshal(data); err != nil || result.BidExchange != 'A'+byte(i) {
			log.Fatal("quote unmarshal failed")
		}
	}
}
