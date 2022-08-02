package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	SaveServices(inputMember InputMember) (Member, error)
	FindByEmailService(loginMember LoginMember) (Member, error)
	IsEmailAvailable(checkEmailInput CheckEmailInput) (bool, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) SaveServices(inputMember InputMember) (Member, error) {
	keyServicesMember := Member{}
	keyServicesMember.Nama = inputMember.Nama
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputMember.Password), bcrypt.MinCost)
	if err != nil {
		return keyServicesMember, err
	} else {
		keyServicesMember.Password = string(passwordHash)
		newServices, err := s.repository.SaveRepository(keyServicesMember)
		if err != nil {
			return newServices, err
		} else {
			return newServices, nil
		}
	}
}

func (s *services) FindByEmailService(loginMember LoginMember) (Member, error) {
	nama := loginMember.Nama
	password := loginMember.Password

	newFindByEmail, err := s.repository.FindByEmail(nama)
	if err != nil {
		return newFindByEmail, err
	} else {
		if newFindByEmail.Id == 0 {
			return newFindByEmail, errors.New("TIdak Ada Email")

		} else {
			err := bcrypt.CompareHashAndPassword([]byte(newFindByEmail.Password), []byte(password))
			if err != nil {
				return newFindByEmail, err
			} else {
				return newFindByEmail, nil
			}
		}

	}

}

func (s *services) IsEmailAvailable(checkEmailInput CheckEmailInput) (bool, error) {
	nama := checkEmailInput.Nama

	newFindByEmail, err := s.repository.FindByEmail(nama)
	if err != nil {
		return false, nil
	} else {
		if newFindByEmail.Id == 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
