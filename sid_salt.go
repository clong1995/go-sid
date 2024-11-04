package sid

import (
	"encoding/binary"
	"fmt"
	"log"
)

// DeSalt 从字符串创建id
func DeSalt(str string, salt uint64) (id Id) {
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

	id = Id(res - salt)
	return
}

// EnSalt 从数字创建id
func EnSalt(num uint64, salt uint64) (id Id) {
	return Id(num + salt)
}
