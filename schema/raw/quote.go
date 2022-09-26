package raw

import (
	"encoding/binary"
	"errors"
	"math"
	"unsafe"
)

var ErrInvalidQuoteLength = errors.New("invalid quote length")

// Quote
// -------------|--------------|-------------|-----------
//  Field       | NASDAQ short | NASDAQ long | NYSE long
// -------------|--------------|-------------|-----------
//  symbol      | [5]byte      | [11]byte    | [11]byte
//  condition   | byte         | byte        | byte
//  bidExchange | byte         | byte        | byte
//  askExchange | byte         | byte        | byte
//  timestamp   | uint64       | uint64      | uint64
//  bidPrice    | uint16       | uint64      | uint64
//  askPrice    | uint16       | uint64      | uint64
//  bidSize     | uint16       | uint32      | uint32
//  askSize     | uint16       | uint32      | uint32
//  nbbo        | bool         | bool        | bool
// -------------|--------------|-------------|-----------

type Quote struct {
	_           byte     //     0 - placeholder for type
	Symbol      [11]byte //  1-11
	Conditions  [2]byte  // 12-13
	BidExchange byte     //    14
	AskExchange byte     //    15
	Timestamp   uint64   // 16-23
	ReceivedAt  uint64   // 24-31 - not in the original Quote message
	BidPrice    float64  // 32-39
	AskPrice    float64  // 40-47
	BidSize     uint32   // 48-51
	AskSize     uint32   // 52-55
	Nbbo        bool     //    56
	Tape        byte     //    57 - not in the original Quote message
}

func (q *Quote) Marshal() []byte {
	b := make([]byte, 58)
	b[57] = q.Tape // go bounds check elimination
	b[0] = 'q'
	copy(b[1:], q.Symbol[:])
	copy(b[12:], q.Conditions[:])
	b[14] = q.BidExchange
	b[15] = q.AskExchange
	binary.LittleEndian.PutUint64(b[16:], q.Timestamp)
	binary.LittleEndian.PutUint64(b[24:], q.ReceivedAt)
	binary.LittleEndian.PutUint64(b[32:], math.Float64bits(q.BidPrice))
	binary.LittleEndian.PutUint64(b[40:], math.Float64bits(q.AskPrice))
	binary.LittleEndian.PutUint32(b[48:], q.BidSize)
	binary.LittleEndian.PutUint32(b[52:], q.AskSize)
	if q.Nbbo {
		b[56] = 1
	}
	return b
}

func (q *Quote) Unmarshal(b []byte) error {
	if len(b) != 58 {
		return ErrInvalidQuoteLength
	}
	q.Tape = b[57] // go bounds check elimination
	q.Symbol = *(*[11]byte)(b[1:])
	q.Conditions = *(*[2]byte)(b[12:])
	q.BidExchange = b[14]
	q.AskExchange = b[15]
	q.Timestamp = binary.LittleEndian.Uint64(b[16:])
	q.ReceivedAt = binary.LittleEndian.Uint64(b[24:])
	q.BidPrice = math.Float64frombits(binary.LittleEndian.Uint64(b[32:]))
	q.AskPrice = math.Float64frombits(binary.LittleEndian.Uint64(b[40:]))
	q.BidSize = binary.LittleEndian.Uint32(b[48:])
	q.AskSize = binary.LittleEndian.Uint32(b[52:])
	q.Nbbo = b[56] == 1
	return nil
}

// Marshal will marshal Quote object into its binary representation.
// The marshaler and unmarshaler node must have the same architecture!
func (q *Quote) MarshalUnsafe() []byte {
	b := make([]byte, 58)
	copy(b, (*(*[58]byte)(unsafe.Pointer(q)))[:])
	b[0] = 'q'
	return b
}

// Unmarshal will unmarshal the byte slice into Quote object.
// The marshaler and unmarshaler node must have the same architecture!
func (q *Quote) UnmarshalUnsafe(b []byte) error {
	if len(b) != 58 {
		return ErrInvalidQuoteLength
	}
	b[0] = 0
	copy((*(*[58]byte)(unsafe.Pointer(q)))[:], b)
	return nil
}
