package config

import (
	"encoding/json"
	"log"
	"os"
)

type SensorType struct {
	Name      string
	Address   uint8
	Alias     uint8
	TimeStamp uint64
	DataType  string
	Value     uint16
}

type SparkPlugB struct {
	TimeStamp uint64
	Metrics   []SensorType
	Seq       int
}

func NewSparkPlugB() *SparkPlugB {
	matriks := &[]SensorType{
		{Name: "wind speed", Address: 8, Alias: 0, DataType: "UInt16"},
		{Name: "wind direction", Address: 10, Alias: 1, DataType: "UInt16"},
		{Name: "barometer", Address: 4, Alias: 2, DataType: "UInt16"},
		{Name: "inside temperature", Address: 5, Alias: 3, DataType: "UInt16"},
		{Name: "outside temperature", Address: 7, Alias: 4, DataType: "UInt16"},
		{Name: "inside humidity", Address: 6, Alias: 5, DataType: "UInt16"},
		{Name: "outside humidity", Address: 19, Alias: 5, DataType: "UInt16"},
		{Name: "rain rate", Address: 24, Alias: 5, DataType: "UInt16"},
		{Name: "solar radiation", Address: 26, Alias: 5, DataType: "UInt16"},
	}
	return &SparkPlugB{TimeStamp: 0, Metrics: *matriks, Seq: 0}
}

type Config struct {
	Port     string
	Speed    int16
	DataBits int8
	StopBits int8
	Broker   string
	ClientId string
	UserName string
	Password string
}

func LoadJson(fin string) (*Config, error) {
	content, err := os.ReadFile(fin)
	if err != nil {
		log.Fatalf("error loading json file: %s\n%s", fin, err)
		return nil, err
	}

	payload := &Config{
		Port:     "rtu:///dev/tnt1",
		Speed:    9600,
		DataBits: 8,
		StopBits: 1,
		Broker:   "localhost:1883",
		ClientId: "MDS_mqtt_sparkplug_sim",
		UserName: "emqx",
		Password: "public",
	}

	err = json.Unmarshal(content, payload)
	if err != nil {
		log.Fatalf("error decode json %s %s", content, err)
		return payload, err
	}
	log.Printf("Config %#v", payload)

	return payload, nil
}
