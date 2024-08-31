package sid

import "encoding/binary"

type Id uint64

// String 获取 string的id
func (id Id) String() string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id.Uint64())
	return encode(b)
}

// Uint64 获取 uint64的id
func (id Id) Uint64() uint64 {
	return uint64(id)
}
