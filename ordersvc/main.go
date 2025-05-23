package main

import (
	"ifapm"
	"net/http"
	"ordersvc/api"
	"ordersvc/grpcclient"
	"protos"
)

func main() {
	ifapm.Infra.Init(ifapm.InfraMysqlDbOption("root:123456@tcp(127.0.0.1:3306)/ordersvc?charset=utf8"))

	client, err := ifapm.NewGrpcClient(":8081")
	if err != nil {
		return
	}
	grpcclient.SkuClient = protos.NewSkuServiceClient(client)

	httpserver := ifapm.NewHttpServer(":8066")
	httpserver.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	httpserver.HandleFunc("/order/add", api.OrderSvc.Add)
	ifapm.EndPoint.Start()
}
