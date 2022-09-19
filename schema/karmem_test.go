package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	karmem "karmem.org/golang"
)

func TestKarmemTradeMarshalUnmarshal(t *testing.T) {
	now := time.Now()
	trade := &KarmemTrade{
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

	// Marshal
	writer := karmem.NewWriter(64)
	_, err := trade.WriteAsRoot(writer)
	require.NoError(t, err)
	data := writer.Bytes()

	// Unmarshal
	reader := karmem.NewReader(data)
	got := &KarmemTrade{}
	got.ReadAsRoot(reader)

	assert.EqualValues(t, got, trade)
}

func TestKarmemQuoteMarshalUnmarshal(t *testing.T) {
	now := time.Now()
	quote := &KarmemQuote{
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

	// Marshal
	writer := karmem.NewWriter(64)
	_, err := quote.WriteAsRoot(writer)
	require.NoError(t, err)
	data := writer.Bytes()

	// Unmarshal
	reader := karmem.NewReader(data)
	got := &KarmemQuote{}
	got.ReadAsRoot(reader)

	assert.EqualValues(t, got, quote)
}
