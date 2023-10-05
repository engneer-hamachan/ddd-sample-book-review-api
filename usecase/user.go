package usecase

import (
	"app/domain/domain_service"
	"app/domain/model/user"
	"app/domain/repository"
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
	"fmt"
)

type UserUseCase interface {
	UserCreate(name string, mail string, password string) (user_id *string, err error)
	UserUpdate(user_id string, name string, mail string, password string) error
	UserDetail(user_id string) (*dto.UserForAuth, error)
}

type userUseCase struct {
	userRepository    repository.UserRepository
	userDomainService domain_service.UserDomainService
	userQueryService  query_service.UserQueryService
}

func NewUserUseCase(ur repository.UserRepository, uds domain_service.UserDomainService, uqs query_service.UserQueryService) UserUseCase {
	return &userUseCase{
		userRepository:    ur,
		userDomainService: uds,
		userQueryService:  uqs,
	}
}

func (uu *userUseCase) UserCreate(name string, mail string, password string) (user_id *string, err error) {
	b, err := uu.userDomainService.IsUserExists(mail)
	if err != nil {
		return nil, err
	}

	if b == false {
		err := fmt.Errorf("%s", "Duplicated User Mail Address")
		return nil, err
	}

	created_user, err := user.Create(name, mail, password)
	if err != nil {
		return nil, err
	}

	user_id, err = uu.userRepository.PostUsers(created_user)
	if err != nil {
		return nil, err
	}

	return user_id, nil
}

func (uu *userUseCase) UserUpdate(user_id string, name string, mail string, password string) error {

	b, err := uu.userDomainService.IsUserEnabled(user_id)
	if err != nil {
		return err
	}

	if b == false {
		err := fmt.Errorf("%s", "user_id is not found")
		return err
	}

	b, err = uu.userDomainService.IsCurrentUserMailDuplicated(user_id, mail)

	if b {
		err := fmt.Errorf("%s", "update mail is Dupulicate")
		return err
	}

	created_user, err := user.New(user_id, name, mail, password)
	if err != nil {
		return err
	}

	err = uu.userRepository.PutUsers(created_user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUseCase) UserDetail(user_id string) (*dto.UserForAuth, error) {
	user, err := uu.userQueryService.GetUserByID(user_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
