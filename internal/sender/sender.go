package sender

import (
	"context"
	"fmt"
	"os"
	sendpb "pcap_statistics/internal/proto/send"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	client sendpb.SendServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: sendpb.NewSendServiceClient(conn),
	}
}

func (c Client) Send(con context.Context, counter []byte, fileName string) (string, error) {
	ctx, cancel := context.WithDeadline(con, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Send(ctx)
	if err != nil {

		return "", err
	}

	if err := stream.Send(&sendpb.SendRequest{Chunk: counter}); err != nil {

		return "", err
	}

	file, err := os.Open(fileName)

	if err != nil {

		return "", err
	}

	for {

		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		if err != nil {
			break
		}

		if err := stream.Send(&sendpb.SendRequest{Chunk: buf[:n]}); err != nil {

			return "", err
		}

	}

	res, err := stream.CloseAndRecv()
	if err != nil {

		return "", err
	}
	fmt.Println("stopped sending")

	return res.GetName(), nil
}
