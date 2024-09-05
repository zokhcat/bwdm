package main

import (
	"github.com/zokhcat/bwdm/modules"
)

func main() {
	interfaceName := "wlo1"
	modules.CapturePackets(interfaceName)
}
