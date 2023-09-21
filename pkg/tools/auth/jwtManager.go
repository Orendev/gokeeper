package auth

import (
	"fmt"
	"time"

	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	AuthorizationKey string = "authorization"
)

var (

	// ErrTokenInvalid означает, что токен не удалось проверить.
	ErrTokenInvalid = errors.New("JWT was invalid")

	// ErrUnexpectedSigningMethod означает, что токен был подписан с использованием неожиданного метода подписи.
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

const TokenExp = time.Hour * 3

type Options struct {
	CryptoKeyJWT     string        `env:"CRYPTO_KEY_JWT" env-default:"supersecretkey"`
	TokenDurationJWT time.Duration `env:"TOKEN_DURATION_JWT"`
}

// Claims — структура утверждений, которая включает стандартные утверждения и
// одно пользовательское UserID
// Её встраивают в структуру утверждений, определённую пользователем.
type Claims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"userID"`
	Role   string    `json:"role"`
}

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

func (manager *JWTManager) Generate(userID uuid.UUID) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedSigningMethod
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
