package service

import (
	"context"
	"errors"
	checkin "ori/microservice/internal"
	mock "ori/microservice/internal/mock/repositories"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCheckinService_UserCheckIn(t *testing.T) {
	userID := checkin.UserID(uuid.New().String())
	friendIDS := []checkin.UserID{
		checkin.NewUserID(uuid.New().String()),
		checkin.NewUserID(uuid.New().String()),
		checkin.NewUserID(uuid.New().String()),
		checkin.NewUserID(uuid.New().String()),
	}
	tests := []struct {
		name           string
		repository     *mock.UserRepository
		ciData         CheckInData
		expectErr      bool
		expectedResult *UserCheckInResult
	}{
		{
			name: "user checkin usecase - happy path",
			repository: &mock.UserRepository{
				GetUserFn: func(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
					return &checkin.User{
						ID:       userID,
						CheckIns: checkin.CheckIns{},
					}, nil
				},
				StoreFn: func(ctx context.Context, u *checkin.User) error {
					return nil
				},
			},
			ciData: CheckInData{
				Friends: friendIDS,
				UserID:  userID,
				Venue: checkin.Venue{
					Name: "Home",
					Location: checkin.Location{
						Latitude:  10,
						Longitude: 10,
					},
				},
			},
			expectErr: false,
			expectedResult: &UserCheckInResult{
				UserID:      userID,
				FriendsWith: friendIDS,
				Venue: checkin.Venue{
					Name: "Home",
					Location: checkin.Location{
						Latitude:  10,
						Longitude: 10,
					},
				},
			},
		},
		{
			name: "error getting user from repository",
			repository: &mock.UserRepository{
				GetUserFn: func(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
					return nil, errors.New("get user error")
				},
			},
			ciData: CheckInData{
				Friends: friendIDS,
				UserID:  userID,
				Venue: checkin.Venue{
					Name: "Home",
					Location: checkin.Location{
						Latitude:  10,
						Longitude: 10,
					},
				},
			},
			expectErr:      true,
			expectedResult: nil,
		},
		{
			name: "error storing user to repository after check-in",
			repository: &mock.UserRepository{
				GetUserFn: func(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
					return &checkin.User{
						ID:       userID,
						CheckIns: checkin.CheckIns{},
					}, nil
				},
				StoreFn: func(ctx context.Context, u *checkin.User) error {
					return errors.New("store user error")
				},
			},
			ciData: CheckInData{
				Friends: friendIDS,
				UserID:  userID,
				Venue: checkin.Venue{
					Name: "Home",
					Location: checkin.Location{
						Latitude:  10,
						Longitude: 10,
					},
				},
			},
			expectErr:      true,
			expectedResult: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCheckinService(tt.repository)
			result, err := service.UserCheckIn(context.Background(), tt.ciData)
			if err != nil == !tt.expectErr {
				t.Fatalf("was not expecting error but received %#v", err)
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Fatalf("was expecting %#v but got %#v", tt.expectedResult, result)
			}
		})
	}

}

func TestCheckinService_GetVenues(t *testing.T) {
	userID := checkin.UserID(uuid.New().String())
	tests := []struct {
		name           string
		repository     *mock.UserRepository
		userID         checkin.UserID
		expectErr      bool
		expectedResult []checkin.Venue
	}{
		{
			name: "user checkin usecase - happy path",
			repository: &mock.UserRepository{
				GetUserFn: func(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
					return &checkin.User{
						ID: userID,
						CheckIns: checkin.CheckIns{
							&checkin.CheckInEntry{
								Venue: checkin.Venue{
									Name: "test venue 1",
									Location: checkin.Location{
										Latitude:  10,
										Longitude: 10,
									},
								},
							},
							&checkin.CheckInEntry{
								Venue: checkin.Venue{
									Name: "test venue 2",
									Location: checkin.Location{
										Latitude:  11,
										Longitude: 11,
									},
								},
							},
						},
					}, nil
				},
				StoreFn: func(ctx context.Context, u *checkin.User) error {
					return nil
				},
			},
			userID:    userID,
			expectErr: false,
			expectedResult: []checkin.Venue{
				checkin.Venue{
					Name: "test venue 1",
					Location: checkin.Location{
						Latitude:  10,
						Longitude: 10,
					},
				},
				checkin.Venue{
					Name: "test venue 2",
					Location: checkin.Location{
						Latitude:  11,
						Longitude: 11,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCheckinService(tt.repository)
			result, err := service.GetVenues(context.Background(), tt.userID)
			if err != nil == !tt.expectErr {
				t.Fatalf("was not expecting error but received %#v", err)
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Fatalf("was expecting %#v but got %#v", tt.expectedResult, result)
			}
		})
	}

}
