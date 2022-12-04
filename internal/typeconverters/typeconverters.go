package typeconverters

import (
	"encoding/binary"
)

// Int64ToBytes converts int64 data to an slice of bytes
func Int64ToBytes(number int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, number)
	return buf[:n]
}

// BytesToInt64 converts slice of bytes data to int64
func BytesToInt64(data []byte) (int64, int) {
	return binary.Varint(data)
}
