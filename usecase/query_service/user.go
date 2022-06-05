package query_service

import (
	"app/usecase/query_service/dto"
)

type UserQueryService interface {
	GetUsers(current_page int) (*dto.Users, error)
	GetUserByID(user_id string) (*dto.UserForAuth, error)
}
