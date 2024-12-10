package auth

type JWTService struct {
	SecretKey string
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
	}
}
