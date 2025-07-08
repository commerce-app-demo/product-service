package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/commerce-app-demo/product-service/internal/config"
	"github.com/commerce-app-demo/product-service/internal/repository/mysql"
	"github.com/commerce-app-demo/product-service/internal/server"
	"github.com/commerce-app-demo/product-service/internal/service"

	productspb "github.com/commerce-app-demo/product-service/proto"
	"google.golang.org/grpc"
)

func main() {
	p := ":50051"
	l, err := net.Listen("tcp", p)

	if err != nil {
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	cfg := config.LoadDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open(cfg.Driver, dsn)

	if err != nil {
		log.Fatalf("failed to open db connection error: %s", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("ping returned error: %s\nPlease check whether database is not running", err)
	}

	productspb.RegisterProductServiceServer(grpcServer, &server.ProductServiceServer{
		ProductService: &service.ProductService{
			Repo: &mysql.ProductRepository{DB: db},
		},
	})

	grpcServer.Serve(l)
}
