package cmd

import (
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
	"log"
	"net"

	"google.golang.org/grpc"
)

func ServeGRPC() {
	// init dependency
	dependency := dependencyInject()

	s := grpc.NewServer()

	tokenvalidation.RegisterTokenValidationServer(s, dependency.TokenValidationAPI)

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
