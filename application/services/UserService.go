package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	v "github.com/duartqx/ddcomments/application/validation"
	m "github.com/duartqx/ddcomments/domains/models"
	r "github.com/duartqx/ddcomments/domains/repositories"
)

type UserService struct {
	userRepository r.IUserRepository
	validator      *v.Validator
}

func GetNewUserService(userRepository r.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		validator:      v.NewValidator(),
	}
}

func (us UserService) Create(user m.User) error {

	if err := us.validator.Struct(user); err != nil {
		return fmt.Errorf("%s", string(*us.validator.JSON(err)))
	}

	if *us.userRepository.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Create(user); err != nil {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	return nil
}
