package uploader

import (
	"log"
	"os"
	"pcap_statistics/internal/process"
	"pcap_statistics/internal/receiver"
	"pcap_statistics/internal/sender"
	"time"

	"github.com/google/gopacket/pcap"
)

func Upload(addr, gaddr string) {
	ch := make(chan []byte)
	s := receiver.NewServer(&ch)
	go s.StartServer(addr)
	chunk := <-ch
	name := time.Now().Format("02-01-2002-59595898")
	file, err := os.Open(name)

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
	ch2 := make(chan []byte)

	go sender.Client{}.StartServer(gaddr, &ch2)
	ch2 <- []byte(process.Stringify(counter))
	<-ch2
	ch2 <- chunk
	mes := <-ch2
	if len(mes) > 0 {
		ch <- []byte("could not save\n" + process.Stringify(counter))
	}
	ch <- []byte(process.Stringify(counter))

}
