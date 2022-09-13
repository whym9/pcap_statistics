package main

import (
	"flag"

	"github.com/whym9/pcap_statistics/internal/worker"
	process "github.com/whym9/pcap_statistics/pkg/process/pcap"
	metrics "github.com/whym9/receiving_service/pkg/metrics/prometheus"
	receiver "github.com/whym9/receiving_service/pkg/receiver/GRPC"
	sender "github.com/whym9/receiving_service/pkg/sender/GRPC"
)

func main() {

	addr1 := *flag.String("address", "localhost:6006", "address of the server")
	addr2 := *flag.String("grpc_address", ":443", "address of the grpc sender client")
	addr3 := *flag.String("metric_address", "8008", "metrics address")
	dir := *flag.String("dir", "files", "directory for saving files")

	flag.Parse()

	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	Promo_Handler := metrics.NewPromoHandler()

	GRPC_Sender := sender.NewGRPCHandler(Promo_Handler, ch2)

	GRPC_Receiver := receiver.NewServer(Promo_Handler, &ch1)
	Pcap_Handler := process.NewPcapHandler(dir, Promo_Handler)

	w := worker.NewWorker(Promo_Handler, GRPC_Receiver, GRPC_Sender, Pcap_Handler, ch1, ch2)
	w.Work(addr1, addr2, addr3)
}
