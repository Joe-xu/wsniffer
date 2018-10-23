package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcapgo"
)

func setupDataSource(in io.Reader) *gopacket.PacketSource {

	r, err := pcapgo.NewReader(in)
	if err != nil {
		log.Fatalf("setupDataSource: %s", err)
	}

	return gopacket.NewPacketSource(r, r.LinkType())
}

func fromFile(fpath string) io.ReadCloser {

	f, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("fromFile: %s", err)
	}

	return f
}

func fromIwcap(interfaceName string) io.ReadCloser {

	// check whether iwcap is exist or not
	_, err := exec.LookPath("iwcap")
	if err != nil {
		log.Fatalf("formIwcap: %s", err)
	}

	iwcapCmd := exec.Command("iwcap", "-i", interfaceName, "-s")

	out, err := iwcapCmd.StdoutPipe()
	if err != nil {
		log.Fatalf("formIwcap: %s", err)
	}

	if err = iwcapCmd.Start(); err != nil {
		log.Fatalf("formIwcap: %s", err)
	}

	return out

}
