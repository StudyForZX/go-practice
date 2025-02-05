package testify

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepository interface {
	CreateUser(user *User) (int, error)
	GetUserById(id int) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string, age int) (*User, error) {
	user := &User{Name: name, Age: age}
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (s *UserService) GetUserById(id int) (*User, error) {
	return s.repo.GetUserById(id)
}
