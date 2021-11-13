package schema

import (
	"log"
	"testing"
	"time"
)

func BenchmarkProtoTrade_MarshalVT(b *testing.B) {
	trade := &ProtoTrade{
		Price:      float64(maxUint64),
		Volume:     maxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Tape:       'A',
	}
	size := 0

	b.ReportAllocs()
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

func BenchmarkProtoTrade_UnmarshalVT(b *testing.B) {
	now := time.Now().UnixNano()
	trade := &ProtoTrade{
		Id:         maxUint64,
		Timestamp:  uint64(now),
		Price:      float64(maxUint64),
		Volume:     maxUint32,
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

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			log.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkProtoQuote_MarshalVT(b *testing.B) {
	quote := &ProtoQuote{
		BidPrice:   float64(maxUint64),
		AskPrice:   float64(maxUint64),
		BidSize:    maxUint32,
		AskSize:    maxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Tape:       'C',
	}
	size := 0

	b.ReportAllocs()
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

func BenchmarkProtoQuote_UnmarshalVT(b *testing.B) {
	now := time.Now().UnixNano()
	quote := &ProtoQuote{
		Timestamp:   uint64(now),
		BidPrice:    float64(maxUint64),
		AskPrice:    float64(maxUint64),
		BidSize:     maxUint32,
		AskSize:     maxUint32,
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

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = result.UnmarshalVT(data); err != nil {
			log.Fatal("quote unmarshal failed")
		}
	}
}
