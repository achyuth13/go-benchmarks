package core

var BITMASK = []byte{
	0b00000001,
	0b00000011,
	0b00000111,
	0b00001111,
	0b00011111,
	0b00111111,
	0b01111111,
	0b11111111,
}

var buf [11]byte
var bitShifts []uint8 = []uint8{7, 7, 7, 7, 7, 7, 7, 7, 7, 1}

func getLSB(x byte, n uint8) byte {
	if n > 8 {
		panic("can extract at max 8 bits from the number")
	}
	return byte(x) & BITMASK[n-1]
}

// It converts the number into a byte
// Then it iterates it, and then gets out 7 bits, adds it into 7 parts of a byte array
// After that if the number continues, it adds a continuation bit(1) or termination bit (0)
// This will be done with little endian enconding, (least significant digits of the number take higher preference)
func EncodeUInt64(x uint64) []byte {
	var i int = 0
	for i = 0; i < len(bitShifts); i++ {
		buf[i] = getLSB(byte(x), bitShifts[i]) | 0b10000000 // marking the continuation bit
		x = x >> bitShifts[i]
		if x == 0 {
			break
		}
	}
	buf[i] = buf[i] & 0b01111111 // marking the termination bit
	return append(make([]byte, 0, i+1), buf[:i+1]...)
}

func DecodeUInt64(vint []byte) uint64 {
	var i int = 0
	var v uint64 = 0
	for i = 0; i < len(vint); i++ {
		b := getLSB(vint[i], 7)
		v = v | uint64(b)<<(7*i)
	}
	return v
}
