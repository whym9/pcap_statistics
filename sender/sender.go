package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"pcap_statistics/process"
	sendpb "pcap_statistics/proto/send"
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

func (c Client) Send(con context.Context, counter []byte, packets []process.Packet) (string, error) {
	ctx, cancel := context.WithDeadline(con, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Send(ctx)
	if err != nil {

		return "", err
	}

	if err := stream.Send(&sendpb.SendRequest{Chunk: counter}); err != nil {
		return "", err
	}

	for _, pack := range packets {

		b, err := json.Marshal(&pack.Ci)

		if err != nil {
			return "", err
		}

		if err := stream.Send(&sendpb.SendRequest{Chunk: b}); err != nil {
			return "", err
		}

		if err := stream.Send(&sendpb.SendRequest{Chunk: pack.Data}); err != nil {
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
