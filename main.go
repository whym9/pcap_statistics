package main

import (
	"flag"

	"github.com/whym9/pcap_statistics/internal/worker"
	metrics "github.com/whym9/pcap_statistics/pkg/metrics/prometheus"
	process "github.com/whym9/pcap_statistics/pkg/process/pcap"
	receiver "github.com/whym9/receiving_service/pkg/receiver/GRPC"
	sender "github.com/whym9/receiving_service/pkg/sender/GRPC"
)

func main() {

	flag.Parse()

	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	Promo_Handler := metrics.NewPromoHandler()

	GRPC_Sender := sender.NewGRPCHandler(Promo_Handler, ch2)

	GRPC_Receiver := receiver.NewServer(Promo_Handler, ch1)
	Pcap_Handler := process.NewPcapHandler()

	w := worker.NewWorker(Promo_Handler, GRPC_Receiver, GRPC_Sender, Pcap_Handler, ch1, ch2)
	w.Work()
}
