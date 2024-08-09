package modbus

import (
	"log"
	"time"

	"config"

	"github.com/simonvetter/modbus"
)

func Init(c *config.Config) (*modbus.ModbusClient, error) {
	log.Printf("ini init\n")
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      c.Port,
		Speed:    uint(c.Speed),
		DataBits: uint(c.DataBits),
		Parity:   modbus.PARITY_NONE,
		StopBits: uint(c.StopBits),
		Timeout:  500 * time.Millisecond,
	})
	client.SetUnitId(uint8(c.UnitId))
	if err != nil {
		return client, err
	}
	err = client.Open()
	if err != nil {
		return client, err
	}
	return client, nil
}

func readReg(c *modbus.ModbusClient, addr uint16, val *uint16) {
	val01, err := c.ReadRegister(addr, modbus.HOLDING_REGISTER)
	if err != nil {
		log.Printf("error reading register: %d, %s", addr, err)
	} else {
		*val = val01
	}
}

func ReadSensor(c *modbus.ModbusClient, dat *config.SparkPlugB) {
	for i := 0; i < len(dat.Metrics); i++ {
		data := &dat.Metrics[i]
		data.TimeStamp = uint64(time.Now().Unix())
		readReg(c, uint16(data.Address), &data.Value)
	}
}
