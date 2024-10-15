package sid

import (
	"encoding/binary"
	"fmt"
	"log"
)

// FromStr 从字符串创建id
func FromStr(str string) (id Id) {
	buf, err := decode(str)
	if err != nil {
		log.Println(err)
		return
	}

	if len(buf) < 8 {
		err = fmt.Errorf("buf len < 8")
		log.Println(err)
		return
	}

	res := binary.LittleEndian.Uint64(buf)

	id = Id(res)
	return
}

// FromNum 从数字创建id
func FromNum(num uint64) (id Id) {
	return Id(num)
}

// New 获取一个id
func New() (id Id) {
	nextID, err := sf.NextID()
	if err != nil {
		log.Println(err)
		return
	}
	id = Id(nextID)
	return
}
