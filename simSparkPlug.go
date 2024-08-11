package main

import (
	"config"
	"log"
	"modbus"
	"mqtt"
	"time"
)

const configFileName = "config.json"

func main() {
	conf, err := config.LoadJson(configFileName)
	if err != nil {
		log.Fatalf("error load %s, %s", configFileName, err)
	}
	log.Print(conf)

	mod, err := modbus.Init(conf)
	if err != nil {
		log.Panicf("error: %s", err)
	}

	mqttpub, err := mqtt.Init(conf)
	if err != nil {
		log.Panicf("error: %s", err)
	}

	data := config.NewSparkPlugB()

	count := 0
	for {
		count += 1
		log.Printf("count: %d", count)
		for i := 0; i < len(*data); i++ {
			(*data)[i].TimeStamp = uint64(time.Now().Unix())
			modbus.ReadSensor(mod, &(*data)[i])
		}
		log.Printf("%+v", data)
		mqtt.Publish(mqttpub, data)
		time.Sleep(1 * time.Second)
	}
}
