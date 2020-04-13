package checkin

type Location struct {
	Longitude float32
	Latitude  float32
}

type Venue struct {
	Name     string
	Location Location
}

type UserRepository interface {
	GetUser(uid UserID) (*User, error)
	Store(u *User) error
}

type UserID string

func NewUserID(uid string) UserID {
	return UserID(uid)
}

type User struct {
	ID       UserID
	CheckIns CheckIns
}

func (u *User) CheckIn(e CheckInEntry) {
	u.CheckIns = append(u.CheckIns, &e)
}

func (u User) LastCheckIn() CheckInEntry {
	return *u.CheckIns[len(u.CheckIns)-1]
}

type CheckIns []*CheckInEntry

type CheckInEntry struct {
	Venue Venue
	With  []UserID
}
