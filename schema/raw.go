package schema

import (
	"encoding/binary"
	"errors"
)

// RawTrade
// -----------|--------------|-------------|------------|-----------
//
//	Field     | NASDAQ short | NASDAQ long | NYSE short | NYSE long
//
// -----------|--------------|-------------|------------|-----------
//
//	timestamp | uint64       | uint64      | uint64     | uint64
//	symbol    | [5]byte      | [11]byte    | [5]byte    | [11]byte
//	tradeId   | uint64       | uint64      | int64      | int64
//	price     | uint16       | uint64      | uint16     | uint64
//	volume    | uint16       | uint32      | uint16     | uint32
//	cond      | [4]byte      | [4]byte     | byte       | [4]byte
//	exchange  | byte         | byte        | byte       | byte
//
// -----------|--------------|-------------|------------|-----------
type RawTrade struct {
	Id         uint64   //  8
	Timestamp  uint64   // 16
	Price      uint64   // 24
	Volume     uint32   // 28
	Conditions [4]byte  // 32
	Symbol     [11]byte // 43
	Exchange   byte     // 44
	Tape       byte     // 45 - not in the original Trade message
	ReceivedAt uint64   // 53 - not in the original Trade message
}

func (t *RawTrade) Marshal() []byte {
	b := make([]byte, 53)

	binary.LittleEndian.PutUint64(b[45:], t.ReceivedAt) // go bounds check elimination
	binary.LittleEndian.PutUint64(b[0:], t.Id)
	binary.LittleEndian.PutUint64(b[8:], t.Timestamp)
	binary.LittleEndian.PutUint64(b[16:], t.Price)
	binary.LittleEndian.PutUint32(b[24:], t.Volume)
	for i := 0; i < 4; i++ {
		b[i+28] = t.Conditions[i]
	}
	for i := 0; i < 11; i++ {
		b[i+32] = t.Symbol[i]
	}
	b[43] = t.Exchange
	b[44] = t.Tape

	return b
}

func (t *RawTrade) Unmarshal(b []byte) error {
	if len(b) != 53 {
		return errors.New("invalid trade length")
	}

	t.ReceivedAt = binary.LittleEndian.Uint64(b[45:]) // go bounds check elimination
	t.Id = binary.LittleEndian.Uint64(b[0:])
	t.Timestamp = binary.LittleEndian.Uint64(b[8:])
	t.Price = binary.LittleEndian.Uint64(b[16:])
	t.Volume = binary.LittleEndian.Uint32(b[24:])
	t.Conditions = *(*[4]byte)(b[28:])
	t.Symbol = *(*[11]byte)(b[32:])
	t.Exchange = b[43]
	t.Tape = b[44]

	return nil
}

// RawQuote
// -------------|--------------|-------------|-----------
//  Field       | NASDAQ short | NASDAQ long | NYSE long
// -------------|--------------|-------------|-----------
//  timestamp   | uint64       | uint64      | uint64
//  symbol      | [5]byte      | [11]byte    | [11]byte
//  bidPrice    | uint16       | uint64      | uint64
//  bidSize     | uint16       | uint32      | uint32
//  bidExchange | byte         | byte        | byte
//  askPrice    | uint16       | uint64      | uint64
//  askSize     | uint16       | uint32      | uint32
//  askExchange | byte         | byte        | byte
//  condition   | byte         | byte        | byte
//  nbbo        | bool         | bool        | bool
// -------------|--------------|-------------|-----------

type RawQuote struct {
	Timestamp   uint64   //  8
	BidPrice    uint64   // 16
	AskPrice    uint64   // 24
	BidSize     uint32   // 28
	AskSize     uint32   // 32
	BidExchange byte     // 33
	AskExchange byte     // 34
	Condition   byte     // 35
	Nbbo        bool     // 36
	Symbol      [11]byte // 47
	Tape        byte     // 48 - not in the original Quote message
	CreatedAt   uint64   // ?? - not in the original Quote message
}

func (q *RawQuote) Marshal() []byte {
	b := make([]byte, 56)

	nbbo := byte(0)
	if q.Nbbo {
		nbbo = 1
	}

	binary.LittleEndian.PutUint64(b[48:], q.CreatedAt) // go bounds check elimination
	binary.LittleEndian.PutUint64(b[0:], q.Timestamp)
	binary.LittleEndian.PutUint64(b[8:], q.BidPrice)
	binary.LittleEndian.PutUint64(b[16:], q.AskPrice)
	binary.LittleEndian.PutUint32(b[24:], q.BidSize)
	binary.LittleEndian.PutUint32(b[28:], q.AskSize)
	b[32] = q.BidExchange
	b[33] = q.AskExchange
	b[34] = q.Condition
	b[35] = nbbo
	for i := 0; i < 11; i++ {
		b[i+36] = q.Symbol[i]
	}
	b[47] = q.Tape

	return b
}

func (q *RawQuote) Unmarshal(b []byte) error {
	if len(b) != 56 {
		return errors.New("invalid quote length")
	}

	q.CreatedAt = binary.LittleEndian.Uint64(b[48:]) // go bounds check elimination
	q.Timestamp = binary.LittleEndian.Uint64(b[0:])
	q.BidPrice = binary.LittleEndian.Uint64(b[8:])
	q.AskPrice = binary.LittleEndian.Uint64(b[16:])
	q.BidSize = binary.LittleEndian.Uint32(b[24:])
	q.AskSize = binary.LittleEndian.Uint32(b[28:])
	q.BidExchange = b[32]
	q.AskExchange = b[33]
	q.Condition = b[34]
	q.Nbbo = b[35] == 1
	q.Symbol = *(*[11]byte)(b[36:])
	q.Tape = b[47]

	return nil
}
