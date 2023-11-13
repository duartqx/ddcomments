package services

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	m "github.com/duartqx/ddcomments/domains/models"
	r "github.com/duartqx/ddcomments/domains/repositories"
)

type UserService struct {
	userRepository r.IUserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository r.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		validator:      validator.New(),
	}
}

func (us UserService) Create(user m.User) error {

	if err := us.validator.Struct(user); err != nil {
		return err
	}

	if *us.userRepository.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf("User exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return err
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}
