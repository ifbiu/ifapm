package skusvc

import (
	"ifapm"
	"protos"
	"skusvc/grpc"
)

func main() {
	ifapm.Infra.Init(
		ifapm.InfraMysqlDbOption("root:123456@tcp(127.0.0.1:3306)/ordersvc?charset=utf8"),
	)
	grpcserver := ifapm.NewGrpcServer(":8081")
	protos.RegisterSkuServiceServer(grpcserver, &grpc.SkuServer{})
	ifapm.EndPoint.Start()
}
