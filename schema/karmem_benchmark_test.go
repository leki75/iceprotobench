//go:build karmem
// +build karmem

package schema

import (
	"log"
	"math"
	"testing"
	"time"

	karmem "karmem.org/golang"
)

func BenchmarkTradeMarshal(b *testing.B) {
	trade := &KarmemTrade{
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

		writer := karmem.NewWriter(64)
		if _, err := trade.WriteAsRoot(writer); err != nil {
			log.Fatal("trade marshal failed")
		}
		data := writer.Bytes()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkTradeUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	trade := &KarmemTrade{
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
	// Marshal
	writer := karmem.NewWriter(64)
	if _, err := trade.WriteAsRoot(writer); err != nil {
		log.Fatal("trade marshal failed")
	}
	data := writer.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Unmarshal
		reader := karmem.NewReader(data)
		trade := &KarmemTrade{}
		trade.ReadAsRoot(reader)
	}
}

func BenchmarkQuoteMarshal(b *testing.B) {
	quote := &KarmemQuote{
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
		writer := karmem.NewWriter(64)
		if _, err := quote.WriteAsRoot(writer); err != nil {
			log.Fatal("quote marshal failed")
		}
		data := writer.Bytes()
		size += len(data)
	}
	b.ReportMetric(float64(size)/float64(b.N), "B/obj")
}

func BenchmarkQuoteUnmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	quote := &KarmemQuote{
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
	// Marshal
	writer := karmem.NewWriter(64)
	if _, err := quote.WriteAsRoot(writer); err != nil {
		log.Fatal("trade marshal failed")
	}
	data := writer.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Unmarshal
		reader := karmem.NewReader(data)
		quote := &KarmemQuote{}
		quote.ReadAsRoot(reader)
	}
}
