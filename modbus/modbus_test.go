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
