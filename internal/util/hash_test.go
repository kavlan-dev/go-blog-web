package util

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "Empty password",
			password: "",
		},
		{
			name:     "Short password",
			password: "short",
		},
		{
			name:     "Normal password",
			password: "normalpassword",
		},
		{
			name:     "Long password",
			password: "verylongpasswordthatexceedstypicalpasswordlengths",
		},
		{
			name:     "Password with special characters",
			password: "password!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed := HashPassword(tt.password)

			// Проверяем что хэш не пустой
			if hashed == "" {
				t.Error("HashPassword() returned empty string")
			}

			// Проверяем что хэш имеет правильную длину (salt + hash в hex)
			// salt - 16 байт, hash - 32 байта, итого 48 байт в hex = 96 символов
			if len(hashed) != 96 {
				t.Errorf("HashPassword() returned hash of length %d, want 96", len(hashed))
			}
		})
	}
}
