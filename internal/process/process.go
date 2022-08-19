package process

import (
	"io"
	"strconv"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type ProcessInterface interface {
	Process(packets []Packet) Protocols
}

type Processor struct{}

type Protocols struct {
	TCP  int `json: "TCP"`
	UDP  int `json: "UDP"`
	IPv4 int `json: "IPv4"`
	IPv6 int `json: "IPv6"`
}

type Capture struct {
	TimeStamp      time.Time     `json: "time"`
	CaptureLength  int           `json: "caplength"`
	Length         int           `json: "length"`
	InterfaceIndex int           `json :  "index"`
	AccalaryData   []interface{} `json: "accalary"`
}

type Packet struct {
	Ci   Capture
	Data []byte
}

var (
	eth layers.Ethernet
	ip4 layers.IPv4
	ip6 layers.IPv6
	tcp layers.TCP
	udp layers.UDP
	dns layers.DNS
)

func (p Processor) Process(handle *pcap.Handle) (Protocols, error) {

	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&eth,
		&ip4,
		&ip6,
		&tcp,
		&udp,
		&dns,
	)

	decoded := make([]gopacket.LayerType, 0, 10)

	counter := Protocols{}

	for {

		data, _, err := handle.ZeroCopyReadPacketData()
		if err == io.EOF {
			break
		}

		if err != nil {
			return Protocols{}, err
		}
		parser.DecodeLayers(data, &decoded)

		for _, layer := range decoded {
			if layer == layers.LayerTypeTCP {
				counter.TCP++
			}
			if layer == layers.LayerTypeUDP {
				counter.UDP++
			}
			if layer == layers.LayerTypeIPv4 {
				counter.IPv4++
			}
			if layer == layers.LayerTypeIPv6 {
				counter.IPv6++
			}
		}

	}

	return counter, nil
}

func Stringify(counter Protocols) string {
	return "TCP: " + strconv.Itoa(counter.TCP) + "\nUDP: " +
		strconv.Itoa(counter.UDP) + "\nIPv4: " +
		strconv.Itoa(counter.IPv4) + "\nIPv6: " +
		strconv.Itoa(counter.IPv6)
}
