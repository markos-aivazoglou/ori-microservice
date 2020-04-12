package main

import (
	"fmt"
	"log"
	"net"
	checkin "ori/microservice/internal"
	"ori/microservice/proto-gen/api_v1/ori/microservice"

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
	fmt.Println(conf)
	repo := checkin.NewUserRepository()
	service := checkin.NewCheckinService(repo)
	server := checkin.NewGRPCCheckinServer(service)
	user := &checkin.User{
		UserID: checkin.UserID("195b5c7f-4bc7-461b-8438-8beb9f9fcd16"),
	}
	_ = repo.Store(user)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	microservice.RegisterCheckinServer(grpcServer, server)
	_ = grpcServer.Serve(lis)
}
