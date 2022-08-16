package uploader

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"pcap_statistics/process"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	uploadpb "pcap_statistics/proto"
	"pcap_statistics/sender"
)

type Server struct {
	uploadpb.UnimplementedUploadServiceServer
	Addr string
}

func NewServer() Server {

	return Server{}
}

func (s Server) Upload(stream uploadpb.UploadService_UploadServer) error {
	processor := process.Processor{}
	conn, err := grpc.Dial(s.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := sender.NewClient(conn)
	defer conn.Close()
	packets := []process.Packet{}

	for {

		req, err := stream.Recv()

		if err == io.EOF {

			counter := processor.Process(packets)

			b, err := json.Marshal(&counter)

			if err != nil {

				return err
			}
			_, err = client.Send(context.Background(), b, packets)
			if err != nil {

				return stream.SendAndClose(&uploadpb.UploadResponse{Name: "Could not save the file"})
			}
			return stream.SendAndClose(&uploadpb.UploadResponse{Name: process.Stringify(counter)})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		ci := process.Capture{}
		bin := req.GetChunk()

		err = json.Unmarshal(bin, &ci)
		if err != nil {
			return err
		}

		req, err = stream.Recv()

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		packets = append(packets, process.Packet{ci, req.GetChunk()})
	}
}
