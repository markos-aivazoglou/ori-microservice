package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	checkin "ori/microservice/internal"
	"ori/microservice/internal/service"

	"github.com/go-chi/chi"
)

type CheckInServer struct {
	service *service.Checkin
}

func NewCheckInServer(service *service.Checkin) *CheckInServer {
	return &CheckInServer{
		service: service,
	}
}

func (s *CheckInServer) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/user/{user-id}/venues", s.getUserVenuesHandler)
	return r
}

type VenueData []checkin.Venue

type userVenues struct {
	Venues VenueData `json:"venues"`
}

func (s *CheckInServer) getUserVenuesHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user-id")
	result, err := s.service.GetVenues(r.Context(), checkin.NewUserID(userID))
	fmt.Println(result, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := userVenues{
		Venues: result,
	}
	err = json.NewEncoder(w).Encode(&v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
