package sid

import (
	"github.com/clong1995/go-config"
	"github.com/sony/sonyflake"
	"log"
	"strconv"
)

var sf *sonyflake.Sonyflake

func init() {
	machineID := config.Value("MACHINE ID")
	if machineID == "" {
		log.Fatalln("MACHINE ID not found")
	}
	atoi, err := strconv.Atoi(machineID)
	if err != nil {
		log.Fatalln(err)
		return
	}
	mid := uint16(atoi)

	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return mid, nil
		},
	})
	if sf == nil {
		log.Fatalln("sid not created")
	}
	log.Printf("sid created %s success\n", machineID)
}
