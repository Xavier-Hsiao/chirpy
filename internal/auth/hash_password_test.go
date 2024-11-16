package auth

import (
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "testpassword123"
	password2 := "testpassword456"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name      string
		password  string
		hash      string
		wantError bool
	}{
		{
			name:      "Correct password",
			password:  password1,
			hash:      hash1,
			wantError: false,
		},
		{
			name:      "Incorrect password",
			password:  "wrongpassword",
			hash:      hash1,
			wantError: true,
		},
		{
			name:      "Password does not match its hash",
			password:  password1,
			hash:      hash2,
			wantError: true,
		},
		{
			name:      "Empty password",
			password:  "",
			hash:      hash1,
			wantError: true,
		},
		{
			name:      "Invalid hash",
			password:  password1,
			hash:      "invalidhash",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantError {
				t.Errorf("CheckPasswordHash() got unexpected result, wantError %t", !tt.wantError)
			}
		})
	}
}
