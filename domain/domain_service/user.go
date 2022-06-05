package domain_service

type UserDomainService interface {
	IsUserExists(mail string) bool
	IsUserEnabled(user_id string) bool
	IsCurrentUserMailDuplicated(user_id string, mail string) bool
}
