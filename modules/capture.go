package modules

import (
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func CapturePackets(interfaceName string, byteSliceChan chan<- []uint64, captureDuration time.Duration) {
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
