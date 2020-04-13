package main

import (
	"context"
	"fmt"
	"log"
	"ori/microservice/proto-gen/api_v1/ori/microservice"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Configuration struct {
	CheckinServiceName string `mapstructure:"checkin_service_name"`
	CheckinServicePort int    `mapstructure:"checkin_service_port"`
}

func init() {
	viper.BindEnv("CHECKIN_SERVICE_NAME")
	viper.BindEnv("CHECKIN_SERVICE_PORT")
}

func main() {
	conf := Configuration{}
	viper.SetEnvPrefix("cis")
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.CheckinServiceName, conf.CheckinServicePort), opts...)
	if err != nil {
		log.Fatalln("cannot connect to grpc server", err)
	}
	defer conn.Close()
	client := microservice.NewCheckinClient(conn)

	req := &microservice.CheckInUserRequest{
		Timestamp:   time.Now().String(),
		FriendsWith: []string{uuid.New().String(), uuid.New().String(), uuid.New().String()},
		Location: &microservice.CheckInUserRequest_Location{
			Longitude: 51.5007,
			Latitude:  0.1246,
		},
		Venue: &microservice.CheckInUserRequest_Venue{
			Name: "Big Ben",
		},
		UserId: "195b5c7f-4bc7-461b-8438-8beb9f9fcd16",
	}
	resp, err := client.CheckInUser(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)

	req = &microservice.CheckInUserRequest{
		Timestamp:   time.Now().String(),
		FriendsWith: []string{uuid.New().String(), uuid.New().String(), uuid.New().String()},
		Location: &microservice.CheckInUserRequest_Location{
			Longitude: 52.5007,
			Latitude:  0.1346,
		},
		Venue: &microservice.CheckInUserRequest_Venue{
			Name: "Other Place",
		},
		UserId: "195b5c7f-4bc7-461b-8438-8beb9f9fcd16",
	}
	resp, err = client.CheckInUser(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	resp2, err := client.GetCheckedInVenues(context.Background(), &microservice.GetCheckedInVenuesRequest{UserId: "195b5c7f-4bc7-461b-8438-8beb9f9fcd16"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp2.GetVenue())
}
