package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"github.com/zokhcat/bwdm/graph"
	"github.com/zokhcat/bwdm/modules"
)

var rootCmd = &cobra.Command{
	Use:   "bwdm",
	Short: "A simple bandwidth monitoring TUI",
	Long:  `A TUI-CLI application to monitor network bandwidth using packet sniffing`,
}

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Sniffing packets for 10 seconds",
	Long:  `Sniffing packets for 10 seconds and present them in a graph visualizer`,
	Run: func(cmd *cobra.Command, args []string) {
		interfaceName, _ := cmd.Flags().GetString("interface")
		ipAddress, _ := cmd.Flags().GetString("ipaddress")
		if interfaceName == "" {
			fmt.Println("Please specify an interface using the -i flag")
			return
		}
		runCapture(interfaceName, ipAddress)
	},
}

var listNetworkCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the network interfaces",
	Long:  "Lists all the network interfaces present",
	Run: func(cmd *cobra.Command, args []string) {
		interfaces, err := pcap.FindAllDevs()
		if err != nil {
			fmt.Println("Error finding network interfaces:", err)
		}

		fmt.Println("Availble interfaces: ")
		for _, iface := range interfaces {
			fmt.Printf("- %s\n", iface.Name)
		}
		fmt.Println("Please specify an interface using the -i flag while using the capture command")
	},
}

func init() {
	rootCmd.AddCommand(captureCmd)
	rootCmd.AddCommand(listNetworkCmd)
	captureCmd.Flags().StringP("interface", "i", "", "Network Interface to sniff packets from")
	captureCmd.Flags().StringP("ip", "p", "", "To sniff packets from a specific IP address")
}

func runCapture(interfaceName string, ipAddress string) {
	byteSliceChan := make(chan []uint64)
	go modules.CapturePackets(interfaceName, byteSliceChan, 10*time.Second, ipAddress)

	byteSlice := <-byteSliceChan
	graph.DrawGraph(byteSlice)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
