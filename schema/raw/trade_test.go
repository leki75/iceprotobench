package raw

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTrade() *Trade {
	now := time.Now()
	return &Trade{
		Id:         1000,
		Timestamp:  uint64(now.UnixNano()),
		Price:      987654321,
		Volume:     123456,
		Conditions: [4]byte{'W', 'X', 'Y', 'Z'},
		Symbol:     *(*[11]byte)([]byte("12345678901")),
		Exchange:   'V',
		Tape:       'A',
		ReceivedAt: uint64(now.Add(-1 * time.Second).UnixNano()),
	}
}

func TestTradeMarshal(t *testing.T) {
	trade := newTrade()
	b := trade.Marshal()
	require.Equal(t, 54, len(b))

	got := &Trade{}
	err := got.Unmarshal(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, trade)
}

func TestTradeMarshalUnsafe(t *testing.T) {
	trade := newTrade()
	b := trade.MarshalUnsafe()
	require.Equal(t, 54, len(b))

	got := &Trade{}
	err := got.UnmarshalUnsafe(b)
	require.NoError(t, err)
	assert.EqualValues(t, got, trade)
}

func TestTradeMarshalCompare(t *testing.T) {
	trade := newTrade()
	b1 := trade.Marshal()
	b2 := trade.MarshalUnsafe()
	assert.EqualValues(t, b1, b2)
}
