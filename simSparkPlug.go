package main

import (
	"config"
	"log"
	"modbus"
	"mqtt"
	"time"
)

const configFileName = "config.json"

//	func NewSparkPlugB() *SensorsDat {
//		matriks := &[]SensorType{
//			{Name: "wind speed", Address: 8, Alias: 0, DataType: "UInt16"},
//			{Name: "wind direction", Address: 10, Alias: 1, DataType: "UInt16"},
//			{Name: "barometer", Address: 4, Alias: 2, DataType: "UInt16"},
//			{Name: "inside temperature", Address: 5, Alias: 3, DataType: "UInt16"},
//			{Name: "outside temperature", Address: 7, Alias: 4, DataType: "UInt16"},
//			{Name: "inside humidity", Address: 6, Alias: 5, DataType: "UInt16"},
//			{Name: "outside humidity", Address: 19, Alias: 5, DataType: "UInt16"},
//			{Name: "rain rate", Address: 24, Alias: 5, DataType: "UInt16"},
//			{Name: "solar radiation", Address: 26, Alias: 5, DataType: "UInt16"},
//		}
//		return &SensorsDat{TimeStamp: 0, Metrics: *matriks, Seq: 0}
//	}

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
	data.Seq = len(data.Metrics)

	count := 0
	for {
		data.TimeStamp = uint64(time.Now().Unix())
		modbus.ReadSensor(mod, data)
		log.Printf("count: %d", count)
		log.Print(data)
		mqtt.Publish(mqttpub, data)
		time.Sleep(1 * time.Second)
	}
}
