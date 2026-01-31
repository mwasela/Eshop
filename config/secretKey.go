package config

import "os"

func GetSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// This is your fallback for local development
		secret = "ksbdks783g4fbyuyvff34i8hfb3i343dvubfiufi3ubfif3nfbi3fn3ff"
	}
	return []byte(secret)
}
