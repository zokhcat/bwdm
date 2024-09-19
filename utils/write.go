package utils

import (
	"fmt"
	"log"
	"os"
)

func writePacketData(file *os.File, packetData []byte) error {
	_, err := file.WriteString(fmt.Sprintf("Packet: %v\n", packetData))
	if err != nil {
		log.Printf("Failed to write packet data to file: %v", err)
		return err
	}
	return nil
}
