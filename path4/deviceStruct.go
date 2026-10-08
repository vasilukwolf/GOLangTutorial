package main

import "fmt"

type Device struct {
	Name   string
	Type   string
	Active bool
}

func (device Device) Info() {
	fmt.Printf("Это устройство %s, типа %s, в состоянии %t \n", device.Name, device.Type, device.Active)
}

func (device *Device) Activate() {
	device.Active = true
}

func (device *Device) Deactivate() {
	device.Active = false
}

func NewDevice(name, deviceType string) *Device {
	return &Device{Name: name, Type: deviceType}
}

func main() {
	lamp := Device{Name: "Лампа", Type: "Освещение"}
	lamp.Info()
	fmt.Println()

	lamp.Activate() // Go сам возьмёт &lamp, потому что получатель — указатель
	lamp.Info()
	fmt.Println()

	sensor := NewDevice("Датчик", "Температура")
	sensor.Activate()
	sensor.Info()
	sensor.Deactivate()
	sensor.Info()
}
