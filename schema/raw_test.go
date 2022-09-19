package schema

import (
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
		ReceivedAt:  uint64(now.Add(-1 * time.Second).UnixNano()),
	}

	b := quote.Marshal()
	require.Equal(t, len(b), 56)

	got := &RawQuote{}
	err := got.Unmarshal(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, quote)
}
