package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {

	// packetSource := setupDataSource(fromIwcap("mon0"))
	packetSource := setupDataSource(fromFile("./testdata/test.pcap"))

	for packet := range packetSource.Packets() {

		fmt.Printf("\n new packet(%d)\n", len(packet.Layers()))

		for _, lay := range packet.Layers() {

			switch layer := lay.(type) {
			case *layers.RadioTap:
				fmt.Printf("%d\n", layer.DBMAntennaSignal)
				// fmt.Println(lay.LayerContents())
			case *layers.Dot11MgmtProbeReq:
				// could fetch SSID that device accessed before

				fmt.Println(packet.Dump())
				os.Exit(0)
			case *layers.Dot11MgmtBeacon:
				fmt.Println(gopacket.LayerString(layer))

			case *layers.Dot11MgmtProbeResp:

				fmt.Println(gopacket.LayerString(layer))

			case *layers.Dot11InformationElement:
				handleInfoElement(layer)
			default:
				fmt.Printf("Unhandle type: %s\n", lay.LayerType())
			}

		}

	}

}
