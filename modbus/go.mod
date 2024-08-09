module modbus

go 1.22.5

replace config => ../config

require (
	config v0.0.0-00010101000000-000000000000
	github.com/simonvetter/modbus v1.6.1
)

require github.com/goburrow/serial v0.1.0 // indirect
