package main

import (
	"fmt"
	"log"
	"net"
	checkin "ori/microservice/internal"
	"ori/microservice/internal/server"
	"ori/microservice/internal/service"
	"ori/microservice/internal/storage/inmem"

	"github.com/spf13/viper"

	"google.golang.org/grpc"
)

type Configuration struct {
	GRPCPort int `mapstructure:"grpc_port"`
}

func init() {
	viper.BindEnv("GRPC_PORT")
}

func main() {
	conf := Configuration{}
	viper.SetEnvPrefix("cis")
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
	// fmt.Println(conf)
	repo := inmem.NewUserRepository()
	service := service.NewCheckinService(repo)
	user := &checkin.User{
		ID: checkin.UserID("195b5c7f-4bc7-461b-8438-8beb9f9fcd16"),
	}
	_ = repo.Store(user)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.NewGRPCCheckinServer(grpcServer, service)
	fmt.Println(grpcServer.GetServiceInfo())
	log.Fatalln(grpcServer.Serve(lis))
}
