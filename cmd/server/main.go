package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	checkin "ori/microservice/internal"
	grpc_server "ori/microservice/internal/servers/grpc"
	http_server "ori/microservice/internal/servers/http"
	"ori/microservice/internal/service"
	"ori/microservice/internal/storage/inmem"

	"github.com/spf13/viper"

	"google.golang.org/grpc"
)

type Configuration struct {
	GRPCPort int `mapstructure:"grpc_port"`
	HTTPPort int `mapstructure:"http_port"`
}

func init() {
	viper.BindEnv("GRPC_PORT")
	viper.BindEnv("HTTP_PORT")
}
func importSampleData(r checkin.UserRepository) {
	user := &checkin.User{
		ID: checkin.UserID("195b5c7f-4bc7-461b-8438-8beb9f9fcd16"),
	}
	r.Store(context.Background(), user)
}
func main() {

	// load config from env
	conf := Configuration{}
	viper.SetEnvPrefix("cis")
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	// initialise dependencies
	repo := inmem.NewUserRepository()
	service := service.NewCheckinService(repo)

	// import some sample user data
	importSampleData(repo)
	fmt.Println(repo)

	// initialize http server
	httpServer := http_server.NewCheckInServer(service)
	router := httpServer.Router()
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), router))
	}()

	// initialize grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpc_server.NewCheckInServer(grpcServer, service)

	log.Fatalln(grpcServer.Serve(lis))
}
