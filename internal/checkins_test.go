package checkin

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestUserCheckIn(t *testing.T) {
	friendsIDS := []UserID{
		NewUserID(uuid.New().String()),
		NewUserID(uuid.New().String()),
		NewUserID(uuid.New().String()),
	}
	userID := NewUserID(uuid.New().String())
	tests := []struct {
		name         string
		user         *User
		checkIn      CheckInEntry
		expectedUser *User
	}{
		{
			name: "check in user",
			user: &User{
				ID: userID,
			},
			checkIn: CheckInEntry{
				Venue: Venue{
					Name: "Test Venue 1",
					Location: Location{
						Latitude:  80.0,
						Longitude: 40.0,
					},
				},
				With: friendsIDS,
			},
			expectedUser: &User{
				ID: userID,
				CheckIns: CheckIns{
					&CheckInEntry{
						Venue: Venue{
							Name: "Test Venue 1",
							Location: Location{
								Latitude:  80.0,
								Longitude: 40.0,
							},
						},
						With: friendsIDS,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.CheckIn(tt.checkIn)
			if !reflect.DeepEqual(tt.user, tt.expectedUser) {
				t.Fatal("user check in incorrect")
			}
		})
	}
}
