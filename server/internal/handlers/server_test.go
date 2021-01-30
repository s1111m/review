package handlers

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "server/pkg/hashservice"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterHashServiceServer(server, &Server{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestCreateHash(t *testing.T) {
	tests := []struct {
		name      string
		strToHash string
		res       *pb.ProtoArrayOfHashes
		hash      string
		errMsg    string
	}{
		{
			"Check hash",
			"hash",
			nil,
			"30163935c002fc4e1200906c3d30a9c4956b4af9f6dcaef1eb4b1fcb8fba69e7a7acdc491ea5b1f2864ea8c01b01580ef09defc3b11b3f183cb21d236f7f1a6b",
			fmt.Sprintf("cannot deposit %v", -1.11),
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHashServiceClient(conn)
	//uc := hashservice.NewHashServiceClient(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := &pb.ProtoStringToHash{Str: tt.strToHash}

			payload := &pb.ProtoArrayOfStrings{}
			payload.StrToConvert = append(payload.StrToConvert, str)

			//			request := &pb.DepositRequest{Amount: tt.amount}
			//payload.RequestId = cwt.Value("request-id").(string)
			response, err := client.CreateHash(context.Background(), payload)

			if err != nil {
				log.Fatal(err)
			}

			if response != nil {
				if response.Hashes[0].Hash != tt.hash {
					t.Error("response: expected", tt.hash, "received", response.Hashes[0].Hash)
				}
			}
		})
	}
}
