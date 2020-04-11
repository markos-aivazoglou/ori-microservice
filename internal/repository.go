package checkin

// We need a user repository
type Repository struct {
	store map[UserID]User
	UserRepository
}

func NewUserRepository() UserRepository {
	return &Repository{
		store: make(map[UserID]User),
	}
}
