module mqtt

go 1.22.5

replace config => ../config

replace modbus => ../modbus

require (
	config v0.0.0
	github.com/eclipse/paho.mqtt.golang v1.5.0
	modbus v0.0.0
)

require (
	github.com/goburrow/serial v0.1.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/simonvetter/modbus v1.6.1 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
)
