package helpers

import (
	"crypto/rand"
	"encoding/binary"
)

func RandInt64Positive() int64 {
	var b [8]byte
	rand.Read(b[:])
	u := int64(binary.LittleEndian.Uint64(b[:]))
	return int64(u & 0x7FFFFFFFFFFFFFFF)
}
