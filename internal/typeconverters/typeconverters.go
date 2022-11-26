package typeconverters

import (
	"encoding/binary"
)

func Int64ToBytes(number int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, number)
	return buf[:n]
}

func BytesToInt64(data []byte) (int64, int) {
	return binary.Varint(data)
}
