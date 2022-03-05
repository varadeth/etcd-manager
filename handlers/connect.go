package handlers

import (
	"context"
	"log"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/varadeth/etcd-manager/swagger_gen/restapi/operations"
	"github.com/varadeth/etcd-manager/swagger_gen/restapi/operations/connect"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var ConnectionObj *clientV3.Client

func SetupHandlers(api *operations.EtcdManagerAPI) {
	api.ConnectConnectHandler = connect.ConnectHandlerFunc(connectEtcd)
}

func connectEtcd(params connect.ConnectParams) middleware.Responder {
	ConnectionObj, err := clientV3.New(clientV3.Config{
		Endpoints:   params.Body.Urls,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("Error : ", err)
	}
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ConnectionObj.Put(ctx, "sample_key", "sample_value")
	cancel()
	return connect.NewConnectOK()
}
