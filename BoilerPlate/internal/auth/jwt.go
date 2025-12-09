package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func NewToken(secret []byte, username string, ttl time.Duration) (string, error) {
    claims := Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
            IssuedAt: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}
