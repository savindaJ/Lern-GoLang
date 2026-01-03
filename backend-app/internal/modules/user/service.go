package user

type UserService struct {
	repo UserRepository
}

func (s *UserService) Register(user *User) error {
	return s.repo.Create(user)
}
