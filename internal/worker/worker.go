package worker

import (
	"fmt"

	"github.com/whym9/pcap_statistics/pkg/process"
	"github.com/whym9/receiving_service/pkg/metrics"
	"github.com/whym9/receiving_service/pkg/receiver"
	"github.com/whym9/receiving_service/pkg/sender"
)

type Worker struct {
	metrics   metrics.Metrics
	receiver  receiver.Receiver
	sender    sender.Sender
	processor process.Process
	ch1       chan []byte
	ch2       chan []byte
}

func NewWorker(m metrics.Metrics, r receiver.Receiver, s sender.Sender, p process.Process, ch1, ch2 chan []byte) Worker {
	return Worker{m, r, s, p, ch1, ch2}
}

func (w Worker) Work(addr1, addr2, addr3 string) {
	go w.sender.StartServer(addr1)
	go w.receiver.StartServer(addr2)
	go w.metrics.StartMetrics(addr3)

	for {
		data := <-w.ch1

		counter, err := w.processor.Process(data)
		if err != nil {
			w.ch1 <- []byte("Could not make statistics")
			continue
		}

		w.ch1 <- []byte(counter)

		w.ch2 <- []byte(counter)
		<-w.ch2
		w.ch2 <- data
		mes := <-w.ch2

		if len(mes) > 0 {
			fmt.Printf("The error occured while saving: %v\n", err)
		}

	}

}
