package raw

import (
	"encoding/binary"
	"errors"
	"math"
	"unsafe"
)

var ErrInvalidTradeLength = errors.New("invalid trade length")

// Trade
// -----------|--------------|-------------|------------|-----------
//	Field     | NASDAQ short | NASDAQ long | NYSE short | NYSE long
// -----------|--------------|-------------|------------|-----------
//	symbol    | [5]byte      | [11]byte    | [5]byte    | [11]byte
//	cond      | [4]byte      | [4]byte     | byte       | [4]byte
//	tradeId   | uint64       | uint64      | int64      | int64
//	timestamp | uint64       | uint64      | uint64     | uint64
//	price     | uint16       | uint64      | uint16     | uint64
//	volume    | uint16       | uint32      | uint16     | uint32
//	exchange  | byte         | byte        | byte       | byte
// -----------|--------------|-------------|------------|-----------

type Trade struct {
	_          byte     //     0 - placeholder for type
	Symbol     [11]byte //  1-11
	Conditions [4]byte  // 12-15
	Id         uint64   // 16-23
	Timestamp  uint64   // 24-31
	ReceivedAt uint64   // 32-39 - not in the original Trade message
	Price      float64  // 40-47
	Volume     uint32   // 48-51
	Exchange   byte     //    52
	Tape       byte     //    53 - not in the original Trade message
}

func (t *Trade) Marshal() []byte {
	b := make([]byte, 54)
	b[53] = t.Tape // go bounds check elimination
	b[0] = 't'
	copy(b[1:], t.Symbol[:])
	copy(b[12:], t.Conditions[:])
	binary.LittleEndian.PutUint64(b[16:], t.Id)
	binary.LittleEndian.PutUint64(b[24:], t.Timestamp)
	binary.LittleEndian.PutUint64(b[32:], t.ReceivedAt)
	binary.LittleEndian.PutUint64(b[40:], math.Float64bits(t.Price))
	binary.LittleEndian.PutUint32(b[48:], t.Volume)
	b[52] = t.Exchange
	return b
}

func (t *Trade) Unmarshal(b []byte) error {
	if len(b) != 54 {
		return ErrInvalidTradeLength
	}
	t.Tape = b[53] // go bounds check elimination
	t.Symbol = *(*[11]byte)(b[1:])
	t.Conditions = *(*[4]byte)(b[12:])
	t.Id = binary.LittleEndian.Uint64(b[16:])
	t.Timestamp = binary.LittleEndian.Uint64(b[24:])
	t.ReceivedAt = binary.LittleEndian.Uint64(b[32:])
	t.Price = math.Float64frombits(binary.LittleEndian.Uint64(b[40:]))
	t.Volume = binary.LittleEndian.Uint32(b[48:])
	t.Exchange = b[52]
	return nil
}

// MarshalUnsafe will marshal Trade object into its binary representation.
// The marshaler and unmarshaler node must have the same architecture!
func (t *Trade) MarshalUnsafe() []byte {
	b := make([]byte, 54)
	copy(b, (*(*[54]byte)(unsafe.Pointer(t)))[:])
	b[0] = 't'
	return b
}

// UnmarshalUnsafe will unmarshal the byte slice into Trade object.
// The marshaler and unmarshaler node must have the same architecture!
func (t *Trade) UnmarshalUnsafe(b []byte) error {
	if len(b) != 54 {
		return ErrInvalidTradeLength
	}
	b[0] = 0
	copy((*(*[54]byte)(unsafe.Pointer(t)))[:], b)
	return nil
}
