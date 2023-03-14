package service

type JwtSecurityService interface {
	GenerateJWTToken(email string) (string, error)
	GetEmail(token string) (string, error)
	GenerateSecretKey() string
	IsValid(token string) bool
	VerifyToken(request string) error
}
