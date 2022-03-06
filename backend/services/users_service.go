package services

import (
	"github.com/ris967/go-react/backend/domain/users"
	"github.com/ris967/go-react/backend/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// パスワード暗号化
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}

	user.Password = string(pwSlice[:])
	
}