package usecase

import (
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
)

type AuthUseCase interface {
	Login(mail string) (*dto.UserForAuth, error)
}

type authUseCase struct {
	authQueryService query_service.AuthQueryService
}

func NewAuthUseCase(aqs query_service.AuthQueryService) AuthUseCase {
	return &authUseCase{
		authQueryService: aqs,
	}
}

func (au authUseCase) Login(mail string) (*dto.UserForAuth, error) {
	user, err := au.authQueryService.GetUserByMailForAuth(mail)
	return user, err
}
