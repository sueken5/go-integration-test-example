package main

import (
	"log"
	"net"

	"github.com/sueken5/go-integration-test-example/pkg/infrastructure/grpc"
	infra_mysql "github.com/sueken5/go-integration-test-example/pkg/infrastructure/persistence/mysql"
	"github.com/sueken5/go-integration-test-example/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := mysqlDB.AutoMigrate(&model.Post{}); err != nil {
		log.Fatal(err)
	}

	repo := infra_mysql.NewPostRepository(mysqlDB)

	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer(lis, repo)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
