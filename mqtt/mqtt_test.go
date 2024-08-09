package mqtt

import (
	"config"
	"log"
	"testing"
	"time"

	"modbus"
)

func TestInit(t *testing.T) {
	config, err := config.LoadJson("../config.json")
	if err != nil {
		log.Panic("error")
	}
	client, err := Init(config)
	if err != nil {
		log.Panic("error")
	}
	client.Disconnect(250)
}

func TestPublish(t *testing.T) {
	conf, err := config.LoadJson("../config.json")
	if err != nil {
		log.Panic("error")
	}
	client, err := Init(conf)
	if err != nil {
		log.Panic("error")
	}
	defer client.Disconnect(250)

	// test with empty/base data
	data := config.NewSparkPlugB()
	Publish(client, data)
}

func TestPubWithModData(t *testing.T) {
	conf, err := config.LoadJson("../config.json")
	if err != nil {
		log.Panic("error")
	}
	client, err := Init(conf)
	if err != nil {
		log.Panic("error")
	}
	defer client.Disconnect(250)

	c, err := modbus.Init(conf)
	if err != nil {
		log.Panicf("error")
		return
	}

	data := config.NewSparkPlugB()
	for i := 0; i < len(*data); i++ {
		(*data)[i].TimeStamp = uint64(time.Now().Unix())
		modbus.ReadSensor(c, &(*data)[i])
	}
	Publish(client, data)
}
