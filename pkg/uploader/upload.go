package uploader

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/whym9/pcap_statistics/pkg/process"
	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/receiving_services/pkg/sender"

	"github.com/google/gopacket/pcap"
)

func Upload(addr, gaddr string) {

	ch := make(chan []byte)
	s := receiver.NewServer(&ch)
	go s.StartServer(addr)
	ch2 := make(chan []byte)

	go sender.Client{}.StartServer(gaddr, &ch2)
	go metrics.PromoHandler{}.StartMetrics("8008")
	for {
		chunk := <-ch
		fmt.Println(len(chunk))
		name := time.Now().Format("02-01-2002-59595898") + ".pcapng"
		file, err := os.Create(name)

		if err != nil {
			log.Fatal(err)
		}

		file.Write(chunk)
		file.Close()
		handle, err := pcap.OpenOffline(name)

		if err != nil {
			log.Fatal(err)
		}

		counter, err := process.Processor{}.Process(handle)

		if err != nil {
			log.Fatal(err)
		}
		ch <- []byte(process.Stringify(counter))
		ch2 <- []byte(process.Stringify(counter))
		<-ch2
		ch2 <- chunk
		mes := <-ch2

		if len(mes) > 0 {
			fmt.Println("Could not save")
		}

	}

}
