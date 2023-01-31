package main

import (
	"embed"

	"github.com/ButterflyGate/logger"
	"github.com/ButterflyGate/piracy_guard/main/config"
	"github.com/ButterflyGate/piracy_guard/main/infras/hardware"
)

//go:embed binaries_for_violater/linux-amd.out
var violaterBinary embed.FS

// created by `$ uuidgen`
const SERIAL_NO = "e1cd4108-c540-4105-9f7e-3873af55a580"

func Check() {
	c := config.NewConfig()
	l := logger.NewLogger(c.LogLevel())

	h, err := hardware.NewHardware()
	if err != nil {
		l.Error(err)
		return
	}
	l.Info(h.GetCPUId())

}
