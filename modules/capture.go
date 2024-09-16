package modules

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func CapturePackets(interfaceName string, byteSliceChan chan<- []uint64, captureDuration time.Duration, ipAddress string) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if ipAddress != "" {
		filter := fmt.Sprintf("host %s", ipAddress)
		if err := handle.SetBPFFilter(filter); err != nil {
			log.Fatalf("Error setting BPF filter: %s", err)
		}
		fmt.Printf("BPF filter set: %s\n", filter)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	var totalBytes uint64 = 0
	var byteSlice []uint64

	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(captureDuration)
	go func() {
		for {
			select {
			case <-ticker.C:
				byteSlice = append(byteSlice, totalBytes)
				totalBytes = 0
			case <-timeout:
				byteSliceChan <- byteSlice
				close(byteSliceChan)
				return
			}
		}
	}()

	for packet := range packetSource.Packets() {
		totalBytes += uint64(len(packet.Data()))
	}
}
