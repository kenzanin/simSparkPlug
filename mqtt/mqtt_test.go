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

	data := config.NewSparkPlugB()
	data.TimeStamp = uint64(time.Now().Unix())
	data.Seq = len(data.Metrics)
	c, err := modbus.Init(conf)
	if err != nil {
		log.Panicf("error")
		return
	}
	modbus.ReadSensor(c, data)
	Publish(client, data)
}
