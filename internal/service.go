package checkin

type CheckinService struct {
	userRepository UserRepository
	// maybe add an event store here
}

func NewCheckinService(ur UserRepository) *CheckinService {
	return &CheckinService{
		userRepository: ur,
	}
}

type CheckInData struct {
	Location Location
	Venue    Venue
	Friends  []UserID
	UserID   UserID
}

func (c *CheckinService) UserCheckIn(d CheckInData) error {
	user, err := c.userRepository.GetUser(d.UserID)
	if err != nil {
		return err
	}

	user.CheckIn(
		CheckinEntry{
			Location: d.Location,
			Venue:    d.Venue,
			With:     d.Friends,
		},
	)
	err = c.userRepository.Store(user)
	if err != nil {
		return err
	}
	// publish event or something
	return nil
}
