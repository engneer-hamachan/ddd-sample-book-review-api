package domain_service

import "app/domain/repository"

type UserDomainService interface {
	IsUserExists(mail string) (bool, error)
	IsUserEnabled(user_id string) (bool, error)
	IsCurrentUserMailDuplicated(user_id string, mail string) (bool, error)
}

type userDomainService struct {
	userRepository repository.UserRepository
}

func NewUserDomainService(ur repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: ur,
	}
}

func (uds *userDomainService) IsUserExists(email string) (bool, error) {
	u, err := uds.userRepository.GetUserByEmail(email)

	if err != nil {
		return false, err
	}

	if u != nil {
		return false, nil
	}

	return true, nil
}

func (uds *userDomainService) IsUserEnabled(user_id string) (bool, error) {
	u, err := uds.userRepository.GetUser(user_id)

	if err != nil {
		return false, err
	}

	if u != nil {
		return false, nil
	}

	return true, nil
}

func (uds *userDomainService) IsCurrentUserMailDuplicated(user_id string, email string) (bool, error) {
	u, err := uds.userRepository.GetUserByEmail(email)

	if err != nil {
		return false, err
	}

	if string(u.GetUserId()) != user_id {
		return true, nil
	}

	return false, nil
}
