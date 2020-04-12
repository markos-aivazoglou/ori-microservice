package checkin

type Location struct {
	Longitude float32
	Latitude  float32
}

type Venue struct {
	Name string
}

type UserRepository interface {
	GetUser(uid UserID) (*User, error)
	Store(u *User) error
}

type UserID string

type User struct {
	UserID   UserID
	Checkins Checkins
}

func (u *User) CheckIn(e *CheckinEntry) {
	u.Checkins = append(u.Checkins, e)
}

func (u User) LastCheckIn() CheckinEntry {
	return *u.Checkins[len(u.Checkins)-1]
}

func (u User) UniqueVenues() []Venue {
	m := make(map[Venue]struct{})
	venues := []Venue{}
	for _, v := range u.Checkins {
		if _, ok := m[v.Venue]; !ok {
			venues = append(venues, v.Venue)
			m[v.Venue] = struct{}{}
		}
	}
	return venues

}

type Checkins []*CheckinEntry

type CheckinEntry struct {
	Location Location
	Venue    Venue
	With     []UserID
}
