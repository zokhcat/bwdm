package modules

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func CapturePackets(interfaceName string) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	// for packet := range packetSource.Packets() {
	// 	fmt.Println(packet)
	// }

	var totalBytes uint64 = 0

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Printf("Bytes in the last second: %d\n", totalBytes)
			totalBytes = 0
		}
	}()

	for packet := range packetSource.Packets() {
		totalBytes += uint64(len(packet.Data()))
	}
}
