package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/varadeth/etcd-manager/swagger_gen/restapi"
	"github.com/varadeth/etcd-manager/swagger_gen/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewEtcdManagerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8080
	defer server.Shutdown()
	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
