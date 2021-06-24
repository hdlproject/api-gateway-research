package main

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type (
	companyServer struct {
		UnimplementedCompanyServiceServer
	}
)

func newCompanyServer() *companyServer {
	return &companyServer{}
}

func (instance *companyServer) Get(ctx context.Context, request *GetCompanyRequest) (*GetCompanyResponse, error) {
	var companies []*Company
	_ = json.Unmarshal(readCompany(), &companies)

	return &GetCompanyResponse{
		Company: companies[request.Id-1],
	}, nil
}

func (instance *companyServer) GetAll(ctx context.Context, request *GetCompaniesRequest) (*GetCompaniesResponse, error) {
	var companies []*Company
	_ = json.Unmarshal(readCompany(), &companies)

	return &GetCompaniesResponse{
		Companies: companies,
	}, nil
}

type (
	grpcServer struct {
		server *grpc.Server
	}
)

func newGRPCServer() *grpcServer {
	zapLogger, _ := zap.NewProduction()

	grpcServer := &grpcServer{
		server: grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(
					grpc_zap.UnaryServerInterceptor(zapLogger),
				),
			),
		),
	}

	RegisterCompanyServiceServer(grpcServer.server, newCompanyServer())

	return grpcServer
}

func (instance *grpcServer) execute() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	if err := instance.server.Serve(lis); err != nil {
		panic(err)
	}
}
