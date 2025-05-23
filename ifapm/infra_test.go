package ifapm

import (
	"context"
	"fmt"
	"net/http"
	"protos"
	"testing"
	"time"
)

func TestInfra_Init(t *testing.T) {
	Infra.Init(
		InfraRdbOption("127.0.0.1:6379"),
		InfraMysqlDbOption("root:123456@tcp(127.0.0.1:3306)/ordersvc"),
	)
}

func TestHttpServer_Start(t *testing.T) {
	s := NewHttpServer(":8088")
	s.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK, Candide\n"))
	})
	s.Start()
	time.Sleep(1 * time.Hour)
}

type helloSvc struct {
	protos.UnimplementedHelloServiceServer
}

func (h *helloSvc) Receive(ctx context.Context, msg *protos.HelloMsg) (*protos.HelloMsg, error) {
	return msg, nil
}

func TestGrpcServer_Start(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8001")
		protos.RegisterHelloServiceServer(s, &helloSvc{})
		err := s.Start()
		if err != nil {
			return
		}
	}()
	client, err := NewGrpcClient("127.0.0.1:8001")
	if err != nil {
		panic(err)
	}
	res, err := protos.NewHelloServiceClient(client).
		Receive(context.TODO(), &protos.HelloMsg{Msg: "Hello Candide"})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Msg)
}
