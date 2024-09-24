package modules

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

func CapturePackets(interfaceName string, byteSliceChan chan<- []uint64, captureDuration time.Duration, ipAddress string, filename string, port string, protocol string, inspect bool) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	writer := pcapgo.NewWriter(file)
	writer.WriteFileHeader(65536, layers.LinkTypeARCNetLinux)

	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if protocol != "" {
		filter := fmt.Sprintf("protocol %s", protocol)
		if err := handle.SetBPFFilter(filter); err != nil {
			log.Fatalf("Error setting BPF filter: %s", err)
		}
		fmt.Printf("BPF filter set: %s\n", filter)
	}

	if port != "" {
		filter := fmt.Sprintf("port %s", port)
		if err := handle.SetBPFFilter(filter); err != nil {
			log.Fatalf("Error setting BPF filter: %s", err)
		}
		fmt.Printf("BPF filter set: %s\n", filter)
	}

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
		err := writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		if err != nil {
			log.Printf("Failed to write packet data to file: %v", err)
		}

		if inspect {
			fmt.Printf("Packet: %s\n", packet.String())
			fmt.Printf("Data: %v\n", packet.Data())
		}

		totalBytes += uint64(len(packet.Data()))
	}
}
