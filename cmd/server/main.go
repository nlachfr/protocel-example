package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/Neakxs/protocel-example/internal/db/sqlite"
	inet "github.com/Neakxs/protocel-example/internal/net"
	"github.com/Neakxs/protocel-example/internal/svc"
	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
	grpcinterceptor "github.com/Neakxs/protocel/authorize/interceptors/grpc"
	"github.com/Neakxs/protocel/options"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

type opts struct {
	apikey string
}

func (o *opts) GetFunctionOverloads() []*options.FunctionOverload {
	return nil
}
func (o *opts) GetVariableOverloads() []*options.VariableOverload {
	return []*options.VariableOverload{
		{
			Name:  "apikey",
			Value: o.apikey,
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8765")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", ".database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	wu := sqlite.NewWorkUnit(db)
	if err := wu.Migrate(context.Background()); err != nil {
		log.Fatal(err)
	}
	librarySvc := svc.NewLibraryService(wu)
	authzInterceptor, err := v1.NewLibraryServiceAuthzInterceptor(&opts{apikey: "mysecret"})
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcinterceptor.NewGRPCUnaryInterceptor(authzInterceptor),
		),
	)
	v1.RegisterLibraryServiceServer(grpcServer, inet.NewLibraryService(librarySvc))
	log.Printf("Serving on %v...\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
