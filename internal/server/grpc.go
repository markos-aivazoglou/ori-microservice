package server

import (
	"context"
	checkin "ori/microservice/internal"
	"ori/microservice/internal/service"
	"ori/microservice/proto-gen/api_v1/ori/microservice"

	"google.golang.org/grpc"
)

type GRPCCheckinServer struct {
	microservice.CheckinServer
	service *service.Checkin
}

func NewGRPCCheckinServer(grpcServer *grpc.Server, service *service.Checkin) microservice.CheckinServer {
	server := &GRPCCheckinServer{
		service: service,
	}
	microservice.RegisterCheckinServer(grpcServer, server)
	return server
}
func (s *GRPCCheckinServer) CheckInUser(ctx context.Context, r *microservice.CheckInUserRequest) (*microservice.CheckInUserResponse, error) {
	location := checkin.Location{
		Longitude: r.GetLocation().GetLongitude(),
		Latitude:  r.GetLocation().GetLatitude(),
	}
	venue := checkin.Venue{
		Name:     r.GetVenue().GetName(),
		Location: location,
	}

	friends := []checkin.UserID{}
	for _, id := range r.GetFriendsWith() {
		friends = append(friends, checkin.UserID(id))
	}
	d := service.CheckInData{
		UserID:  checkin.UserID(r.GetUserId()),
		Venue:   venue,
		Friends: friends,
	}
	result, err := s.service.UserCheckIn(d)
	if err != nil {
		return nil, err
	}
	return &microservice.CheckInUserResponse{
		UserId:       string(result.UserID),
		FriendsCount: int32(len(result.FriendsWith)),
		VenueName:    result.Venue.Name,
	}, err
}

func (s *GRPCCheckinServer) GetCheckedInVenues(ctx context.Context, in *microservice.GetCheckedInVenuesRequest) (*microservice.GetCheckedInVenuesResponse, error) {
	venues, err := s.service.GetVenues(checkin.UserID(in.GetUserId()))
	if err != nil {
		return nil, err
	}
	vr := []*microservice.GetCheckedInVenuesResponse_Venue{}
	for _, venue := range venues {
		v := &microservice.GetCheckedInVenuesResponse_Venue{
			Name: venue.Name,
			Location: &microservice.GetCheckedInVenuesResponse_Venue_Location{
				Latitude:  venue.Location.Latitude,
				Longitude: venue.Location.Longitude,
			},
		}
		vr = append(vr, v)
	}
	response := &microservice.GetCheckedInVenuesResponse{
		Venue: vr,
	}

	return response, nil
}
