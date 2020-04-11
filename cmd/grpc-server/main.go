package main

import (
	checkin "ori/microservice/internal"
)

func main() {
	//  read env variables etc

	repo := checkin.NewUserRepository()
	service := checkin.NewCheckinService(repo)
	server := checkin.NewGRPCCheckinServer(service)
	_ = server
}
