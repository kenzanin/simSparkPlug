package config

import (
	"log"
	"testing"
)

func TestLoadJson(t *testing.T) {
	config, err := LoadJson("../config.json")
	if err != nil {
		log.Fatalf("test error loading json file")
	}
	t.Logf("port:%s broker:%s", config.Port, config.Broker)
}
