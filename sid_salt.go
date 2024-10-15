package sid

import (
	"encoding/binary"
	"fmt"
	"log"
)

// deSalt 从字符串创建id
func deSalt(str string, salt uint64) (id Id) {
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

// enSalt 从数字创建id
func enSalt(num uint64, salt uint64) (id Id) {
	return Id(num + salt)
}
