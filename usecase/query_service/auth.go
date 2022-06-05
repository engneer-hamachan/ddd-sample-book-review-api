package query_service

import (
	"app/usecase/query_service/dto"
)

type AuthQueryService interface {
	GetUserByMailForAuth(mail string) (*dto.UserForAuth, error)
}
