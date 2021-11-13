package schema

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawTradeMarshalUnmarshal(t *testing.T) {
	now := time.Now()
	trade := &RawTrade{
		Id:         1000,
		Timestamp:  uint64(now.UnixNano()),
		Price:      987654321,
		Volume:     123456,
		Conditions: [4]byte{'@', 0, 0, 0},
		Symbol:     *(*[11]byte)([]byte("AAPL       ")),
		Exchange:   'V',
		Tape:       'A',
		ReceivedAt: uint64(now.Add(-1 * time.Second).UnixNano()),
	}

	b := trade.Marshal()
	require.Equal(t, len(b), 53)

	got := &RawTrade{}
	err := got.Unmarshal(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, trade)
}

func TestRawQuoteMarshalUnmarshal(t *testing.T) {
	now := time.Now()
	quote := &RawQuote{
		Timestamp:   uint64(now.UnixNano()),
		BidPrice:    123456789,
		AskPrice:    987654321,
		BidSize:     123456,
		AskSize:     654321,
		BidExchange: 'B',
		AskExchange: 'A',
		Condition:   '@',
		Nbbo:        true,
		Symbol:      *(*[11]byte)([]byte("AAPL       ")),
		Tape:        'C',
		CreatedAt:   uint64(now.Add(-1 * time.Second).UnixNano()),
	}

	b := quote.Marshal()
	require.Equal(t, len(b), 56)

	got := &RawQuote{}
	err := got.Unmarshal(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, quote)
}

func BenchmarkRawTrade_Marshal(b *testing.B) {
	trade := &RawTrade{
		Price:      maxUint64,
		Volume:     maxUint32,
		Conditions: [4]byte{'A', 'B', 'C', 'D'},
		Symbol:     *(*[11]byte)([]byte("12345678901")),
		Tape:       'A',
	}
	size := 0

	b.ReportAllocs()
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

func BenchmarkRawTrade_Unmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	trade := &RawTrade{
		Id:         maxUint64,
		Timestamp:  now,
		Price:      maxUint64,
		Volume:     maxUint32,
		Conditions: [4]byte{'A', 'B', 'C', 'D'},
		Symbol:     *(*[11]byte)([]byte("12345678901")),
		Exchange:   '!',
		Tape:       'A',
		ReceivedAt: now,
	}
	data := trade.Marshal()
	result := &RawTrade{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data[43] = 'A' + byte(i)
		if err := result.Unmarshal(data); err != nil || result.Exchange != 'A'+byte(i) {
			log.Fatal("trade unmarshal failed")
		}
	}
}

func BenchmarkRawQuote_Marshal(b *testing.B) {
	quote := &RawQuote{
		BidPrice:  maxUint64,
		AskPrice:  maxUint64,
		BidSize:   maxUint32,
		AskSize:   maxUint32,
		Condition: '@',
		Symbol:    *(*[11]byte)([]byte("12345678901")),
		Tape:      'C',
	}
	size := 0

	b.ReportAllocs()
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

func BenchmarkRawQuote_Unmarshal(b *testing.B) {
	now := uint64(time.Now().UnixNano())
	quote := &RawQuote{
		Timestamp:   now,
		BidPrice:    maxUint64,
		AskPrice:    maxUint64,
		BidSize:     maxUint32,
		AskSize:     maxUint32,
		BidExchange: '!',
		AskExchange: '!',
		Condition:   '@',
		Nbbo:        false,
		Symbol:      *(*[11]byte)([]byte("12345678901")),
		Tape:        'C',
		CreatedAt:   now,
	}
	data := quote.Marshal()
	result := &RawQuote{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data[32] = 'A' + byte(i)
		if err := result.Unmarshal(data); err != nil || result.BidExchange != 'A'+byte(i) {
			log.Fatal("quote unmarshal failed")
		}
	}
}
