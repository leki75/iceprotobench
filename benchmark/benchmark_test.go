package benchmark

import (
	"math"
	"time"

	"github.com/leki75/iceprotobench/schema"
	"github.com/leki75/iceprotobench/schema/raw"
)

var (
	symbol          = "12345678901"
	tradeConditions = []byte{'A', 'B', 'C', 'D'}
	quoteConditions = []byte{'E', 'F'}
)

func createProtoTrade() *schema.ProtoTrade {
	now := uint64(time.Now().UnixNano())
	return &schema.ProtoTrade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: tradeConditions,
		Symbol:     symbol,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: int64(now),
	}
}

func createProtoQuote() *schema.ProtoQuote {
	now := uint64(time.Now().UnixNano())
	return &schema.ProtoQuote{
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
}

func createKarmemTrade() *schema.KarmemTrade {
	sym := *(*[11]byte)([]byte(symbol))
	tc := *(*[4]byte)(tradeConditions)
	now := uint64(time.Now().UnixNano())
	return &schema.KarmemTrade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: tc,
		Symbol:     sym,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
}

func createKarmemQuote() *schema.KarmemQuote {
	sym := *(*[11]byte)([]byte(symbol))
	qc := *(*[2]byte)(quoteConditions)
	now := uint64(time.Now().UnixNano())
	return &schema.KarmemQuote{
		Timestamp:   now,
		BidPrice:    math.MaxFloat64,
		AskPrice:    math.MaxFloat64,
		BidSize:     math.MaxUint32,
		AskSize:     math.MaxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Conditions:  qc,
		Nbbo:        false,
		Symbol:      sym,
		Tape:        'C',
		ReceivedAt:  now,
	}
}

func createRawTrade() *raw.Trade {
	sym := *(*[11]byte)([]byte(symbol))
	tc := *(*[4]byte)(tradeConditions)
	now := uint64(time.Now().UnixNano())
	return &raw.Trade{
		Id:         math.MaxUint64,
		Timestamp:  now,
		Price:      math.MaxFloat64,
		Volume:     math.MaxUint32,
		Conditions: tc,
		Symbol:     sym,
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
}

func createRawQuote() *raw.Quote {
	sym := *(*[11]byte)([]byte(symbol))
	qc := *(*[2]byte)(quoteConditions)
	now := uint64(time.Now().UnixNano())
	return &raw.Quote{
		Timestamp:   now,
		BidPrice:    math.MaxFloat64,
		AskPrice:    math.MaxFloat64,
		BidSize:     math.MaxUint32,
		AskSize:     math.MaxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Conditions:  qc,
		Nbbo:        false,
		Symbol:      sym,
		Tape:        'C',
		ReceivedAt:  now,
	}
}
