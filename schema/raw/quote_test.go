package raw

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newQuote() *Quote {
	now := time.Now()
	return &Quote{
		Timestamp:   uint64(now.UnixNano()),
		BidPrice:    123456789,
		AskPrice:    987654321,
		BidSize:     123456,
		AskSize:     654321,
		BidExchange: 'B',
		AskExchange: 'A',
		Conditions:  [2]byte{'Y', 'Z'},
		Nbbo:        true,
		Symbol:      *(*[11]byte)([]byte("12345678901")),
		Tape:        'C',
		ReceivedAt:  uint64(now.Add(-1 * time.Second).UnixNano()),
	}
}

func TestQuoteMarshal(t *testing.T) {
	quote := newQuote()

	b := quote.Marshal()
	require.Equal(t, 58, len(b))

	got := &Quote{}
	err := got.Unmarshal(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, quote)
}

func TestQuoteMarshalUnsafe(t *testing.T) {
	quote := newQuote()

	b := quote.MarshalUnsafe()
	require.Equal(t, 58, len(b))

	got := &Quote{}
	err := got.UnmarshalUnsafe(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, quote)
}

func TestQuoteMarshalCompare(t *testing.T) {
	quote := newQuote()
	b1 := quote.Marshal()
	b2 := quote.MarshalUnsafe()
	assert.EqualValues(t, b1, b2)
}
