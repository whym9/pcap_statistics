package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/gopacket/pcap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pcap_statistics/internal/metrics"
	"pcap_statistics/internal/process"
	uploadpb "pcap_statistics/internal/proto"
	"pcap_statistics/internal/sender"
)

type Server struct {
	uploadpb.UnimplementedUploadServiceServer
	Addr string
}

func NewServer() Server {

	return Server{}
}

func (s Server) Upload(stream uploadpb.UploadService_UploadServer) error {
	metrics.RecordMetrics()
	conn, err := grpc.Dial(s.Addr, grpc.WithInsecure())
	if err != nil {

		log.Fatalln(err)
	}

	client := sender.NewClient(conn)
	defer conn.Close()
	name := "internal/files/" + time.Now().Format("02-01-1789-8989") + ".pcapng"
	file, err := os.Create(name)
	if err != nil {

		log.Fatal(err)
	}
	file.Close()
	chunk := []byte{}
	count := 0
	for {

		req, err := stream.Recv()

		if err == io.EOF {

			break
		}
		if err != nil {

			return status.Error(codes.Internal, err.Error())
		}

		bin := req.GetChunk()
		count += len(bin)
		chunk = append(chunk, bin...)

		if err != nil {

			return status.Error(codes.Internal, err.Error())
		}

	}
	fmt.Println(count)
	err = ioutil.WriteFile(name, chunk, 600)
	if err != nil {

		log.Fatal(err)
	}

	handle, err := pcap.OpenOffline(name)
	if err != nil {

		return err
	}
	counter, err := process.Processor{}.Process(handle)
	if err != nil {

		return err
	}

	b, err := json.Marshal(&counter)
	fmt.Println(string(b))
	_, err = client.Send(context.Background(), b, name)
	if err != nil {
		return stream.SendAndClose(&uploadpb.UploadResponse{Name: process.Stringify(counter) + "\n Could not save to Database"})
	}
	return stream.SendAndClose(&uploadpb.UploadResponse{Name: process.Stringify(counter)})
}
