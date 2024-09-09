package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"github.com/zokhcat/bwdm/modules"
)

var rootCmd = &cobra.Command{
	Use:   "bwdm",
	Short: "A simple bandwidth monitoring TUI",
	Long:  `A TUI-CLI application to monitor network bandwidth using packet sniffing`,
}

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Start sniffing packets",
	Long:  `Start capturing packets and calculate bandwidth usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		interfaceName, _ := cmd.Flags().GetString("interface")
		if interfaceName == "" {
			interfaces, err := pcap.FindAllDevs()
			if err != nil {
				fmt.Println("Error finding network interfaces:", err)
				return
			}
			fmt.Println("Available interfaces:")
			for _, iface := range interfaces {
				fmt.Printf("- %s\n", iface.Name)
			}
			fmt.Println("Please specify an interface using the -i flag")
			return
		}
		runCapture(interfaceName)
	},
}

func init() {
	rootCmd.AddCommand(captureCmd)
	captureCmd.Flags().StringP("interface", "i", "", "Network Interface to sniff packets from")
}

func runCapture(interfaceName string) {
	byteSliceChan := make(chan []uint64)
	go modules.CapturePackets(interfaceName, byteSliceChan, 10*time.Second) // taking 10 as an example, going to set it as flag

	bytesSlice := <-byteSliceChan
	fmt.Println(bytesSlice)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
