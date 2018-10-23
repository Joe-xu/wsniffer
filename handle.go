package main

import (
	"fmt"

	"github.com/google/gopacket/layers"
)

// Information Element AKA Tagged parameter
func handleInfoElement(ele *layers.Dot11InformationElement) {

	switch ele.ID {
	case layers.Dot11InformationElementIDSSID:
		if ele.LayerType() == layers.LayerTypeDot11MgmtProbeReq {
			fmt.Printf("ESSID: %s\n", ele.Info)
			break
		}

		fmt.Printf("SSID: %s\n", ele.Info)
	default:
		fmt.Printf("Unhandle element: %s\n", ele.ID)
	}

}
