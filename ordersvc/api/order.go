package api

import (
	"context"
	"github.com/google/uuid"
	"ifapm"
	"net/http"
	"ordersvc/grpcclient"
	"protos"
	"strconv"
)

type orderSvc struct {
}

var OrderSvc = &orderSvc{}

func (os *orderSvc) Add(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var (
		uid, _   = strconv.Atoi(values.Get("uid"))
		skuid, _ = strconv.Atoi(values.Get("sku_id"))
		num, _   = strconv.Atoi(values.Get("num"))
	)

	skuMsg, err := grpcclient.SkuClient.DecreaseStock(context.TODO(), &protos.Sku{
		Id:  int64(skuid),
		Num: int32(num),
	})
	if err != nil {
		return
	}

	_, err = ifapm.Infra.Db.ExecContext(context.TODO(), "insert into t_order(order_id, sku_id, num, price, uid) values (?,?,?,?,?)", uuid.New().String(), skuid, num, int(skuMsg.Price)*num, uid)
	if err != nil {
		ifapm.Logger.Error(context.TODO(), "createOrder", map[string]interface{}{
			"uid":    uid,
			"sku_id": skuid,
		}, err)
		ifapm.HttpStatus.Error(w, err.Error(), nil)
	}
	ifapm.HttpStatus.Ok(w)
}
