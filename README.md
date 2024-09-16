# Bandwidth Monitoring TUI(bwdm)

_bwdm_ is a simple CLI- and TUI-based bandwidth monitoring tool that allows users to monitor traffic by sniffing packets from a specified network interface. It displays real-time network using a graph, helping users visualize bandwidth consumption

## Features

- _Packet Sniffing_: Capture and monitor network traffic from a specific interface.
- _Graph Visualization_: Displays real-time network bandwidth in a graphical format for easy analysis.
- _Network Interface Listing_: Lists all network interfaces for quick selection.

## Installation

### Prerequisites

- [go](https://go.dev/) (v1.18+)
- [libpcap](https://www.tcpdump.org/) , required for packet caputring
- [git](https://git-scm.com/)

### Clone the repository

```
git clone https://github.com/zokhcat/bwdm.git
cd bwdm
```

### Build the Application

```
go build -o bwdm
```

## Usage

### List available network interfaces

Before capturing traffic, you can list all interfaces on your system:

```
./bwdm list
```

The output will be something like:

```
Available interfaces:
- eno1
- wlo1
- lo
Please specify an interface using the -i flag while using the capture command.
```

### Start capturing traffic(for 10 seconds)

To capture network traffic from a specific interface for 10 seconds, use the `capture` command with `-i` flag:

```
./bwdm capture -i <interface-name>
```

For example:

```
./bwdm capture -i wlo1
```

This will sniff the packets on the `wlo1` interface for 10 seconds and visualize the network bandwidth in a graph.

### Options:

`capture`: Sniff packets on a specified network interface for 10 seconds and display network usage in a graphical interface

#### Flags:

`p` or `ip`: to sniff packets from a specific IP address

`list`: List all available network interfaces

## Future Todos:

- [x] Packet Filtering such for specific IP address
- [ ] Log the packet capture into a file
- [ ] Configurable capture duration(Maybe)
