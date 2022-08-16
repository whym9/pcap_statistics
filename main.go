package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	uploadpb "pcap_statistics/proto"

	"pcap_statistics/sender"
	"pcap_statistics/uploader"
)

var client sender.Client

func main() {

	addr := *flag.String("address", "localhost:5005", "address of the server")
	saddr := *flag.String("sender_address", ":443", "address of the grpc sender client")
	flag.Parse()
	lis, err := net.Listen("tcp", addr)
	fmt.Println("GRPC server has started")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	uplSrv := uploader.NewServer()
	uplSrv.Addr = saddr

	rpcSrv := grpc.NewServer()

	uploadpb.RegisterUploadServiceServer(rpcSrv, uplSrv)

	log.Fatal(rpcSrv.Serve(lis))

}
