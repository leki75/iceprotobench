//go:build proto
// +build proto

package schema

import (
	"log"
	"math"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
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
		trade.Id = uint64(i)
		trade.Timestamp = uint64(time.Now().UnixNano())
		trade.Exchange = 'A' + int32(i%26)

		data, err := proto.Marshal(trade)
		if err != nil {
			log.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	trade := &ProtoTrade{
		Id:         math.MaxUint64,
		Timestamp:  uint64(time.Now().UnixNano()),
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
	}
	data, err := proto.Marshal(trade)
	if err != nil {
		log.Fatal("trade marshal failed")
	}
	result := &ProtoTrade{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = proto.Unmarshal(data, result); err != nil {
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
		exchange := 'A' + int32(i%26)
		quote.Timestamp = uint64(time.Now().UnixNano())
		quote.BidExchange = exchange
		quote.AskExchange = exchange
		quote.Nbbo = (i % 2) == 0

		data, err := proto.Marshal(quote)
		if err != nil {
			log.Fatal("trade marshal failed")
		}
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	quote := &ProtoQuote{
		Timestamp:   uint64(time.Now().UnixNano()),
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
	}
	data, err := proto.Marshal(quote)
	if err != nil {
		log.Fatal("quote marshal failed")
	}
	result := &ProtoQuote{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = proto.Unmarshal(data, result); err != nil {
			log.Fatal("trade unmarshal failed")
		}
	}
}
