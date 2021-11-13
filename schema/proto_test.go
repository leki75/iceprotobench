package schema

import (
	"log"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

const (
	maxUint64 = ^uint64(0)
	maxUint32 = ^uint32(0)
)

var (
	conditions = []byte{'A', 'B', 'C', 'D'}
	symbol     = "12345678901"
)

func BenchmarkProtoTrade_Marshal(b *testing.B) {
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

func BenchmarkProtoTrade_Unmarshal(b *testing.B) {
	trade := &ProtoTrade{
		Id:         maxUint64,
		Timestamp:  uint64(time.Now().UnixNano()),
		Price:      float64(maxUint64),
		Volume:     maxUint32,
		Conditions: conditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
	}
	data, err := trade.MarshalVT()
	if err != nil {
		log.Fatal("trade marshal failed")
	}
	result := &ProtoTrade{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = proto.Unmarshal(data, result); err != nil {
			log.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkProtoQuote_Marshal(b *testing.B) {
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

func BenchmarkProtoQuote_Unmarshal(b *testing.B) {
	quote := &ProtoQuote{
		Timestamp:   uint64(time.Now().UnixNano()),
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
	}
	data, err := quote.MarshalVT()
	if err != nil {
		log.Fatal("quote marshal failed")
	}
	result := &ProtoQuote{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = proto.Unmarshal(data, result); err != nil {
			log.Fatal("trade unmarshal failed")
		}
	}
}
