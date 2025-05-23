package grpc

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"protos"
	"skusvc/dao"
)

type SkuServer struct {
	protos.UnimplementedSkuServiceServer
}

func (s *SkuServer) DecreaseStock(ctx context.Context, sku *protos.Sku) (*protos.Sku, error) {
	info := dao.SkuDao.Get(ctx, sku.Id)
	if len(info) == 0 {
		return nil, status.Error(codes.NotFound, "sku not found")
	}
	decrRes, err := dao.SkuDao.Decr(ctx, sku.Id, sku.Num)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if affected, _ := decrRes.RowsAffected(); affected == 0 {
		return nil, status.Error(codes.PermissionDenied, "sku not decreased")
	}
	return &protos.Sku{
		Name:  cast.ToString(info["name"]),
		Num:   cast.ToInt32(info["num"]) - sku.Num,
		Id:    cast.ToInt64(info["id"]),
		Price: cast.ToInt32(info["price"]),
	}, nil
}
