package config

import "os"

func JWTSecret() string {
    s := os.Getenv("JWT_SECRET")
    if s == "" {
        s = "dev-secret-replace-in-prod"
    }
    return s
}
