package main

import (
	"flag"
	"fmt"
  "log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

func main() {
  bluetooth()
  commandPtr, err := getCommand()
  if err != nil {
    log.Fatalln("Error: ", err)
  }

  switch (*commandPtr) {
  case "on":
    fmt.Println("on")
  case "off":
    fmt.Println("off")
  }
}

func getCommand() (*string, error) {
  flag.Parse()
  flags := flag.Args()

  if len(flags) == 0 {
    return nil, &getCommandError{}
  }

  command := flags[0]

  if command != "on" && command != "off" {
    return nil, &getCommandError{}
  }

  return &command, nil
}

type getCommandError struct {
}

func(ptr *getCommandError) Error() string {
  return "Invalid command"
}

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func onDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
  fmt.Printf("\nPeripheral ID:%s, NAME:(%s)\n", p.ID(), p.Name())
	fmt.Println("  Local Name        =", a.LocalName)
	fmt.Println("  TX Power Level    =", a.TxPowerLevel)
	fmt.Println("  Manufacturer Data =", a.ManufacturerData)
	fmt.Println("  Service Data      =", a.ServiceData)
}

func bluetooth() {
  d, err := gatt.NewDevice(option.DefaultClientOptions...)
  if err != nil {
    fmt.Println("fuck you", err)
    return
  }

  d.Handle(gatt.PeripheralDiscovered(onDiscovered))
  d.Init(onStateChanged)
  select{}
}
