//go:build protovt
// +build protovt

package schema

import (
	"log"
	"math"
	"testing"
	"time"
)

var (
	conditions = []byte{'A', 'B', 'C', 'D'}
	symbol     = "12345678901"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := &ProtoTrade{
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Tape:       'A',
	}
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now().UnixNano()
		trade.Id = uint64(i)
		trade.Timestamp = uint64(now)
		trade.Exchange = 'A' + int32(i%26)
		trade.ReceivedAt = now

		data, err := trade.MarshalVT()
		if err != nil {
			log.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	now := time.Now().UnixNano()
	trade := &ProtoTrade{
		Id:         math.MaxUint64,
		Timestamp:  uint64(now),
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
	data, err := trade.MarshalVT()
	if err != nil {
		log.Fatal("trade marshal failed")
	}
	result := &ProtoTrade{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			log.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := &ProtoQuote{
		BidPrice:   math.MaxFloat64,
		AskPrice:   math.MaxFloat64,
		BidSize:    math.MaxUint32,
		AskSize:    math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Tape:       'C',
	}
	size := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now().UnixNano()
		exchange := 'A' + int32(i%26)
		quote.Timestamp = uint64(now)
		quote.BidExchange = exchange
		quote.AskExchange = exchange
		quote.Nbbo = (i % 2) == 0
		quote.ReceivedAt = now

		data, err := quote.MarshalVT()
		if err != nil {
			log.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	now := time.Now().UnixNano()
	quote := &ProtoQuote{
		Timestamp:   uint64(now),
		BidPrice:    math.MaxFloat64,
		AskPrice:    math.MaxFloat64,
		BidSize:     math.MaxUint32,
		AskSize:     math.MaxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Conditions:  conditions,
		Nbbo:        false,
		Symbol:      symbol,
		Tape:        'C',
		ReceivedAt:  now,
	}
	data, err := quote.MarshalVT()
	if err != nil {
		log.Fatal("quote marshal failed")
	}
	result := &ProtoQuote{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			log.Fatal("quote unmarshal failed")
		}
	}
}
