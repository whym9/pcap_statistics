package main

import (
	"flag"

	"pcap_statistics/internal/uploader"
)

func main() {

	addr := *flag.String("address", "localhost:5005", "address of the server")
	gaddr := *flag.String("grpc_address", ":443", "address of the grpc sender client")
	flag.Parse()

	uploader.Upload(addr, gaddr)
}
