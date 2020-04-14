package service

import (
	"context"
	"log"
	checkin "ori/microservice/internal"
)

type Checkin struct {
	userRepository checkin.UserRepository
}

func NewCheckinService(ur checkin.UserRepository) *Checkin {
	return &Checkin{
		userRepository: ur,
	}
}

type CheckInData struct {
	Venue   checkin.Venue
	Friends []checkin.UserID
	UserID  checkin.UserID
}
type UserCheckInResult struct {
	UserID      checkin.UserID
	Venue       checkin.Venue
	FriendsWith []checkin.UserID
}

func (c *Checkin) UserCheckIn(ctx context.Context, d CheckInData) (*UserCheckInResult, error) {
	user, err := c.userRepository.GetUser(ctx, d.UserID)
	if err != nil {
		return nil, err
	}

	user.CheckIn(
		checkin.CheckInEntry{
			Venue: d.Venue,
			With:  d.Friends,
		},
	)
	err = c.userRepository.Store(ctx, user)
	if err != nil {
		return nil, err
	}
	log.Printf("check in for user %s created", user.ID)
	log.Printf("most recent checkin %v", user.LastCheckIn())
	result := &UserCheckInResult{
		UserID:      user.ID,
		Venue:       user.LastCheckIn().Venue,
		FriendsWith: user.LastCheckIn().With,
	}
	return result, nil
}

func (c *Checkin) GetVenues(ctx context.Context, uid checkin.UserID) ([]checkin.Venue, error) {
	user, err := c.userRepository.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return c.uniqueVenuesFromCheckins(user.CheckIns), nil

}

func (c Checkin) uniqueVenuesFromCheckins(ch checkin.CheckIns) []checkin.Venue {
	m := make(map[checkin.Venue]struct{})
	venues := []checkin.Venue{}
	for _, c := range ch {
		if _, ok := m[c.Venue]; !ok {
			venues = append(venues, c.Venue)
			m[c.Venue] = struct{}{}
		}
	}
	return venues

}
