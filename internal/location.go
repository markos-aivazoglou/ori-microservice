package checkin

type Location struct {
	Longitude float32
	Latitude  float32
}

type Venue struct {
	Name string
}

type UserRepository interface {
	GetUser(uid UserID) (User, error)
	Store(u User) error
}

type UserID string

type User struct {
	UserID   UserID
	Checkins Checkins
}

func (u *User) CheckIn(e CheckinEntry) {
	u.Checkins = append(u.Checkins, e)
}

type Checkins []CheckinEntry

type CheckinEntry struct {
	Location Location
	Venue    Venue
	With     []UserID
}
