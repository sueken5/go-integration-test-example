package sample_test

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/sueken5/go-integration-test-example/pkg/apis/sample"
	infra_grpc "github.com/sueken5/go-integration-test-example/pkg/infrastructure/grpc"
	infra_mysql "github.com/sueken5/go-integration-test-example/pkg/infrastructure/persistence/mysql"
	"github.com/sueken5/go-integration-test-example/pkg/model"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sampleClient sample.SampleClient
	mysqlDB      *gorm.DB
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	var err error
	dsn := "root:root@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	if err := mysqlDB.AutoMigrate(&model.Post{}); err != nil {
		os.Exit(1)
	}

	repo := infra_mysql.NewPostRepository(mysqlDB)

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		os.Exit(1)
	}

	srv := infra_grpc.NewServer(lis, repo)

	go func() {
		if err := srv.Run(); err != nil {
			os.Exit(1)
		}
	}()

	conn, err := grpc.DialContext(
		ctx,
		lis.Addr().String(),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close()

	sampleClient = sample.NewSampleClient(conn)

	ret := m.Run()

	os.Exit(ret)
}
