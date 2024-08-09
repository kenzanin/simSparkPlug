package modbus

import (
	"config"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	config, err := config.LoadJson("../config.json")
	if err != nil {
		log.Fatalf("error loading json")
		return
	}
	client, err := Init(config)
	if err != nil || client == nil {
		t.Errorf("%s|%s", "error di create handle", err)
	}
	defer client.Close()
}

func TestReadSensor(t *testing.T) {
	conf, err := config.LoadJson("../config.json")
	if err != nil {
		log.Fatalf("error loading json")
		return
	}
	client, err := Init(conf)
	if err != nil || client == nil {
		t.Errorf("%s|%s", "error di create handle", err)
	}
	defer client.Close()
	dat := config.NewSparkPlugB()
	ReadSensor(client, dat)
}
