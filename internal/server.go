package checkin

import (
	"context"
	"ori/microservice/proto-gen/api_v1/ori/microservice"
)

type GRPCCheckinServer struct {
	microservice.CheckinServer
	service *CheckinService
}

func NewGRPCCheckinServer(s *CheckinService) microservice.CheckinServer {
	return &GRPCCheckinServer{
		service: s,
	}
}
func (s *GRPCCheckinServer) UserCheckIn(ctx context.Context, r *microservice.CheckInRequest) (*microservice.CheckInResponse, error) {
	venue := Venue{
		Name: r.GetVenue().GetName(),
	}
	location := Location{
		Longitude: r.GetLocation().GetLongitude(),
		Latitude:  r.GetLocation().GetLatitude(),
	}

	friends := []UserID{}
	for _, id := range r.GetFriendsWith() {
		friends = append(friends, UserID(id))
	}
	d := CheckInData{
		UserID:   UserID(r.GetUserId()),
		Location: location,
		Venue:    venue,
		Friends:  friends,
	}
	err := s.service.UserCheckIn(d)
	return &microservice.CheckInResponse{}, err
}
